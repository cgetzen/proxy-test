FROM golang:1.19-alpine AS build
WORKDIR /build

COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download

COPY src/*.go ./
RUN go build -o /out
EXPOSE 8000

FROM alpine
COPY --from=build /out /server
COPY assets /assets
COPY config /config
ENTRYPOINT "/server"
