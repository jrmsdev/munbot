FROM alpine:3

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.07.03"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN apk upgrade --no-cache -U

ARG MUNBOT_UID
ARG MUNBOT_GID

RUN addgroup -g ${MUNBOT_GID} munbot \
	&& adduser -h /home/munbot -g munbot -s /bin/sh -G munbot -D -u ${MUNBOT_UID} munbot

USER munbot:munbot
WORKDIR /home/munbot

ENV USER munbot
ENV HOME /home/munbot

ARG MUNBOT_UMASK

RUN printf 'umask %s\n' "${MUNBOT_UMASK}" >>/home/munbot/.profile
RUN printf "export PS1='%s '\n" '\u:\W\$' >>/home/munbot/.profile
RUN printf "export PATH=/usr/local/bin:/usr/bin:/bin\n" >>/home/munbot/.profile

CMD exec /bin/sh -l
