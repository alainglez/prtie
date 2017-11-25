FROM alpine
MAINTAINER Alain Gonzalez <alain@caldo.io>
COPY prtie-linux-amd64 /prtie
COPY templates /templates/

ENTRYPOINT ["/prtie"]
