FROM golang:1.22.1

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./key.pem /usr/certs/key.pem

COPY . .

RUN go build -v -o /usr/local/bin/app ./cmd/cert_service/main.go

CMD app
