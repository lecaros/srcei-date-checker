/*
Copyright © 2023 José Lecaros Cisterna jose.lecaros AT gmail DOT com
*/
package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var tramite string
var region int

var rootCmd = &cobra.Command{
	Use:     "srcei-date-checker",
	Version: "0.1.1",
	Short:   "Obtener la cita más cercana para el trámite y región indicado.",
	Long: `Obtener la cita más cercana para el trámite y región indicado.
Ejemplo:
srcei-date-checker -t 2-5 -r 13

Buscará la cita más próxima para Primera obtención de Cédula de Identidad para extranjero en la Región Metropolitana.

Para obtener la lista de regiones y trámites disponibles, ejecutar el comando sin opciones.

Trámite a realizar. Valores posibles:
- Primera obtención de Cédula de Identidad para chileno: 2-4
- Primera obtención de Cédula de Identidad para extranjero: 2-5
- Reimpresión de Cédula: 2-22
- Renovación de Cédula Chileno/a: 2-6

`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		oficinasFromRegion := findOficinasFromRegion()

		// traverse oficinasFromRegion and return the best date
		FindTheNearestDate(oficinasFromRegion)
	},
}

func FindTheNearestDate(oficinas []Oficinas) {
	var bestDate time.Time
	var bestOficina string
	var bestDateAndOffice []BestOfficeDate

	println("Buscando la cita más cercana para el trámite " + tramite + " en la región " + strconv.Itoa(region) + "...")
	for _, oficina := range oficinas {

		url := createURL("https://agenda.registrocivil.cl/api/backend/horas/", tramite, oficina.CodigoOficina)
		resp, _ := http.Get(url)

		var availableDates AvailableDates

		if err := json.NewDecoder(resp.Body).Decode(&availableDates); err != nil {
			println("Error al obtener las fechas disponibles para la oficina " + oficina.NombreOficina)
			println(err.Error())
			os.Exit(1)
		}
		//get only the first date of the office
		if len(availableDates.Horas) > 0 {
			date, _ := time.Parse("2006-01-02 15:04:05", availableDates.Horas[0].FechaHora)
			//print office and date
			println("Oficina: " + oficina.NombreOficina + " - Fecha: " + date.Format("2006-01-02 15:04:05"))
			bestDateAndOffice = append(bestDateAndOffice, BestOfficeDate{oficina.NombreOficina, date})
			if bestDate.IsZero() || date.Before(bestDate) {
				bestDate = date
				bestOficina = oficina.NombreOficina
			}
		}

	}

	println("===================================== Fechas más cercanas ordenadas. =====================================")
	//sort bestDateAndOffice by date
	sort.Slice(bestDateAndOffice, func(i, j int) bool {
		return bestDateAndOffice[i].Date.Before(bestDateAndOffice[j].Date)
	})
	//print ordered bestDateAndOffice
	for _, bestOfficeDate := range bestDateAndOffice {
		println("Oficina: " + bestOfficeDate.Oficina + " - Fecha: " + bestOfficeDate.Date.Format("2006-01-02 15:04:05"))
	}

	println("La mejor fecha es: " + bestDate.Format("2006-01-02 15:04") + " en la oficina " + bestOficina)
}

func createURL(baseUrl string, tramite string, region int) string {
	var timeString = time.Now().Format("200601021504")
	return baseUrl + strconv.Itoa(region) + "/" + strings.Split(tramite, "-")[0] + "/" + strings.Split(tramite, "-")[1] + "/" + timeString
}

func findOficinasFromRegion() []Oficinas {
	url := createURL("https://agenda.registrocivil.cl/api/backend/oficinas/region/", tramite, region)

	resp, _ := http.Get(url)

	defer resp.Body.Close()

	var cOficinas []Oficinas

	if err := json.NewDecoder(resp.Body).Decode(&cOficinas); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	//traverse cOficinas and save the CodifoOficina
	var oficinas []int
	for _, oficina := range cOficinas {
		oficinas = append(oficinas, oficina.CodigoOficina)
	}
	return cOficinas

}

type Oficinas struct {
	CodigoOficina int         `json:"codigo_oficina"`
	TipoOficina   string      `json:"tipo_oficina"`
	FechaHora     string      `json:"fecha_hora"`
	NombreOficina string      `json:"nombre_oficina"`
	Direccion     string      `json:"direccion"`
	Latitud       interface{} `json:"latitud"`
	Longitud      interface{} `json:"longitud"`
	CodigoComuna  int         `json:"codigo_comuna"`
	NombreComuna  string      `json:"nombre_comuna"`
	NombreRegion  string      `json:"nombre_region"`
}

type AvailableDates struct {
	Code  int `json:"code"`
	Horas []struct {
		CodigoComuna       int    `json:"codigo_comuna"`
		CodigoTipoTramite  int    `json:"codigo_tipo_tramite"`
		CodigoOficinaHoras int    `json:"codigo_oficina_horas"`
		CodigoOficina      int    `json:"codigo_oficina"`
		Fecha              string `json:"fecha"`
		Hora               string `json:"hora"`
		FechaHora          string `json:"fecha_hora"`
		NombreOficina      string `json:"nombre_oficina"`
		Direccion          string `json:"direccion"`
		Cantidad           int    `json:"cantidad"`
	} `json:"horas"`
}

type BestOfficeDate struct {
	Oficina string
	Date    time.Time
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.srcei-date-checker.yaml)")
	rootCmd.PersistentFlags().StringVarP(&tramite, "tramite", "t", "", "")
	rootCmd.PersistentFlags().IntVarP(&region, "region", "r", 0, "Región donde se realizará el trámite.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}
