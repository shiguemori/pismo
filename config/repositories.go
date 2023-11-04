package config

import "pismo/repositories"

func (i *Infra) SetupAccountRepository() repositories.AccountsRepository {
	return repositories.NewAccountsRepository(i.DB)
}

func (i *Infra) SetupOperationTypeRepository() repositories.OperationTypesRepository {
	return repositories.NewOperationTypesRepository(i.DB)
}

func (i *Infra) SetupTransactionRepository() repositories.TransactionsRepository {
	return repositories.NewTransactionsRepository(i.DB)
}
