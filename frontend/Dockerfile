FROM alpine:edge

ENV TERM dumb
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

ENV DATABASE_HOST db 
ENV DATABASE_USER appuser
ENV DATABASE_PASS apppass
ENV DATABASE_PORT 3306
ENV DATABASE_DB appdb

RUN set -x && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
    apk update && \
    apk upgrade

RUN apk add --update --no-cache \
    wget \
    bash \
    curl \
    git \
    mysql-client \
    go \
    && rm -rf /var/cache/apk/* \
    && mkdir -p ${GOPATH}/src ${GOPATH}/bin 

ADD files/apps ${GOPATH}/src
ADD files/etc/start.sh /start.sh

RUN chmod 755 /start.sh 

EXPOSE 5000 

# Start Supervisord
ENTRYPOINT ["/start.sh"]



