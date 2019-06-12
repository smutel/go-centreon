package centreonweb

import (
	"go-centreon/client"
	"io"
	"net/http"
	"net/url"

	pkgerrors "github.com/pkg/errors"
)

const centreon_api_path string = "/centreon/api/index.php"

type centreonwebClient struct {
	MainClient *client.Client

	ConfigQuery  *url.Values
	ConfigHeader *http.Header

	AuthQuery  *url.Values
	AuthInput  *url.Values
	AuthHeader *http.Header
	AuthToken  string
}

type centreonwebConfigInput struct {
	Action string
	Object string
	Values string
}

func New(centreonURL string, insecure bool, username string, password string) (*centreonwebClient, error) {
	client, err := client.New(centreonURL, insecure)

	if err != nil {
		return nil, err
	}

	configQuery := &url.Values{}
	configQuery.Set("action", "action")
	configQuery.Add("object", "centreon_clapi")

	configHeader := &http.Header{}
	configHeader.Set("Content-Type", "application/json")

	authQuery := &url.Values{}
	authQuery.Set("action", "authenticate")

	authInput := &url.Values{}
	authInput.Set("username", username)
	authInput.Add("password", password)

	authHeader := &http.Header{}
	authHeader.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	return &centreonwebClient{
		MainClient:   client,
		ConfigQuery:  configQuery,
		ConfigHeader: configHeader,
		AuthQuery:    authQuery,
		AuthInput:    authInput,
		AuthHeader:   authHeader,
	}, nil
}

func (c *centreonwebClient) Commands() *commandsClient {
	return &commandsClient{c}
}

func (c *centreonwebClient) centreonApiRequest(action string, object string, values string) (io.ReadCloser, error) {
	err := c.login()
	if err != nil {
		return nil, err
	}

	input := &centreonwebConfigInput{
		Action: action,
		Object: object,
		Values: values,
	}
	body, err := input.toAPI()
	if err != nil {
		return nil, err
	}

	reqInputs := client.RequestInput{
		Method: http.MethodPost,
		Path:   centreon_api_path,
		Query:  c.ConfigQuery,
		Header: c.ConfigHeader,
		Body:   body,
	}

	respReader, err := c.MainClient.ExecuteRequest(reqInputs)
	if err != nil {
		return nil, err
	} else {
		return respReader, nil
	}
}

func (input *centreonwebConfigInput) toAPI() (map[string]interface{}, error) {
	params := 2
	if input.Values != "" {
		params += 1
	}
	result := make(map[string]interface{}, params)

	if input.Action != "" {
		result["action"] = input.Action
	} else {
		return nil, pkgerrors.New("action is mandatory to send request to centreon API")
	}

	if input.Object != "" {
		result["object"] = input.Object
	} else {
		return nil, pkgerrors.New("object is mandatory to send request to centreon API")
	}

	if input.Values != "" {
		result["values"] = input.Values
	}

	return result, nil
}
