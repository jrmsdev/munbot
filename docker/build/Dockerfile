FROM jrmsdev/munbot:base

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.07.03"

USER root:root
WORKDIR /root

RUN chown -R munbot:munbot /usr/lib/go/pkg/$(go env GOOS)_$(go env GOARCH)

USER munbot:munbot
WORKDIR /go/src/munbot

ENV USER munbot
ENV HOME /home/munbot

WORKDIR /go/src/munbot

ADD --chown=munbot bin/entrypoint.sh /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
