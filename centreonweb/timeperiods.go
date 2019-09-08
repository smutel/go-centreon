package centreonweb

import (
	"encoding/json"

	pkgerrors "github.com/pkg/errors"
)

const timeperiodObject string = "TP"

// Const with days to use in setparam function of timeperiod
const (
	TimeperiodSunday    string = "sunday"
	TimeperiodMonday    string = "monday"
	TimeperiodTuesday   string = "tuesday"
	TimeperiodWednesday string = "wednesday"
	TimeperiodThursday  string = "thursday"
	TimeperiodFriday    string = "friday"
	TimeperiodSaturday  string = "saturday"
)

// ClientTimeperiods is used to store the client to interact with the Centreon API
type ClientTimeperiods struct {
	CentClient *ClientCentreonWeb
}

// Timeperiods is an array of Timeperiod to store the answer from Centreon API
type Timeperiods struct {
	Cmd []Timeperiod `json:"result"`
}

// Timeperiod struct is used to store parameters of a timeperiod
type Timeperiod struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	Sunday    string `json:"sunday"`
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
}

// Show lists available timeperiods
func (c *ClientTimeperiods) Show(name string) ([]Timeperiod, error) {
	respReader, err := c.CentClient.centreonAPIRequest("show", timeperiodObject, name)
	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var cmds Timeperiods
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&cmds)

	return cmds.Cmd, nil
}

// Get returns a specific timeperiod
func (c *ClientTimeperiods) Get(name string) (Timeperiod, error) {
	var objFound Timeperiod
	objs, err := c.Show(name)
	if err != nil {
		return objFound, err
	}

	for _, c := range objs {
		if c.Name == name {
			objFound = c
		}
	}

	if objFound.ID == "" {
		return objFound, pkgerrors.New("object " + name + " not found")
	}

	return objFound, nil
}

// Exists returns true if the timeperiod exists, false otherwise
func (c *ClientTimeperiods) Exists(name string) (bool, error) {
	objExists := false

	objs, err := c.Show(name)
	if err != nil {
		return objExists, err
	}

	for _, c := range objs {
		if c.Name == name {
			objExists = true
		}
	}

	return objExists, nil
}

// Add adds a new timeperiod
func (c *ClientTimeperiods) Add(tp Timeperiod) error {
	values := tp.Name + ";" + tp.Alias

	respReader, err := c.CentClient.centreonAPIRequest("add", timeperiodObject, values)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Del removes the specified timeperiod
func (c *ClientTimeperiods) Del(name string) error {
	respReader, err := c.CentClient.centreonAPIRequest("del", timeperiodObject, name)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setparam is used to change a specific parameters for a timeperiod
func (c *ClientTimeperiods) Setparam(name string, param string, value string) error {
	values := name + ";" + param + ";" + value

	respReader, err := c.CentClient.centreonAPIRequest("setparam", timeperiodObject, values)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setexception is used to add or change an exception for this timeperiod
func (c *ClientTimeperiods) Setexception(name string, param string, value string) error {
	values := name + ";" + param + ";" + value

	respReader, err := c.CentClient.centreonAPIRequest("setexception", timeperiodObject, values)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Delexception is used to remove an exception for this timeperiod
func (c *ClientTimeperiods) Delexception(name string, param string) error {
	values := name + ";" + param

	respReader, err := c.CentClient.centreonAPIRequest("delexception", timeperiodObject, values)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}