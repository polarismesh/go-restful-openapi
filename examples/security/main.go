package main

import (
	"encoding/json"
	"fmt"
	restful "github.com/emicklei/go-restful/v3"
	restfulspec "github.com/polarismesh/go-restful-openapi/v2"
	"log"
)

func main() {
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	swagger := restfulspec.BuildSwagger(config)
	spec, err := json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(spec))
}
