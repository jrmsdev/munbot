FROM jrmsdev/munbot:base

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.06.27"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN apk --no-cache add go sudo

ENV GOPATH /root/go
ENV GOBIN /usr/local/bin

RUN apk add --no-cache git && \
	go get -v golang.org/x/tools/cmd/godoc && \
	go generate -v golang.org/x/tools/godoc/static && \
	go install -v -i golang.org/x/tools/cmd/godoc && \
	apk del --no-cache --purge git && \
	rm -rf /root/go /root/.cache

RUN printf '\n%%munbot ALL=(ALL) NOPASSWD: ALL\n' >>/etc/sudoers

RUN mkdir -vp -m 0750 /go && mkdir -vp -m 0750 /go/src && \
	mkdir -vp -m 0750 /go/src/munbot && chown -vR munbot:munbot /go

VOLUME /go/src/munbot

ARG MUNBOT_UMASK

RUN printf 'umask %s\n' "${MUNBOT_UMASK}" >>/home/munbot/.profile
RUN printf 'export GOPATH=/go\n' >>/home/munbot/.profile

ADD --chown=munbot entrypoint.sh /usr/local/bin/entrypoint.sh

USER munbot:munbot
WORKDIR /go/src/munbot

ENV USER munbot
ENV HOME /home/munbot

ENV GOBIN ''
ENV GOPATH /go
RUN go env && go version

ENTRYPOINT /usr/local/bin/entrypoint.sh
