package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// RenderHome Rendering the Home Page
func RenderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/index.html")
}

// StoreMessage function will verify the Captcha
func StoreMessage(response http.ResponseWriter, request *http.Request) {
	var storeMessageRequest StoreMessageRequest
	var captchaVerifyResponse interface{}
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&storeMessageRequest)
    fmt.Println("🚀 ~ file: routes-handlers.go ~ line 27 ~ funcStoreMessage ~ storeMessageRequest : ", storeMessageRequest)
	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {

		if storeMessageRequest.Message == "" {
			errorResponse.Code = http.StatusBadRequest
			errorResponse.Message = "Message can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {

			captchaPayloadRequest := url.Values{}
			captchaPayloadRequest.Set("secret", os.Getenv("CAPTCHA_SECRET_KEY"))
			captchaPayloadRequest.Set("response", storeMessageRequest.Token)

			verifyCaptchaRequest, _ := http.NewRequest("POST", os.Getenv("VERIFY_CAPTCHA_GOOGLE_API"), strings.NewReader(captchaPayloadRequest.Encode()))
			verifyCaptchaRequest.Header.Add("content-type", "application/x-www-form-urlencoded")
			verifyCaptchaRequest.Header.Add("cache-control", "no-cache")

			verifyCaptchaResponse, _ := http.DefaultClient.Do(verifyCaptchaRequest)

			decoder := json.NewDecoder(verifyCaptchaResponse.Body)
			decoderErr := decoder.Decode(&captchaVerifyResponse)
            fmt.Println("🚀 ~ file: routes-handlers.go ~ line 50 ~ funcStoreMessage ~ captchaVerifyResponse : ", captchaVerifyResponse)

			defer verifyCaptchaResponse.Body.Close()

			if decoderErr != nil {
				returnErrorResponse(response, request, errorResponse)
				return
			}

			var successResponse = SuccessResponse{
				Code:     http.StatusOK,
				Message:  "Your request is verified",
				Response: captchaVerifyResponse,
			}
			successJSONResponse, jsonError := json.Marshal(successResponse)

			if jsonError != nil {
				returnErrorResponse(response, request, errorResponse)
				return
			}

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(successResponse.Code)
			response.Write(successJSONResponse)
		}

	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage ErrorResponse) {
	httpResponse := &ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}
