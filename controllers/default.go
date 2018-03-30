package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type user struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}

// URLMapping ...
func (c *MainController) URLMapping() {
	//	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	//	c.Mapping("GetAll", c.GetAll)
	//	c.Mapping("Put", c.Put)
	//	c.Mapping("Delete", c.Delete)
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MainController) GetOne() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	fmt.Println("get smt")

	c.ServeJSON()
	//c.TplName = "index.tpl"
}

func (this *MainController) Post() {
	u := user{}
	if err := this.ParseForm(&u); err != nil {

	}

}
