# Use an official Golang runtime as a parent image
FROM golang:1.9-alpine3.6 AS build_stage

# The base image does not have Git, but we need it
# to fetch WebSocket library, install it
RUN apk --no-cache --update add git

# Download all required packages
# Build will be performed by Go in the following steps
RUN ["go-wrapper", "download", "github.com/gorilla/websocket", "gopkg.in/mgo.v2"]

# Add the source files that will be compiled in the next step
ADD . ${GOPATH}/src/github.com/mrclayman/rest-and-go/gameserver/

# Set the working directory and compile and install the server
# binaries, and then remove the source files all in one swoop
WORKDIR ${GOPATH}/src
RUN CGO_ENABLED=0 GOOS=linux go build -a github.com/mrclayman/rest-and-go/gameserver && \
    go install github.com/mrclayman/rest-and-go/gameserver

# Using alpine image will let us reduce
# the final image size dramatically
FROM alpine:3.6 AS package_stage

# RUN apk --no-cache --update add bash
RUN apk --no-cache --update add ca-certificates

WORKDIR /server

COPY --from=build_stage /go/bin/gameserver .

# Make port 8000 available to the world outside this container
EXPOSE 8000

# Run the server when the image is instantiated
CMD ["./gameserver"]
# CMD [ "/bin/bash" ]
