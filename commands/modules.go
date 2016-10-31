package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Sirupsen/logrus"

	"../lib"
)

func checkError(message string, err error, info string) {
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":           err,
			"additional-info": info,
		}).Fatal(message)
	}
}

func ListModulesCommand() {
	var modules []Module
	body, err := lib.HTTPGet("v1/modules/")
	checkError("HTTP ERROR", err, body)

	err = json.Unmarshal([]byte(body), &modules)
	checkError("JSON format error", err, body)

	printModules(modules)
}

func ShowModulesCommand(moduleIds []string) {
	var modules []Module
	for _, id := range moduleIds {
		var module Module
		body, err := lib.HTTPGet("v1/modules/" + id + "/")
		checkError("HTTP ERROR", err, body)

		err = json.Unmarshal([]byte(body), &module)
		checkError("JSON format error", err, body)

		modules = append(modules, module)
	}
	printModules(modules)
}

func AddModuleCommand(id string, password string, projetId int, name string) {
	registerInfo := ModuleRegisterInfo{
		ID:       id,
		Password: password,
		Project:  projetId,
		Name:     name,
	}
	var module Module

	reqJSON, _ := json.Marshal(registerInfo)

	body, err := lib.HTTPPost("v1/modules/", string(reqJSON))
	checkError("HTTP ERROR", err, body)

	err = json.Unmarshal([]byte(body), &module)
	checkError("JSON format error", err, body)

	printModules([]Module{module})
}

func printModules(modules []Module) {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 4, ' ', 0)
	fmt.Fprintln(w, "ID\tProject\tOnline\tName")
	for _, v := range modules {
		fmt.Fprintf(w, "%s\t%d\t%t\t%s\n", v.ID, v.Project, v.IsOnline, v.Name)
	}
	w.Flush()
}

type Module struct {
	ID       string
	Name     string
	Project  int
	IsOnline bool `json:"is_online"`
}

type ModuleRegisterInfo struct {
	ID       string `json:"register_id"`
	Password string `json:"register_password"`
	Name     string `json:"name"`
	Project  int    `json:"project"`
}
