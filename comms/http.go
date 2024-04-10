package comms

import (
	"net/http"
)

func MakeRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}