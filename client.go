package alsgoclient

import (
	"encoding/json"
	"errors"

	"github.com/riftbit/ALS-Go/httpmodels"
)

//Client client object
type Client struct {
	Config
}

func (s *Client) processRequest(method string, data interface{}) ([]byte, error) {
	if s.IsAsync {
		go sendRequest(s.Uri, s.Login, s.Password, s.Timeout, method, data)
		return nil, nil
	} else {
		data, err := sendRequest(s.Uri, s.Login, s.Password, s.Timeout, method, data)
		return data, err
	}
}

//Add - call Log.Add
func (s *Client) Add(args httpmodels.RequestLogAdd) (httpmodels.ResponseLogAdd, error) {
	var resultStruct rLogAdd
	data, err := s.processRequest("Log.Add", args)
	if err != nil {
		return httpmodels.ResponseLogAdd{}, err
	}
	if err := json.Unmarshal(data, &resultStruct); err != nil {
		var errAnswer rpcErrorAnswer
		if err := json.Unmarshal(data, &errAnswer); err != nil {
			return httpmodels.ResponseLogAdd{}, err
		} else {
			return httpmodels.ResponseLogAdd{}, errors.New(errAnswer.Error.Code + " - " + errAnswer.Error.Message)
		}
		return httpmodels.ResponseLogAdd{}, err
	}
	return resultStruct.Result, nil
}

//AddCustom - call Log.AddCustom
func (s *Client) AddCustom(args httpmodels.RequestLogAddCustom) (httpmodels.ResponseLogAdd, error) {
	var resultStruct rLogAdd
	data, err := s.processRequest("Log.AddCustom", args)
	if err != nil {
		return httpmodels.ResponseLogAdd{}, err
	}
	if err := json.Unmarshal(data, &resultStruct); err != nil {
		var errAnswer rpcErrorAnswer
		if err := json.Unmarshal(data, &errAnswer); err != nil {
			return httpmodels.ResponseLogAdd{}, err
		} else {
			return httpmodels.ResponseLogAdd{}, errors.New(errAnswer.Error.Code + " - " + errAnswer.Error.Message)
		}
		return httpmodels.ResponseLogAdd{}, err
	}
	return resultStruct.Result, nil
}

//Get - call Log.Get
func (s *Client) Get(args httpmodels.RequestLogGetLog) (httpmodels.ResponseLogGet, error) {
	var resultStruct rLogGet
	data, err := s.processRequest("Log.Get", args)
	if err != nil {
		return httpmodels.ResponseLogGet{}, err
	}
	if err := json.Unmarshal(data, &resultStruct); err != nil {
		var errAnswer rpcErrorAnswer
		if err := json.Unmarshal(data, &errAnswer); err != nil {
			return httpmodels.ResponseLogGet{}, err
		} else {
			return httpmodels.ResponseLogGet{}, errors.New(errAnswer.Error.Code + " - " + errAnswer.Error.Message)
		}
		return httpmodels.ResponseLogGet{}, err
	}
	return resultStruct.Result, nil
}

//GetCount - call Log.GetCount
func (s *Client) GetCount(args httpmodels.RequestLogGetCount) (httpmodels.ResponseLogGetCount, error) {
	var resultStruct rLogGetCount
	data, err := s.processRequest("Log.GetCount", args)
	if err != nil {
		return httpmodels.ResponseLogGetCount{}, err
	}
	if err := json.Unmarshal(data, &resultStruct); err != nil {
		var errAnswer rpcErrorAnswer
		if err := json.Unmarshal(data, &errAnswer); err != nil {
			return httpmodels.ResponseLogGetCount{}, err
		} else {
			return httpmodels.ResponseLogGetCount{}, errors.New(errAnswer.Error.Code + " - " + errAnswer.Error.Message)
		}
		return httpmodels.ResponseLogGetCount{}, err
	}
	return resultStruct.Result, nil
}

//GetCategories - call Log.GetCategories
func (s *Client) GetCategories() (httpmodels.ResponseLogGetCategories, error) {
	var resultStruct rLogGetCategories
	data, err := s.processRequest("Log.GetCategories", struct{}{})
	if err != nil {
		return httpmodels.ResponseLogGetCategories{}, err
	}
	if err := json.Unmarshal(data, &resultStruct); err != nil {
		var errAnswer rpcErrorAnswer
		if err := json.Unmarshal(data, &errAnswer); err != nil {
			return httpmodels.ResponseLogGetCategories{}, err
		} else {
			return httpmodels.ResponseLogGetCategories{}, errors.New(errAnswer.Error.Code + " - " + errAnswer.Error.Message)
		}
		return httpmodels.ResponseLogGetCategories{}, err
	}
	return resultStruct.Result, nil
}

//Remove - call Log.Remove
func (s *Client) Remove(args httpmodels.RequestLogRemoveLog) (httpmodels.ResponseLogRemoveLog, error) {
	var resultStruct rLogRemoveLog
	data, err := s.processRequest("Log.Remove", args)
	if err != nil {
		return httpmodels.ResponseLogRemoveLog{}, err
	}
	if err := json.Unmarshal(data, &resultStruct); err != nil {
		var errAnswer rpcErrorAnswer
		if err := json.Unmarshal(data, &errAnswer); err != nil {
			return httpmodels.ResponseLogRemoveLog{}, err
		} else {
			return httpmodels.ResponseLogRemoveLog{}, errors.New(errAnswer.Error.Code + " - " + errAnswer.Error.Message)
		}
		return httpmodels.ResponseLogRemoveLog{}, err
	}
	return resultStruct.Result, nil
}

//RemoveCategory - call Log.RemoveCategory
func (s *Client) RemoveCategory(args httpmodels.RequestLogRemoveCategory) (httpmodels.ResponseLogRemoveCategory, error) {
	var resultStruct rLogRemoveCategory
	data, err := s.processRequest("Log.RemoveCategory", args)
	if err != nil {
		return httpmodels.ResponseLogRemoveCategory{}, err
	}
	if err := json.Unmarshal(data, &resultStruct); err != nil {
		var errAnswer rpcErrorAnswer
		if err := json.Unmarshal(data, &errAnswer); err != nil {
			return httpmodels.ResponseLogRemoveCategory{}, err
		} else {
			return httpmodels.ResponseLogRemoveCategory{}, errors.New(errAnswer.Error.Code + " - " + errAnswer.Error.Message)
		}
		return httpmodels.ResponseLogRemoveCategory{}, err
	}
	return resultStruct.Result, nil
}

//Transfer - call Log.Transfer
func (s *Client) Transfer(args httpmodels.RequestLogTransferLog) (httpmodels.ResponseLogTransferLog, error) {
	var resultStruct rLogTransferLog
	data, err := s.processRequest("Log.Transfer", args)
	if err != nil {
		return httpmodels.ResponseLogTransferLog{}, err
	}
	if err := json.Unmarshal(data, &resultStruct); err != nil {
		var errAnswer rpcErrorAnswer
		if err := json.Unmarshal(data, &errAnswer); err != nil {
			return httpmodels.ResponseLogTransferLog{}, err
		} else {
			return httpmodels.ResponseLogTransferLog{}, errors.New(errAnswer.Error.Code + " - " + errAnswer.Error.Message)
		}
		return httpmodels.ResponseLogTransferLog{}, err
	}
	return resultStruct.Result, nil
}

//ModifyTTL - call Log.ModifyTTL
func (s *Client) ModifyTTL(args httpmodels.RequestLogModifyTTL) (httpmodels.ResponseLogModifyTTL, error) {
	var resultStruct rLogModifyTTL
	data, err := s.processRequest("Log.ModifyTTL", args)
	if err != nil {
		return httpmodels.ResponseLogModifyTTL{}, err
	}
	if err := json.Unmarshal(data, &resultStruct); err != nil {
		var errAnswer rpcErrorAnswer
		if err := json.Unmarshal(data, &errAnswer); err != nil {
			return httpmodels.ResponseLogModifyTTL{}, err
		} else {
			return httpmodels.ResponseLogModifyTTL{}, errors.New(errAnswer.Error.Code + " - " + errAnswer.Error.Message)
		}
		return httpmodels.ResponseLogModifyTTL{}, err
	}
	return resultStruct.Result, nil
}
