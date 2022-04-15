package animals

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AnimalsController struct {
	Ctx iris.Context
}

func (c *AnimalsController) Get() mvc.View {
	var animalsListView = mvc.View{
		Name: "/animals/list.html",
		Data: iris.Map{"message": "🐭🐮🐯🐰🐲🐍🐴🐑🐵🐔🐶🐗"},
	}

	return animalsListView
}