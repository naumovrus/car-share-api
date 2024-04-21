FROM golang:1.21.4-alpine3.18 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -a -installsuffix cgo -o /main cmd/api/main.go


EXPOSE 55111

FROM alpine:3
COPY --from=builder main /app/main
ENTRYPOINT ["/app/main"]