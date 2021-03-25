FROM golang:latest as builder
WORKDIR /app
COPY ./crud-service .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main ./cmd/main

###### Start a new stage ######

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]


