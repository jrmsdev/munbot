FROM jrmsdev/munbot:base

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.07.03"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN apk update --no-cache && apk add --no-cache go

RUN mkdir -vp -m 0750 /go/src/munbot && chown -vR munbot:munbot /go

VOLUME /go/src/munbot

USER munbot:munbot
WORKDIR /go/src/munbot

ENV USER munbot
ENV HOME /home/munbot

RUN printf "export PS1='%s '\n" '\u@build:\W\$' >>/home/munbot/.profile
RUN printf "export PATH=/go/bin:/usr/local/bin:/usr/bin:/bin\n" >>/home/munbot/.profile

ENV GOPATH /go
ENV GOENV /go/src/munbot/go.env

RUN go env && go version

ADD --chown=munbot bin/entrypoint.sh /usr/local/bin/
ENTRYPOINT /usr/local/bin/entrypoint.sh
