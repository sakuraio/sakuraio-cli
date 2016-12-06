package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kamijin-fanta/sakuraio-cli/lib"
)

func ListProjectsCommand() {

	var projects []Project
	body, err := lib.HTTPGet("v1/projects/")
	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}
	err = json.Unmarshal([]byte(body), &projects)

	if err != nil {
		fmt.Println("[JSON format error] ", err)
		fmt.Println(body)
		os.Exit(1)
	}
	printProjects(projects)
}

func ShowProjectsCommand(projectIDs []string) {
	for _, id := range projectIDs {
		var project Project
		body, err := lib.HTTPGet("v1/projects/" + id + "/")
		checkError("HTTP ERROR", err, body)

		err = json.Unmarshal([]byte(body), &project)
		checkError("JSON format error", err, body)

		printProjects([]Project{project})
		fmt.Print("\n")
		ListServiceFilterProjectCommand(id)
		fmt.Print("\n")
		ListModulesFilterProjectCommand(id)
	}

}

func AddProjectCommand(projectName string) {
	project := Project{
		ID:   0,
		Name: projectName,
	}
	reqJSON, _ := json.Marshal(project)

	res, err := lib.HTTPPost("v1/projects/", string(reqJSON))
	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(res), &project)
	if err != nil {
		fmt.Println("[JSON format error]", err)
		os.Exit(1)
	}

	printProjects([]Project{project})
}

func DeleteProject(forceRemove bool, deleteProjectID string) {
	if forceRemove == false {
		ShowProjectsCommand([]string{deleteProjectID})
		if lib.YesOrNo("Remove?") == false {
			fmt.Println("Abort...")
			os.Exit(1)
		}
	}
	res, err := lib.HTTPDelete("v1/projects/" + deleteProjectID + "/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
	fmt.Println("Success")
}

func printProjects(projects []Project) {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
	fmt.Fprintln(w, "ID\tProjectName")
	for _, v := range projects {
		fmt.Fprintf(w, "%d\t%s\n", v.ID, v.Name)
	}
	w.Flush()
}

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
