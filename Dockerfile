FROM golang:1.25-alpine AS builder
WORKDIR /workspace
COPY go.mod go.mod
RUN go mod download
COPY cmd/ cmd/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} GO111MODULE=on go build -a -o server cmd/main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/server .
USER nonroot:nonroot

ENTRYPOINT ["/server"]