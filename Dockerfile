FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o viewtracker .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/viewtracker .

EXPOSE 6969
ENV GIN_MODE release
ENV PORT 6969


CMD ["./viewtracker"]
