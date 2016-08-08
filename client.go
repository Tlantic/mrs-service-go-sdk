package mrs_service_go_sdk

import (

	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
	"net/http/httputil"
	"bytes"
	"errors"
)

func NewClient(endpoint string,  organization string, application string) (*Client, error){
	if(organization == "" || application == ""){
		return &Client{}, errors.New("ORGANIZATION_OR_APPLICATION_IS_EMPTY")
	}

	return &Client{
		&http.Client{},
		organization,
		application,
		endpoint,
		nil,
	}, nil
}

func (c *Client) SetLog(log io.Writer) error {
	c.Log = log
	return nil
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json")

	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	resp, err := c.client.Do(req)
	c.log(req, resp)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &ErrorResponse{Response: resp}
		data, err := ioutil.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			json.Unmarshal(data, errResp)
		}

		return errResp
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}


func (c *Client) NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		var b []byte
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequest(method, url, buf)
}


func (c *Client) log(req *http.Request, resp *http.Response) {
	if c.Log != nil {
		reqDump, _ := httputil.DumpRequestOut(req, true)
		respDump, _ := httputil.DumpResponse(resp, true)

		c.Log.Write([]byte("Request: " + string(reqDump) + "\nResponse: " + string(respDump) + "\n\n"))

	}
}