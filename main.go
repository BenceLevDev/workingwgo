package main

import (
	"fmt"
	"io/ioutil" // Kérés testének (Body) kiolvasásához
	"log"       // Naplózáshoz (konzolra íráshoz)
	"net/http"  // HTTP szerver funkciókhoz
)

// A program belépési pontja
func main() {
	// 1. Útvonal regisztrálása: "/" (gyökér útvonal)
	// A függvény megkapja a válasz írót (rw) és a kérés objektumot (r)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World") // Naplózza a kérés érkezését

		// Kérés testének (r.Body) kiolvasása és eltárolása a 'd' változóban
		d, err := ioutil.ReadAll(r.Body) 

		// Hibakezelés: ha a kiolvasás sikertelen volt (err nem nil)
		if err != nil {
			// HTTP 400 Bad Request státusz küldése a válaszban ('rw'-n keresztül)
			http.Error(rw, "oops", http.StatusBadRequest)
			return // Kilépés a függvényből
		}

		// A válasz testének összeállítása és elküldése: "Hello " + beolvasott adat (d)
		fmt.Fprintf(rw, "Hello %s", d)
	})

	// 2. Másik útvonal regisztrálása: "/goodbye"
	// Itt a paramétereket nem használjuk, ezért a nevüket elhagyjuk.
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// A szerver indítása a 9090-es porton. 
	// A 'nil' azt jelenti, hogy az alapértelmezett routert (DefaultServeMux) használjuk.
	log.Fatal(http.ListenAndServe(":9090", nil))
}