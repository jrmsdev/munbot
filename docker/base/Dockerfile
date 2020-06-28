FROM alpine:3

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.06.27"

USER root:root
WORKDIR /root

RUN apk --no-cache -U upgrade

ARG MUNBOT_UID
ARG MUNBOT_GID

RUN addgroup -g ${MUNBOT_GID} munbot && \
	adduser -h /home/munbot -g munbot -s /bin/sh -G munbot -D -u ${MUNBOT_UID} munbot

USER munbot:munbot
WORKDIR /home/munbot

ENV USER munbot
ENV HOME /home/munbot

RUN printf "export PS1='%s '\n" '\u:\W\$' >.profile
RUN printf "export PATH=/usr/local/bin:/usr/bin:/bin\n" >>.profile

CMD exec /bin/sh -l
