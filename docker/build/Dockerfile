FROM munbot/master:base

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.07.04"

USER munbot:munbot
WORKDIR /go/src/munbot

ENV USER munbot
ENV HOME /home/munbot

WORKDIR /go/src/munbot

ADD --chown=munbot bin/entrypoint.sh /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
