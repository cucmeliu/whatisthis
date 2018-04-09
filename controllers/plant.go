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
	c.Ctx.WriteString(models.GetAccessToken())

}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.Plant	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *PlantController) Post() {
	var p models.Plant

	token := models.GetAccessToken()

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &p); err == nil {
		var v models.PlantResult
		c.RecogPlant(token, &p, &v)
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
		fmt.Println("Unmarshal Request Body: ", err)
	}

	c.ServeJSON()
}

func (this *PlantController) RecogPlant(access_token string, rec_rst *models.Plant, v *models.PlantResult) {
	urlstr := "https://aip.baidubce.com/rest/2.0/image-classify/v1/plant?access_token=" + access_token
	params := make(map[string]string)
	params["image"] = rec_rst.Img_base64

	resp_str, err := utils.HttpPost(urlstr, params)
	if err != nil {
		fmt.Println("post err, ", err)
	}
	fmt.Println("resp: ", resp_str)

	// var v *models.PlantResult
	if err := json.Unmarshal([]byte(resp_str), &v); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println("Unmarshal Response, ", err)
		panic(err)
	}
}
