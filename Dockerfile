FROM golang:1.18.8-alpine3.16 AS BuildStage

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

RUN go build -o /main main.go


FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /main /main

EXPOSE 8080

ENTRYPOINT ["/main", "server"]
