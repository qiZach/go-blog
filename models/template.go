package models

import (
	"html/template"
	"log"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func InitTemplate(templateDir string) HtmlTemplate {
	tp := readTemplate(
		[]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templateDir,
	)
	// 拿到以上所有页面的解析
	var htmlTemplate HtmlTemplate
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate

}

func readTemplate(templates []string, templateDir string) []TemplateBlog {

	var tbs []TemplateBlog

	// 循环依次解析： index,category,custom,detail,login,pigeonhole,writing
	for _, view := range templates {
		viewName := view + ".html"                   // index.html category.html
		t := template.New(viewName)                  // new 一个模板
		home := templateDir + "home.html"            // 拼接模板路径 /template/home.html
		header := templateDir + "layout/header.html" // /template/layout/header.html
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		// 模板中需要的函数
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		// 将所有的页面进行解析
		// templateDir+viewName /template/index.html
		t, err := t.ParseFiles(templateDir+viewName, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Println("配置模板出错： ", err)
		}
		// 创建一个TemplateBlog, 一个页面
		var tb TemplateBlog
		// 放入template.Template
		tb.Template = t
		// 添加到数组末位
		tbs = append(tbs, tb)
	}
	return tbs
}
