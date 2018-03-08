package norris

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

// Client represents a new Chuck Norris fact client.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// Fact represents a new Chuck Norris fact.
type Fact struct {
	Value    string   `json:"value"`
	Category []string `json:"category,omitempty"`
}

// Category represents a Chuck Norris fact category.
type Category string

// Categories returns a collection of chuck norris fact categories.
func (c *Client) Categories() ([]Category, error) {
	url := fmt.Sprintf("%s/jokes/categories", c.baseURL)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var categories []Category

	err = json.Unmarshal(body, &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// Fact fetches a fresh Chuck Norris fact.
func (c *Client) Fact() (Fact, error) {
	url := fmt.Sprintf("%s/jokes/random", c.baseURL)

	var fact Fact

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fact, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fact, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fact, err
	}

	err = json.Unmarshal(body, &fact)
	if err != nil {
		return fact, err
	}

	return fact, nil
}

// Categorized fetches a fresh Chuck Norris fact for a particular category.
func (c *Client) Categorized(category Category) (Fact, error) {
	var fact Fact

	rawURL := fmt.Sprintf("%s/jokes/random", c.baseURL)
	endpoint, err := url.Parse(rawURL)
	if err != nil {
		return fact, err
	}

	if category != "" {
		params := url.Values{}
		params.Add("category", string(category))
		endpoint.RawQuery = params.Encode()
	}

	request, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return fact, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fact, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fact, err
	}

	err = json.Unmarshal(body, &fact)
	if err != nil {
		return fact, err
	}

	return fact, nil
}

// New returns a new chuck norris client.
func New(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 35 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 10 * time.Second,
				MaxIdleConnsPerHost: 500,
			},
		},
	}
}
