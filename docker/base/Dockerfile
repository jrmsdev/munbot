FROM alpine:3

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.08.23"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN apk upgrade --no-cache -U && apk add --no-cache go

ARG MUNBOT_UID
ARG MUNBOT_GID

RUN addgroup -g ${MUNBOT_GID} munbot \
	&& adduser -h /home/munbot -g munbot -s /bin/sh -G munbot -D -u ${MUNBOT_UID} munbot

RUN mkdir -vp -m 0750 /munbot/src/master && chown -vR munbot:munbot /munbot

VOLUME /munbot/src/master

USER munbot:munbot
WORKDIR /home/munbot

ENV USER munbot
ENV HOME /home/munbot

ARG MUNBOT_UMASK

RUN printf 'umask %s\n' "${MUNBOT_UMASK}" >>/home/munbot/.profile
RUN printf "export PS1='%s '\n" '\h:\W\$' >>/home/munbot/.profile
RUN printf "export PATH=/munbot/bin:/usr/local/bin:/usr/bin:/bin\n" >>/home/munbot/.profile

ENV GOPATH /munbot
ENV GOENV /munbot/src/master/go.env

RUN go env && go version

CMD exec /bin/sh -l
