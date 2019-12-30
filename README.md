# go-slack

A super minimaistic Slackchat incoming webhook client for Go

Usage:

```go
import "github.com/roffe/go-slack"

func main() {
    msg := slack.Message{
        Attachements: []slack.Attachment{
            slack.Attachment{
                Fallback: "Text message for fallback for text clients",
                Color:    "good",
                Text:     "This *supports* _markdown_",
                Footer:   "send by go-slack",
            },
        },
    }
    resp, err := msg.SetURL(slackURL)
    if err != nil {
        panic(err)
    }
    fmt.Println("Slack: ", resp)
}
```
