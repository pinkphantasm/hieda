FROM golang:1.21.3

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# TODO: Change `<service_name>` to the name of your service.
RUN go build -v -o /usr/local/bin/app ./cmd/<service_name>/main.go

CMD app
