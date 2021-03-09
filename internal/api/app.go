package api

import (
	"fmt"

	"github.com/lonmarsDev/ER-rest-service-go/pkg/log"
	"github.com/realpamisa/RestAPI/internal/api/handler"
	"github.com/realpamisa/RestAPI/internal/api/routes"
	"github.com/realpamisa/RestAPI/internal/api/server"
)

const Tag = "Rest Api"

func Init() {
	fmt.Println("Listening [::]:3000")
	srv := server.NewRestAPI()
	srv.EnableLogger()
	// mysql.Init()
	handler := handler.Init()
	route := routes.Init()
	route.SetHandler(handler).SetServer(&srv).Build()
	if err := srv.Serve(); err != nil {
		log.Error(Tag, err.Error())
	}
}
