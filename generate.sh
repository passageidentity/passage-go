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
.components.schemas.CreateMagicLinkRequest.properties.channel \
.components.schemas.CreateMagicLinkRequest.properties.email \
.components.schemas.CreateMagicLinkRequest.properties.language \
.components.schemas.CreateMagicLinkRequest.properties.magic_link_path \
.components.schemas.CreateMagicLinkRequest.properties.phone \
.components.schemas.CreateMagicLinkRequest.properties.redirect_url \
.components.schemas.CreateMagicLinkRequest.properties.send \
.components.schemas.CreateMagicLinkRequest.properties.ttl \
.components.schemas.CreateMagicLinkRequest.properties.type \
.components.schemas.CreateMagicLinkRequest.properties.user_id \
.components.schemas.MagicLinkType \
.components.schemas.MagicLinkChannel \
.components.schemas.UserInfo.properties.user_metadata \
.components.schemas.UserInfo.properties.webauthn_types"
for field in $fields; do
  jq "$field |= . + {\"x-go-type-skip-optional-pointer\": true}" openapi.json > tmp.json && mv tmp.json openapi.json
done

# Rename component to avoid name clash with generated struct
jq ".components.schemas.ListPaginatedUsersResponse |= . + {\"x-go-name\": \"PaginatedUsersResponse\"}" openapi.json > tmp.json && mv tmp.json openapi.json

# JSON string of key-value pairs
transforms='{
    "Active": "StatusActive",
    "CreateMagicLinkRequest": "CreateMagicLinkBody",
    "CreateUserRequest": "CreateUserBody",
    "Inactive": "StatusInactive",
    "MagicLinkTypeLogin": "LoginType",
    "MagicLinkTypeVerifyIdentifier": "VerifyIdentifierType",
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
    "UserMetadataFieldTypeString": "StringMD"
}'

# Function to perform replacements
command -v gorename &> /dev/null || go install golang.org/x/tools/cmd/gorename@v0.24.0
replace() {
    local in=$1 out=$2
    gorename -from "\"github.com/passageidentity/passage-go\".$in" -to "$out" -force
}

# Un-apply transforms
echo "$transforms" | jq -r 'to_entries[] | "\(.key) \(.value)"' | while read -r key value; do
    replace "$value" "$key"
done

# Run codegen
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 \
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

