package config

import "pismo/repositories"

func (i *Infra) SetupAccountRepository() repositories.AccountsRepository {
	return repositories.NewAccountsRepository(i.DB)
}
