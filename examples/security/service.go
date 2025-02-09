package main

import restful "github.com/emicklei/go-restful/v3"
import restfulspec "github.com/polarismesh/go-restful-openapi/v2"

type ExampleService struct{}
type ExampleReq struct{}
type ExampleResp struct{}
type ResponseMsg struct{}

func (s ExampleService) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/example").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"example"}

	ws.Route(ws.POST("/example").To(s.create).
		Doc("create example thing").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(SecurityDefinitionKey, OAISecurity{Name: "jwt"}).
		Writes(ExampleResp{}).
		Reads(ExampleReq{}).
		Returns(200, "OK", &ExampleResp{}).
		Returns(404, "NotFound", &ResponseMsg{}).
		Returns(500, "InternalServerError", &ResponseMsg{}))

	return ws

}
func (s ExampleService) create(*restful.Request, *restful.Response) {}
