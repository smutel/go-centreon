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

	var cmds []centreonweb.Command
	cmds, err = c.Commands().Show("")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Commands count:", len(cmds))
	}

	var cmd centreonweb.Command
	cmd, err = c.Commands().Get("check_centreon_dummy")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Command check_centreon_dummy found with ID:", cmd.ID)
	}

	cmdExists, err := c.Commands().Exists("check_centreon_dummy")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if cmdExists == true {
			fmt.Println("Command check_centreon_dummy exists")
		} else {
			fmt.Println("Command check_centreon_dummy does not exists")
		}
	}

	cmd = centreonweb.Command{
		Name: "check-host-alive",
		Type: "check",
		Line: "$USER1$/check_ping -H $HOSTADDRESS$ -w 3000.0,80% -c 5000.0,100% -p 1",
	}
	err = c.Commands().Add(cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Command added: " + cmd.Name)

	err = c.Commands().Setparam("check-host-alive", "type", "check")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Set type (check) of command " + cmd.Name)

	err = c.Commands().Del("check-host-alive")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Command " + cmd.Name + " deleted")

	os.Exit(0)
}
