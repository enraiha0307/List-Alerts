FROM golang:1.16-alpine

LABEL maintainer="akanksha"

WORKDIR /go/src/gerrit.ext.net.nokia.com/AANM/sandbox-team3/blr-intern-projects/Akanksha/list_alerts

COPY . .

RUN  go build -o /list_alerts cmd/list_alerts/main.go

EXPOSE 8080

CMD [ "/list_alerts" ]