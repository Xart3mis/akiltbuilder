package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetClientCanvas() *fyne.Container{
  server_ip := widget.NewEntry()

  return container.NewVBox(widget.NewLabel("Client Configuration"), widget.NewForm(widget.NewFormItem("Server IP", server_ip)))
}
