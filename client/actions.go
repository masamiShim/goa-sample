// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "goa simple sample": actions Resource Client
//
// Command:
// $ goagen
// --design=design
// --out=$(GOPATH)\goa-sample
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// IDActionsPath computes a request path to the ID action of actions.
func IDActionsPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/v1/actions/%s", param0)
}

// 複数アクション(:id)
func (c *Client) IDActions(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewIDActionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewIDActionsRequest create the request corresponding to the ID action endpoint of the actions resource.
func (c *Client) NewIDActionsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// HelloActionsPath computes a request path to the hello action of actions.
func HelloActionsPath() string {

	return fmt.Sprintf("/api/v1/actions/hello")
}

// 挨拶します
func (c *Client) HelloActions(ctx context.Context, path string, name string) (*http.Response, error) {
	req, err := c.NewHelloActionsRequest(ctx, path, name)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewHelloActionsRequest create the request corresponding to the hello action endpoint of the actions resource.
func (c *Client) NewHelloActionsRequest(ctx context.Context, path string, name string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("name", name)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PingActionsPath computes a request path to the ping action of actions.
func PingActionsPath() string {

	return fmt.Sprintf("/api/v1/actions/ping")
}

// サーバーとの疎通確認
func (c *Client) PingActions(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewPingActionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPingActionsRequest create the request corresponding to the ping action endpoint of the actions resource.
func (c *Client) NewPingActionsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}