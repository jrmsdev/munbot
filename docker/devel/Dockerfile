FROM jrmsdev/munbot:base

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.07.01"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN apk --no-cache add sudo go go-doc

ENV GOPATH /root/go
ENV GOBIN /usr/local/bin

RUN apk add --no-cache git && \
	go get -u -v golang.org/x/tools/cmd/godoc && \
	go generate -v golang.org/x/tools/godoc/static && \
	go install -v -i golang.org/x/tools/cmd/godoc && \
	go get -u -v github.com/go-bindata/go-bindata/... && \
	apk del --no-cache --purge git && \
	rm -rf /root/go /root/.cache

ADD ca/CA.pl.patch /tmp/CA.pl.patch

RUN apk add --no-cache patch perl openssl && \
	mv -v /etc/ssl/openssl.cnf /etc/ssl/openssl.cnf.orig && \
	ln -vs /go/src/munbot/docker/devel/ca/openssl.cnf /etc/ssl/openssl.cnf && \
	cp -va /etc/ssl/misc/CA.pl /etc/ssl/misc/CA.pl.orig && \
	patch /etc/ssl/misc/CA.pl /tmp/CA.pl.patch && \
	rm -vf /tmp/CA.pl.patch && \
	apk del --no-cache --purge patch

RUN printf '\n%%munbot ALL=(ALL) NOPASSWD: ALL\n' >>/etc/sudoers

RUN ln -vs /usr/share/doc/go/doc /usr/lib/go/doc

RUN mkdir -vp -m 0750 /go && mkdir -vp -m 0750 /go/src && \
	mkdir -vp -m 0750 /go/src/munbot && chown -vR munbot:munbot /go

VOLUME /go/src/munbot

ARG MUNBOT_UMASK

RUN printf 'umask %s\n' "${MUNBOT_UMASK}" >>/home/munbot/.profile
RUN printf 'export GOPATH=/go\n' >>/home/munbot/.profile

ADD --chown=munbot bin/*.sh /usr/local/bin/

RUN mkdir -vp -m 0750 /godoc/vendor && chown -vR munbot:munbot /godoc

USER munbot:munbot
WORKDIR /go/src/munbot

ENV USER munbot
ENV HOME /home/munbot

ENV GOBIN ''
ENV GOPATH /go
ENV GOENV /go/src/munbot/go.env

RUN go env && go version

RUN mkdir -vp -m 0750 ~/.config/munbot/master/ssl

ENTRYPOINT /usr/local/bin/entrypoint.sh
