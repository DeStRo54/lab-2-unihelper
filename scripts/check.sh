#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."
CA_FILE=secrets/server.crt
BASE_URL=https://localhost:8443

curl --fail --silent --show-error --cacert "$CA_FILE" "$BASE_URL/readyz"
curl --fail --silent --show-error --cacert "$CA_FILE" "$BASE_URL/health"
curl --fail --silent --show-error --cacert "$CA_FILE" "$BASE_URL/api/livez" >/dev/null
curl --fail --silent --show-error --cacert "$CA_FILE" "$BASE_URL/" >/dev/null
printf '\nTLS and readiness checks passed.\n'
