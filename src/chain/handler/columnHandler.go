package handler

import (
	"github.com/unidoc/unioffice/document"
	"table_desc/src/chain"
)

type ColumnHandler struct {
	Next chain.Handler
}

func (h *ColumnHandler) SetNext(next chain.Handler) {
	h.Next = next
}

func (h *ColumnHandler) Handle(hp *chain.HandlerParam) {
	data := hp.Ctx.QueryColumnInfo(hp.TableName, hp.Param.Scheme, hp.Db)
	writeColWord(hp, data)
	// 写入文件
	if h.Next != nil {
		h.Next.Handle(hp)
	}
}

func writeColWord(hp *chain.HandlerParam, data []map[string]string) {
	if len(data) == 0 {
		return
	}
	doc, err := document.Open(hp.Path)
	if err != nil {
		table := doc.AddTable()
		table.Properties().SetWidthAuto()
		headerRow := table.AddRow()
		headerRow.AddCell().AddParagraph().AddRun().AddText("columnName")
		headerRow.AddCell().AddParagraph().AddRun().AddText("dataType")
		headerRow.AddCell().AddParagraph().AddRun().AddText("dataLength")
		headerRow.AddCell().AddParagraph().AddRun().AddText("columnComment")
		for _, colInfo := range data {
			dataRow := table.AddRow()
			dataRow.AddCell().AddParagraph().AddRun().AddText(colInfo["colName"])
			dataRow.AddCell().AddParagraph().AddRun().AddText(colInfo["dataType"])
			dataRow.AddCell().AddParagraph().AddRun().AddText(colInfo["dataLength"])
			dataRow.AddCell().AddParagraph().AddRun().AddText(colInfo["colComment"])
		}
		err = doc.SaveToFile(hp.Path)
		if err != nil {
			panic(err)
		}
	}
	panic(err)
}
