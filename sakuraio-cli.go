package main

import (
	"fmt"
	"os"
	"path"

	"github.com/Sirupsen/logrus"
	colorable "github.com/mattn/go-colorable"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/sakuraio/sakuraio-cli/commands"
	"github.com/sakuraio/sakuraio-cli/commands/service"
	"github.com/sakuraio/sakuraio-cli/lib"
)

var (
	_, commandName = path.Split(os.Args[0])

	app = kingpin.New(commandName, "sakuraio client command")

	appToken  = app.Flag("api-token", "API Token").String()
	appSecret = app.Flag("api-secret", "API Secret").String()

	/////// auth
	authCmd          = app.Command("auth", "Authentication")
	authConfigToken  = authCmd.Arg("token", "API Token").String()
	authConfigSecret = authCmd.Arg("secret", "API Secret").String()

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
	addServiceCmd      = servicesCmd.Command("add", "Add Service")
	addServiceType     = addServiceCmd.Arg("type", "Service type").Required().String()
	addServiceProject  = addServiceCmd.Arg("project id", "Project ID").Required().Int()
	addServiceOptions  = addServiceCmd.Arg("option", "Service option").Strings()
)

/////// data store
var (
	dataStoreCmd         = servicesCmd.Command("datastore", "Data Store Service")
	DataStoreChannelsCmd = dataStoreCmd.Command("channels", "Get channel data")
	DataStoreMessagesCmd = dataStoreCmd.Command("messages", "Get message data")
)
var dataStoreChannelOption = service.DataStoreChannelOptions{
	Module:    DataStoreChannelsCmd.Flag("module", "Module ID").Short('m').Default("").String(),
	Size:      DataStoreChannelsCmd.Flag("size", "Fetch Size").Short('s').Default("100").String(),
	Order:     DataStoreChannelsCmd.Flag("order", "Order asc/desc").String(),
	Token:     DataStoreChannelsCmd.Flag("token", "Service Token").String(),
	Cursor:    DataStoreChannelsCmd.Flag("cursor", "Cursor").String(),
	After:     DataStoreChannelsCmd.Flag("after", "Datetime range from").String(),
	Before:    DataStoreChannelsCmd.Flag("before", "Datetime range to").String(),
	Channel:   DataStoreChannelsCmd.Flag("channel", "Channel").String(),
	Project:   DataStoreChannelsCmd.Flag("project", "Project ID").String(),
	RawOutput: DataStoreChannelsCmd.Flag("raw", "Raw JSON output").Default("false").Bool(),
	Recursive: DataStoreChannelsCmd.Flag("recursive", "Collect recursive").Short('r').Default("false").Bool(),
	MaxReq:    DataStoreChannelsCmd.Flag("max-req", "Max request count in --recursive").Default("100").Int(),
}
var dataStoreMessageOption = service.DataStoreMessagesOption{
	Module:    DataStoreMessagesCmd.Flag("module", "Module ID").Short('m').Default("").String(),
	Size:      DataStoreMessagesCmd.Flag("size", "Fetch Size").Short('s').Default("100").String(),
	Order:     DataStoreMessagesCmd.Flag("order", "Order asc/desc").String(),
	Cursor:    DataStoreMessagesCmd.Flag("cursor", "Cursor").String(),
	After:     DataStoreMessagesCmd.Flag("after", "Datetime range from").String(),
	Before:    DataStoreMessagesCmd.Flag("before", "Datetime range to").String(),
	Project:   DataStoreMessagesCmd.Flag("project", "Project ID").String(),
	RawOutput: DataStoreMessagesCmd.Flag("raw", "Raw JSON output").Default("false").Bool(),
	Token:     DataStoreMessagesCmd.Flag("token", "Service Token").String(),
	Recursive: DataStoreMessagesCmd.Flag("recursive", "Collect recursive").Short('r').Default("false").Bool(),
	MaxReq:    DataStoreMessagesCmd.Flag("max-req", "Max request count in --recursive").Default("100").Int(),
}

var (
	websocketCmd       = servicesCmd.Command("websocket", "Websocket Service").Alias("ws")
	websocketListenCmd = websocketCmd.Command("listen", "Listen to Websocket")
)
var websocketListneOption = service.WebsocketListenOptions{
	Project: websocketListenCmd.Flag("project", "Project ID").String(),
	// RawOutput: websocketListenCmd.Flag("raw", "Raw JSON output").Default("false").Bool(),
	Token: websocketListenCmd.Flag("token", "Service Token").String(),
}

/////// END Flags

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
}

func main() {

	app.Version(commands.Version)
	app.UsageTemplate(kingpin.CompactUsageTemplate)
	app.VersionFlag.Action(func(c *kingpin.ParseContext) error {
		fmt.Println("PreAction")
		fmt.Println(*c)

		return nil
	})
	parseResult := kingpin.MustParse(app.Parse(os.Args[1:]))
	lib.OverrideSettings.APIToken = *appToken
	lib.OverrideSettings.APISecret = *appSecret

	switch parseResult {
	case authCmd.FullCommand(): // Auth Command
		commands.AuthConfigCommand(*authConfigToken, *authConfigSecret)

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
	case addServiceCmd.FullCommand():
		commands.AddServiceCommand(*addServiceType, *addServiceProject, *addServiceOptions)

	case DataStoreChannelsCmd.FullCommand(): // Service Data Store Command
		service.DataStoreChannelsCommand(dataStoreChannelOption)
	case DataStoreMessagesCmd.FullCommand():
		service.DataStoreMessagesCmd(dataStoreMessageOption)

	case websocketListenCmd.FullCommand(): // Service Websocket Command
		service.WebsocketListenCommand(websocketListneOption)
	}
}
