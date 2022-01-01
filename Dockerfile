FROM golang:1.17
WORKDIR /
COPY . . 
RUN go mod download
EXPOSE 8080
CMD go run ./server/server.go $PORT