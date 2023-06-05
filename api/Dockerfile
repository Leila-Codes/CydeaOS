FROM golang:1.19-alpine AS builder
LABEL authors="Leila-Codes"

RUN mkdir /build /build/bin/
WORKDIR "/build"
ADD ./go.mod ./go.sum ./
RUN go mod download
ADD . .
RUN go build -o /build/bin/api-server .

FROM alpine
RUN mkdir /build
COPY --from=builder /build/bin/api-server /bin/api-server
ENTRYPOINT ["/bin/api-server"]