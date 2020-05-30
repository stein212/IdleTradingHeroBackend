# https://blog.ivorscott.com/ultimate-go-react-development-setup-with-docker#docker-basics

FROM golang:1.14-alpine as dev

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

FROM alpine:latest as prod

WORKDIR /app

COPY --from=dev /go/src/app/main .

EXPOSE 8080

CMD ["./main"]