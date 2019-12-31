# go-slack

A super minimaistic Slackchat webhook client for Go

Usage:

```go
import (
    "fmt"
    "github.com/roffe/go-slack"
)
func main() {
    webhookURL := "https://hooks.slack.com/services/..."
    msg := slack.Message{
        Blocks: []slack.Block{
            slack.Block{
                Type: slack.Section,
                Text: &slack.Text{
                    Type: slack.Markdown,
                    Text: "A part of a *message*",
                },
            },
            slack.Block{
                Type: slack.Divider,
            },
            slack.Block{
                Type: slack.Section,
                Text: &slack.Text{
                    Type: slack.Markdown,
                    Text: "some more text",
                },
            },
        },
    }
    resp, err := msg.Send(webhookURL)
    if err != nil {
        panic(err)
    }
    fmt.Println("Slack: ", resp)
}
```
