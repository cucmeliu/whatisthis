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
	//	c.Mapping("GetOne", c.Get)
	//	c.Mapping("GetAll", c.GetAll)
	//	c.Mapping("Put", c.Put)
	//	c.Mapping("Delete", c.Delete)
}

func (this *PlantController) Get() {
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
	token := models.GetToken()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		var rec models.Plant
		c.RecogPlant(token, v.Img_base64, &rec)
		fmt.Println()
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
		//		} else {
		//			c.Data["json"] = err.Error()
		//		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (this *PlantController) RecogPlant(access_token, image_base64 string, rec_rst *models.Plant) {
	url := "https://aip.baidubce.com/rest/2.0/image-classify/v1/plant?access_token=" + access_token
	params := make(map[string]string)
	params["image"] = image_base64

	resp_str, err := utils.HttpPost(url, params)
	_ = json.Unmarshal([]byte(resp_str), &rec_rst)

	if err == nil {
		// fmt.Println("encoded err", err)
		panic(err)
	}

}
