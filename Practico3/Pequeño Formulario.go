package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func main() {
	form := app.New()
	ventana := form.NewWindow("Pequeño Formulario")
	ventana.Resize(fyne.NewSize(300, 250))
	LDni := widget.NewLabel("DNI")
	LNombre := widget.NewLabel("Nombre")
	LApellido := widget.NewLabel("Apellido")
	LNyA := widget.NewLabel("Nombre y Apellido:")
	LModificar := canvas.NewText("Modificar", color.RGBA{255, 0, 0, 255})

	TDni := widget.NewEntry()
	TDni.Resize(fyne.NewSize(100, 50))

	TApellido := widget.NewEntry()

	TNombre := widget.NewEntry()

	TDni.Validator = validation.NewRegexp(`^[0-9]*$`, "Solo se permiten números")
	TApellido.Validator = validation.NewRegexp(`^[a-zA-Z]*$`, "Solo se permiten letras")
	TNombre.Validator = validation.NewRegexp(`^[a-zA-Z]*$`, "Solo se permiten letras")

	TGuardar := widget.NewButton("Guardar", func() {
		if TApellido.Validate() == nil && TNombre.Validate() == nil {
			LModificar.Text = TNombre.Text + " " + TApellido.Text
			LModificar.Refresh()
		}
	})
	TEliminar := widget.NewButton("Eliminar", func() {
		LModificar.Text = "Modificar"
		LModificar.Refresh()
	})

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
