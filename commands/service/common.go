package service

import "github.com/Sirupsen/logrus"
import "github.com/sakuraio/sakuraio-cli/commands"

func checkError(message string, err error, info string) {
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":           err,
			"additional-info": info,
		}).Fatal(message)
	}
}
func GetToken(token string, projectID string, serviceType string) string {
	if token != "" {
		return token
	}
	if projectID == "" {
		logrus.Fatal("You need to specify Token or ProjectID")
	}

	logrus.Warn("Project ID arg don't use production env! If you specified a token, the request will be faster.")
	services := commands.GetServiceFromProject(projectID, serviceType)

	if len(services) == 0 {
		logrus.Fatal("Not found service ServiceType:" + serviceType)
	}
	return services[0].Token
}
