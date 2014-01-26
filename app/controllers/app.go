package controllers

import (
	"SanWenJia/app/models"
	"fmt"
	"github.com/robfig/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	email := c.Session["email"]
	nickName := c.Session["nickName"]

	manager, err := models.NewDbManager()
	if err != nil {
		fmt.Println("链接数据库失败")
	}
	defer manager.Close()

	// 获取摘抄最新15条数据
	c.RenderQuotations(manager)

	// 获取名人语录最新5条数据
	c.RenderWitticismQuotation(manager)

	// 获取古诗词最新15条数据
	c.RenderAncientPoems(manager)

	// 获取现代诗最新15条数据
	c.RenderModernPoems(manager)

	// 获取散文最新15条数据
	c.RenderEssays(manager)

	// 获取微小说最新15条数据
	c.RenderHintFictions(manager)

	return c.Render(email, nickName)
}

func (c App) Add() revel.Result {
	email := c.Session["email"]
	nickName := c.Session["nickName"]

	return c.Render(email, nickName)
}

func (c *App) RenderQuotations(manager *models.DbManager) error {
	qs, err := manager.GetAllQuotation()

	// 只取前15条摘录用于首页的显示
	var quotations []models.Quotation
	var count = len(qs)
	var more bool
	if count > 15 {
		more = true
		quotations = qs[(count - 15):]
	} else {
		more = false
		quotations = qs
	}

	c.RenderArgs["quotations"] = quotations
	c.RenderArgs["moreQuotation"] = more
	return err
}

func (c *App) RenderWitticismQuotation(manager *models.DbManager) error {
	ws, err := manager.GetAllQuotation()

	// 只取前5条名人语录用于首页的显示
	var witticisms []models.Quotation
	var count = len(ws)
	var more bool
	if count > 5 {
		more = true
		witticisms = ws[(count - 5):]
	} else {
		more = false
		witticisms = ws
	}

	c.RenderArgs["witticisms"] = witticisms
	c.RenderArgs["moreMitticism"] = more
	return err
}

func (c *App) RenderAncientPoems(manager *models.DbManager) error {
	poems, err := manager.GetAllAncientPoem()
	c.RenderArgs["ancientPoems"] = poems

	return err
}

func (c *App) RenderModernPoems(manager *models.DbManager) error {
	ps, err := manager.GetAllModernPoem()

	// 只取前15篇现代诗用于首页的显示
	var poems []models.ModernPoem
	var count = len(ps)
	var more bool
	if count > 15 {
		more = true
		poems = ps[(count - 15):]
	} else {
		more = false
		poems = ps
	}

	c.RenderArgs["modernPoems"] = poems
	c.RenderArgs["moreModernPoem"] = more
	return err
}

func (c *App) RenderEssays(manager *models.DbManager) error {
	es, err := manager.GetAllEssay()

	// 只取前15篇散文用于首页的显示
	var essays []models.Essay
	var count = len(es)
	var more bool
	if count > 15 {
		more = true
		essays = es[(count - 15):]
	} else {
		more = false
		essays = es
	}

	c.RenderArgs["essays"] = essays
	c.RenderArgs["moreEssay"] = more
	return err
}

func (c *App) RenderHintFictions(manager *models.DbManager) error {
	hs, err := manager.GetAllHintFiction()

	// 只取前15篇现代诗用于首页的显示
	var hintFictions []models.HintFiction
	var count = len(hs)
	var more bool
	if count > 15 {
		more = true
		hintFictions = hs[(count - 15):]
	} else {
		more = false
		hintFictions = hs
	}

	c.RenderArgs["hintFictions"] = hintFictions
	c.RenderArgs["moreHintFiction"] = more
	return err
}

func (c App) AboutUs() revel.Result {
	email := c.Session["email"]
	nickName := c.Session["nickName"]
	return c.Render(email, nickName)
}

func (c App) Donate() revel.Result {
	email := c.Session["email"]
	nickName := c.Session["nickName"]
	return c.Render(email, nickName)
}
