package main

import (
	"fmt"
	"go-centreon/centreonweb"
	"os"
	"strconv"
)

const (
	DEFAULT_CENTREON_URL      string = "http://127.0.0.1"
	DEFAULT_CENTREON_SSL      bool   = false
	DEFAULT_CENTREON_USER     string = "admin"
	DEFAULT_CENTREON_PASSWORD string = "centreon"
)

func main() {
	centreon_url, set := os.LookupEnv("CENTREON_URL")
	if !set {
		fmt.Println("Variable CENTREON_URL not set, using default value (" + DEFAULT_CENTREON_URL + ")")
		centreon_url = DEFAULT_CENTREON_URL
	}

	ssl, err := strconv.ParseBool(os.Getenv("CENTREON_ALLOW_UNVERIFIED_SSL"))
	if err != nil {
		fmt.Println("No boolean found in variable CENTREON_ALLOW_UNVERIFIED_SSL, using default value (" + strconv.FormatBool(DEFAULT_CENTREON_SSL) + ")")
		ssl = DEFAULT_CENTREON_SSL
	}

	centreon_user, set := os.LookupEnv("CENTREON_USER")
	if !set {
		fmt.Println("Variable CENTREON_USER not set, using default value (" + DEFAULT_CENTREON_USER + ")")
		centreon_user = DEFAULT_CENTREON_USER
	}

	centreon_password, set := os.LookupEnv("CENTREON_PASSWORD")
	if !set {
		fmt.Println("Variable CENTREON_PASSWORD not set, using default value (" + DEFAULT_CENTREON_PASSWORD + ")")
		centreon_password = DEFAULT_CENTREON_PASSWORD
	}

	c, err := centreonweb.New(centreon_url, ssl, centreon_user, centreon_password)

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

	cmd := centreonweb.Command{
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
