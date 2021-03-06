package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// Markdown string
	Markdown = "mrkdwn"
	// Divider works like a <hr>
	Divider = "divider"
	// Section string
	Section = "section"
	// Context string
	Context = "context"
	// PlainText string
	PlainText = "plain_text"
	// Image string
	Image = "image"
	// Button type
	Button = "button"
)

// Message is our slackchat message
type Message struct {
	url          string
	Text         string       `json:"text,omitempty"`
	Blocks       []Block      `json:"blocks,omitempty"`
	Attachements []Attachment `json:"attachments,omitempty"`
	ThreadTS     string       `json:"thread_ts,omitempty"`
	Markdown     bool         `json:"mrkdwn,omitempty"`
}

// Attachment is the structure of a message attachement
type Attachment struct {
	Fallback string `json:"fallback,omitempty"`
	Color    string `json:"color,omitempty"`
	Text     string `json:"text,omitempty"`
	Footer   string `json:"footer,omitempty"`
}

// Block type
type Block struct {
	Type      string     `json:"type"`
	Text      *Text      `json:"text,omitempty"`
	Accessory *Accessory `json:"accessory,omitempty"`
	Elements  []*Element `json:"elements,omitempty"`
	BlockID   string     `json:"block_id,omitempty"`
	Fields    []*Field   `json:"fields,omitempty"`
}

// Accessory type
type Accessory struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url,omitempty"`
	Text     *Text  `json:"text,omitempty"`
	AltText  string `json:"alt_text,omitempty"`
	URL      string `json:"url,omitempty"`
}

// Element type
type Element struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url,omitempty"`
	AltText  string `json:"alt_text,omitempty"`
	Emoji    bool   `json:"emoji,omitempty"`
	Text     string `json:"text,omitempty"`
}

// Field type
type Field struct {
	Type     string `json:"type"`
	Text     *Text  `json:"text,omitempty"`
	Emoji    bool   `json:"emoji,omitempty"`
	Verbatim bool   `json:"verbatim,omitempty"`
}

// Text type
type Text struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

// SetURL sets the webhook url
func (m *Message) setURL(webhookURL string) error {
	parsed, err := url.Parse(webhookURL)
	if err != nil {
		return fmt.Errorf("Failed to parse url: %s", err)
	}
	m.url = parsed.String()
	return nil
}

// JSON outputs the json string
func (m *Message) JSON() ([]byte, error) {
	b, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to json marshal %s", err)
	}
	return b, nil
}

// Send the message
func (m *Message) Send(webhookURL string) (string, error) {
	if err := m.setURL(webhookURL); err != nil {
		return "", err
	}

	b, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("Failed to json marshal %s", err)
	}
	r := bytes.NewReader(b)

	resp, err := http.Post(m.url, "application/json", r)
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

// AddBlock to the message
func (m *Message) AddBlock(b Block) {
	m.Blocks = append(m.Blocks, b)
}
