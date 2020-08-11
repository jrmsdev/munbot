FROM munbot/master:base

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.08.08"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN apk add -U --no-cache sudo go-doc

ENV GOPATH /root/go
ENV GOBIN /usr/local/bin

RUN apk add --no-cache git \
	&& go get -u -v golang.org/x/tools/cmd/godoc \
	&& go generate -v golang.org/x/tools/godoc/static \
	&& go install -v -i golang.org/x/tools/cmd/godoc \
	&& ln -vs /usr/share/doc/go/doc /usr/lib/go/doc \
	&& apk del --no-cache --purge git \
	&& rm -rf ${GOPATH} /root/.cache

RUN apk add -U --no-cache openssh-keygen

RUN printf '\n%%munbot ALL=(ALL) NOPASSWD: ALL\n' >>/etc/sudoers

ADD --chown=munbot bin/*.sh /usr/local/bin/

USER munbot:munbot
WORKDIR /munbot/src/master

ENV USER munbot
ENV HOME /home/munbot

ENV GOBIN ''
ENV GOPATH /munbot
ENV GOENV /munbot/src/master/go.env

RUN go env && go version

ADD --chown=munbot tmp/go.env .
ADD --chown=munbot tmp/go.mod .
ADD --chown=munbot tmp/go.sum .
ADD --chown=munbot tmp/vendor/ ./vendor/
RUN go mod download \
	&& (rm -rf ./vendor $(go env GOCACHE); rm -vf go.mod go.sum go.env)

EXPOSE 6060
EXPOSE 6490
EXPOSE 6492

ENTRYPOINT /usr/local/bin/entrypoint.sh
