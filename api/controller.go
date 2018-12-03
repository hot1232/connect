package api

import (
	"github.com/jinzhu/gorm"
	"github.com/emicklei/go-restful"
	"sync"
)

type BaseController struct {
	db *gorm.DB
	ws *restful.WebService
	watcher *sync.Map
}


func NewBaseController(db *gorm.DB) *BaseController{
	ws := new(restful.WebService)
	ws.Path("/api/v1").Consumes(restful.MIME_JSON,restful.MIME_XML).Produces(restful.MIME_JSON,restful.MIME_XML)
	baseController := &BaseController{
		ws:ws,
		db:db,
		watcher:new(sync.Map),
	}

	return baseController
}
