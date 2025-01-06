FROM alpine:3.13

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD http://github.com/pressly/goose/releases/download/v3.24.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD migrations/*.sql migrations_chat/migrations/
ADD migrations_chat.sh migrations_chat/
ADD .env .

RUN chmod +x migrations_chat/migrations_chat.sh

ENTRYPOINT ["bash","migrations_chat/migrations_chat.sh"]