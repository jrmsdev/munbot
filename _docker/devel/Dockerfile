FROM munbot/master:base

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.08.04"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN apk update --no-cache \
	&& apk add --no-cache sudo go-doc openssl ca-certificates

ENV GOPATH /root/go
ENV GOBIN /usr/local/bin

RUN apk add --no-cache git \
	&& go get -u -v golang.org/x/tools/cmd/godoc \
	&& go generate -v golang.org/x/tools/godoc/static \
	&& go install -v -i golang.org/x/tools/cmd/godoc \
	&& ln -vs /usr/share/doc/go/doc /usr/lib/go/doc \
	&& apk del --no-cache --purge git \
	&& rm -rf ${GOPATH} /root/.cache

RUN printf '\n%%munbot ALL=(ALL) NOPASSWD: ALL\n' >>/etc/sudoers

RUN mkdir -vp -m 0750 /etc/ssl/munbot && chown -v munbot:munbot /etc/ssl/munbot

EXPOSE 6060
EXPOSE 6492

ADD --chown=munbot bin/*.sh /usr/local/bin/

USER munbot:munbot
WORKDIR /munbot/src/master

ENV USER munbot
ENV HOME /home/munbot

ENV GOBIN ''
ENV GOPATH /munbot
ENV GOENV /munbot/src/master/go.env

RUN go env && go version

ADD --chown=munbot tmp/go.mod .
ADD --chown=munbot tmp/go.sum .
ADD --chown=munbot tmp/vendor/ ./vendor/
RUN go mod download \
	&& go get -u -v github.com/FiloSottile/mkcert \
	&& (rm -rf ./vendor; rm -vf go.mod go.sum)

ENV CAROOT /etc/ssl/munbot
ENV TRUST_STORES ''
RUN mkcert devel.munbot.local \
	&& mkcert -client user.munbot.local \
	&& mkcert -client client.munbot.local

ENTRYPOINT /usr/local/bin/entrypoint.sh
