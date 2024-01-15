package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.RouterGroup, parentUrl string) {

	r.GET(parentUrl+"/", List)
	r.GET(parentUrl+"/spec", Spec)
	r.GET(parentUrl+"/:id", Detail)
	r.POST(parentUrl, Store)
	r.DELETE(parentUrl+"/:id", Delete)
	r.PUT(parentUrl+"/:id", Update)

}
