package main

// ErrorResponse is struct for sending error message with code.
type ErrorResponse struct {
	Code    int
	Message string
}

// SuccessResponse is struct for sending error message with code.
type SuccessResponse struct {
	Code     int
	Message  string
	Response interface{}
}

// StoreMessageRequest is struct for requests of /storeMessage API
type StoreMessageRequest struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// CaptchaPayloadRequest is struct for payload of verify captcha API
type CaptchaPayloadRequest struct {
	secret   string
	response string
}
