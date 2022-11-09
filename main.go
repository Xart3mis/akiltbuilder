package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&dark{})

	w := a.NewWindow("AKILT Builder")

	w.Resize(fyne.NewSize(800, 450))

	var server_ip string

	server_ip_entry := widget.NewEntry()

	submit_btn := widget.NewButton("Start Build", func() {
		CloneAKILT()
		scanner, runner := BuildwithServerIP(server_ip)

		runner.Starter()
		for (*scanner).Scan() {
			fmt.Println(scanner.Text())
		}

		runner.Waiter()

		fmt.Println(server_ip)
	})

	form := widget.NewForm(widget.NewFormItem("Server IP", server_ip_entry), widget.NewFormItem("Done", submit_btn))

	server_ip_entry.OnSubmitted = func(s string) {
		fmt.Println(s)
	}

	server_ip_entry.Validator = ValidateRegexWithFormSubmitDisable(`^((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]?|0)\.){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]?|0)$`, "Invalid IP", submit_btn)

	server_ip_entry.OnChanged = func(s string) {
		server_ip = s
	}

	w.SetContent(container.NewVBox(form))
	w.ShowAndRun()
}
