// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "goa simple sample": Application Media Types
//
// Command:
// $ goagen
// --design=design
// --out=$(GOPATH)\goa-sample
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// example Response (default view)
//
// Identifier: application/vnd.integer+json; view=default
type Integer struct {
	// id
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
}

// DecodeInteger decodes the Integer instance encoded in resp body.
func (c *Client) DecodeInteger(resp *http.Response) (*Integer, error) {
	var decoded Integer
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
