package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Jira configuration - Replace with your own Jira credentials and URL
const (
    jiraProject = ""
	jiraURL = ""
	jiraUsername = ""
	jiraToken = ""
)

type Issue struct {
	Key string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
		Assignee Assignee
		Description string `json:"description"`
	} `json:"fields"`
}

type Assignee struct {
	DisplayName string `json:"displayName"`
}

func ListTickets() {
	url := fmt.Sprintf("%s/rest/api/2/search", jiraURL)

    query := fmt.Sprintf("project = '%v' AND status != Closed ORDER BY created DESC", jiraProject)
	payload := []byte(`{"jql":"` + query + `"}`)

	response, err := sendRequest("POST", url, payload)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(response, &data); err != nil {
		log.Fatal(err)
	}

	issues := data["issues"].([]interface{})

	fmt.Println("Open Jira Tickets:")
	for _, issue := range issues {
		issueData := issue.(map[string]interface{})
		fields := issueData["fields"].(map[string]interface{})
		fmt.Printf("[%s] %s - %s\n", issueData["key"], fields["summary"], fields["assignee"].(map[string]interface{})["displayName"])
	}
}

func ViewTicket(ticketID string) {
	url := fmt.Sprintf("%s/rest/api/2/issue/%s", jiraURL, ticketID)

	response, err := sendRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	var issue Issue
	if err := json.Unmarshal(response, &issue); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ticket: %s\n", issue.Key)
	fmt.Printf("Summary: %s\n", issue.Fields.Summary)
	fmt.Printf("Assignee: %s\n", issue.Fields.Assignee.DisplayName)
	fmt.Printf("Description:\n%s\n", issue.Fields.Description)
}

func sendRequest(method, url string, payload []byte) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// Basic authentication using username and password
	// auth := base64.StdEncoding.EncodeToString([]byte(jiraUsername + ":" + jiraPassword))
	req.Header.Add("Authorization", "Bearer "+jiraToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

