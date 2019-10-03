FROM alpine:latest

RUN mkdir -p /alertmgrcfg/bin
WORKDIR /alertmgrcfg
COPY build/bin/alertmgr-config-operator-linux-amd64 bin/alertmgr-config-operator
COPY alertmgrcfg /etc/alertmgrcfg/

RUN chmod +x bin/alertmgr-config-operator

ENTRYPOINT [ "/alertmgrcfg/bin/alertmgr-config-operator" ]
