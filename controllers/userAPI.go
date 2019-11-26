package controllers

func (c *UserController) API_Profile() {
	type user struct {
		Userid string
		Hobby []string
	}

	u := user{
		"jason",
		[]string{"chess", "code"},
	}

	c.Data["json"] = u
	c.ServeJSON()

}
