package main

import (
	"github.com/micro-plat/fsync/services/file"
	"github.com/micro-plat/hydra/component"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *fsync) init() {
	r.Initializing(func(c component.IContainer) error {
		//appconf.func#//
		//#appconf.func//

		//db.init#//
		//#db.init//

		//cache.init#//
		//#cache.init//

		//queue.init#//
		//#queue.init//

		//login.router#//
		//#login.router//

		//service.router#//
		r.Micro("/file/upload", file.NewUploadHandler)
		//#service.router//

		return nil
	})
}
