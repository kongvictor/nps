package controllers

import (
	"os"
	"strings"

	"ehang.io/nps/lib/common"
	"ehang.io/nps/lib/file"
	"github.com/astaxie/beego/logs"
)

type GlobalController struct {
	BaseController
}

func (s *GlobalController) Index() {
	//if s.Ctx.Request.Method == "GET" {
	//
	//	return
	//}
	s.Data["menu"] = "global"
	s.SetInfo("global")
	s.display("global/index")

	global := file.GetDb().GetGlobal()
	if global == nil {
		return
	}
	s.Data["globalBlackIpList"] = strings.Join(global.BlackIpList, "\r\n")
}

// 添加全局黑名单IP
func (s *GlobalController) Save() {
	//global, err := file.GetDb().GetGlobal()
	//if err != nil {
	//	return
	//}
	if s.Ctx.Request.Method == "GET" {
		s.Data["menu"] = "global"
		s.SetInfo("save global")
		s.display()
	} else {

		t := &file.Glob{BlackIpList: RemoveRepeatedElement(strings.Split(s.getEscapeString("globalBlackIpList"), "\r\n"))}

		if err := file.GetDb().SaveGlobal(t); err != nil {
			s.AjaxErr(err.Error())
		}
		s.AjaxOk("save success")
	}
}

func (s *GlobalController) ShowLog() {
	s.Data["menu"] = "global"
	s.SetInfo("global")
	s.display("global/ShowLog")

	content, err := os.ReadFile(common.GetLogPath())
	if err != nil {
		logs.Error("Unable to read log file" + common.GetLogPath() + " error:" + err.Error())
		return
	}
	global := file.GetDb().GetGlobal()
	if global == nil {
		return
	}
	s.Data["logcontent"] = string(content)
}
