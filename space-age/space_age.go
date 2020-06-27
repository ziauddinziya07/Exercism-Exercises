package space

import (
	"fmt"
	"strconv"
)

type Planet string


func Age ( sec float64, pla string ) ( res float64 ) {

	var temp float64
	
	switch pla {
		// temp = Earth-Period * Planet-Period
		case "Mercury" : temp = 7600543.82
		case "Venus" : temp = 19414149.1
		case "Earth" : temp = 31557600.0
		case "Mars" : temp = 59354032.7
		case "Jupiter" : temp = 374355659.0
		case "Saturn" : temp = 929292363.0
		case "Uranus" : temp = 2651370019.0
		case "Neptune" : temp = 5200418560.0
	}

	res, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", sec / temp), 64)
	return 
}