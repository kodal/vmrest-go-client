generate:
	swagger-codegen generate -i http://localhost:8697/json/swagger.json -c config.json -l go