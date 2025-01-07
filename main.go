package main

import (
	"math"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func increase_in_price_percent(new_price, old_price float64) float64 {

	return ((new_price - old_price) / new_price) * 100

}

func percentage_loss_per_unit(old_profit_percentage, new_price, old_price float64) float64 {
	return math.Abs(old_profit_percentage - increase_in_price_percent(new_price, old_price))
}

func percentage_gain_of_changed_unit(old_profit_percentage, new_price, old_price float64) float64 {
	return old_profit_percentage + increase_in_price_percent(new_price, old_price)
}

func percentage_of_units_to_to_change_its_price(old_profit_percentage, new_price, old_price float64) float64 {
	return percentage_loss_per_unit(old_profit_percentage, new_price, old_price) / percentage_gain_of_changed_unit(old_profit_percentage, new_price, old_price)
}

func number_of_changed_units_to_break_even(old_profit_percentage, new_price, old_price, bought_units_with_old_price float64) float64 {
	return percentage_of_units_to_to_change_its_price(old_profit_percentage, new_price, old_price) * bought_units_with_old_price
}

func number_of_changed_units_to_provide_enough_capital_to_buy_an_amount_of(old_profit_percentage, new_price, old_price, bought_units_with_old_price, amount_to_buy_with_new_price float64) float64 {
	ratio_of_new_amount_to_old_amount := amount_to_buy_with_new_price / bought_units_with_old_price
	buying_factor := 1 + ratio_of_new_amount_to_old_amount
	result := number_of_changed_units_to_break_even(old_profit_percentage, new_price, old_price, bought_units_with_old_price) * buying_factor
	if result < 0 {
		return 0
	}
	return result
}

func main() {

	const (
		app_title     = "Khiba"
		window_height = 600
		window_width  = 800
	)
	go func() {
		window := new(app.Window)
		window.Option(app.Title(app_title))
		window.Option(app.Size(unit.Dp(window_width), unit.Dp(window_height)))

		var ops op.Ops

		var calculate_button widget.Clickable

		theme := material.NewTheme()

		for {
			event := window.Event()

			switch typ := event.(type) {
			case app.FrameEvent:
				{
					gtx := app.NewContext(&ops, typ)
					btn := material.Button(theme, &calculate_button, "Calculate")
					btn.Layout(gtx)
					typ.Frame(gtx.Ops)
				}
			case app.DestroyEvent:
				os.Exit(0)
			}
		}
	}()

	app.Main()

}
