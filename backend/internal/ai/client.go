package ai

import (
    "bytes"
    "encoding/json"
    "net/http"
    "time"
)

type Detection struct {
    Label      string  `json:"label"`
    Confidence float64 `json:"confidence"`
}

type InferResponse struct {
    Detections []Detection `json:"detections"`
}

type Client struct {
    base string
    http *http.Client
}

func NewClient(baseURL string) *Client {
    return &Client{base: baseURL, http: &http.Client{Timeout: 10 * time.Second}}
}

func (c *Client) Infer(imageBase64 string) (*InferResponse, error) {
    req := map[string]string{"image": imageBase64}
    b, _ := json.Marshal(req)
    resp, err := c.http.Post(c.base+"/infer", "application/json", bytes.NewReader(b))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    var r InferResponse
    json.NewDecoder(resp.Body).Decode(&r)
    return &r, nil
}
