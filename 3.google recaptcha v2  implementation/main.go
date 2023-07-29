package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func verifyRecaptcha(response string, remoteIP string) (bool, error) {
    fmt.Println("🚀 ~ file: main.go ~ line 16 ~ funcverifyRecaptcha ~ remoteIP : ", remoteIP)
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
		url.Values{
			"secret":   {"6LcEllglAAAAAMnAPLdjVx1C-g0Qdl6MCpUI-P3K"},
			"response": {response},
			// "remoteip": {remoteIP},
		})

	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
    fmt.Println("🚀  body : ", body)
	if err != nil {
		return false, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
    fmt.Println("🚀 result : ", result)
	if err != nil {
		return false, err
	}

	if result["success"] != true {
		return false, errors.New("reCAPTCHA verification failed")
	}

	return true, nil
}


func submitHandler(w http.ResponseWriter, r *http.Request) {
	recaptchaResponse := r.FormValue("g-recaptcha-response")
	remoteIP := r.RemoteAddr

	verified, err := verifyRecaptcha(recaptchaResponse, remoteIP)
    fmt.Println("🚀  verified : ", verified)
	if err != nil {
    fmt.Println("🚀 ~ file: main.go ~ line 52 ~ funcsubmitHandler ~ err : ", err)
		// Handle the error
	}

	if !verified {
		// Handle the case where the reCAPTCHA was not verified
        fmt.Println("🚀  Handle the case where the reCAPTCHA was not verified")
		// http.Redirect(w, r, "https://example.com", http.StatusSeeOther)
	} else {
		// Proceed with the form submission
        fmt.Println("🚀  Proceed with the form submission")
	}
}


func main() {

	// godotenv.Load()

	log.Println("Server will start at http://localhost:8000/")

	route := mux.NewRouter()

	log.Println("Loadeding Routes...")

	route.HandleFunc("/", RenderHome)

	route.HandleFunc("/submit", submitHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", route))
}

func RenderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "./index.html")
}