package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Mi primer Forms")
	myWindow.Resize(fyne.NewSize(500, 300))

	LApellido := widget.NewLabel("Apellido")
	LNombre := widget.NewLabel("Nombre")

	inApellido := widget.NewEntry()
	inNombre := widget.NewEntry()
	outCuadro := canvas.NewText("", color.Black)
	outCuadro.Move(fyne.NewPos(0, 0))
	bgColor := color.RGBA{R: 200, G: 200, B: 200, A: 255}
	background := canvas.NewRectangle(bgColor)
	background.SetMinSize(fyne.NewSize(100, 100))
	contOut := container.NewWithoutLayout(background, outCuadro)

	BGuardar := widget.NewButton("Guardar", func() {
		text := inApellido.Text + " " + inNombre.Text
		outCuadro.Text = text
		outCuadro.Refresh()
	})
	BEliminar := widget.NewButton("Eliminar", func() {
		outCuadro.Text = ""
	})
	BSalir := widget.NewButton("Salir", func() {
		myApp.Quit()
	})

	ctrlS := &desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: fyne.KeyModifierControl}
	myWindow.Canvas().AddShortcut(ctrlS, func(shortcut fyne.Shortcut) {
		myApp.Quit()
	})

	apellido := container.NewGridWithColumns(
		2, LApellido, inApellido,
	)
	nombre := container.NewGridWithColumns(
		2, LNombre, inNombre,
	)
	botones := container.NewHBox(
		BGuardar, BEliminar,
	)
	salir := container.NewVBox(
		contOut, BSalir,
	)
	column := container.NewVBox(
		apellido, nombre, botones,
	)
	content := container.NewHBox(
		column,
		salir,
	)
	myWindow.CenterOnScreen()
	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}
