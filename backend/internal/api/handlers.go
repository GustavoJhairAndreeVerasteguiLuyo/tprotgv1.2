package api

import (
    "net/http"

    "perimeter/internal/ai"
    "perimeter/internal/db"
    "perimeter/internal/ws"

    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *db.DB, hub *ws.Hub, aiClient *ai.Client) {
    api := r.Group("/api")
    {
        api.GET("/events", func(c *gin.Context) {
            events, _ := db.ListEvents(50)
            c.JSON(http.StatusOK, events)
        })

        api.POST("/infer", func(c *gin.Context) {
            var req struct{ Image string `json:"image"` }
            if err := c.BindJSON(&req); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "bad payload"})
                return
            }

            res, err := aiClient.Infer(req.Image)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            c.JSON(http.StatusOK, res)
        })

        api.GET("/ws", func(c *gin.Context) {
            hub.ServeWS(c.Writer, c.Request)
        })
    }
}
