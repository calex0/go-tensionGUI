package main

import (
	"strconv"

	"github.com/andlabs/ui"
)

func diagnostico(max float32, min float32) string {

	var diagnostico = ""

	if max < 120 && min < 80 {
		diagnostico = "Normal"
	}
	if (max >= 120 && max < 129) && (min < 80) {
		diagnostico = "Elevada"
	}
	if (max >= 130 && max < 139) || (min >= 80 && min <= 90) {
		diagnostico = "Alta. Hipertension nivel 1"
	}
	if max >= 140 || min >= 90 {
		diagnostico = "Alta. Hipertension nivel 2"
	}
	if (max > 180 && min > 120) || (max > 180 || min > 120) {
		diagnostico = "CRISIS DE HIPERTENSIÓN\nConsulte a su médico de inmediato"
	}

	return diagnostico

}

func main() {
	err := ui.Main(func() {

		window := ui.NewWindow("Tensión arterial", 400, 320, false)
		window.SetMargined(true)
		window.OnClosing(func(*ui.Window) bool {
			return true
		})

		maximaBox := ui.NewHorizontalBox()
		maximaBox.SetPadded(true)
		maximaBox.Append(ui.NewLabel("Máxima"), false)
		maxima := ui.NewEntry()
		maxima.SetText("120")
		maximaBox.Append(maxima, true)

		minimaBox := ui.NewHorizontalBox()
		minimaBox.SetPadded(true)
		minimaBox.Append(ui.NewLabel("Mínima"), false)
		minima := ui.NewEntry()
		minima.SetText("80")
		minimaBox.Append(minima, true)

		labelCuota := ui.NewLabel("")

		buttonBox := ui.NewHorizontalBox()
		buttonBox.SetPadded(true)
		buttonBox.Append(ui.NewLabel(""), true)

		cancel := ui.NewButton("Cancelar")
		cancel.OnClicked(func(*ui.Button) {
			ui.Quit()
		})

		buttonBox.Append(cancel, false)

		calcular := ui.NewButton("Diagnosticar")
		calcular.OnClicked(func(*ui.Button) {

			var maxima1, _ = strconv.ParseFloat(maxima.Text(), 64)
			var minima1, _ = strconv.ParseFloat(minima.Text(), 64)

			maxima2 := float32(maxima1)
			minima2 := float32(minima1)

			var res = diagnostico(maxima2, minima2)

			labelCuota.SetText("Diagnóstico: " + res)

		})
		buttonBox.Append(calcular, false)

		layout := ui.NewVerticalBox()
		layout.SetPadded(true)
		layout.Append(maximaBox, false)
		layout.Append(minimaBox, false)

		layout.Append(labelCuota, false)

		layout.Append(buttonBox, false)
		window.SetChild(layout)
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
