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

	var tps []centreonweb.Timeperiod
	tps, err = c.Timeperiods().Show("")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Timeperiods count:", len(tps))
	}

	var tmp centreonweb.Timeperiod
	tmp, err = c.Timeperiods().Get("24x7")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Timeperiod 24x7 found with ID:", tmp.ID)
	}

	tmpExists, err := c.Timeperiods().Exists("nonworkhours")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if tmpExists == true {
			fmt.Println("Timeperiod nonworkhours exists")
		} else {
			fmt.Println("Timeperiod nonworkhours does not exists")
		}
	}

	tmp = centreonweb.Timeperiod{
		Name:   "onlymonday",
		Alias:  "Only the monday",
		Monday: "00:00-24:00",
	}
	err = c.Timeperiods().Add(tmp)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Timeperiod added: " + tmp.Name)

	err = c.Timeperiods().Setparam(tmp.Name, centreonweb.TimeperiodMonday, tmp.Monday)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Set time period for Monday in " + tmp.Name)

	err = c.Timeperiods().Setexception(tmp.Name, "january 1", "00:00-02:00")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Set exception for " + tmp.Name)

	err = c.Timeperiods().Delexception(tmp.Name, "january 1")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Del exception for " + tmp.Name)

	err = c.Timeperiods().Del(tmp.Name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Timeperiod " + tmp.Name + " deleted")

	os.Exit(0)
}
