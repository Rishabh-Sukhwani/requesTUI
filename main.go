package main

import (
	"github.com/rivo/tview"
	"example.com/gotui/handlers"
)

func main() {
	
	app := tview.NewApplication()
	
	form := tview.NewForm().
			AddInputField("URL", "", 500, nil, nil).
			AddDropDown("Method", []string{"GET", "POST", "PATCH", "DELETE"}, 0, nil).
			AddTextArea("Headers", "", 0, 7, 0, nil).
			AddTextArea("Body", "", 0, 10, 0, nil)

	bufferBox := tview.NewTextView()
	
	textView := tview.NewTextView().
				SetDynamicColors(true).
				SetRegions(true).
				SetWordWrap(true)

	form.AddButton("Submit", func() {
		handlers.PrintInputValues(form, textView)
	})
	
	flex := tview.NewFlex().
			AddItem(form, 0, 4, true).
			AddItem(bufferBox, 0, 1, false).
			AddItem(textView, 0, 4, false)


	form.SetBorder(true).SetTitle("POSTMAN").SetTitleAlign(tview.AlignCenter)
	textView.SetTitle("RESPONSE:").SetTitleAlign(tview.AlignCenter)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
