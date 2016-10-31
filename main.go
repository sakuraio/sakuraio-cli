package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	colorable "github.com/mattn/go-colorable"
	"gopkg.in/alecthomas/kingpin.v2"

	"./commands"
	"./commands/service"
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
	deleteProjectCmd   = projectsCmd.Command("remove", "Remove Project").Alias("delete").Alias("rm")
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

	/////// service
	servicesCmd        = app.Command("service", "Service management")
	listServicesCmd    = servicesCmd.Command("list", "List of services").Alias("ls")
	showServiceCmd     = servicesCmd.Command("show", "show service")
	showServiceIDs     = showServiceCmd.Arg("id", "Show projectID").Required().Strings()
	deleteServiceCmd   = servicesCmd.Command("remove", "Remove Project").Alias("delete").Alias("rm")
	deleteServiceForce = deleteServiceCmd.Flag("force", "Force remove").Short('f').Bool()
	deleteServiceID    = deleteServiceCmd.Arg("ID", "Remove service ID").Required().String()

	// outgoing webhook
	outgoingWebhookCmd        = servicesCmd.Command("outgoingwebhook", "Service Outgoing Webhook")
	addOutgoingWebhookCmd     = outgoingWebhookCmd.Command("add", "Add service")
	addOutgoingWebhookProject = addOutgoingWebhookCmd.Arg("project", "Project ID").Required().Int()
	addOutgoingWebhookName    = addOutgoingWebhookCmd.Arg("name", "Service Name").Required().String()
	addOutgoingWebhookURL     = addOutgoingWebhookCmd.Arg("url", "Dest URL").Required().String()
	addOutgoingWebhookSecret  = addOutgoingWebhookCmd.Arg("secret", "secret").Required().String()
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

	case listServicesCmd.FullCommand(): // Service Command
		commands.ListServicesCommand()
	case showServiceCmd.FullCommand():
		commands.ShowServicesCommand(*showServiceIDs)
	case deleteServiceCmd.FullCommand():
		commands.DeleteServiceCommand(*deleteServiceForce, *deleteServiceID)

	case addOutgoingWebhookCmd.FullCommand(): // Outgoin Webhook
		service.OutgoingWebhookAddCommand(*addOutgoingWebhookProject, *addOutgoingWebhookName, *addOutgoingWebhookURL, *addOutgoingWebhookSecret)

	}
}
