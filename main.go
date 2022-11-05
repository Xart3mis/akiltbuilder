package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main(){
  a := app.New()
  a.Settings().SetTheme(&dark{})  
  
  w := a.NewWindow("AKILT Builder")

  w.SetContent(container.NewVBox())
  w.ShowAndRun()
}
