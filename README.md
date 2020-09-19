# Monitoring IoT devices data with Golang, Google Cloud Platform and Grafana

![cover](images/cover.png)

This repostiry is a complement to my medium article. If you wanna follow step by step, check [my writing](medium.com).

## Architecture

This is how our infrastructure works:

![architecture](images/GCP.png)

## Golang

Running:

```bash
$ go run pubsub.go
```

Running Dockerfile:

```bash
$ docker build . -t metrics-exporter

$ docker run -p 2112:2112 metrics-exporter
```

## Prometheus

TODO:

## Grafana

TODO:
