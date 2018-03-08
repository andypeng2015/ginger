#!/bin/sh

ACCEPT_HEADER="Accept: application/vnd.github.jean-grey-preview+json"
TOKEN_HEADER="Authorization: token ${GITHUB_TOKEN}"
RELEASE_ENDPOINT="https://api.github.com/repos/ysugimoto/ginger/releases"

RELEASE_ID=$(curl -X POST -H "${ACCEPT_HEADER}" -H "${TOKEN_HEADER}" -d "{\"tag_name\": \"${CIRCLE_TAG}\", \"name\": \"${CIRCLE_TAG}\"}" "${RELEASE_ENDPOINT}" | jq .id)
RELEASE_URL="https://uploads.github.com/repos/ysugimoto/ginger/releases/${RELEASE_ID}/assets"

for FILE in `ls ./dist`; do
  curl -v -X POST -H "${ACCEPT_HEADER}" -H "${TOKEN_HEADER}" -H "Content-Type: application/octet-stream" --data-binary "@dist/${FILE}" "${RELEASE_URL}?name=${FILE}"
done
