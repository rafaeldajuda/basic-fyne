package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	result = 0
	number = 0
	s      = ""
)

func main() {

	// START FYNE
	appFyne := app.New()

	// CONFIG MAIN WINDOW
	mainWindow := appFyne.NewWindow("Calculator")
	mainWindow.SetMaster()
	mainWindow.SetFixedSize(true)

	// SET CONTETS
	lblResult := widget.NewLabel("0")

	btnZero := widget.NewButton("0", func() {
		lblResult.SetText(lblZero(lblResult.Text) + "0")
	})

	btnOne := widget.NewButton("1", func() {
		lblResult.SetText(lblZero(lblResult.Text) + "1")
	})
	btnTwo := widget.NewButton("2", func() {
		lblResult.SetText(lblZero(lblResult.Text) + "2")
	})
	btnThree := widget.NewButton("3", func() {
		lblResult.SetText(lblZero(lblResult.Text) + "3")
	})

	btnFour := widget.NewButton("4", func() {
		lblResult.SetText(lblZero(lblResult.Text) + "4")
	})
	btnFive := widget.NewButton("5", func() {
		lblResult.SetText(lblZero(lblResult.Text) + "5")
	})
	btnSix := widget.NewButton("6", func() {
		lblResult.SetText(lblZero(lblResult.Text) + "6")
	})

	btnSeven := widget.NewButton("7", func() {
		lblResult.SetText(lblZero(lblResult.Text) + "7")
	})
	btnEight := widget.NewButton("8", func() {
		lblResult.SetText(lblZero(lblResult.Text) + "8")
	})
	btnNine := widget.NewButton("9", func() {
		lblResult.SetText(lblZero(lblResult.Text) + "9")
	})

	btnPlus := widget.NewButton("+", func() {
		s = "+"
		number, _ = strconv.Atoi(lblResult.Text)
		lblResult.SetText("0")
		calc()
	})

	btnMinus := widget.NewButton("-", func() {
		s = "-"
		number, _ = strconv.Atoi(lblResult.Text)
		lblResult.SetText("0")
		calc()
	})

	btnBackSapce := widget.NewButton("<-", func() {
		if len(lblResult.Text) != 0 {
			lblResult.SetText(lblResult.Text[0 : len(lblResult.Text)-1])

			if len(lblResult.Text) == 0 {
				lblResult.SetText("0")
			}
		}

	})

	btnReset := widget.NewButton("reset", func() {
		lblResult.SetText("0")
		result = 0
		number = 0
		s = ""
	})

	btnCalc := widget.NewButton("calc", func() {
		lblResult.SetText(strconv.Itoa(result))
	})

	lineOne := container.NewHBox(lblResult)
	lineTwo := container.NewHBox(btnNine, btnEight, btnSeven, btnPlus)
	lineThree := container.NewHBox(btnSix, btnFive, btnFour, btnMinus)
	lineFour := container.NewHBox(btnThree, btnTwo, btnOne, btnBackSapce)
	lineSix := container.NewHBox(btnZero, btnCalc, btnReset)

	vNumberBox := container.NewVBox(lineOne, lineTwo, lineThree, lineFour, lineSix)
	hMainBox := container.NewHBox(vNumberBox)

	mainWindow.SetContent(hMainBox)

	// RUN
	mainWindow.ShowAndRun()
}

func lblZero(number string) string {
	if number == "0" {
		return ""
	}

	return number
}

func calc() {

	if s == "+" {
		result = result + number
	} else if s == "-" {
		result = result - number
	} else {
		result = number
	}

	fmt.Println("number", number)
	fmt.Println("sinal", s)
	fmt.Println("result", result)

}
