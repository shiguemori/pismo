package utils

const (
	IDCannotBeEmpty         = "ID cannot be empty"
	ErrorConvertingIDToUint = "Error converting ID to uint"
	NotFound                = "Not found"
	ErrorBindingJSON        = "Error binding JSON"
	DeletedSuccessfully     = "deleted successfully"
)

// Error message response
// swagger:model utils.Response
type Response struct {
	Message string `json:"message"`
}
