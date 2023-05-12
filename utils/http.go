package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/qinsheng99/go-domain-web/common/logger"
)

var transport = &http.Transport{
	MaxIdleConns:        250,
	MaxIdleConnsPerHost: 250,
	IdleConnTimeout:     120 * time.Second,
	DisableKeepAlives:   false,
	TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
}

type ReqImpl interface {
	CustomRequest(url, method string, bytesData interface{}, headers map[string]string, u url.Values, try bool, data interface{}) ([]byte, error)
}

type request struct {
	c *http.Client
}

func NewRequest(t *http.Transport) ReqImpl {
	if t == nil {
		t = transport
	}

	return &request{c: &http.Client{Transport: t}}
}

// MainRequest 所有公用的http请求
func (r *request) mainRequest(url, method string, bytesData interface{}, headers map[string]string) (resByte []byte, err error) {
	var req *http.Request
	var resp *http.Response
	err = Do(func(attempt int) (retry bool, err error) {
		req, err = http.NewRequest(method, url, r.getBody(bytesData))
		if err != nil {
			logger.Log.Errorf("reqURL:%s ;http new request err: %v", url, err)
			return attempt < 3, err
		}

		if http.MethodPost == method {
			req.Header.Set("Content-Type", "application/json")
		}

		for key, item := range headers {
			req.Header.Set(key, item)
		}

		resp, err = r.c.Do(req)
		if err != nil || resp == nil {
			return attempt < 3, err
		}

		defer resp.Body.Close()
		resByte, err = io.ReadAll(resp.Body)

		if resp.StatusCode > http.StatusMultipleChoices || resp.Body == nil {
			logger.Log.Error(fmt.Sprintf("statusCode is %d ,data : %s", resp.StatusCode, string(resByte)))
			return attempt < 3, errors.New(fmt.Sprintf("statusCode is %d ,data : %s", resp.StatusCode, string(resByte)))
		}

		return attempt < 3, err
	})
	return
}

func (r *request) CustomRequest(url, method string, bytesData interface{}, headers map[string]string, u url.Values, try bool, data interface{}) ([]byte, error) {
	var (
		bys []byte
		err error
	)

	if try {
		bys, err = r.mainRequest(r.getUrl(url, u), strings.ToUpper(method), bytesData, headers)
	} else {
		bys, err = r.noTryRequest(r.getUrl(url, u), strings.ToUpper(method), bytesData, headers)
	}

	if err != nil {
		return nil, err
	}

	if data == nil {
		return bys, nil
	}

	return bys, json.NewDecoder(bytes.NewReader(bys)).Decode(data)
}

// noTryRequest 所有公用的http请求无重试
func (r *request) noTryRequest(url, method string, bytesData interface{}, headers map[string]string) (resByte []byte, err error) {
	req, err := http.NewRequest(method, url, r.getBody(bytesData))
	if err != nil {
		logger.Log.Errorf("url:%s ;http new request err: %v", url, err)
		return
	}

	if http.MethodPost == method {
		req.Header.Set("Content-Type", "application/json")
	}

	for key, item := range headers {
		req.Header.Set(key, item)
	}

	resp, err := r.c.Do(req)
	if err != nil || resp == nil {
		logger.Log.Warnf("url:%s ;client Do err: %v", url, err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode > http.StatusMultipleChoices || resp.Body == nil {
		logger.Log.Errorf("url:%s ;http new request err: %v", url, resp.Status)
		return nil, errors.New(fmt.Sprintf("request err %s", resp.Status))
	}

	resByte, err = io.ReadAll(resp.Body)

	return
}

func (r *request) getBody(bytesData interface{}) io.Reader {
	var body = io.Reader(nil)
	switch t := bytesData.(type) {
	case []byte:
		body = bytes.NewReader(t)
	case string:
		body = strings.NewReader(t)
	case *strings.Reader:
		body = t
	case *bytes.Buffer:
		body = t
	default:
		body = nil
	}
	return body
}

func (r *request) getUrl(u string, values url.Values) string {
	path, err := url.Parse(u)
	if err != nil {
		return u
	}

	if len(values) > 0 {
		q := path.Query()

		for s, value := range values {
			for _, v := range value {
				q.Add(s, v)
			}
		}
		path.RawQuery = q.Encode()
	}
	return path.String()
}
