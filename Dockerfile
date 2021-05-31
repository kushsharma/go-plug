FROM alpine:3.13

COPY goplug /usr/bin/goplug
COPY plugin-sql /usr/bin/plugin-sql

CMD ["goplug", "--plugin", "./plugin-sql"]