FROM golang

ADD . go/src/github.com/MaximeHeckel/basic_web_server

RUN go install github.com/MaximeHeckel/basic_web_server

ENTRYPOINT /go/bin/basic_web_server

EXPOSE 80
