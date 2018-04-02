package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/cucmeliu/whatisthis/models"
	"github.com/cucmeliu/whatisthis/utils"
)

type PlantController struct {
	beego.Controller
}

func (c *PlantController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.Get)
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
func (c *PlantController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	fmt.Println("get smt")
	//c.TplName = "views/index.tpl"
	c.Ctx.WriteString(models.GetAccessToken())

	//c.ServeJSON()
	//	var param models.Param_rec
	//	param.User_token = this.GetString("user_token")
	//	param.Src_base64 = this.GetString("src_base64")
	//	if param.User_token == "" || param.Src_base64 == "" {
	//		this.Ctx.WriteString("param is wrong")
	//		return
	//	}

}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.Plant	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *PlantController) Post() {
	var v models.Plant
	token := models.GetAccessToken()

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		c.RecogPlant(token, &v)
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
		c.Ctx.WriteString(v.Img_base64)
	} else {
		c.Data["json"] = err.Error()
		fmt.Println("err.", err)
	}

	c.ServeJSON()
}

func (this *PlantController) RecogPlant(access_token string, rec_rst *models.Plant) {
	url := "https://aip.baidubce.com/rest/2.0/image-classify/v1/plant?access_token=" + access_token
	params := make(map[string]string)
	params["image"] = rec_rst.Img_base64

	resp_str, err := utils.HttpPost(url, params)
	fmt.Println("resp: ", resp_str)
	_ = json.Unmarshal([]byte(resp_str), &rec_rst)

	if err == nil {
		fmt.Println("encoded err", err)
		panic(err)
	}

}
