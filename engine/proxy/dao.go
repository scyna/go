package proxy

import "github.com/scyna/go/scyna"

func (proxy *Proxy) saveContext(ctx *scyna.Context) {
	ctx.Save() //FIXME: direct save
}
