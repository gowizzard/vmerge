# With this file you can build the docker image
# We load the language golang:alpine from docker hub load, change the workdir, load the modules & build the application
# After that we create a scratch image that based on alpine:latest, copy the files from the build image & start the application
FROM golang:alpine AS build

WORKDIR /tmp/src/

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o action

FROM alpine:latest AS production

WORKDIR /app/

COPY --from=build /tmp/src/action .

ENTRYPOINT ["/action"]