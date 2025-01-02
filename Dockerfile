# syntax=docker/dockerfile:1

FROM golang:1.22-bookworm as builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux  go build -o /docker-gs-ping ./cmd/api


# deployment image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates


WORKDIR /root/
COPY --from=builder /docker-gs-ping /docker-gs-ping

# Copy the .env file
COPY .env /root/.env

CMD [ "/docker-gs-ping" ]

EXPOSE 8000

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose

# Run