FROM golang:1.16

ENV config=docker

WORKDIR /app

COPY .. /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 5000

ENTRYPOINT CompileDaemon --build="go build cmd/user/main.go" --command=./main