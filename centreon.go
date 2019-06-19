package main

import (
	"fmt"
	"os"

	"github.com/smutel/go-centreon/centreonweb"
)

func main() {
	c, err := centreonweb.New("http://10.182.19.122:8080", true, "admin", "centreon")
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

	err = c.Commands().Setparam("check-host-alive", "type", "check")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = c.Commands().Del("check-host-alive")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
