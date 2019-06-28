# go-centreon

[![Build Status](https://travis-ci.org/smutel/go-centreon.svg?branch=master)](https://travis-ci.org/smutel/go-centreon)
[![Go Report Card](https://goreportcard.com/badge/github.com/smutel/go-centreon)](https://goreportcard.com/report/github.com/smutel/go-centreon)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

`go-centreon` is a client SDK for Go applications to access centreon APIs.  
For now this project is focus on centreon web application.

## Usage

First of all you need to init a `centreonweb` object with the URL of centreon 
web application, a boolean to bypass the SSL check, the login and the password 
to use to connection to the API. This user should be setup in centreon to be 
able to reach the centreon API.

Then you can access each configuration object by an interface (Commands, Hosts, 
...). 

```go
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
  }
  
  fmt.Println("Commands count:", len(cmds))
}

```

## Contributing to this project

To contribute to this project I suggest to use the [centreon style guides](https://github.com/centreon/centreon/blob/master/CONTRIBUTING.md#centreon-style-guides).  
This project is still in progress but you can check the spreadsheet below to 
take one task to be implemented.

## Status of this project

### API coverage

|        Object Name        |    Object Code    |    Status    |
|:-------------------------:|:-----------------:|:------------:|
|            ACL            |        ACL        | **NOT DONE** |
|         Action ACL        |     ACLACTION     | **NOT DONE** |
|         ACL Groups        |      ACLGROUP     | **NOT DONE** |
|          Menu ACL         |      ACLMENU      | **NOT DONE** |
|        Resource ACL       |    ACLRESOURCE    | **NOT DONE** |
| Real time Ackcnowledgment | RTACKNOWLEDGEMENT | **NOT DONE** |
|      Centreon Broker      |   CENTBROKERCFG   | **NOT DONE** |
|          Commands         |        CMD        |    *DONE*    |
|          Contacts         |      CONTACT      | **NOT DONE** |
|     Contact Templates     |     CONTACTTPL    | **NOT DONE** |
|       Contact Groups      |         CG        | **NOT DONE** |
|        Dependencies       |        DEP        | **NOT DONE** |
|         Downtimes         |      DOWNTIME     | **NOT DONE** |
|    Real Time Downtimes    |     RTDOWNTIME    | **NOT DONE** |
|         Engine CFG        |     ENGINECFG     | **NOT DONE** |
|       Host Templates      |        HTPL       | **NOT DONE** |
|           Hosts           |        HOST       | **NOT DONE** |
|      Host Categories      |         HC        | **NOT DONE** |
|        Host Groups        |         HG        | **NOT DONE** |
|    Host Group Services    |     HGSERVICE     | **NOT DONE** |
|    Instances (Pollers)    |      INSTANCE     | **NOT DONE** |
|     LDAP Configuration    |        LDAP       | **NOT DONE** |
|        Resource CFG       |    RESOURCECFG    | **NOT DONE** |
|     Service Templates     |        STPL       | **NOT DONE** |
|          Services         |      SERVICE      | **NOT DONE** |
|       Service Groups      |         SG        | **NOT DONE** |
|     Service Categories    |         SC        | **NOT DONE** |
|          Settings         |      Settings     | **NOT DONE** |
|        Time Periods       |         TP        | **NOT DONE** |
|           Traps           |        TRAP       | **NOT DONE** |
|          Vendors          |       VENDOR      | **NOT DONE** |

## Examples

You can find some examples in the examples folder.  
Each example can be executed directly with command go run.  
You can set different environment variables for your test:
* CENTREON_URL to define the URL (and optionally the port) | DEFAULT=http://127.0.0.1
* CENTREON_ALLOW_UNVERIFIED_SSL to avoid checking the SSL certs (true or false) | DEFAULT=false
* CENTREON_USER to define the user | DEFAULT=admin
* CENTREON_PASSWORD to define the password | DEFAULT=centreon

```bash
$ export CENTREON_URL="http://10.164.48.254:8080"
$ export CENTREON_ALLOW_UNVERIFIED_SSL="true"
$ export CENTREON_USER="admin"
$ export CENTREON_PASSWORD="centreon"
$ go run examples/commands.go
Commands count: 101
Command check_centreon_dummy found with ID: 59
Command check_centreon_dummy exists
Command added: check-host-alive
Set type (check) of command check-host-alive
Command check-host-alive deleted
```

## Known bugs which can impact this framework

* Issue [7621](https://github.com/centreon/centreon/issues/7621)
