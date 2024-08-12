package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
	outCuadro := widget.NewMultiLineEntry()
	outCuadro.Resize(fyne.NewSize(100, 100))

	BGuardar := widget.NewButton("Guardar", func() {
		outCuadro.SetText(inApellido.Text + " " + inNombre.Text)
	})
	BEliminar := widget.NewButton("Eliminar", func() {
		outCuadro.SetText("")
	})
	BSalir := widget.NewButton("Salir", func() {
		myApp.Quit()
	})
	BSalir.Resize(fyne.NewSize(40, 70))

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
		outCuadro, BSalir,
	)
	column := container.NewVBox(
		apellido, nombre, botones,
	)
	content := container.NewGridWithColumns(2, column, salir)
	myWindow.CenterOnScreen()
	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}
