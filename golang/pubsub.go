package main

import (
        "context"
        "encoding/json"
        "log"
        "net/http"
        "strings"
        "sync"

        "cloud.google.com/go/pubsub"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promauto"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Payload receveid from pub/sub
type Payload struct {
        Temperature float64 `json:"temperature"`
}

// Gauge metric temperature for Prometheus
var (
        temperatureGauge = promauto.NewGauge(prometheus.GaugeOpts{
                Name: "temperature_gauge",
                Help: "Temperature in °C",
        })
)

// Pulling messages from pub/sub
func pullMessages() {
        log.Println("STARTED PULLING MESSAGES")

        ctx := context.Background()

        // Set your $PROJECT_ID
        client, err := pubsub.NewClient(ctx, "temperature-grafana")
        if err != nil {
                log.Fatal(err)
        }

        // Set your $SUBSCRIPTION
        subID := "temperature-subscription"
        var mu sync.Mutex

        sub := client.Subscription(subID)
        cctx, cancel := context.WithCancel(ctx)
        err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
                mu.Lock()
                defer mu.Unlock()

                log.Print("Got message: " + string(msg.Data))

                if !containsInitPayload(string(msg.Data)) {
                        var t Payload

                        err := json.Unmarshal(msg.Data, &t)
                        if err != nil {
                                panic(err)
                        }

                        temperatureGauge.Set(t.Temperature)
                }

                msg.Ack()
        })
        if err != nil {
                cancel()
                log.Fatal(err)
        }
        cancel()
}

func containsInitPayload(payload string) bool {
        if strings.Contains(payload, "esp32-connected") {
                return true
        }
        return false
}

func main() {
        go pullMessages()

        log.Println("STARTED PROMETHEUS")

        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":2112", nil)
}