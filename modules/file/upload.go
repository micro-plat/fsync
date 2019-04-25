
package file

import "github.com/micro-plat/hydra/component"

type IUpload interface {
}

type Upload struct {
	c component.IContainer
}

func NewUpload(c component.IContainer) *Upload {
	return &Upload{
		c: c,
	}
}
