/*

WEB SERVER che utilizza gorilla

[2023-11-29] Max

*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/maxstrambini/goutils"
)

var version = "1.2.0"
var start = time.Now()

var progressiveId = 10000

var appName = "webservermax"
var maxLogSizeMB = 10
var maxLogRotations = 100

func main() {

	//setting up the log:
	//#fix_path_LINUX_Windows:
	logname := ""
	if runtime.GOOS == "windows" {
		fmt.Println("WINDOWS OS")
		logname = fmt.Sprintf(".\\log\\%s.log", appName)
	} else {
		fmt.Println("NON WINDOWS OS")
		logname = fmt.Sprintf("./log/%s.log", appName)
	}

	fmt.Printf("SETTING LOG '%s' ...\n", logname)

	//using max rotating log to write only to files with rotation
	mw := goutils.NewMaxRotateWriter(logname, maxLogSizeMB*1024*1024, true, maxLogRotations) //filename string, maxBytes int, rotateFilesByNumber bool, maxRotatedFilesByNumber int
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
	} else {
		log.Printf("Config 'root_folder': '%s'", conf.RootFolder)
	}

	//qui servo pagine statiche:
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(conf.RootFolder))))
	//router.PathPrefix("/html/").Handler(http.StripPrefix("/html/", http.FileServer(http.Dir("./html/"))))

	log.Printf("Config 'server_port': '%v'", conf.ServerPort)

	serverPortDef := fmt.Sprintf(":%d", conf.ServerPort)
	log.Printf("Serving on port '%v' from root: '%s' \n(for example call 'http://localhost%s/index.html' to get './html/index.html')", serverPortDef, conf.RootFolder, serverPortDef)

	//WITHOUT CORS:
	//log.Fatal(http.ListenAndServe(serverPortDef, router))
	//WITH CORS:
	//log.Fatal(http.ListenAndServe(serverPortDef, handlers.CORS()(router)))

	loggedRouter := handlers.LoggingHandler(mw, router)
	err := http.ListenAndServe(serverPortDef, handlers.CORS()(loggedRouter))
	if err != nil {
		log.Printf("ListenAndServe has returned error: %v", err)
	}
	log.Printf("Server now exit!")
}
