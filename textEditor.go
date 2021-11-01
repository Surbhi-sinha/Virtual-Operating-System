package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	// "github.com/sirupsen/logrus/hooks/writer"
)

func texteditor(w fyne.Window) {
	

	var count int =1;

	content:= container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Our text Editor"),
		),
	)

// ADD BUTTON
	content.Add(widget.NewButton("Add new file" , func(){
		content.Add(widget.NewLabel("File" + strconv.Itoa(count)))
		count++
	}))
	// TEXT INPUT FIELD
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")

	input.Resize(fyne.NewSize(400 , 400))

	saveButton := widget.NewButton("Save file" , func(){
		// func NewFileSave(callback func(fyne.URIWriteCloser, error), parent fyne.Window) *FileDialog
		SaveFileName := dialog.NewFileSave(
			func( uc fyne.URIWriteCloser , _ error){
				textData := []byte(input.Text)

				uc.Write(textData)
			} , w)

			SaveFileName.SetFileName("File" + strconv.Itoa(count) + ".txt")
			SaveFileName.Show();

	})


	openButton := widget.NewButton("Open file" , func(){
		openfileDialog := dialog.NewFileOpen(
			func (r fyne.URIReadCloser , _ error)  {
				ReadData ,_ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New file" , ReadData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))
					w.SetContent(container.NewScroll(viewData))
					w.Resize(fyne.NewSize(700 ,400))
					w.Show()
			},w)

			openfileDialog.SetFilter(storage.NewExtensionFileFilter([] string{".txt"}))

			openfileDialog.Show()
	})
	// setting our content inside the container

	tContainer := container.NewVBox(
				content,
				input,
				container.NewHBox(
					saveButton,
					openButton,
				),
			)
		
	
		// window show and rin
		w.SetContent(container.NewBorder(panelContent,nil,nil,nil,tContainer),)
	// window show and rin
	w.Show()

}