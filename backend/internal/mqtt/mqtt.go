package mqtt

import (
    "encoding/json"
    "log"

    "perimeter/internal/ai"
    "perimeter/internal/db"
    "perimeter/internal/ws"

    paho "github.com/eclipse/paho.mqtt.golang"
)

type Client struct{ cli paho.Client }

func NewClient(broker string, hub *ws.Hub, aiClient *ai.Client, db *db.DB) (*Client, error) {
    opts := paho.NewClientOptions().AddBroker(broker)
    opts.SetClientID("perimeter-backend")
    c := paho.NewClient(opts)
    if token := c.Connect(); token.Wait() && token.Error() != nil { return nil, token.Error() }

    handler := func(client paho.Client, msg paho.Message) {
        var payload map[string]interface{}
        if err := json.Unmarshal(msg.Payload(), &payload); err == nil {
            if _, ok := payload["image"]; ok {
                img := payload["image"].(string)
                res, err := aiClient.Infer(img)
                if err == nil && len(res.Detections) > 0 {
                    db.InsertEvent("site1", "intrusion", "detected")
                    hub.Publish([]byte("intrusion: site1"))
                }
            }
        }
    }

    if token := c.Subscribe("site/+/sensors/+", 0, handler); token.Wait() && token.Error() != nil {
        return nil, token.Error()
    }
    return &Client{cli: c}, nil
}

func (c *Client) Disconnect() { c.cli.Disconnect(250) }
