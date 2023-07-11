package coin_cap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			Transport: loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
			Timeout: timeout,
		},
	}, nil
}

func (c *Client) GetAssets() (AssetsResponse, error) {
	resp, err := c.client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status:", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AssetsResponse{}, err
	}

	var r AssetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return AssetsResponse{}, err
	}

	return r, nil
}

func (c *Client) GetAsset(name string) (AssetResponse, error) {
	url := fmt.Sprintf("https://api.coincap.io/v2/assets/%s", name)
	resp, err := c.client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status:", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AssetResponse{}, err
	}

	var r AssetResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return AssetResponse{}, err
	}

	return r, nil
}
