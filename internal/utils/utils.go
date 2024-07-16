package utils

import (
	"fmt"
	"net/url"
	"path"
)

func GetFileNameFromURL(fileURL string) (string, error) {
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}
	return path.Base(parsedURL.Path), nil //Get the last element from path
}

func GetAPIUrlFromIdInstance(idInstance string) (string, error) {
	if len(idInstance) < 4 {
		return "", fmt.Errorf("INCORRECT IdInstance")
	}
	return idInstance[:4], nil
}