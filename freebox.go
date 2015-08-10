package freebox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
)

type ApiVersion struct {
	UID        string `json:"uid",omitempty`
	DeviceName string `json:"device_name",omitempty`
	Version    string `json:"api_version",omitempty`
	BaseURL    string `json:"api_base_url",omitempty`
	DeviceType string `json:"device_type",omitempty`
}

type Client struct {
	URL     string
	TrackId int
	Token   string

	apiVersion ApiVersion
	client     *http.Client
}

func New() *Client {
	return &Client{
		URL:     "http://mafreebox.free.fr/",
		Token:   "",
		TrackId: 42,
		client:  &http.Client{},
	}
}

func (c *Client) GetApiResource(resource string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.URL, resource)
	logrus.Debugf("GET %q", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("Status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

func (c *Client) Connect() error {
	body, err := c.GetApiResource("api_version")
	if err != nil {
		return err
	}
	logrus.Debugf("API response: %s", string(body))

	err = json.Unmarshal(body, &c.apiVersion)
	if err != nil {
		return err
	}

	logrus.Debugf("API version: UID=%q DeviceName=%q Version=%q BaseURL=%q DeviceType=%q", c.apiVersion.UID, c.apiVersion.DeviceName, c.apiVersion.Version, c.apiVersion.BaseURL, c.apiVersion.DeviceType)

	return nil
}
