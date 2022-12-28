# greenhouse-coffee

Runs calculations on growing coffee in a greenhouse, regardless of where you may live.

## Get Started

First copy `weather_data.go.example` to `weather_data.go`:

```shell
cp weather_data.go.example weather_data.go
```

Then, fill in the `weather_data.go` file with your own data. You will need to provide the high and lows for your location for each month of the year. You can find this data for example on [weather.com](https://weather.com/).

Once that's done, you can run the program with:

```shell
go run main.go
```

A series of charts will open in your browser, showing the results of the calculations, as well as various statements printed out to the terminal.

You can change the values in `main.go` to change the type of coffee plant, the target crop yield per year, and the thermal efficiency of the greenhouse.