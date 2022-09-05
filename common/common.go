package common

import (
	"go-blog/config"
	"go-blog/models"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		// 耗时，放入协程中加载
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			// 直接挂掉
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}
