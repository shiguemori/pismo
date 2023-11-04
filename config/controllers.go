package config

import "pismo/controllers"

func (i *Infra) SetupAccountController() controllers.AccountsController {
	return controllers.NewAccountsController(i.SetupAccountService())
}

func (i *Infra) SetupOperationTypeController() controllers.OperationTypesController {
	return controllers.NewOperationTypeController(i.SetupOperationTypeService())
}

func (i *Infra) SetupTransactionController() controllers.TransactionsController {
	return controllers.NewTransactionsController(i.SetupTransactionsService())
}
