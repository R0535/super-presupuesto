package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

type Page struct {
	Title string
	Body  []byte
}

const Port = ":5500"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", rootPage)
	router.HandleFunc("/cuenta/nueva/{fetchCuentaName}", nuevacuenta).Methods("GET", "POST")

	fmt.Println("Serving @ http://127.0.0.1" + Port)
	log.Fatal(http.ListenAndServe(Port, router))
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Felicidades, Conectado"))
}

func nuevacuenta(w http.ResponseWriter, r *http.Request) {
	fetchCuentaName := mux.Vars(r)["fetchCuentaName"]
	guid := xid.New()

	newAccount := Cuenta{
		id:    guid.String(),
		fecha: time.Now(),
		name:  fetchCuentaName,
	}

	jsonAccount, err := json.Marshal(newAccount.id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonAccount)
	}
	//Aqui deberia ir la insercion en la base de datos
	//w.Write([]byte("ID: " + newAccount.id + " Fecha: " + newAccount.fecha.Local().String() + " name: " + newAccount.name))
}

type Cuenta struct {
	id      string
	name    string
	fecha   time.Time
	Packs   []Pack
	Equipo  Equipo
	Totales Totales
}

type Pack struct {
	id          int
	Actividades []Actividad
}

type Actividad struct {
	id          int
	Titulo      string
	Categoria   string
	Semanas     string
	Descripcion string
}

type PasivoEstatico struct {
	id     int
	Titulo string
	Precio string
}

type Equipo struct {
	id       int
	Personas []Persona
}

type Persona struct {
	id              int
	Nombre          string
	Puesto          string
	SalarioPromedio uint64
}

type Totales struct {
	id         int
	Top        uint32
	TopDolares uint32

	Mid        uint32
	MidDolares uint32

	Bot        uint32
	BotDolares uint32
}
