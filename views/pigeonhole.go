package views

import (
	"go-blog/common"
	"go-blog/config"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole

	pigeonhole.WriteData(w, config.Cfg.Viewer)
}
