package utils

import (
	"fmt"
	"log"
	"net/http"
)

func HandleError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func handleResponse(res *http.Response) error {
	switch res.StatusCode {

	case http.StatusOK:
		return nil

	case http.StatusBadRequest:
		return fmt.Errorf("The request headers were malformed. Check your cookie.")

	case http.StatusUnauthorized, http.StatusForbidden:
		return fmt.Errorf("Unauthorized / Forbidden from accessing file. Your cookie is probably incorrect/expired.")

	case http.StatusTooManyRequests:
		return fmt.Errorf("Too many requests made to the AOC server. Wait a bit or watch the DDOS.")

	case http.StatusNotFound:
		return fmt.Errorf("The url doesn't exist.")

	default:
		return fmt.Errorf("Some 500 series error.")
	}
}
