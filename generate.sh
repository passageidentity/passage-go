# #!/bin/bash
set -euo pipefail

output_file="passage.gen.go"

# Skip pointer on optional fields
fields="\
.components.schemas.CreateUserRequest.properties.email \
.components.schemas.CreateUserRequest.properties.phone \
.components.schemas.CreateUserRequest.properties.user_metadata \
.components.schemas.UpdateUserRequest.properties.email \
.components.schemas.UpdateUserRequest.properties.phone \
.components.schemas.UpdateUserRequest.properties.user_metadata \
.components.schemas.CreateMagicLinkRequest.properties.language \
.components.schemas.MagicLinkType \
.components.schemas.UserInfo.properties.user_metadata \
.components.schemas.UserInfo.properties.webauthn_types"
for field in $fields; do
  jq "$field |= . + {\"x-go-type-skip-optional-pointer\": true}" openapi.json > tmp.json && mv tmp.json openapi.json
done

# JSON string of key-value pairs
transforms='{
    "Active": "StatusActive",
    "CreateMagicLinkRequest": "CreateMagicLinkBody",
    "CreateUserRequest": "CreateUserBody",
    "Inactive": "StatusInactive",
    "Login": "LoginType",
    "MagicLinkChannel": "ChannelType",
    "MagicLinkChannelEmail": "EmailChannel",
    "MagicLinkChannelPhone": "PhoneChannel",
    "Pending": "StatusPending",
    "UpdateUserRequest": "UpdateBody",
    "UserInfo": "User",
    "UserMetadataFieldTypeBoolean": "BooleanMD",
    "UserMetadataFieldTypeDate": "DateMD",
    "UserMetadataFieldTypeEmail": "EmailMD",
    "UserMetadataFieldTypeInteger": "IntegerMD",
    "UserMetadataFieldTypePhone": "PhoneMD",
    "UserMetadataFieldTypeString": "StringMD",
    "VerifyIdentifier": "VerifyIdentifierType"
}'

# Function to perform replacements
command -v gorename &> /dev/null || go install golang.org/x/tools/cmd/gorename@latest
replace() {
    local in=$1 out=$2
    gorename -from "\"github.com/passageidentity/passage-go\".$in" -to "$out" -force
}

# Un-apply transforms
echo "$transforms" | jq -r 'to_entries[] | "\(.key) \(.value)"' | while read -r key value; do
    replace "$value" "$key"
done

# Run codegen
go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.16.2 \
  -generate types,client \
  -package passage \
  -o "$output_file" \
  -initialism-overrides \
  openapi.json

# Replace initialisms with uppercase versions
perl -pe 's/(Url|Uri|Ttl|Id|Rsa|Ip)(s?)(?=\b|[A-Z])/\U$1\E$2/g' "$output_file" > tmp.txt && mv tmp.txt "$output_file" 

# Apply transforms
echo "$transforms" | jq -r 'to_entries[] | "\(.key) \(.value)"' | while read -r key value; do
    replace "$key" "$value"
done

