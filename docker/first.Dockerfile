# Dockerfile
FROM golang:1.14

WORKDIR /go/src/food-encyclopedia
COPY . .

# ENV GOBIN /Users/gavinjampani/go/bin

# RUN go get -d -v ./...
RUN go get
RUN cd /go/src/food-encyclopedia
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
# RUN go build .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -a -tags netgo -ldflags '-s -w' .
# RUN GOOS=linux GOARCH=amd64 go build -o app .

EXPOSE 8080

# CMD ["/Users/gavinjampani/go/src/src"]
CMD ["food-encyclopedia"]
