FROM munbot/master:base

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.07.10"

USER root:root
WORKDIR /root

ENV USER root
ENV HOME /root

RUN apk update --no-cache \
	&& apk add --no-cache sudo go-doc patch perl openssl ca-certificates

ADD ca/CA.pl.patch /tmp/CA.pl.patch
RUN mv -v /etc/ssl/openssl.cnf /etc/ssl/openssl.cnf.orig \
	&& ln -vs /go/src/munbot/docker/devel/ca/openssl.cnf /etc/ssl/openssl.cnf \
	&& cp -va /etc/ssl/misc/CA.pl /etc/ssl/misc/CA.pl.orig \
	&& patch /etc/ssl/misc/CA.pl /tmp/CA.pl.patch \
	&& rm -vf /tmp/CA.pl.patch

ENV GOPATH /root/go
ENV GOBIN /usr/local/bin

RUN apk add --no-cache git \
	&& go get -u -v golang.org/x/tools/cmd/godoc \
	&& go generate -v golang.org/x/tools/godoc/static \
	&& go install -v -i golang.org/x/tools/cmd/godoc \
	&& apk del --no-cache --purge git \
	&& rm -rf ${GOPATH} /root/.cache

RUN printf '\n%%munbot ALL=(ALL) NOPASSWD: ALL\n' >>/etc/sudoers

RUN ln -vs /usr/share/doc/go/doc /usr/lib/go/doc

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
RUN go mod download && (rm -rf ./vendor; rm -vf go.mod go.sum)

RUN mkdir -vp -m 0750 /home/munbot/.config/munbot/default/ssl
VOLUME /home/munbot/.config/munbot/default/ssl

# maybe? go get -u -v github.com/go-bindata/go-bindata/...

ENTRYPOINT /usr/local/bin/entrypoint.sh
