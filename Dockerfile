FROM debian:stretch-slim

WORKDIR /

COPY _output/bin/scheduler-framework-redis /usr/local/bin

CMD ["./scheduler-framework-redis"]
