package config

import "pismo/services"

func (i *Infra) SetupAccountService() services.AccountsService {
	return services.NewAccountsService(i.SetupAccountRepository())
}

func (i *Infra) SetupOperationTypeService() services.OperationTypesService {
	return services.NewOperationTypesService(i.SetupOperationTypeRepository())
}

func (i *Infra) SetupTransactionsService() services.TransactionsService {
	return services.NewTransactionsService(i.SetupTransactionRepository())
}
