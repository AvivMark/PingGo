package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	hostModel "github.com/AvivMark/PingGo/models/host"
	routes "github.com/AvivMark/PingGo/routes"
)

var PORT string = "5000"        // APP PORT
var JsonFilePath = "hosts.json" // JSON file path
var DebugMode = false
var Hosts []hostModel.Host
// ///////// UI ROUTES FUNCTIONS
func homePage(w http.ResponseWriter, r *http.Request) {
}

// ///////// Utils functions

// Function creates 100 demo servers for tests only
func Get100Servers() []hostModel.Host {
	hostsOverload := []hostModel.Host{}
	for i := 1; i < 121; i++ {
		numStr := strconv.Itoa(i)
		name := "server-" + numStr
		groupNum := strconv.Itoa(i / 12)
		newHost := hostModel.Host{
			ID:       numStr,
			Hostname: name,
			HostIP:   "172.17.17." + numStr,
			IsAlive:  false,
			Group:    "Group" + groupNum,
			PingData: "",
		}
		hostsOverload = append(hostsOverload, newHost)
	}
	return hostsOverload
}

// ROUTE TO  RELOAD HOSTS FROM JSON FILE
func refresh(w http.ResponseWriter, r *http.Request) {
	Hosts = hostModel.GetHostsFromJson(JsonFilePath)
	log.Printf("EndpointHit: Refreshed app!")
}
// Routes main Declaration function
func handleRequest() {
	r := mux.NewRouter()

	// Util routes
	r.HandleFunc("/refresh", refresh)

	// Host routes
	r.HandleFunc("/host", routes.CreateHost).Methods("POST")
	r.HandleFunc("/hostUpdate", routes.UpdateHost).Methods("PUT")
	r.HandleFunc("/host/{ID}", routes.DeleteHost).Methods("DELETE")
	r.HandleFunc("/host/{ID}", routes.GetHost)
	r.HandleFunc("/hostAvailable/{ID}", routes.GetHostWithPing)

	// Hosts routes
	r.HandleFunc("/hosts", routes.ReturnAllHosts)
	r.HandleFunc("/hostsAvailable", routes.ReturnAllHostsWithPing)

	// Groups Routes
	r.HandleFunc("/getGroupHosts/{GroupName}", routes.GetGroupHosts)
	r.HandleFunc("/getGroupAvailable/{GroupName}", routes.GetGroupAvailable)
	r.HandleFunc("/getGroups", routes.GetGroups)

	//UI Routes
	// Set up a route for the reverse proxy.
	r.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":"+PORT, r))
}

// //////////////////////////////////////////////////////////////

// MAIN FUNCTION
func main() {

	// Create Test Mode in run
	if len(os.Args) > 1 {
		debugArg := strings.ToLower(os.Args[1])
		if debugArg == "test" {
			Hosts = Get100Servers()
			DebugMode = true
		}
	} else {
		Hosts = hostModel.GetHostsFromJson(JsonFilePath)
	}

	if DebugMode {
		fmt.Println("--------------Debug MODE----------------------")
	}
	fmt.Println("PingGo - The watchdog for your servers")
	fmt.Println("API CREATED BY AVIV MARK RUNNING ON PORT:" + PORT)
	fmt.Println("----------------------------------------------")

	handleRequest()
}
