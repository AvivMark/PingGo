package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	hostModel "github.com/AvivMark/PingGo/models/host"
	"github.com/AvivMark/PingGo/utils"
)

// Groups Functions
func GetGroupsList() []string {
	var Groups []string
	for _, host := range Hosts {
		inSlice := utils.Contains(Groups, host.Group)
		if inSlice == false {
			Groups = append(Groups, host.Group)
		}
	}
	return Groups
}

func FindGroupHosts(groupName string) []hostModel.Host {
	var GroupHosts []hostModel.Host = []hostModel.Host{}

	for _, host := range Hosts {
		if host.Group == groupName {

			GroupHosts = append(GroupHosts, host)
		}
	}
	return GroupHosts
}

func GetGroups(w http.ResponseWriter, r *http.Request) {
	GroupsList := GetGroupsList()
	log.Printf("EndpointHit: getGroups")
	json.NewEncoder(w).Encode(GroupsList)
}

func GetGroupHosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["GroupName"]

	groupsHosts := FindGroupHosts(key)
	log.Printf("EndpointHit: getGroupHosts for group: " + key)
	json.NewEncoder(w).Encode(groupsHosts)
}

func GetGroupAvailable(w http.ResponseWriter, r *http.Request) {
	log.Printf("EndpointHit: getGroupAvailable")
	vars := mux.Vars(r)
	key := vars["GroupName"]
	groupHosts := FindGroupHosts(key)
	var wg sync.WaitGroup
	wg.Wait()
	for i := range groupHosts {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			hostModel.PingHost(&groupHosts[i])
		}(i)
	}
	log.Printf("Got Ping data for group")
	json.NewEncoder(w).Encode(groupHosts)
}
