#!/usr/bin/env bash

set -euo pipefail

: "${BASE_URL:?Set BASE_URL to the trusted CyberStrikeAI origin}"
: "${CYBERSTRIKE_TOKEN:?Set CYBERSTRIKE_TOKEN to a temporary bearer token}"

case "${BASE_URL}" in
  http://127.0.0.1:* | http://localhost:* | https://*)
    ;;
  *)
    echo "BASE_URL must use HTTPS or loopback HTTP" >&2
    exit 2
    ;;
esac

search_term="${1:-}"
limit="${LIMIT:-20}"

case "${limit}" in
  '' | *[!0-9]*)
    echo "LIMIT must be an integer from 1 to 100" >&2
    exit 2
    ;;
esac

if ((limit < 1 || limit > 100)); then
  echo "LIMIT must be an integer from 1 to 100" >&2
  exit 2
fi

curl --fail-with-body \
  --silent \
  --show-error \
  --get \
  --connect-timeout 5 \
  --max-time 30 \
  --retry 2 \
  --retry-connrefused \
  --header "Authorization: Bearer ${CYBERSTRIKE_TOKEN}" \
  --data-urlencode "search=${search_term}" \
  --data-urlencode "limit=${limit}" \
  --data-urlencode "offset=0" \
  --data-urlencode "sort_by=updated_at" \
  "${BASE_URL%/}/api/conversations"
