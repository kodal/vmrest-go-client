#!/bin/sh

swagger-codegen generate -i http://localhost:8697/json/swagger.json -c config.json -l go

sed -i '' 's#jsonCheck = regexp.MustCompile("(?i:\[application|text\]/json)")#jsonCheck = regexp.MustCompile("(?i:(?:application|text)/json|application/vnd\.vmware\.vmw\.rest-v1\\\\+json)")#' client.go

sed -i '' 's#strings.Contains(contentType, "application/json")#jsonCheck.MatchString(contentType)#' client.go