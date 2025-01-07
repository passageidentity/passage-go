# #!/bin/bash
set -euo pipefail

if [ -z "$1" ]; then
  echo "Required generator file is missing."
  exit 1
fi

file="$1"

# Run codegen
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 \
  -config cfg.yml \
  $file
