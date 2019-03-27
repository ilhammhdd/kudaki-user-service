FROM golang:1.11-alpine AS build-env

RUN apk update
RUN apk upgrade
RUN apk add --no-cache curl
RUN apk add --no-cache git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/ilhammhdd/kudaki-user-service/
COPY . /go/src/github.com/ilhammhdd/kudaki-user-service/
RUN dep ensure
RUN go build -o kudaki_user_service_app

FROM alpine
COPY --from=build-env /go/src/github.com/ilhammhdd/kudaki-user-service/kudaki_user_service_app .

ENTRYPOINT ./kudaki_user_service_app