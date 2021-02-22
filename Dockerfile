FROM golang:1.15.4-alpine3.12 AS base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/europeanaSearch

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Run Unit tests recursively
RUN CGO_ENABLED=0 go test -v ./...

RUN go build -o ./out/europeanaSearch .

FROM alpine:3.12
RUN apk add ca-certificates

COPY --from=base /tmp/europeanaSearch/out/europeanaSearch /app/europeanaSearch

EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/europeanaSearch"]