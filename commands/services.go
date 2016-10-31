package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"../lib"
)

func fetchServices() []Service {
	var services []Service
	body, err := lib.HTTPGet("v1/services/")
	checkError("[ERROR]", err, "")
	err = json.Unmarshal([]byte(body), &services)

	checkError("[JSON format error]", err, body)
	return services
}

func ListServicesCommand() {
	printService(fetchServices())
}

func ListServiceFilterProjectCommand(project string) {
	services := fetchServices()

	var inProject = []Service{}
	for _, v := range services {
		if strconv.Itoa(v.Project) == project {
			inProject = append(inProject, v)
		}
	}

	printService(inProject)
}

func ShowServicesCommand(projectIDs []string) {
	var services = []Service{}
	for _, id := range projectIDs {
		var service Service
		body, err := lib.HTTPGet("v1/services/" + id + "/")
		checkError("HTTP ERROR", err, body)

		err = json.Unmarshal([]byte(body), &service)
		checkError("JSON format error", err, body)

		services = append(services, service)
	}
	printService(services)
}

func DeleteServiceCommand(forceRemove bool, deleteServicdID string) {
	if forceRemove == false {
		ShowServicesCommand([]string{deleteServicdID})
		if lib.YesOrNo("Remove?") == false {
			fmt.Println("Abort...")
			os.Exit(1)
		}
	}
	res, err := lib.HTTPDelete("v1/services/" + deleteServicdID + "/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
	fmt.Println("Success")
}

func AddService(request []byte) {
	var service = Service{}
	body, err := lib.HTTPPost("v1/services/", string(request))
	checkError("HTTP ERROR", err, body)

	err = json.Unmarshal([]byte(body), &service)
	checkError("JSON format error", err, body)

	printService([]Service{service})
}

func printService(services []Service) {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
	fmt.Fprintln(w, "ID\tProject\tType\tToken\tServiceName")
	for _, v := range services {
		fmt.Fprintf(w, "%d\t%d\t%s\t%s\t%s\n", v.ID, v.Project, v.Type, v.Token, v.Name)
	}
	w.Flush()
}

type Service struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Project int    `json:"project"`
	Token   string `json:"token"`
}
