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

ENV KAFKA_BROKERS=178.62.107.160
ENV DB_PATH=tcp(178.62.107.160:3306)
ENV DB_USERNAME=root
ENV DB_PASSWORD=mysqlrocks
ENV DB_NAME=kudaki_user
ENV MAIL=service@kudaki.id
ENV MAIL_PASSWORD=OlahragaOtak2K19!
ENV MAIL_HOST=mail.privateemail.com
ENV VERIFICATION_PRIVATE_KEY=./verification_private.pem
ENV VERIFICATION_PUBLIC_KEY=./verification_public.pem

COPY --from=build-env /go/src/github.com/ilhammhdd/kudaki-user-service/kudaki_user_service_app .

ENTRYPOINT ./kudaki_user_service_app