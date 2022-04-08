FROM golang:latest

COPY ./src/go.mod /src/go.mod
COPY ./src/go.sum /src/go.sum
WORKDIR /src
RUN ls
RUN go mod download

COPY ./src /src

CMD go run main.go