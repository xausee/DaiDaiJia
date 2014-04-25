package app

import (
	"github.com/revel/revel"
	"strings"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.ActionInvoker,           // Invoke the action.
	}

	// 注册模板里的字符串相加函数
	revel.TemplateFuncs["concat"] = func(a, b string) string { return a + b }

	// 注册模板里的字符串拆分函数
	revel.TemplateFuncs["prev"] = func(input, key string) string {
		if input == key {
			return input
		} else if strings.Contains(input, key) {
			return strings.SplitN(input, key, 2)[0]
		} else {
			return "no substring found"
		}
	}

	revel.TemplateFuncs["next"] = func(input, key string) string {
		if input == key {
			return input
		} else if strings.Contains(input, key) {
			return strings.SplitN(input, key, 2)[1]
		} else {
			return "no substring found"
		}
	}
}
