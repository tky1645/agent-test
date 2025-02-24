# Build stage
FROM golang:1.24 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o app .

# Final stage
FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/app .
EXPOSE 8080
CMD [ "./app" ]