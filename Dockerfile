FROM golang:1.17.5-alpine as builder

WORKDIR /app 

COPY /app /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM busybox

WORKDIR /app

COPY --from=builder /app/app /usr/bin/

EXPOSE 3030

ENTRYPOINT ["app"]