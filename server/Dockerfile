FROM golang:alpine AS builder

# Move to working directory /build
WORKDIR /build

# Build ENV vars
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Use /dist as resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM scratch

COPY --from=builder /dist/main /
COPY --from=builder /build/templates ./templates

EXPOSE 3000

# Command to run
ENTRYPOINT ["/main"]
