package handler

var (
	InMemDB            = make(map[string]string)
	KeyError           = "The 'key' is required."
	KeyNotFoundError   = "The key '%s' could not be found."
	ValueError         = "The 'value' is required."
	SetResponsePattern = "The value '%s' is stored with the key '%s'."
	FlushResponse      = "All data has been deleted."
)

// ApiResponse represents the response that the endpoints return.
type ApiResponse struct {
	Error  string `json:"error,omitempty"`
	Result string `json:"result,omitempty"`
}

// LoginRequest represent the request body of set request
type LoginRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
