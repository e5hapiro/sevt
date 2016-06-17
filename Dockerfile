# Start from a Debian image with the latest Alpine version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

# Copy the local package files to the container's workspace.

RUN apk add --no-cache git \
    && go get github.com/golang/example/outyet \
    && apk del git


ADD . /go/src/github.com/golang/example/outyet

RUN apk add --no-cache git \
    && go get github.com/CiscoZeus/go-zeusclient \
    && apk del git

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/outyet

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/outyet

# Document that the service listens on port 9090.
EXPOSE 9090

ENV ZEUS_TOKEN=8dcccc00
