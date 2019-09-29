package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/smutel/go-centreon/centreonweb"
)

const (
	defaultCentreonURL      string = "http://127.0.0.1"
	defaultCentreonSsl      bool   = false
	defaultCentreonUser     string = "admin"
	defaultCentreonPassword string = "centreon"
)

func main() {
	centreonURL, set := os.LookupEnv("CENTREON_URL")
	if !set {
		fmt.Println("Variable CENTREON_URL not set, using default value (" + defaultCentreonURL + ")")
		centreonURL = defaultCentreonURL
	}

	ssl, err := strconv.ParseBool(os.Getenv("CENTREON_ALLOW_UNVERIFIED_SSL"))
	if err != nil {
		fmt.Println("No boolean found in variable CENTREON_ALLOW_UNVERIFIED_SSL, using default value (" + strconv.FormatBool(defaultCentreonSsl) + ")")
		ssl = defaultCentreonSsl
	}

	centreonUser, set := os.LookupEnv("CENTREON_USER")
	if !set {
		fmt.Println("Variable CENTREON_USER not set, using default value (" + defaultCentreonUser + ")")
		centreonUser = defaultCentreonUser
	}

	centreonPassword, set := os.LookupEnv("CENTREON_PASSWORD")
	if !set {
		fmt.Println("Variable CENTREON_PASSWORD not set, using default value (" + defaultCentreonPassword + ")")
		centreonPassword = defaultCentreonPassword
	}

	c, err := centreonweb.New(centreonURL, ssl, centreonUser, centreonPassword)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var hosts []centreonweb.Host
	hosts, err = c.Hosts().Show("")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Hosts count:", len(hosts))
	}

	var host centreonweb.Host
	host, err = c.Hosts().Get("Centeon-central")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Host Centeon-central found with ID:", host.ID)
	}

	hostExists, err := c.Hosts().Exists("Server")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if hostExists == true {
			fmt.Println("Host Server exists")
		} else {
			fmt.Println("Host Server does not exists")
		}
	}

	host = centreonweb.Host{
		Name:    "Server",
		Alias:   "New server",
		Address: "192.168.0.101",
	}
	err = c.Hosts().Add(host, "Central")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Host added: " + host.Name)

	err = c.Hosts().Setparam("Server", "notes", "A new server is there")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Set notes of host " + host.Name)

	hp, err := c.Hosts().Getparam("Server", "notes")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if len(hp) > 0 {
		fmt.Println("Notes of host "+host.Name+" are:", hp[0].Notes)
	}

	i, err := c.Hosts().Getinstance(host.Name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Host "+host.Name+" is monitored by poller:", i)

	macro := centreonweb.HostMacro{
		Name:        "WARNING",
		Value:       "80",
		IsPassword:  "0",
		Description: "Warning threshold",
	}

	err = c.Hosts().Setmacro(host.Name, macro)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Macro " + macro.Name + "added to host " + host.Name)

	hm, err := c.Hosts().Getmacro(host.Name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, m := range hm {
		fmt.Println(m.Name + "=" + m.Value)
	}

	err = c.Hosts().Delmacro(host.Name, macro.Name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Macro " + macro.Name + " deleted from host " + host.Name)

	err = c.Hosts().AddTemplate(host.Name, "App-DB-MySQL")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Template App-DB-MySQL added to host " + host.Name)

	tpls, err := c.Hosts().Gettemplates(host.Name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Templates:")
	for _, tpl := range tpls {
		fmt.Println(tpl.Name)
	}

	err = c.Hosts().Applytemplates(host.Name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = c.Hosts().Deltemplate(host.Name, "App-DB-MySQL")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Template App-DB-MySQL deleted from host " + host.Name)

	err = c.Hosts().Del(host.Name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Server " + host.Name + " deleted")

	os.Exit(0)
}
