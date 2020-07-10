FROM munbot/master:base

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.07.10"

USER munbot:munbot
WORKDIR /munbot/src/master

ENV USER munbot
ENV HOME /home/munbot

WORKDIR /munbot/src/master

ADD --chown=munbot bin/entrypoint.sh /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
