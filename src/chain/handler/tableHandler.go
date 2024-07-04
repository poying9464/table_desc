package handler

import (
	"fmt"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"table_desc/src/chain"
)

type TableHandler struct {
	Next chain.Handler
}

func (h *TableHandler) SetNext(next chain.Handler) {
	h.Next = next
}

func (h *TableHandler) Handle(hp *chain.HandlerParam) {
	info := hp.Ctx.QueryTableInfo(hp.Param.Scheme, hp.Db)
	hp.Allows = false
	writeHeader(hp)
	if info != nil {
		// 将info写入文件中 path为文件路径
		for _, data := range info {
			writeWord(data, hp)
			hp.TableName = data["tableName"]
			h.Next.Handle(hp)
		}
		hp.TableName = ""
	}
	hp.Allows = true
	if h.Next != nil {
		h.Next.Handle(hp)
	}
}

func writeHeader(hp *chain.HandlerParam) {

	doc := document.New()

	// 添加一个段落
	p := doc.AddParagraph()

	// 设置段落样式
	run := p.AddRun()
	run.Properties().SetBold(true)                   // 加粗
	run.Properties().SetFontFamily("Arial")          // 设置字体
	run.Properties().SetSize(14 * measurement.Point) // 设置字体大小
	// 写入文本
	scheme := hp.Param.Scheme
	run.AddText(fmt.Sprintf("%s数据库信息", scheme))

	err := doc.SaveToFile(hp.Path)
	if err != nil {
		panic(err)
	}
}

func writeWord(info map[string]string, hp *chain.HandlerParam) {
	// 因为只能word只能覆盖写，所以需要先读取出之前的内容
	// 创建一个新的Word文档
	doc, _ := document.Open(hp.Path)

	tp := doc.AddParagraph()
	tp.AddRun().AddText(info["tableName"] + "	" + info["comments"])
	// 保存文档
	if err := doc.SaveToFile(hp.Path); err != nil {
		panic(err)
	}

}
