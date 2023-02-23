package handlers

import (
	"github.com/rivo/tview"
	"example.com/gotui/services"
)

func PrintInputValues(form *tview.Form, textView *tview.TextView) {
	url := form.GetFormItemByLabel("URL").(*tview.InputField).GetText()
	headers := form.GetFormItemByLabel("Headers").(*tview.TextArea).GetText()
	body := form.GetFormItemByLabel("Body").(*tview.TextArea).GetText()

	dropdown := form.GetFormItemByLabel("Method").(*tview.DropDown)
	_, method := dropdown.GetCurrentOption()
	
	switch method {
	case "GET":
		services.GetRequest(url, headers, textView)
	case "POST":
		services.PostRequest(url, headers, body, textView)
	case "PATCH":
		services.PatchRequest(url, headers, body, textView)
	case "DELETE":
		services.DeleteRequest(url, headers, body, textView)
	default:
		textView.SetText("Invalid method selected")
	}
	
}