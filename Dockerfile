############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/
COPY . .
# Fetch dependencies.
RUN go mod tidy
RUN go mod verify
# Build the binary.
RUN go build -o /go/bin/gotell
# Copy Support Files
RUN cp -r ./utilities /go/bin/
RUN chmod -R 777 ./utilities
RUN rm /go/bin/utilities/.DS_STORE
############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /go/bin/gotell /go/bin/gotell
COPY --from=builder /go/bin/utilities /go/bin/utilities

EXPOSE 9002
ENTRYPOINT ["/go/bin/gotell"]