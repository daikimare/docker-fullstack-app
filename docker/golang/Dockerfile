FROM golang:alpine

RUN mkdir -p /usr/src/app/go/src
COPY go-app/src /go/src/

WORKDIR /usr/src/app/go/src

# RUN go mod download
RUN apk update \
	&& apk add --no-cache git \
	&& go get github.com/gin-gonic/gin \
	&& go get github.com/jinzhu/gorm \
	&& go get github.com/go-sql-driver/mysql

EXPOSE 8000

# CMD ["/usr/src/app/go/main"]
