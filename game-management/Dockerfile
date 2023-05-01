FROM golang:1.19-alpine AS builder
LABEL authors="Leila-Codes"

RUN mkdir /build /bin/
WORKDIR "/build"
ADD ./go.mod ./go.sum ./
RUN go mod download
ADD . .
RUN go build -o /bin/game-manager .

FROM alpine
RUN mkdir /build
COPY --from=builder /bin/game-manager /bin/game-manager
ENTRYPOINT ["/bin/game-manager"]