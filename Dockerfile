FROM alpine:latest
MAINTAINER  Piotr Kowalczuk <p.kowalczuk.priv@gmail.com>

ARG BUILD_DATE
ARG VCS_REF

LABEL org.label-schema.build-date=$BUILD_DATE \
	org.label-schema.docker.dockerfile="Dockerfile" \
	org.label-schema.license="ASL" \
	org.label-schema.name="charon" \
	org.label-schema.url="https://github.com/piotrkowalczuk/charon" \
	org.label-schema.vcs-ref=$VCS_REF \
	org.label-schema.vcs-type="git" \
	org.label-schema.vcs-url="https://github.com/piotrkowalczuk/charon"

COPY ./bin /usr/local/bin/
COPY ./scripts/docker-entrypoint.sh /

EXPOSE 8080

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["charond"]