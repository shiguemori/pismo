package utils

import "encoding/json"

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

func ToJSON(o interface{}) string {
	bytes, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
