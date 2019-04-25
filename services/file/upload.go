package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/utility"
)

type UploadHandler struct {
	container component.IContainer
	fields    []string
	root      string
	url       string
}

func NewUploadHandler(root string, url string) func(container component.IContainer) (u *UploadHandler) {
	return func(container component.IContainer) (u *UploadHandler) {
		return &UploadHandler{
			container: container,
			root:      root,
			url:       url,
			fields:    []string{"filename"},
		}
	}
}

//Handle .
func (u *UploadHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------------上传文件-------------------")

	ctx.Log.Info("1. 校验参数")
	if err := ctx.Request.Check(u.fields...); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 读取文件内容")
	fileName := ctx.Request.GetString("filename")
	f, err := ctx.Request.Http.Get()
	uf, _, err := f.FormFile("file")
	if err != nil {
		return fmt.Errorf("无法读取上传的文件:%s(err:%v)", fileName, err)
	}
	defer uf.Close()

	name := fmt.Sprintf("%s%s", utility.GetGUID(), filepath.Ext(fileName))
	path := filepath.Join(u.root, name)
	urlPath := fmt.Sprintf("%s/%s", strings.Trim(u.url, "/"), name)

	ctx.Log.Info("3. 构建文件名")
	nf, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("创建文件失败:%s(err:%v)", path, err)
	}
	defer nf.Close()

	ctx.Log.Info("4. 保存文件:", path)
	_, err = io.Copy(nf, uf)
	if err != nil {
		return fmt.Errorf("保存文件(%s)失败:%v", path, err)
	}

	ctx.Log.Info("5. 返回结果")
	return map[string]string{
		"url": urlPath,
	}
}
