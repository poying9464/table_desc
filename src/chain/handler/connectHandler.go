package handler

import (
	"table_desc/src/chain"
)

type ConnectHandler struct {
	Next chain.Handler
}

func (h *ConnectHandler) SetNext(next chain.Handler) {
	h.Next = next
}

func (h *ConnectHandler) Handle(hp *chain.HandlerParam) {
	connect := hp.Ctx.Connect(hp.Param)
	hp.Db = connect
	if h.Next != nil {
		h.Next.Handle(hp)
	}
}
