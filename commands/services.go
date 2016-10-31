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
	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}
	err = json.Unmarshal([]byte(body), &services)

	if err != nil {
		fmt.Println("[JSON format error] ", err)
		fmt.Println(body)
		os.Exit(1)
	}
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

// func ShowServiceCommand(projectIDs []string) {
// 	var projects []Project
// 	for _, id := range projectIDs {
// 		var project Project
// 		body, err := lib.HTTPGet("v1/projects/" + id + "/")
// 		checkError("HTTP ERROR", err, body)

// 		err = json.Unmarshal([]byte(body), &project)
// 		checkError("JSON format error", err, body)

// 		projects = append(projects, project)
// 	}
// 	printProjects(projects)
// }

// func AddServiceCommand(projectName string) {
// 	project := Project{
// 		ID:   0,
// 		Name: projectName,
// 	}
// 	reqJSON, _ := json.Marshal(project)

// 	res, err := lib.HTTPPost("v1/projects/", string(reqJSON))
// 	if err != nil {
// 		fmt.Println("[ERROR]", err)
// 		os.Exit(1)
// 	}

// 	err = json.Unmarshal([]byte(res), &project)
// 	if err != nil {
// 		fmt.Println("[JSON format error]", err)
// 		os.Exit(1)
// 	}

// 	printProjects([]Project{project})
// }

// func DeleteService(forceRemove bool, deleteProjectID string) {
// 	if forceRemove == false {
// 		ShowProjectsCommand([]string{deleteProjectID})
// 		reader := bufio.NewReader(os.Stdin)
// 		fmt.Print("Remove? [yN]")
// 		str, _ := reader.ReadString('\n')
// 		lower := strings.ToLower(str)
// 		if strings.HasPrefix(lower, "y") == false {
// 			fmt.Println("Abort...")
// 			os.Exit(1)
// 		}
// 	}
// 	res, err := lib.HTTPDelete("v1/projects/" + deleteProjectID + "/")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(res)
// 	fmt.Println("Success")
// }

func printService(services []Service) {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
	fmt.Fprintln(w, "ID\tProject\tType\tToken\tName")
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
