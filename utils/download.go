package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func downloadInput(url string, dst io.Writer) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	sessionKey := os.Getenv("SESSION")
	if sessionKey == "" {
		return fmt.Errorf("Session Cookie not found in .env file")
	}

	var adventOfCodeCookie = http.Cookie{
		Name:        "session",
		Value:       sessionKey,
		Domain:      "adventofcode.com",
		Path:        "/",
		Secure:      true,
		HttpOnly:    true,
		Partitioned: false,
		SameSite:    http.SameSiteNoneMode,
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&adventOfCodeCookie)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	err = handleResponse(res)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	_, err = io.Copy(dst, res.Body)
	if err != nil {
		return err
	}
	return nil
}
