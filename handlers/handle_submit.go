package handlers

import (
	"github.com/rivo/tview"
	"example.com/gotui/services"
	
)

func PrintInputValues(form *tview.Form, textView *tview.TextView) {
	url := form.GetFormItemByLabel("URL").(*tview.InputField).GetText()
	headers := form.GetFormItemByLabel("Headers").(*tview.TextArea).GetText()

	textView.SetText(headers)
	services.GetRequest(url, headers, textView)
}