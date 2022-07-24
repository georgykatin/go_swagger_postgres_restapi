package main

import (
	"fmt"
	"log"
	_ "net"
	"net/http"
	"os"
	"practiceproject/controllers"
	"practiceproject/database"
	_ "practiceproject/docs"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

//@title CRUD API
//@version 1.0
//@description API server

// @contact.name API Support
// @contact.url http://demo.com/support

//@host 127.0.0.1:8080
//@BasePath /

var DB *gorm.DB
var AppConfig *Config

type Config struct {
	ConnectionString string `json:"connection_string"`
	Port             int    `json:"port"`
}

func LoadAppConfig() {
	log.Println("Loading server configurations")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
func setupRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/person/getpersons", controllers.GetPersons).Methods(http.MethodGet)
	r.HandleFunc("/person/getpersonbyid/{id}", controllers.GetPersonById).Methods(http.MethodGet)
	r.HandleFunc("/person/create", controllers.CreatePerson).Methods(http.MethodPost)
	r.HandleFunc("/person/update", controllers.Update).Methods(http.MethodPatch)
	r.HandleFunc("/person/delete", controllers.DeletePerson).Methods(http.MethodDelete)
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return r
}

func main() {
	r := setupRouter()

	fmt.Fprintf(os.Stderr, "hfhdhdhhh\n")
	database.Connect("")
	fmt.Fprintf(os.Stderr, "da4w4ga4ashh\n")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

//	/*_ = r.Run(":8000")
//	log.Println("router")
//
//	LoadAppConfig()
//*/
//	database.Connect(AppConfig.ConnectionString)
//
//	router := mux.NewRouter().StrictSlash(true)
//
//	x := fmt.Sprintf("Starting server on port %s", AppConfig.Port)
//	log.Println(x)
//	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
//*/
//	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	api := operations.NewExample1API(swaggerSpec)
//	server := restapi.NewServer(api)
//	defer server.Shutdown()
//
//	parser := flags.NewParser(server, flags.Default)
//	parser.ShortDescription = "Example service"
//	parser.LongDescription = "Example service"
//	server.ConfigureFlags()
//	for _, optsGroup := range api.CommandLineOptionsGroups {
//		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
//		if err != nil {
//			log.Fatalln(err)
//		}
//	}
//
//	if _, err := parser.Parse(); err != nil {
//		code := 1
//		if fe, ok := err.(*flags.Error); ok {
//			if fe.Type == flags.ErrHelp {
//				code = 0
//			}
//		}
//		os.Exit(code)
//	}
//
//	server.ConfigureAPI()
//
//	if err := server.Serve(); err != nil {
//		log.Fatalln(err)
//	}
//
//}
