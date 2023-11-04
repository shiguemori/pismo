package config

import "pismo/controllers"

func (i *Infra) SetupAccountController() controllers.AccountsController {
	return controllers.NewAccountsController(i.SetupAccountService())
}
