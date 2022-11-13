#!/bin/sh
openssl req  -new  -newkey rsa:2048  -nodes  -keyout lgcapi.key  -out lgcapi.csr
openssl x509 -req  -days 365  -in lgcapi.csr  -signkey lgcapi.key  -out lgcapi.crt
exit 0
