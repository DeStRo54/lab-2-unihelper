#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."
mkdir -p secrets
chmod 700 secrets

if [[ ! -s secrets/postgres_password.txt ]]; then
  openssl rand -hex 32 > secrets/postgres_password.txt
fi
chmod 600 secrets/postgres_password.txt

if [[ ! -s secrets/server.crt || ! -s secrets/server.key ]]; then
  openssl req -x509 -newkey rsa:3072 -sha256 -nodes -days 365 \
    -keyout secrets/server.key \
    -out secrets/server.crt \
    -subj "/CN=localhost" \
    -addext "subjectAltName=DNS:localhost,IP:127.0.0.1,IP:::1" \
    -addext "basicConstraints=critical,CA:TRUE" \
    -addext "keyUsage=critical,digitalSignature,keyEncipherment,keyCertSign" \
    -addext "extendedKeyUsage=serverAuth"
fi
chmod 600 secrets/server.key secrets/postgres_password.txt
chmod 644 secrets/server.crt

printf 'Local secrets are ready in %s/secrets\n' "$PWD"
