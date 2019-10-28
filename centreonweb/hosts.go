package centreonweb

import (
	"encoding/json"

	pkgerrors "github.com/pkg/errors"
)

const hostObject string = "HOST"

// ClientHosts is used to store the client to interact with the Centreon API
type ClientHosts struct {
	CentClient *ClientCentreonWeb
}

// Hosts is an array of Host to store the answer from Centreon API
type Hosts struct {
	Host []Host `json:"result"`
}

// Host struct is used to store parameters of a host
type Host struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	Address  string `json:"address"`
	Activate string `json:"activate"`
}

// HostParams is an array of HostParam to store the answer from Centreon API
type HostParams struct {
	HostParam []HostParam `json:"result"`
}

// HostParam struct is used to store result of function getparam
type HostParam struct {
	Coords2D                   string `json:"2d_coords"`
	Coords3D                   string `json:"3d_coords"`
	ActionURL                  string `json:"action_url"`
	Activate                   string `json:"activate"`
	ActiveChecksEnables        string `json:"active_checks_enabled"`
	Address                    string `json:"address"`
	Alias                      string `json:"alias"`
	CheckCommand               string `json:"check_command"`
	CheckCommandArguments      string `json:"check_command_arguments"`
	CheckInterval              string `json:"check_interval"`
	CheckFreshness             string `json:"check_freshness"`
	CheckPeriod                string `json:"check_period"`
	ContactAdditiveInheritance string `json:"contact_additive_inheritance"`
	CgAdditiveInheritance      string `json:"cg_additive_inheritance"`
	EventHandler               string `json:"event_handler"`
	EventHandlerArguments      string `json:"event_handler_arguments"`
	EventHandlerEnables        string `json:"event_handler_enabled"`
	FirstNotificationDelay     string `json:"first_notification_delay"`
	FlapDetectionEnabled       string `json:"flap_detection_enabled"`
	FlapDetectionOptions       string `json:"flap_detection_options"`
	HostHighFlapThreshold      string `json:"host_high_flap_threshold"`
	HostLowFlapThreshold       string `json:"host_low_flap_threshold"`
	IconImage                  string `json:"icon_image"`
	IconImageAlt               string `json:"icon_image_alt"`
	MaxCheckAttempts           string `json:"max_check_attempts"`
	Name                       string `json:"name"`
	Notes                      string `json:"notes"`
	NotesURL                   string `json:"notes_url"`
	NotificationsEnabled       string `json:"notifications_enabled"`
	NotificationInterval       string `json:"notification_interval"`
	NotificationOptions        string `json:"notification_options"`
	NotificationPeriod         string `json:"notification_period"`
	RecoveryNotificationsDelay string `json:"recovery_notification_delay"`
	ObsessOverHost             string `json:"obsess_over_host"`
	PassiveChecksEnabled       string `json:"passive_checks_enabled"`
	ProcessPerfData            string `json:"process_perf_data"`
	RetainNonstatusInformation string `json:"retain_nonstatus_information"`
	RetainStatusInformation    string `json:"retain_status_information"`
	RetryCheckInterval         string `json:"retry_check_interval"`
	SnmpCommunity              string `json:"snmp_community"`
	SnmpVersion                string `json:"snmp_version"`
	StalkingOptions            string `json:"stalking_options"`
	StatusmapImage             string `json:"statusmap_image"`
	HostNotificationOptions    string `json:"host_notification_options"`
	Timezone                   string `json:"timezone"`
}

// Instances is an array of Instance to store the answer from Centreon API
type Instances struct {
	Instance []Instance `json:"result"`
}

// Instance struct is used to store parameters of an instance
type Instance struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// HostMacros is an array of HostMacro to store the answer from Centreon API
type HostMacros struct {
	HostMacro []HostMacro `json:"result"`
}

// HostMacro struct is used to store parameters of a macro
type HostMacro struct {
	Name        string `json:"macro name"`
	Value       string `json:"macro value"`
	IsPassword  string `json:"is_password"`
	Description string `json:"description"`
}

// HostTemplates is an array of HostTemplate to store the answer from
// Centreon API
type HostTemplates struct {
	HostTemplate []HostTemplate `json:"result"`
}

// HostTemplate struct is used to store parameters of a macro
type HostTemplate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Show lists available hosts
func (c *ClientHosts) Show(name string) ([]Host, error) {
	respReader, err := c.CentClient.centreonAPIRequest("show", hostObject,
		name)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var hosts Hosts
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&hosts)

	return hosts.Host, nil
}

// Get returns a specific host
func (c *ClientHosts) Get(name string) (Host, error) {
	var objFound Host

	if name == "" {
		return objFound, pkgerrors.New("name parameter cannot be empty when " +
			"calling Get function")
	}

	cmds, err := c.Show(name)
	if err != nil {
		return objFound, err
	}

	for _, c := range cmds {
		if c.Name == name {
			objFound = c
		}
	}

	if objFound.ID == "" {
		return objFound, pkgerrors.New("host " + name + " not found")
	}

	return objFound, nil
}

// Exists returns true if the host exists, false otherwise
func (c *ClientHosts) Exists(name string) (bool, error) {
	objExists := false

	if name == "" {
		return objExists, pkgerrors.New("name parameter cannot be empty when " +
			"calling Exists function")
	}

	cmds, err := c.Show(name)
	if err != nil {
		return objExists, err
	}

	for _, c := range cmds {
		if c.Name == name {
			objExists = true
		}
	}

	return objExists, nil
}

// Add adds a new host
func (c *ClientHosts) Add(host Host, instance string) error {
	values := host.Name + ";" + host.Alias + ";" + host.Address + ";;" + instance + ";"

	respReader, err := c.CentClient.centreonAPIRequest("add", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Del removes the specified host
func (c *ClientHosts) Del(name string) error {
	respReader, err := c.CentClient.centreonAPIRequest("del", hostObject,
		name)

	if name == "" {
		return pkgerrors.New("name parameter cannot be empty when calling Del " +
			"function")
	}

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setparam is used to change a specific parameters for a host
func (c *ClientHosts) Setparam(name string, param string,
	value string) error {

	if name == "" || param == "" || value == "" {
		return pkgerrors.New("name, param or value parameters cannot be empty " +
			"when calling Setparam function")
	}

	values := name + ";" + param + ";" + value

	respReader, err := c.CentClient.centreonAPIRequest("setparam", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Getparam is used to retrieve a specific parameters for a host
func (c *ClientHosts) Getparam(name string, param string) ([]HostParam, error) {

	if name == "" || param == "" {
		return nil, pkgerrors.New("name or param parameters cannot be empty when " +
			"calling Getparam function")
	}

	// Workaround to be sure we have an array as return value
	values := name + ";" + "name|" + param

	respReader, err := c.CentClient.centreonAPIRequest("getparam", hostObject,
		values)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var hps HostParams
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&hps)

	return hps.HostParam, nil
}

// Setinstance is used to set the instance (poller) for a host
func (c *ClientHosts) Setinstance(name string, instance string) error {

	if name == "" || instance == "" {
		return pkgerrors.New("name or instance parameters cannot be empty " +
			"when calling Setinstance function")
	}

	values := name + ";" + instance

	respReader, err := c.CentClient.centreonAPIRequest("setinstance", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Getinstance is used to retrieve the instance (poller) of a host
func (c *ClientHosts) Getinstance(name string) (string, error) {
	if name == "" {
		return "", pkgerrors.New("name parameter cannot be empty when calling " +
			"Getinstance function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("showinstance",
		hostObject, name)

	if err != nil {
		return "", err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var is Instances
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&is)

	if len(is.Instance) != 1 {
		return "", pkgerrors.New("too much results from API (!=1) when calling " +
			"Getinstance function")
	}

	return is.Instance[0].Name, nil
}

// Setmacro is used to add or to update a macro linked to a host
func (c *ClientHosts) Setmacro(hostName string, macro HostMacro) error {
	if hostName == "" || macro.Name == "" || macro.Value == "" ||
		macro.Description == "" {
		return pkgerrors.New("hostName or macro.Name or macro.Value or " +
			"macro.Description parameters cannot be empty when calling" +
			"Setmacro function")
	}

	if macro.IsPassword != "0" && macro.IsPassword != "1" {
		return pkgerrors.New("macro.IsPassword parameters should be equal to 0 " +
			"or 1 in Setmacro function")
	}

	values := hostName + ";" + macro.Name + ";" + macro.Value + ";" +
		macro.IsPassword + ";" + macro.Description

	respReader, err := c.CentClient.centreonAPIRequest("setmacro", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Getmacro is used to get a macro linked to a host
func (c *ClientHosts) Getmacro(hostName string) ([]HostMacro, error) {
	if hostName == "" {
		return nil, pkgerrors.New("hostName parameter cannot be empty when " +
			"calling Getmacro function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("getmacro",
		hostObject, hostName)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var ms HostMacros
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&ms)

	return ms.HostMacro, nil
}

// Delmacro is used to delete a macro linked to a host
func (c *ClientHosts) Delmacro(hostName string, macroName string) error {

	if hostName == "" || macroName == "" {
		return pkgerrors.New("hostName or macroName parameter cannot be " +
			"empty when calling Delmacro function")
	}

	values := hostName + ";" + macroName

	respReader, err := c.CentClient.centreonAPIRequest("delmacro", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Gettemplates is used to get a template linked to a host
func (c *ClientHosts) Gettemplates(hostName string) ([]HostTemplate, error) {
	if hostName == "" {
		return nil, pkgerrors.New("hostName parameter cannot be empty when " +
			"calling Gettemplate function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("gettemplate",
		hostObject, hostName)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var ms HostTemplates
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&ms)

	return ms.HostTemplate, nil
}

// AddTemplate is used to link a template to a host
func (c *ClientHosts) AddTemplate(hostName string, template string) error {
	if hostName == "" || template == "" {
		return pkgerrors.New("hostName or template parameters cannot be empty " +
			"when calling Setmacro function")
	}

	values := hostName + ";" + template

	respReader, err := c.CentClient.centreonAPIRequest("addtemplate", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Deltemplate is used to delete a template linked to a host
func (c *ClientHosts) Deltemplate(hostName string, templateName string) error {

	if hostName == "" || templateName == "" {
		return pkgerrors.New("hostName or templateName parameter cannot be " +
			"empty when calling Deltemplate function")
	}

	values := hostName + ";" + templateName

	respReader, err := c.CentClient.centreonAPIRequest("deltemplate", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Applytemplates is used to apply templates linked to a host
func (c *ClientHosts) Applytemplates(hostName string) error {

	if hostName == "" {
		return pkgerrors.New("hostName parameter cannot be empty when calling " +
			"Applytemplate function")
	}

	values := hostName

	respReader, err := c.CentClient.centreonAPIRequest("applytpl", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}
