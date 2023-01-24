package models

type DefaultResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
}
