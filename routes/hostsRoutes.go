package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	hostModel "github.com/AvivMark/PingGo/models/host"

)

var Hosts []hostModel.Host //list of the hosts

func ReturnAllHostsWithPing(w http.ResponseWriter, r *http.Request) {
	log.Printf("EndpointHit: returnAllHostsWithPing")
	var wg sync.WaitGroup
	for i := range Hosts {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			hostModel.PingHost(&Hosts[i])
		}(i)
	}
	wg.Wait()
	log.Printf("EndpointHit: returned Data!")
	json.NewEncoder(w).Encode(Hosts)
}

func ReturnAllHosts(w http.ResponseWriter, r *http.Request) {
	log.Printf("EndpointHit: returnAllHosts")
	json.NewEncoder(w).Encode(Hosts)
}

func GetHostWithPing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ID"]
	hostFound := hostModel.Host{}
	for _, host := range Hosts {
		if host.ID == key {
			hostFound = host
		}
		if host.HostIP == key {
			hostFound = host
		}
		if host.Hostname == key {
			hostFound = host
		}
	}
	if hostFound.HostIP != "" {
		hostModel.PingHost(&hostFound)
		log.Printf("EndpointHit: getHostWithPing for host %s", hostFound.Hostname)
		json.NewEncoder(w).Encode(hostFound)
	} else {
		w.WriteHeader(404)
	}

}

// Host List Edit Functions
func GetHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ID"]
	hostFound := hostModel.Host{}
	for _, host := range Hosts {
		if host.ID == key {
			hostFound = host
		}
		if host.HostIP == key {
			hostFound = host
		}
		if host.Hostname == key {
			hostFound = host
		}
	}
	if hostFound.HostIP != "" {
		log.Printf("EndpointHit: getHost for host %s", hostFound.Hostname)
		json.NewEncoder(w).Encode(hostFound)
	} else {
		w.WriteHeader(404)
	}

}

func CreateHost(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var host hostModel.Host
	json.Unmarshal(reqBody, &host)
	Hosts = append(Hosts, host)
	log.Printf("EndpointHit: added Host!")
	json.NewEncoder(w).Encode(host)
}

func DeleteHost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["ID"]

	for i, host := range Hosts {
		if host.ID == key {
			Hosts = append(Hosts[:i], Hosts[i+1:]...)
		}
	}
	log.Printf("EndpointHit: deleted Host!")
}

func UpdateHost(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var hostToUpdate hostModel.Host
	json.Unmarshal(reqBody, &hostToUpdate)

	for i, host := range Hosts {
		if host.HostIP == hostToUpdate.HostIP {
			Hosts[i].Hostname = hostToUpdate.Hostname
			Hosts[i].IsAlive = hostToUpdate.IsAlive
			Hosts[i].Group = hostToUpdate.Group
		}
		if host.Hostname == hostToUpdate.Hostname {
			Hosts[i].HostIP = hostToUpdate.HostIP
			Hosts[i].IsAlive = hostToUpdate.IsAlive
			Hosts[i].Group = hostToUpdate.Group
		}
		if host.ID == hostToUpdate.ID {
			Hosts[i].HostIP = hostToUpdate.HostIP
			Hosts[i].Hostname = hostToUpdate.Hostname
			Hosts[i].IsAlive = hostToUpdate.IsAlive
			Hosts[i].Group = hostToUpdate.Group

		}
	}

	log.Printf("EndpointHit: Updated Host!")
	json.NewEncoder(w).Encode(Hosts)

}
