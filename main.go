/*

WEB SERVER e APP SERVER per la app di scrittura metadati di ARAMCO

[2016-11-08] Max

*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var version = "1.0.0"
var start = time.Now()

var progressiveId = 10000

func main() {

	//setting up the log:
	log.Printf("SETTING LOG 'web-server-max.log' ...\n")
	f, err := os.OpenFile("web-server-max.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("Error opening file: %v", err)
	}
	defer f.Close()
	////to log to file only:
	//log.SetOutput(f)
	//to log to stdout AND file:
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	//log set up completed.

	log.Printf("Version %s\n", version)

	log.Println("Reading config ...")
	ReadConfig() //read into var conf
	log.Println("Done with config")

	router := mux.NewRouter().StrictSlash(true)

	if len(conf.RootFolder) == 0 {
		conf.RootFolder = "./html/"
		log.Printf("Config 'root_folder' is empty: forced to '%s'", conf.RootFolder)
	}

	//qui servo pagine statiche:
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(conf.RootFolder))))

	serverPortDef := fmt.Sprintf(":%d", conf.ServerPort)
	log.Printf("Serving: '%s' in '%s", serverPortDef, conf.RootFolder)

	//WITHOUT CORS:
	//log.Fatal(http.ListenAndServe(serverPortDef, router))
	//WITH CORS:
	//log.Fatal(http.ListenAndServe(serverPortDef, handlers.CORS()(router)))

	loggedRouter := handlers.LoggingHandler(mw, router)
	http.ListenAndServe(serverPortDef, handlers.CORS()(loggedRouter))
}
