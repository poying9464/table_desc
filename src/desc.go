package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	"log"
	"os"
	"strconv"
	"strings"
	"table_desc/src/chain"
	"table_desc/src/chain/handler"
	queryContext "table_desc/src/db/context"
	"table_desc/src/db/dm"
	"table_desc/src/db/mysql"
	"table_desc/src/db/oracle"
	"table_desc/src/entity"
	"table_desc/src/ui"
	"time"
)

func init() {
	paths := findfont.List()
	for _, path := range paths {
		if strings.Contains(path, "HYYuJinLi-45J-2") {
			_ = os.Setenv("FYNE_FONT", path)
		}
	}
}

func main() {
	myApp := app.New()

	myWindow := myApp.NewWindow("table desc")
	myWindow.Resize(fyne.NewSize(500, 250))
	var dbType string
	newSelect := widget.NewSelect([]string{"mysql", "dm", "oracle"}, func(s string) {
		dbType = s
	})
	newSelect.PlaceHolder = "Select db type"
	// 创建表单项
	formItems := []*widget.FormItem{
		widget.NewFormItem("db type", newSelect),
		widget.NewFormItem("host", ui.CreateNormalEntry("Enter host address")),
		widget.NewFormItem("port", ui.CreateNormalEntry("Enter port")),
		widget.NewFormItem("username", ui.CreateNormalEntry("Entry username")),
		widget.NewFormItem("password", ui.CreateNormalEntry("Entry password")),
		widget.NewFormItem("scheme", ui.CreateNormalEntry("Enter scheme or dbname")),
	}
	form := widget.NewForm(formItems...)
	content := container.NewVBox(form)

	subBtn := ui.CreateNormalButton("提价", func() {
		// 当点击确定时，加载loading窗口
		progressValue := binding.NewFloat()
		progressBar := widget.NewProgressBarWithData(progressValue)
		ladingContent := container.NewVBox(
			widget.NewLabel("Loading..."),
			progressBar,
		)
		myWindow.Hide()
		myWindow.SetContent(ladingContent)

		// 将表单的参数赋值到connectParam中
		connectParam := entity.ConnectParam{
			Username: formItems[3].Widget.(*widget.Entry).Text,
			Password: formItems[4].Widget.(*widget.Entry).Text,
			Scheme:   formItems[5].Widget.(*widget.Entry).Text,
			Host:     formItems[1].Widget.(*widget.Entry).Text,
		}
		connectParam.Port, _ = strconv.Atoi(formItems[2].Widget.(*widget.Entry).Text)
		var opContext *queryContext.OpContext
		// 根据类型选择不同的数据库操作对象
		switch dbType {
		case "dm":
			opContext = queryContext.NewOpContext(&dm.Operation{})
		case "mysql":
			opContext = queryContext.NewOpContext(&mysql.Operation{})
		case "oracle":
			opContext = queryContext.NewOpContext(&oracle.Operation{})
		}
		connectHandler := &handler.ConnectHandler{
			Next: &handler.TableHandler{
				Next: &handler.ColumnHandler{
					Next: &handler.CloseHandler{},
				},
			},
		}
		param := chain.HandlerParam{
			Param:     connectParam,
			Path:      fmt.Sprintf("%s%s%s", "table_desc", time.Now().Format("20060102150405"), ".doc"),
			TableName: "",
			Ctx:       opContext,
			Allows:    true,
		}
		connectHandler.Handle(&param)
		myApp.Quit()
	})
	container.New(layout.NewCenterLayout(), subBtn)

	// 将submitBtn 添加到表单中
	// form.AppendItem(submitBtn)

	content = container.New(layout.NewVBoxLayout(), content, subBtn)
	border := container.NewBorder(
		ui.CreateSpacer(fyne.NewSize(20, 20)),
		ui.CreateSpacer(fyne.NewSize(20, 20)),
		ui.CreateSpacer(fyne.NewSize(5, 5)),
		ui.CreateSpacer(fyne.NewSize(5, 5)),
		content)

	// border.Resize(fyne.NewSize(500, 250))

	myWindow.SetContent(border)
	myWindow.SetOnClosed(func() {
		log.Println("app closed")
		myApp.Quit()
	})
	myWindow.ShowAndRun()
}
