package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AppList struct {
	Applications []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"applications"`
}

type Deployment struct {
	Revision    string
	Changelog   string
	Description string
	User        string
	Timestamp   string
}

func getApplicationID(applicationName string, apiKey string) (int, error) {
	var appList AppList
	url := fmt.Sprintf("%s/%s", apiURL, "applications.json")
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("An application called %s was not found", applicationName)
	}
	req.Header.Add("X-Api-Key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("An application called %s was not found", applicationName)
	}

	err = json.NewDecoder(resp.Body).Decode(&appList)
	if err != nil {
		return 0, fmt.Errorf("An application called %s was not found", applicationName)
	}

	for _, application := range appList.Applications {
		if application.Name == applicationName {
			return application.ID, nil
		}
	}
	return 0, fmt.Errorf("An application called %s was not found", applicationName)
}

func recordDeployment(applicationID int, revision, changelog, description, user, apiKey string) error {
	t := time.Now()
	timestamp := fmt.Sprintf("%d-%d-%dT%d:%d:%dZ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	d := Deployment{
		Revision:    revision,
		Changelog:   changelog,
		Description: description,
		User:        user,
		Timestamp:   timestamp,
	}
	url := fmt.Sprintf("%s/applications/%d/deployments.json", apiURL, applicationID)
	body, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("Couldn't marshal deployment: %s", err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("Couldn't create POST request: %s", err)
	}
	req.Header.Add("X-Api-Key", apiKey)
	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("Couldn't send POST request: %s", err)
	}
	return nil
}
