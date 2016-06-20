# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/dhhenderson/sevt

RUN apk add --no-cache git \
    && go get github.com/CiscoZeus/go-zeusclient \
    && apk del git

# Build inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/dhhenderson/sevt


# Run the command by default when the container starts.
ENTRYPOINT /go/bin/sevt

EXPOSE 8080
