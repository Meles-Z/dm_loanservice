package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"dm_loanservice/drivers/utils"
)

const (
	successCode            string = "00"
	successGetUserAuthCode string = "01"
	generalError           string = "99"
)

const (
	HeaderContentType   = "Content-Type"
	MIMEApplicationJSON = "application/json"
)

type Response struct {
	Code    string `valid:"Required" json:"code"`
	Message string `valid:"Required" json:"message"`
}

type DefaultResponse struct {
	Data interface{} `json:"data"`
	Response
}

func encodeDefaultResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusOK)
	if errResp, ok := response.(*utils.ApplicationError); ok {
		return responseJSONApplicationError(errResp, w)
	}
	if errResp, ok := response.(*utils.Error); ok {
		return responseJSONUtilError(errResp, w)
	}
	if errResp, ok := response.(error); ok {
		return responseJSONUtilErr(errResp, w)
	}
	return responseJSONOK(w, response, "")
}

func responseJSONUtilError(err *utils.Error, w http.ResponseWriter) error {
	resp := DefaultResponse{
		Response: Response{
			Code:    generalError,
			Message: "error",
		},
		Data: struct{}{},
	}
	if len(err.Message) > 0 {
		resp.Message = err.Message
	}
	if len(err.Code) > 0 {
		resp.Code = err.Code
	}
	return json.NewEncoder(w).Encode(resp)
}

func responseJSONOK(w http.ResponseWriter, response interface{}, code string) error {
	responseCode := successCode
	if code != "" {
		responseCode = code
	}
	if response == nil {
		response = struct{}{}
	}
	resp := DefaultResponse{
		Response: Response{
			Code:    responseCode,
			Message: "Success",
		},
		Data: response,
	}
	return json.NewEncoder(w).Encode(resp)
}

func responseJSONApplicationError(err *utils.ApplicationError, w http.ResponseWriter) error {
	resp := DefaultResponse{
		Response: Response{
			Code:    generalError,
			Message: "error",
		},
		Data: struct{}{},
	}
	if len(err.Message) > 0 {
		resp.Message = err.Message
	}
	if len(err.Code) > 0 {
		resp.Code = err.Code
	}
	return json.NewEncoder(w).Encode(resp)
}

func ProviderEncodeResponseWithCode(_ context.Context, w http.ResponseWriter, response interface{}, code string) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	errResp, ok := response.(*utils.ApplicationError)
	if ok {
		return responseJSONApplicationError(errResp, w)
	}
	return responseJSONOK(w, response, code)
}

func responseJSONUtilErr(err error, w http.ResponseWriter) error {
	resp := DefaultResponse{
		Response: Response{
			Code:    generalError,
			Message: err.Error(),
		},
		Data: struct{}{},
	}
	return json.NewEncoder(w).Encode(resp)
}
