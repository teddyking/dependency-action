FROM golang:1.13 as builder

LABEL "com.github.actions.name"="dependency"
LABEL "com.github.actions.description"="Downloads and extracts dependencies required by workflow jobs."
LABEL "com.github.actions.icon"="arrow-down"
LABEL "com.github.actions.color"="white"

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd/ cmd/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o dependency-action cmd/dependency-action/main.go

FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=builder /workspace/dependency-action .
ENTRYPOINT ["/dependency-action"]

