package aria2

import "encoding/json"

const (
	methodAddUri = "aria2.addUri"
)

type _RpcRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type _RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type _RpcResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	Id      string          `json:"id"`
	Error   *_RpcError      `json:"error"`
	Result  json.RawMessage `json:"result"`
}

type _ResultVersion struct {
	Version         string   `json:"version"`
	EnabledFeatures []string `json:"enabledFeatures"`
}
