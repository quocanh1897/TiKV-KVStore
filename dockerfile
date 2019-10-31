FROM golang:alpine
RUN apk update && apk add --no-cache git

WORKDIR /

EXPOSE 10002

CMD ping localhost