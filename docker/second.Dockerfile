# Dockerfile
# FROM golang:1.14
FROM alpine:latest

WORKDIR /

COPY static static
COPY food-encyclopedia  .

EXPOSE 8080

CMD ["./food-encyclopedia "]