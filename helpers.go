package alsgoclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/riftbit/ALS-Go/httpmodels"
)

type rpcParams struct {
	Id      string      `json:"id"`
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type rpcError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type rpcAnswer struct {
	Id      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

type rpcErrorAnswer struct {
	rpcAnswer
	Error rpcError `json:"error"`
}

type rLogAdd struct {
	rpcAnswer
	Result httpmodels.ResponseLogAdd `json:"result"`
}

type rLogGet struct {
	rpcAnswer
	Result httpmodels.ResponseLogGet `json:"result"`
}

type rLogGetCount struct {
	rpcAnswer
	Result httpmodels.ResponseLogGetCount `json:"result"`
}

type rLogGetCategories struct {
	rpcAnswer
	Result httpmodels.ResponseLogGetCategories `json:"result"`
}

type rLogRemoveLog struct {
	rpcAnswer
	Result httpmodels.ResponseLogRemoveLog `json:"result"`
}

type rLogRemoveCategory struct {
	rpcAnswer
	Result httpmodels.ResponseLogRemoveCategory `json:"result"`
}

type rLogTransferLog struct {
	rpcAnswer
	Result httpmodels.ResponseLogTransferLog `json:"result"`
}

type rLogModifyTTL struct {
	rpcAnswer
	Result httpmodels.ResponseLogModifyTTL `json:"result"`
}

func sendRequest(url string, login string, password string, timeout int, method string, Data interface{}) ([]byte, error) {
	client := http.Client{Timeout: time.Duration(time.Duration(timeout) * time.Millisecond)}
	rand.Seed(time.Now().Unix())
	toSend := rpcParams{
		Id:      "request_" + string(rand.Intn(9999999)),
		Jsonrpc: "2.0",
		Method:  method,
		Params:  Data,
	}
	jsonedData, err := json.Marshal(toSend)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonedData))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(login, password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Unexpected backend error")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)
	if strings.Contains(bodyStr, "error") {
		var errStruct rpcErrorAnswer
		_ = json.Unmarshal(body, &errStruct)
		return body, errors.New(errStruct.Error.Message)
	}
	return body, nil
}
