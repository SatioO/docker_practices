FROM golang:1.15rc2-alpine

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/satioO/gogogo

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/gogogo .

# This container exposes port 8081 to the outside world
EXPOSE 8081

CMD ["./out/gogogo"]