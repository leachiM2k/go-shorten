FROM alpine

# proper reaping / signal forwarding init
RUN apk add --no-cache tini ca-certificates
ENTRYPOINT ["/sbin/tini", "--"]

ENV APP_ENV=production

ADD ./dist/linux/amd64/* /
ADD config.yml /

CMD ["/shorten"]
