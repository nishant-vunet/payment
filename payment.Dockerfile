FROM golang:1.17.7-alpine as builder
RUN apk add --no-cache ca-certificates git
RUN apk add build-base

# Set the Current Working Directory inside the container
WORKDIR /app/payment

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download && go mod tidy

COPY . .

ARG SKAFFOLD_GO_GCFLAGS

# Build the Go app
#RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app /app/payment/cmd/paymentsvc
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /app /app/payment/cmd/paymentsvc

RUN apk add --no-cache ca-certificates \
    busybox-extras net-tools bind-tools


ENV GOTRACEBACK=single

EXPOSE 80

ENTRYPOINT ["/app/paymentsvc"]
