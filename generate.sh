curl -s http://localhost:8697/json/swagger.json -o src/vmrest.json

tail -c +4 src/vmrest.json > src/vmrest_clean.json

npx swagger2openapi src/vmrest_clean.json -o src/vmrest_oas3.json # use json for yaml converts 'on' to true

yq -P '
  .components.schemas.NicIpStackAll.properties.nics = {"type":"array","items":{"$ref":"#/components/schemas/NicIpStack"}} |
  .info.description |= sub(" build-[0-9]+"; "")
' src/vmrest_oas3.json > src/vmrest_result.json

docker run --rm -v "${PWD}:/out" -v "${PWD}/src:/src" openapitools/openapi-generator-cli:latest-release generate \
    -i /src/vmrest_result.json \
    --skip-validate-spec \
    -c /src/config.yaml \
    -g go \
    -o /out

sed -i '' 's#`(?i:\(.*\)json)`#`(?i:\1json|application/vnd\.vmware\.vmw\.rest-v1\+json)`#' client.go # sed for Mac BSD, on Linux may not work
