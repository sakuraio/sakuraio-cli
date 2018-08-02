package commands

import (
	"fmt"
	"os"

	"github.com/sakuraio/sakuraio-cli/lib"
)

func AuthConfigCommand(token string, secret string) {
	if len(token) == 0 && len(secret) == 0 {
		setting := lib.GetSetting()
		fmt.Printf("[Current API Key] token: %s  secret: ***************\n", setting.APIToken)
		return
	}
	if len(secret) == 0 {
		fmt.Printf("Need 'secret' arg")
		os.Exit(1)
		return
	}

	userSetting, _ := lib.GetUserSetting()
	userSetting.APIToken = token
	userSetting.APISecret = secret

	err := lib.WriteSetting(userSetting)
	if err != nil {
		fmt.Println("Config write error")
		os.Exit(1)
		return
	}

	fmt.Printf("[Set Token] token: %s  secret: ***************\n", token)
}
