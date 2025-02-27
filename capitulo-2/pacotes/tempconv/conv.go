package tempconv

// CTof converte uma temperatura em Celsius para Fahrenheit
func CTof(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converte uma temperatura em Fahrenheit para Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
