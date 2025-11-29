package mqtt

import (
    "encoding/json"
    "log" // <-- import agregado y usado

    "perimeter/internal/ai"
    "perimeter/internal/db"
    "perimeter/internal/ws"

    paho "github.com/eclipse/paho.mqtt.golang"
)

type Client struct{ cli paho.Client }

func NewClient(broker string, hub *ws.Hub, aiClient *ai.Client, db *db.DB) (*Client, error) {
    log.Println("Initializing MQTT client for broker:", broker) // <- logging

    opts := paho.NewClientOptions().AddBroker(broker)
    opts.SetClientID("perimeter-backend")
    c := paho.NewClient(opts)

    if token := c.Connect(); token.Wait() && token.Error() != nil {
        log.Println("Error connecting to MQTT broker:", token.Error())
        return nil, token.Error()
    }
    log.Println("MQTT client connected successfully!")

    handler := func(client paho.Client, msg paho.Message) {
        log.Printf("Received message on topic %s\n", msg.Topic())
        var payload map[string]interface{}
        if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
            log.Println("Failed to parse message payload:", err)
            return
        }
        if _, ok := payload["image"]; ok {
            img := payload["image"].(string)
            res, err := aiClient.Infer(img)
            if err != nil {
                log.Println("AI inference error:", err)
                return
            }
            if len(res.Detections) > 0 {
                db.InsertEvent("site1", "intrusion", "detected")
                hub.Publish([]byte("intrusion: site1"))
                log.Println("Intrusion detected and published to hub")
            }
        }
    }

    if token := c.Subscribe("site/+/sensors/+", 0, handler); token.Wait() && token.Error() != nil {
        log.Println("Subscription error:", token.Error())
        return nil, token.Error()
    }
    log.Println("Subscribed to MQTT topics successfully")

    return &Client{cli: c}, nil
}

func (c *Client) Disconnect() {
    log.Println("Disconnecting MQTT client...")
    c.cli.Disconnect(250)
    log.Println("MQTT client disconnected")
}

