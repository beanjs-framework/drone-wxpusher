FROM alpine:3.9

LABEL maintainer="beanjs <502554248@qq.com>"

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /drone/src

COPY wxpusher /usr/local/bin/wxpusher

ENTRYPOINT ["wxpusher"]