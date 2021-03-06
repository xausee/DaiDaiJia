package controllers

import (
	"SanWenJia/app/models"
	"fmt"
	"github.com/revel/revel"
	"strconv"
)

type Witticism struct {
	*revel.Controller
}

func (this *Witticism) Index() revel.Result {
	manager, err := models.NewDbManager()
	if err != nil {
		fmt.Println("链接数据库失败")
	}
	defer manager.Close()

	witticisms, err := manager.GetAllWitticism()

	this.RenderArgs["userid"] = this.Session["userid"]
	this.RenderArgs["nickName"] = this.Session["nickName"]
	this.RenderArgs["avatarUrl"] = this.Session["avatarUrl"]
	this.RenderArgs["witticisms"] = witticisms

	return this.Render()
}

func (this *Witticism) Add() revel.Result {
	this.RenderArgs["userid"] = this.Session["userid"]
	this.RenderArgs["nickName"] = this.Session["nickName"]
	this.RenderArgs["avatarUrl"] = this.Session["avatarUrl"]

	return this.Render()
}

func (this *Witticism) Edit(id string) revel.Result {
	manager, err := models.NewDbManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()
	originalWitticism, _ := manager.GetWitticismById(id)

	this.RenderArgs["userid"] = this.Session["userid"]
	this.RenderArgs["nickName"] = this.Session["nickName"]
	this.RenderArgs["avatarUrl"] = this.Session["avatarUrl"]
	this.RenderArgs["originalWitticism"] = originalWitticism

	return this.Render()
}

func (this *Witticism) PostAdd(witticism *models.Witticism) revel.Result {
	this.Validation.Required(witticism.Content).Message("慧语内容不能为空")
	this.Validation.Required(witticism.Author).Message("作者不能为空")

	fmt.Println("慧语被容： ", witticism.Content)
	fmt.Println("作者： ", witticism.Author)

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		return this.Redirect((*Witticism).Add)
	}

	manager, err := models.NewDbManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	err = manager.AddWitticism(witticism)
	if err != nil {
		// q.Validation.Keep()
		// q.FlashParams()
		this.Flash.Error(err.Error())
		return this.Redirect((*Witticism).Add)
	}

	// 返回到管理员首页
	return this.Redirect((*Admin).Index)
}

func (this *Witticism) PostEdit(originalWitticismID string, newWitticism *models.Witticism) revel.Result {
	this.Validation.Required(newWitticism.Content).Message("内容不能为空")
	this.Validation.Required(newWitticism.Author).Message("作者不能为空")

	fmt.Println("内容： ", newWitticism.Content)
	fmt.Println("作者： ", newWitticism.Author)

	if this.Validation.HasErrors() {
		this.Validation.Keep()
		this.FlashParams()
		fmt.Println("error in validation ")
		return this.Redirect((*Witticism).Edit)
	}

	manager, err := models.NewDbManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()

	err = manager.UpdateWitticism(originalWitticismID, newWitticism)
	if err != nil {
		this.Validation.Keep()
		this.FlashParams()
		this.Flash.Error(err.Error())
		fmt.Println("error in update Witticism ")
		return this.Redirect((*Witticism).Edit)
	}

	return this.Redirect((*Witticism).Index)
}

func (this *Witticism) Show(id string) revel.Result {
	manager, err := models.NewDbManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()
	witticism, _ := manager.GetWitticismById(id)
	fmt.Println("作者： ", witticism.Author)
	// if err != nil {
	// 	this.Flash.Error(err.Error())
	// 	//return this.Redirect((*Essay).Add)
	// }

	// session 里保存的是字符串，所以需要类型转换
	var admin = false
	var role models.Role
	roleStr, isExists := this.Session[models.CSessionRole]

	if isExists {
		value, _ := strconv.Atoi(roleStr)
		role = models.Role(value)

		if role == models.AdminRole {
			admin = true
		}
	}

	this.RenderArgs["admin"] = admin
	this.RenderArgs["userid"] = this.Session["userid"]
	this.RenderArgs["nickName"] = this.Session["nickName"]
	this.RenderArgs["avatarUrl"] = this.Session["avatarUrl"]
	this.RenderArgs["witticism"] = witticism

	return this.Render()
}

func (this *Witticism) Delete(id string) revel.Result {
	manager, err := models.NewDbManager()
	if err != nil {
		this.Response.Status = 500
		return this.RenderError(err)
	}
	defer manager.Close()
	err = manager.DeleteWitticismById(id)

	this.RenderArgs["userid"] = this.Session["userid"]
	this.RenderArgs["nickName"] = this.Session["nickName"]
	this.RenderArgs["avatarUrl"] = this.Session["avatarUrl"]

	return this.Redirect((*Witticism).Index)
}
