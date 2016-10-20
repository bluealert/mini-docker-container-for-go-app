FROM scratch
MAINTAINER bluealert

ADD tls-ca-bundle.pem /etc/ssl/certs/
ADD hello /

CMD ["/hello"]
