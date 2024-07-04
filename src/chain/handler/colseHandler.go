package handler

import (
	"table_desc/src/chain"
)

type CloseHandler struct {
	Next chain.Handler
}

func (h *CloseHandler) SetNext(next chain.Handler) {
	h.Next = next
}

func (h *CloseHandler) Handle(hp *chain.HandlerParam) {
	if hp.Allows {
		hp.Ctx.Close(hp.Db)
		if h.Next != nil {
			h.Next.Handle(hp)
		}
	}
}
