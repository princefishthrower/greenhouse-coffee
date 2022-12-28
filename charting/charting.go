package charting

import (
	"coffee/constants"
	"fmt"
	"os"
	"reflect"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func getField(input interface{}, field string) (interface{}, error) {
	val := reflect.ValueOf(input)
	if !val.IsValid() {
		return nil, fmt.Errorf("invalid value")
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("value is not a struct")
	}

	fieldVal := val.FieldByName(field)
	if !fieldVal.IsValid() {
		return nil, fmt.Errorf("field not found")
	}

	return fieldVal.Interface(), nil
}

func getMonths(monthlyData []constants.MonthlyData) []string {
	months := make([]string, 0)
	for i := 0; i < len(monthlyData); i++ {
		months = append(months, monthlyData[i].Month)
	}
	return months
}

// generate random data for bar chart
func generateBarItems(monthlyData []constants.MonthlyData, keyValue string) []opts.BarData {
	// dynamically get the value from the monthlyData array
	items := make([]opts.BarData, 0)
	for i := 0; i < len(monthlyData); i++ {
		value, err := getField(monthlyData[i], keyValue)
		if err != nil {
			fmt.Println(err)
		}
		items = append(items, opts.BarData{Value: value})
	}
	return items
}

func BuildMonthlyEnergyCostChart(monthlyData []constants.MonthlyData) {
	// create a new bar instance
	bar := charts.NewBar()

	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Monthly Coffee Growing Costs",
			Subtitle: "Greenhouse Coffee - Grow Coffee Sustainably Anywhere",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Month",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "EUR/Month",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Right: "80%"}),
		charts.WithColorsOpts(opts.Colors{"blue"}),
	)

	// Put data into instance - x axis is the Month property from monthlyData
	bar.SetXAxis(getMonths(monthlyData)).
		AddSeries("Energy Cost", generateBarItems(monthlyData, "EnergyCost")).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
			}),
		)

	// Where the magic happens
	f, _ := os.Create("monthly_energy_cost.html")
	bar.Render(f)
}

func BuildMonthlyEnergyUsageChart(monthlyData []constants.MonthlyData) {
	// create a new bar instance
	bar := charts.NewBar()

	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Monthly Energy Usage",
			Subtitle: "Greenhouse Coffee - Grow Coffee Sustainably Anywhere",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Month",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "kWh/Month",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Right: "80%"}),
		charts.WithColorsOpts(opts.Colors{"pink"}),
	)

	// Put data into instance - x axis is the Month property from monthlyData
	bar.SetXAxis(getMonths(monthlyData)).
		AddSeries("Energy Usage", generateBarItems(monthlyData, "EnergyUsage")).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
			}),
		)

	// Where the magic happens
	f, _ := os.Create("monthly_energy_usage.html")
	bar.Render(f)
}

func BuildTotalsChart(totalEnergyKWh float64, totalCostEUR float64) {
	// Create a new horizontal bar chart
	bar := charts.NewBar()

	// Set the chart properties
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Energy Consumption and Cost Totals",
		Subtitle: "Greenhouse Coffee - Grow Coffee Sustainably Anywhere",
	}))

	// Add the data series
	bar.SetXAxis([]string{"Total"}).
		AddSeries("Total Energy Consumption", []opts.BarData{{Value: totalEnergyKWh}}).
		AddSeries("Total Energy Cost", []opts.BarData{{Value: totalCostEUR}})

	// Render the chart
	// Where the magic happens
	f, _ := os.Create("totals.html")
	bar.Render(f)
}
