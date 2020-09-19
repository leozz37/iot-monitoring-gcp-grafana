FROM golang

COPY . /app
WORKDIR /app

ENV GOOGLE_APPLICATION_CREDENTIALS=/app/resources/service-account-key.json

RUN go mod download

CMD ["go", "run", "pubsub.go"]