FROM golang:1.17
RUN mkdir /src
WORKDIR ./src
COPY . . 
RUN go mod download
EXPOSE 8080
CMD ["go","run","./server/server.go"]