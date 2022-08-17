FROM golang:1.16-alpine


WORKDIR /app

RUN apk update && apk add make

COPY ./ /app/
 

RUN go mod download

RUN make build


EXPOSE 8080

FROM alpine:latest

WORKDIR /root

COPY --from=0 /app/bin/ps-notification-service ./

CMD [ "./ps-notification-service", "start" ]