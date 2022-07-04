package main

type celsius float64
type fahrenheit float64
type kelvin float64

func main() {

}

func (c celsius) toFahrenheit() fahrenheit {
	f := fahrenheit((c * 9.0 / 5.0) + 32.0)
	return f
}

func (c celsius) toKelvin() kelvin {
	k := kelvin(c + 273.15)
	return k
}
