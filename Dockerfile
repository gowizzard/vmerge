# With this file you can build the docker image
# We load the language golang:alpine from docker hub load, change the workdir & build the application
# After that we create a prodution image that based on alpine:latest
# Copy the files from the build image & start the application
FROM golang:alpine AS build

WORKDIR /tmp/src/

COPY go.mod .

COPY . .
RUN go build -o vmerge

FROM alpine:latest AS production

WORKDIR /action

COPY --from=build /tmp/src/vmerge .

ENTRYPOINT ["/action/vmerge"]