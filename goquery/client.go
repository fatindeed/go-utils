// Package goquery provides a goquery client.
package goquery

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

// NewClient returns a new GoqueryClient.
func NewClient(c *resty.Client) *GoqueryClient {
	return &GoqueryClient{Client: c}
}

// GoqueryClient is client to get goquery document.
type GoqueryClient struct {
	*resty.Client
}

// GetDoc get goquery document with url.
func (c *GoqueryClient) GetDoc(uri string) (*goquery.Document, error) {
	return c.GetDocWithQuery(uri, nil)
}

// GetDoc get goquery document with url.
func (c *GoqueryClient) GetDocWithQuery(uri string, query url.Values) (*goquery.Document, error) {
	resp, err := c.Client.R().SetQueryParamsFromValues(query).Get(uri)
	if err != nil {
		fullurl := uri
		if q := query.Encode(); q != "" {
			fullurl += "?" + q
		}
		return nil, fmt.Errorf("get %s failed: %w", fullurl, err)
	}
	r := bytes.NewReader(resp.Body())
	return goquery.NewDocumentFromReader(r)
}
