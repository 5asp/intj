FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.io"

WORKDIR /app 

COPY ../../go.mod .
COPY config/minio/app.yaml .

RUN go mod download

COPY . .

RUN go mod tidy
RUN go build -v -o minio cmd/minio/main.go

FROM scratch

WORKDIR /app


COPY --from=builder /app/minio .

COPY --from=builder /app/app.yaml .

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


EXPOSE 9000

ENTRYPOINT ["/app/minio"]