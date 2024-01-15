package temp

import (
	"flag"
	"fmt"
)

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var val float64
	fmt.Sscanf(s, "%f%s", &val, &unit)
	switch unit {
	case "C", "c":
		f.Celsius = Celsius(val)
		return nil
	case "F", "f":
		f.Celsius = FToC(Fahrenheit(val))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
