package utils

import (
	"net/http"
)

type ResponseSuccess struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ReturnJsonResponse function for returning data in JSON format
func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(resMessage)
}
