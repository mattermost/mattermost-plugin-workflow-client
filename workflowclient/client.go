package workflowclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PluginAPI interface {
	PluginHTTP(*http.Request) *http.Response
}

type Client struct {
	http.Client
}

type pluginAPIRoundTripper struct {
	api PluginAPI
}

func (p *pluginAPIRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp := p.api.PluginHTTP(req)
	if resp == nil {
		return nil, fmt.Errorf("Failed to make interplugin request")
	}
	return resp, nil
}

func NewClientPlugin(api PluginAPI) *Client {
	client := &Client{}
	client.Transport = &pluginAPIRoundTripper{api}
	return client
}

func (c *Client) NotifyWorkflow(callbackURL string, params ActivateParameters) error {
	return c.NotifyWorkflows([]string{callbackURL}, params)
}

func (c *Client) NotifyWorkflows(callbackURLs []string, params ActivateParameters) error {
	out, err := json.Marshal(&params)
	if err != nil {
		return err
	}

	for _, callbackURL := range callbackURLs {
		req, err := http.NewRequest("POST", callbackURL, bytes.NewReader(out))
		if err != nil {
			return err
		}

		resp, err := c.Do(req)
		if err != nil {
			return err
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			respBody, _ := ioutil.ReadAll(resp.Body)
			return fmt.Errorf("Error response from workflow plugin. Error: %v, %v", resp.StatusCode, string(respBody))
		}
	}

	return nil
}
