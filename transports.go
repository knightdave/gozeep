package gozeep

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Transport object handles all communication to the SOAP server.
type Transport struct {
	cache            []byte
	loadTimeout      time.Duration
	operationTimeout int
	session          *http.Client
}

// NewTransport constructs new instance of Transport
func NewTransport() *Transport {
	t := new(Transport)
	t.loadTimeout = 300
	tr := &http.Transport{
		Proxy:              http.ProxyFromEnvironment,
		MaxIdleConns:       10,
		IdleConnTimeout:    t.loadTimeout * time.Second,
		DisableCompression: false,
	}
	t.session = &http.Client{Transport: tr}
	return t
}

// Get is proxy to http.Get function
// TODO make possibility to run default get without params and headers
func (t *Transport) Get(address string, params map[string]string, headers map[string]string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", address, nil)

	req.Header.Add("Content-Type", "text/xml")
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	res, err := t.session.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, nil
}

// Post is proxy to http.Get function
// TODO make possibility to run default get without params and headers
func (t *Transport) Post(address string, message string, headers map[string]string) (*http.Response, error) {
	var m = []byte(message)
	req, _ := http.NewRequest("POST", address, bytes.NewBuffer(m))

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := t.session.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, nil
}

// PostXML the envelope xml element to the given address with the headers.
func (t *Transport) PostXML(address string, envelope Envelope, headers map[string]string) (*http.Response, error) {
	message := etreeToString(envelope)
	res, err := t.Post(address, message, headers)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Load the content from the given URL
// TODO implement cache
func (t *Transport) Load(urladdr string) ([]byte, error) {

	u, err := url.Parse(urladdr)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "http" || u.Scheme == "https" {
		resp, err := http.Get(urladdr)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return content, nil
	} else if u.Scheme == "file" {
		urladdr = urladdr[7:]
		content, err := ioutil.ReadFile(urladdr)
		if err != nil {
			return nil, err
		}
		return content, nil
	}

	return nil, fmt.Errorf("Wrong url scheme, only http/https/file")
}

// Settings context manager to temporarily overrule options.
// TODO implement temporary overrule
