package api

import (
	"github.com/marshhu/ma-novel-crawler/api/controllers"
	"github.com/marshhu/ma-novel-crawler/service"
)

//控制器工厂
type CtrlFactory struct {
	NovelCtrl *controllers.NovelController
}

var CtrlFactoryInstance = &CtrlFactory{
	NovelCtrl: &controllers.NovelController{NovelService: &service.NovelService{}},
}
