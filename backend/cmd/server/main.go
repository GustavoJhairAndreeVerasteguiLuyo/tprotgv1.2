package main

import (
    "log"
    "os"

    "perimeter/internal/api"
    "perimeter/internal/ai"
    "perimeter/internal/db"
    "perimeter/internal/mqtt"
    "perimeter/internal/ws"

    "github.com/gin-gonic/gin"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    dbURL := os.Getenv("DATABASE_URL")
    dbConn, err := db.New(dbURL)
    if err != nil {
        log.Fatalf("db init: %v", err)
    }
    defer dbConn.Close()

    aiURL := os.Getenv("AI_SERVICE_URL")
    aiClient := ai.NewClient(aiURL)

    hub := ws.NewHub()
    go hub.Run()

    mqttBroker := os.Getenv("MQTT_BROKER")
    if mqttBroker == "" {
        mqttBroker = "tcp://mosquitto:1883"
    }
    m, err := mqtt.NewClient(mqttBroker, hub, aiClient, dbConn)
    if err != nil {
        log.Fatalf("mqtt init: %v", err)
    }
    defer m.Disconnect()

    r := gin.Default()
    r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
    api.RegisterRoutes(r, dbConn, hub, aiClient)

    log.Printf("listening on :%s", port)
    r.Run(":" + port)
}
