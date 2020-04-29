package aria2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/xid"
	"go.dead.blue/cli115/internal/pkg/util"
	"net/http"
)

type Client struct {
	hc *http.Client

	endpoint string
	token    string
}

func (c *Client) send(req *_RpcRequest, result interface{}) (err error) {
	// Fill common fields
	req.Jsonrpc = "2.0"
	req.Id = xid.New().String()
	// Make http request
	httpReq, _ := http.NewRequest(http.MethodPost, c.endpoint,
		bytes.NewReader(util.MustMarshal(req)))
	httpResp, err := c.hc.Do(httpReq)
	if err != nil {
		return
	}
	defer util.QuietlyClose(httpResp.Body)
	// Parse response
	d, resp := json.NewDecoder(httpResp.Body), _RpcResponse{}
	if err = d.Decode(&resp); err != nil {
		return err
	}
	if resp.Error != nil {
		err = errors.New(resp.Error.Message)
	} else {
		if result != nil {
			err = json.Unmarshal(resp.Result, result)
		}
	}
	return
}

func (c *Client) GetVersion() (ver string, err error) {
	req, result := _RpcRequest{
		Method: "aria2.getVersion",
		Params: []string{c.token},
	}, _ResultVersion{}
	if err = c.send(&req, &result); err == nil {
		ver = result.Version
	}
	return
}

func (c *Client) AddTask(url string, options map[string]interface{}) error {
	req := _RpcRequest{
		Method: methodAddUri,
		Params: []interface{}{
			c.token,
			[]string{url},
			options,
		},
	}
	return c.send(&req, nil)
}

func New(endpoint, token string) *Client {
	return &Client{
		hc: &http.Client{},
		// RPC endpoint
		endpoint: endpoint,
		// RPC token
		token: fmt.Sprintf("token:%s", token),
	}
}
