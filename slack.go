package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Message is our slackchat message
type Message struct {
	url          string
	Text         string       `json:"text,omitempty"`
	Attachements []Attachment `json:"attachments,omitempty"`
}

// Attachment is the structure of a message attachement
type Attachment struct {
	Fallback string `json:"fallback,omitempty"`
	Color    string `json:"color,omitempty"`
	Text     string `json:"text,omitempty"`
	Footer   string `json:"footer,omitempty"`
}

// SetURL sets the webhook url
func (m *Message) SetURL(webhookURL string) error {
	parsed, err := url.Parse(webhookURL)
	if err != nil {
		return fmt.Errorf("Failed to parse url: %s", err)
	}
	m.url = parsed.String()
	return nil
}

// Send the message
func (m *Message) Send() (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("Failed to json marshal %s", err)
	}
	r := bytes.NewReader(b)

	request, err := http.NewRequest("POST", m.url, r)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", response), nil
}
