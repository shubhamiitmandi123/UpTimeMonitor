# FROM golang:1.15.2-alpine3.12
FROM golang:alpine as builder

LABEL maintainer="Shubham <shubham.choudhary@razorpay.com>"

RUN apk update && apk add --no-cache git

WORKDIR /UpTimeMonitor

COPY . .

RUN go mod download

# RUN go build .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o moniter .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /UpTimeMonitor/moniter .
COPY --from=builder /UpTimeMonitor/.env . 

EXPOSE 8080


CMD ["./moniter"]