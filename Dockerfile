FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go build -ldflags "-w -s"

FROM alpine:latest
RUN apk add -U --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/subji ./subji
COPY resources /app/resources
COPY config.production.json ./config.production.json
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
ENV GOYAVE_ENV=production
EXPOSE 8080
CMD dockerize -wait tcp://postgres:5432 -- ./subji
