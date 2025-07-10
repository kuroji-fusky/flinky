FROM golang:1.24.3

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build cmd/server

CMD ["go", "run", "cmd/server/main.go"]
EXPOSE 4000