ARG KAFKA_BROKERS
ARG DB_PATH
ARG DB_USERNAME
ARG DB_PASSWORD
ARG DB_NAME
ARG MAIL
ARG MAIL_PASSWORD
ARG MAIL_HOST
ARG VERIFICATION_PRIVATE_KEY
ARG VERIFICATION_PUBLIC_KEY
ARG GRPC_ADDRESS
ARG GRPC_PORT
ARG GATEWAY_HOST

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

ENV KAFKA_BROKERS=$KAFKA_BROKERS
ENV DB_PATH=$DB_PATH
ENV DB_USERNAME=$DB_USERNAME
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_NAME=$DB_NAME
ENV MAIL=$MAIL
ENV MAIL_PASSWORD=$MAIL_PASSWORD
ENV MAIL_HOST=$MAIL_HOST
ENV VERIFICATION_PRIVATE_KEY=$VERIFICATION_PRIVATE_KEY
ENV VERIFICATION_PUBLIC_KEY=$VERIFICATION_PUBLIC_KEY
ENV GRPC_ADDRESS=$GRPC_ADDRESS
ENV GRPC_PORT=$GRPC_PORT
ENV GATEWAY_HOST=$GATEWAY_HOST

RUN echo ${MAIL_HOST}

COPY --from=build-env /go/src/github.com/ilhammhdd/kudaki-user-service/kudaki_user_service_app .

ENTRYPOINT ./kudaki_user_service_app