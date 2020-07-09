package admin

/** 蜘蛛池统计面板 **/
type Statistics struct{ Main }

// @router /statistics/index [get,post]
func (c *Statistics) Index() {
	c.journalService.CachedHandleSetDebug() //todo 功能做好需要删除掉
	items := c.journalService.CachedHandleAnalysisWeek()
	dayRes := c.journalService.CachedHandleGetDya()
	class := c.journalService.CachedHandleAnalysisClass(dayRes)
	c.Data["class"] = class
	c.Data["week"] = items
	c.Data["autoWeek"] = c.automaticService.AnalysisWeek()
	c.Data["monitorTags"] = c.journalService.CachedHandleGetMonitorTags()
	c.Data["headMenu"] = c.menus()
	c.Data["footerMenu"] = c.footer()
	c.Data["hotDomain"] = c.journalService.HotDomain()
}

/** 获取顶部导航 **/
func (c *Statistics) menus() []map[string]interface{} {
	menu := []map[string]interface{}{
		{"title": "总蜘蛛量", "count": c.journalService.SumAll(), "urls": ""},
		{"title": "总推送量", "count": c.automaticService.SumAll(), "urls": ""},
		{"title": "蜘蛛池数", "count": c.arachnidService.AliveAllNum(), "urls": ""},
		{"title": "挂载域名", "count": c.domainService.AliveAllNum(), "urls": ""},
	}
	return menu
}

/** 获取底部导航 **/
func (c *Statistics) footer() []map[string]interface{} {
	menu := []map[string]interface{}{
		{"title": "日蜘蛛量", "count": c.journalService.SumDay(), "urls": ""},
		{"title": "日推送量", "count": c.automaticService.SumDay(), "urls": ""},
	}
	return menu
}

/** 获取热门站点 **/
func (c *Statistics) hotSet() []map[string]interface{} {
	menu := []map[string]interface{}{
		{"title": "日蜘蛛量", "count": c.journalService.SumDay(), "urls": ""},
		{"title": "日推送量", "count": c.automaticService.SumDay(), "urls": ""},
	}
	return menu
}
