FROM golang:1.19-alpine AS builder
RUN mkdir /src /bin
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/ktopic-init .

FROM alpine:latest
RUN mkdir /bin
COPY --from=builder /bin/ktopic-init /bin/ktopic-init
ENTRYPOINT ["/bin/ktopic-init"]