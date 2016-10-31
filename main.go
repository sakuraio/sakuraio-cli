package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	colorable "github.com/mattn/go-colorable"

	"./commands"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("sakuraio", "sakuraio client command")

	/////// auth
	authCmd          = app.Command("auth", "Authentication")
	authConfigToken  = authCmd.Arg("token", "API Token").String()
	authConfigSecret = authCmd.Arg("secret", "API Secret").String()

	/////// data store
	dataStoreCmd  = app.Command("datastore", "Data Store Service")
	dataStoreSize = dataStoreCmd.Flag("size", "Fetch Size").Short('s').Default("100").Int()

	/////// project
	projectsCmd        = app.Command("project", "Project management")
	listProjectsCmd    = projectsCmd.Command("list", "List of projects").Alias("ls")
	showProjectsCmd    = projectsCmd.Command("show", "Lookup projects")
	showProjectIDs     = showProjectsCmd.Arg("id", "Show projectID").Required().Strings()
	addProjectCmd      = projectsCmd.Command("add", "Add Project")
	addProjectName     = addProjectCmd.Arg("name", "project name").Required().String()
	deleteProjectCmd   = projectsCmd.Command("remove", "Add Project").Alias("delete").Alias("rm")
	deleteProjectForce = deleteProjectCmd.Flag("force", "Project force remove").Short('f').Bool()
	deleteProjectID    = deleteProjectCmd.Arg("ID", "Remove project ID").Required().String()

	/////// module
	modulesCmd       = app.Command("module", "Modules managemet")
	listModulesCmd   = modulesCmd.Command("list", "List of modules").Alias("ls")
	showModulesCmd   = modulesCmd.Command("show", "Lookup modules")
	showModuleIDs    = showModulesCmd.Arg("ID", "Modlue ID list").Required().Strings()
	addModuleCmd     = modulesCmd.Command("add", "Add module")
	addModuleID      = addModuleCmd.Arg("id", "Register ID").Required().String()
	addModulePW      = addModuleCmd.Arg("password", "Register Password").Required().String()
	addModuleProject = addModuleCmd.Arg("project-id", "Project ID").Required().Int()
	addModuleName    = addModuleCmd.Arg("name", "Module name").String()

	// service
	servicesCmd     = app.Command("service", "Sservice management")
	listServicesCmd = servicesCmd.Command("list", "List of services").Alias("ls")
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
}

func main() {

	app.UsageTemplate(kingpin.CompactUsageTemplate)
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case authCmd.FullCommand(): // Auth Command
		commands.AuthConfigCommand(*authConfigToken, *authConfigSecret)

	case dataStoreCmd.FullCommand(): // Service Data Store Command
		commands.DataStoreCommand(*dataStoreSize)

	case listProjectsCmd.FullCommand(): // Project Command
		commands.ListProjectsCommand()
	case showProjectsCmd.FullCommand():
		commands.ShowProjectsCommand(*showProjectIDs)
	case addProjectCmd.FullCommand():
		commands.AddProjectCommand(*addProjectName)
	case deleteProjectCmd.FullCommand():
		commands.DeleteProject(*deleteProjectForce, *deleteProjectID)

	case listModulesCmd.FullCommand(): // Module Command
		commands.ListModulesCommand()
	case showModulesCmd.FullCommand():
		commands.ShowModulesCommand(*showModuleIDs)
	case addModuleCmd.FullCommand():
		commands.AddModuleCommand(*addModuleID, *addModulePW, *addModuleProject, *addModuleName)
	case listServicesCmd.FullCommand():
		commands.ListServicesCommand()
	}
}
