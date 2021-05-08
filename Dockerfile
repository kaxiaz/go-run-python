FROM golang:1.12.7-alpine3.10 AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY . ./
ENV GO111MODULE=on
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors
RUN go get github.com/lib/pq
RUN go get github.com/spf13/viper
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./main.go

FROM alpine:3.10
RUN apk update

# Install python/pip
ENV PYTHONUNBUFFERED=1
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip setuptools

WORKDIR /usr/bin
COPY ./config/development.yaml ./config/
COPY --from=build /go/src/app/bin /go/bin

EXPOSE 4110
ENTRYPOINT /go/bin/test
