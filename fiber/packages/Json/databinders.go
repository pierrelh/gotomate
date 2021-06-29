package json

// CreateJsonDatabinder Define the CreateJson parameters
type CreateJsonDatabinder struct {
	Path        string
	PathVarName string
	PathIsVar   bool
	Output      string
}

// JsonToDictionaryDatabinder Define the JsonToDictionary parameters
type JsonToDictionaryDatabinder struct {
	JsonVarName string
	Output      string
}

// SaveJsonDatabinder Define the SaveJson parameters
type SaveJsonDatabinder struct {
	JsonVarName string
	Path        string
	PathVarName string
	PathIsVar   bool
}
