package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Profile() {
	c.Data["userid"] = "lenjos"
	c.Data["tag"] = "jason"
	c.Data["hobby"] = []string{"chess", "code"}
	c.TplName = "user/profile.html"
}
