#!/bin/bash

set -e

export NO_PROXY=localhost,127.0.0.1

ServerAddr="127.0.0.1:12345"
QueryString="user_id=testuserid&user_key=testuserkey"

curl \
  -X PUT \
  --progress-bar \
  -F "file=@test.txt" \
  "http://$ServerAddr/ft/uploads/03cfd743661f07975fa2f1220c5194cbaff48451?$QueryString"

