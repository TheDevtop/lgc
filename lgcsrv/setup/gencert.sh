#!/bin/sh
openssl req  -new  -newkey rsa:2048  -nodes  -keyout selfsigned.key  -out selfsigned.csr
openssl x509 -req  -days 365  -in selfsigned.csr  -signkey selfsigned.key  -out selfsigned.crt
exit 0
