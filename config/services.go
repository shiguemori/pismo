package config

import "pismo/services"

func (i *Infra) SetupAccountService() services.AccountsService {
	return services.NewAccountsService(i.SetupAccountRepository())
}
