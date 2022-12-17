package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

var ColorBlack = color.RGBA{0, 0, 0, 255}
var ColorCyan = color.RGBA{0, 55, 55, 255}
var ColorLightCyan = color.RGBA{0, 200, 200, 255}
var ColorWhite = color.RGBA{255, 255, 255, 255}
var Bill = 0.0
var Tips = 0
var NumPeople = 1

func main() {
	a := app.New()
	w := a.NewWindow("Tips Calculator")
	icon, _ := fyne.LoadResourceFromPath("icon.png")
	w.SetIcon(icon)
	// right side
	labelTotal := canvas.NewText("Total", ColorLightCyan)
	labelTotalNum := canvas.NewText("$0.00", ColorWhite)
	labelTotal.TextStyle = fyne.TextStyle{Italic: true}
	labelTotalNum.TextStyle = fyne.TextStyle{Bold: true}

	labelTipAmount := canvas.NewText(fmt.Sprintf("Tip per person"), ColorLightCyan)
	labelTipAmountNum := canvas.NewText("$0.00", ColorWhite)
	labelTipAmount.TextStyle = fyne.TextStyle{Italic: true}
	labelTipAmountNum.TextStyle = fyne.TextStyle{Bold: true}

	labelTotalPerson := canvas.NewText("Total per person", ColorLightCyan)
	labelTotalPersonNum := canvas.NewText("$0.00", ColorWhite)
	labelTotalPerson.TextStyle = fyne.TextStyle{Italic: true}
	labelTotalPersonNum.TextStyle = fyne.TextStyle{Bold: true}

	labelTipAmountNum.TextSize = 2 * labelTipAmount.TextSize
	labelTotalNum.TextSize = 2 * labelTotal.TextSize
	labelTotalPersonNum.TextSize = 2 * labelTotalPerson.TextSize

	// left side
	labelBill := canvas.NewText("Bill", ColorCyan)
	labelSelectTip := canvas.NewText("Select Tip %", ColorCyan)
	labelNumPeople := canvas.NewText("Number of people", ColorCyan)

	entryBill := widget.NewEntry()
	entryBill.PlaceHolder = "0"
	entryCustomTip := widget.NewEntry()
	entryCustomTip.PlaceHolder = "Custom"
	entryNumPeople := widget.NewEntry()
	entryNumPeople.PlaceHolder = "1"

	entryBill.OnChanged = func(string) {
		Bill = GetFloatNumFromEntry(entryBill, 0)
		CalculateTips(labelTotalNum, labelTipAmountNum, labelTotalPersonNum, Bill, Tips, NumPeople)
	}

	entryNumPeople.OnChanged = func(string) {
		NumPeople = GetIntNumFromEntry(entryNumPeople, 1)
		CalculateTips(labelTotalNum, labelTipAmountNum, labelTotalPersonNum, Bill, Tips, NumPeople)
	}

	btn5Tip := widget.NewButton("", nil)
	btn10Tip := widget.NewButton("", nil)
	btn15Tip := widget.NewButton("", nil)
	btn20Tip := widget.NewButton("", nil)
	btn25Tip := widget.NewButton("", nil)

	button5Tip := ColoredButton("5%", ColorCyan, btn5Tip)
	button10Tip := ColoredButton("10%", ColorCyan, btn10Tip)
	button15Tip := ColoredButton("15%", ColorCyan, btn15Tip)
	button20Tip := ColoredButton("20%", ColorCyan, btn20Tip)
	button25Tip := ColoredButton("25%", ColorCyan, btn25Tip)

	buttons := []*fyne.Container{button5Tip, button10Tip, button15Tip, button20Tip, button25Tip}

	btn5Tip.OnTapped = func() {
		Tips = 5
		CalculateTips(labelTotalNum, labelTipAmountNum, labelTotalPersonNum, Bill, Tips, NumPeople)
		ButtonsOnClick(0, buttons)
	}
	btn10Tip.OnTapped = func() {
		Tips = 10
		CalculateTips(labelTotalNum, labelTipAmountNum, labelTotalPersonNum, Bill, Tips, NumPeople)
		ButtonsOnClick(1, buttons)
	}
	btn15Tip.OnTapped = func() {
		Tips = 15
		CalculateTips(labelTotalNum, labelTipAmountNum, labelTotalPersonNum, Bill, Tips, NumPeople)
		ButtonsOnClick(2, buttons)
	}
	btn20Tip.OnTapped = func() {
		Tips = 20
		CalculateTips(labelTotalNum, labelTipAmountNum, labelTotalPersonNum, Bill, Tips, NumPeople)
		ButtonsOnClick(3, buttons)
	}
	btn25Tip.OnTapped = func() {
		Tips = 25
		CalculateTips(labelTotalNum, labelTipAmountNum, labelTotalPersonNum, Bill, Tips, NumPeople)
		ButtonsOnClick(4, buttons)
	}

	entryCustomTip.OnChanged = func(string) {
		Tips = GetIntNumFromEntry(entryCustomTip, 0)
		CalculateTips(labelTotalNum, labelTipAmountNum, labelTotalPersonNum, Bill, Tips, NumPeople)
		ButtonsOnClick(-1, buttons)
	}

	// buttons
	containerTips := container.NewGridWithRows(
		2,
		container.NewGridWithColumns(3, button5Tip, button10Tip, button15Tip),
		container.NewGridWithColumns(3, button20Tip, button25Tip, entryCustomTip),
	)

	// all left side
	leftContainer := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		labelBill,
		entryBill, layout.NewSpacer(),
		labelSelectTip,
		containerTips, layout.NewSpacer(),
		labelNumPeople,
		entryNumPeople,
		layout.NewSpacer(),
	)

	rightContainer := container.New(
		layout.NewMaxLayout(),
		canvas.NewRectangle(ColorCyan),
		container.New(&myLayout{},
			labelTotal, labelTotalNum,
			labelTipAmount, labelTipAmountNum,
			labelTotalPerson, labelTotalPersonNum,
		),
	)

	leftContainer.Resize(fyne.NewSize(300, 300))
	rightContainer.Resize(fyne.NewSize(100, 300))
	w.Resize(fyne.NewSize(400, 300))

	w.SetContent(container.NewHBox(leftContainer, rightContainer))
	w.ShowAndRun()
}

// calculate at any action
func CalculateTips(labelTotal *canvas.Text, labelTipAmountPerson *canvas.Text, labelTotalPerson *canvas.Text, bill float64, tips int, numpeople int) {
	tipWithBill := bill * (float64(tips) / 100)
	billWithTips := bill * (float64(100+tips) / 100)
	tipsPerPerson := tipWithBill / float64(numpeople)
	totalPerPerson := billWithTips / float64(numpeople)

	labelTotal.Text = fmt.Sprintf("$ %.2f", billWithTips)
	labelTipAmountPerson.Text = fmt.Sprintf("$ %.2f", tipsPerPerson)
	labelTotalPerson.Text = fmt.Sprintf("$ %.2f", totalPerPerson)

	labelTotal.Refresh()
	labelTipAmountPerson.Refresh()
	labelTotalPerson.Refresh()
}

func ColoredButton(label string, c color.RGBA, b *widget.Button) *fyne.Container {
	rect := canvas.NewRectangle(c)
	rect.SetMinSize(fyne.NewSize(100, 40))
	text := canvas.NewText(label, ColorLightCyan)
	text.Alignment = fyne.TextAlignCenter
	button := container.New(
		layout.NewMaxLayout(),
		rect,
		b,
		text,
	)
	return button
}

func ButtonsOnClick(change int, buttons []*fyne.Container) {

	for i, btn := range buttons {
		if i == change {
			btn.Objects[0].(*canvas.Rectangle).FillColor = ColorLightCyan
			btn.Objects[2].(*canvas.Text).Color = ColorCyan
			btn.Refresh()
			continue
		}
		btn.Objects[0].(*canvas.Rectangle).FillColor = ColorCyan
		btn.Objects[2].(*canvas.Text).Color = ColorLightCyan
		btn.Refresh()
	}

}

func GetIntNumFromEntry(entry *widget.Entry, minVal int) int {
	num, err := strconv.Atoi(entry.Text)
	if err != nil || num < minVal {
		entry.Text = ""
		entry.SetPlaceHolder(fmt.Sprintf("%v", minVal))
		entry.Refresh()
		return 1
	}
	return num
}

func GetFloatNumFromEntry(entry *widget.Entry, minVal float64) float64 {
	num, err := strconv.ParseFloat(entry.Text, 64)
	if err != nil || num < minVal {
		entry.Text = ""
		entry.SetPlaceHolder(fmt.Sprintf("%.1f", minVal))
		entry.Refresh()
		return 0
	}
	return num
}

type myLayout struct {
}

func (d *myLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(350, 250)
}

func (d *myLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(20, 45)
	generalSize := fyne.Size{20, 45}
	for i, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		if i%2 == 0 {
			pos = fyne.NewPos(20, generalSize.Height+size.Height/2+45)
		} else {
			pos = fyne.NewPos(containerSize.Width-size.Width-20, generalSize.Height+45)
			generalSize.Height += size.Height
		}
		o.Move(pos)

	}
}
