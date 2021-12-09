package utils

var (
	InMemStore       = make(map[string]string)
	KeyRequired      = "Key field is required"
	KeyNotFound      = "Key '%s' not found"
	FlushResponse    = "All datas has been deleted"
	ValueRequired    = "'Value' is required"
	SetResponseError = "Value '%s' is stored to '%s'"
)