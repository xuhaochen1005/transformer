FROM 192.168.88.1:8888/library/centos:8.3.2011

WORKDIR /root

COPY bin/transformer /root/

COPY config/ /root/config/

ENV ENV pro

CMD []

ENTRYPOINT ["/root/transformer"]
