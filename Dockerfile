FROM golang
MAINTAINER Maxime Heckel <heckelmaxime@gmail.com> Borja Burgos <borja@tutum.co>

ADD . go/src/github.com/MaximeHeckel/basic_web_server

RUN go get gopkg.in/mgo.v2
RUN go install github.com/MaximeHeckel/basic_web_server

ENV NAME world

ENTRYPOINT /go/bin/basic_web_server

EXPOSE 80
