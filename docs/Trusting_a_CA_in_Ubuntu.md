On Debian and Ubuntu you have to copy the certificate.pem to /usr/local/share/ca-certificates/certificate.crt and then run dpkg-reconfigure ca-certificates. /etc/ssl/certs is managed by that command.
