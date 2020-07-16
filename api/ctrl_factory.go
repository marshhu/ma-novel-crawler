package api

import (
	"ma-novel-crawler/api/controllers"
	"ma-novel-crawler/service"
)

//控制器工厂
type CtrlFactory struct {
	NovelCtrl *controllers.NovelController
}

var CtrlFactoryInstance = &CtrlFactory{
	NovelCtrl: &controllers.NovelController{NovelService: &service.NovelService{}},
}
