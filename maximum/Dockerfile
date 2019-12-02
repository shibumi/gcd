FROM golang:latest as build
# Disable CGO for not using dynamic libs
ENV CGO_ENABLED=0
# Set go workdir
WORKDIR /go/src/app/
# Copy project directory to /go/src/app/
COPY . .
# build project
RUN go build
FROM alpine:latest AS release
# Copy binary to app folder in release
COPY --from=build /go/src/app/maximum /app/maximum
# Set user and port
USER 9999:999
EXPOSE 8080
# Execute maximum on start
ENTRYPOINT ["/app/maximum"]
