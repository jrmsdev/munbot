FROM alpine:3

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.07.03"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN apk upgrade --no-cache -U && apk add --no-cache go

ARG MUNBOT_UID
ARG MUNBOT_GID

RUN addgroup -g ${MUNBOT_GID} munbot \
	&& adduser -h /home/munbot -g munbot -s /bin/sh -G munbot -D -u ${MUNBOT_UID} munbot

RUN mkdir -vp -m 0750 /go/src/munbot && chown -vR munbot:munbot /go

VOLUME /go/src/munbot

USER munbot:munbot
WORKDIR /home/munbot

ENV USER munbot
ENV HOME /home/munbot

ARG MUNBOT_UMASK

RUN printf 'umask %s\n' "${MUNBOT_UMASK}" >>/home/munbot/.profile
RUN printf "export PS1='%s '\n" '\h:\W\$' >>/home/munbot/.profile
RUN printf "export PATH=/go/bin:/usr/local/bin:/usr/bin:/bin\n" >>/home/munbot/.profile

ENV GOPATH /go
ENV GOENV /go/src/munbot/go.env

RUN go env && go version

CMD exec /bin/sh -l
