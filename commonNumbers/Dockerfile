FROM golang:latest as build
# Copy project directory to app
COPY . /app
# build project
RUN go build
FROM scratch AS release
# Copy binary to app folder in release
COPY --from=build app/commonNumbers /app/commonNumbers
# Set user and port
USER 9999:999
EXPOSE 8080
# Execute commonNumbers on start
ENTRYPOINT ["/app/commonNumbers"]

