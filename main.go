package main

import (
	"flag"
	"fmt"

	"github.com/Andreassm99/funtemps/conv"
	//"github.com/Andreassm99/funtemps/funfacts"
)

// Definerer flag-variablene i hoved-"scope"
var fahrenheit float64
var celsius float64
var kelvin float64
var out string
var funfacts string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahrenheit, "F", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

	flag.Parse()

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println(fahrenheit, out, funfacts)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))

	// Eksempel på enkel logikk

	er := "er "
	f := "°F"
	c := "°C"
	k := "K"

	// FahrenheitToCelsius
	if out == "C" && isFlagPassed("F") {

		svar := conv.FarhenheitToCelsius(fahrenheit)
		fmt.Printf("%.9g %s %s ", fahrenheit, f, er)
		if svar == float64(int(svar)) {

			fmt.Printf("%d %s\n", int(svar), c)
		} else {
			fmt.Printf("%.2f %s\n", svar, c)
		}

		//fmt.Println(fahrenheit, f, erlik, conv.FarhenheitToCelsius(fahrenheit), c)

	}

	//CelsiusToFahrenheit
	if out == "F" && isFlagPassed("C") {

		svar := conv.CelsiusToFahrenheit(celsius)
		fmt.Printf("%.9g %s %s", celsius, c, er)
		if svar == float64(int(svar)) {

			fmt.Printf("%d %s\n", int(svar), f)
		} else {
			fmt.Printf("%.2f %s\n", svar, f)
		}

		//fmt.Println(celsius, c, erlik, conv.CelsiusToFahrenheit(celsius), f)
	}

	//FahrenheitToKelvin
	if out == "K" && isFlagPassed("F") {

		svar := conv.FahrenheitToKelvin(fahrenheit)
		fmt.Printf("%.9g %s %s", fahrenheit, f, er)
		if svar == float64(int(svar)) {

			fmt.Printf("%d %s\n", int(svar), k)
		} else {
			fmt.Printf("%.2f %s\n", svar, k)
		}

		//fmt.Println(fahrenheit, f, erlik, conv.FahrenheitToKelvin(fahrenheit), k)
	}

	//KelvinToFahrenheit
	if out == "F" && isFlagPassed("K") {

		svar := conv.KelvinToFahrenheit(kelvin)
		fmt.Printf("%.9g %s %s", kelvin, k, er)
		if svar == float64(int(svar)) {
			fmt.Printf("%d %s\n", int(svar), f)
		} else {
			fmt.Printf("%.2f %s\n", svar, f)
		}

		//fmt.Println(kelvin, k, erlik, conv.KelvinToFahrenheit(kelvin), f)
	}

	//CelsiusToKelvin
	if out == "K" && isFlagPassed("C") {

		svar := conv.CelsiusToKelvin(celsius)
		fmt.Printf("%.9g %s %s", celsius, c, er)
		if svar == float64(int(svar)) {
			fmt.Printf("%d %s\n", int(svar), k)
		} else {
			fmt.Printf("%.2f %s\n", svar, k)
		}

		//fmt.Println(celsius, c, erlik, conv.CelsiusToKelvin(celsius), k)
	}

	//KelvinToCelsius
	if out == "C" && isFlagPassed("K") {

		svar := conv.KelvinToCelsius(kelvin)
		fmt.Printf("%.9g %s %s", kelvin, k, er)
		if svar == float64(int(svar)) {
			fmt.Printf("%d %s\n", int(svar), c)
		} else {
			fmt.Printf("%.2f %s\n", svar, c)
		}

		//fmt.Println(kelvin, k, erlik, conv.KelvinToCelsius(kelvin), c)
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
