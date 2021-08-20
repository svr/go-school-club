FROM golang:1.17.0-alpine3.14 AS build

ENV PORT 3000

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go test -v && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server .



FROM scratch

WORKDIR /app
COPY --chown=10001:10001 --from=build /app/server .
USER 10001

EXPOSE ${PORT}

ENTRYPOINT ["./server"]