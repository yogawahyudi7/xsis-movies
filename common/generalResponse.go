package common

type StatusResponse struct {
	Code  int
	Error error
}

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
