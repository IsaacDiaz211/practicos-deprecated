package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	form := app.New()
	ventana := form.NewWindow("Pequeño Formulario")

	LDni := widget.NewLabel("DNI")
	LNombre := widget.NewLabel("Nombre")
	LApellido := widget.NewLabel("Apellido")
	LNyA := widget.NewLabel("Nombre y Apellido:")
	LModificar := canvas.NewText("Modificar", color.RGBA{255, 0, 0, 255})

	TDni := widget.NewEntry()
	TApellido := widget.NewEntry()
	TNombre := widget.NewEntry()

	TGuardar := widget.NewButton("Guardar", func() {})
	TEliminar := widget.NewButton("Eliminar", func() {})

	columLf := container.NewGridWithRows(
		5, LNyA, LDni, LApellido, LNombre, TGuardar,
	)
	columRg := container.NewGridWithRows(
		5, LModificar, TDni, TApellido, TNombre, TEliminar,
	)
	content := container.NewHBox(columLf, columRg)
	ventana.SetContent(content)
	ventana.ShowAndRun()
}
