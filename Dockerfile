# docker build . -t inventory-service 

# docker run \
# --name inventory-service \
# -p 8801:8801 \
# -e POSTGRES_DB_CONN_STR=host=localhost store=postgres password=password dbname=inventory-service port=5432 sslmode=disable TimeZone=Europe/Berlin \
# inventory-service

FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o inventory-service .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/inventory-service .

# Export necessary port
EXPOSE 8801

FROM scratch

COPY --from=builder /dist/inventory-service /
COPY ./config/config.json /config/config.json

# Command to run the executable
ENTRYPOINT ["/inventory-service"]