#!/bin/zsh

if [[ -z $2 ]]; then
  echo "Usage: $0 ssl path"
  exit 1
fi

SSLNAME=$1
SSLDIR=$2

# 生成CA证书
openssl req -x509 -newkey rsa:4096 -days 3650 -nodes \
        -keyout "${SSLDIR}/${SSLNAME}.key" \
        -out "${SSLDIR}/${SSLNAME}.crt" \
        -config "${SSLDIR}/${SSLNAME}.conf" \

cp -rf "${SSLDIR}/${SSLNAME}.crt" "${SSLDIR}/${SSLNAME}.pem"
