FROM quay.io/bitnami/kubectl:1.13

COPY ./mulan-kube /usr/bin/mulan-kube

CMD /usr/bin/mulan-kube
