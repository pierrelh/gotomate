package dictionary

// CreateDictionaryDatabinder Define the CreateDictionary parameters
type CreateDictionaryDatabinder struct {
	Output string
}

// CreateEntryDatabinder Define the CreateEntry parameters
type CreateEntryDatabinder struct {
	DictVarName  string
	Key          string
	KeyVarName   string
	KeyIsVar     bool
	Value        string
	ValueVarName string
	ValueIsVar   bool
}

// DictionaryToJsonDatabinder Define the DictionaryToJson parameters
type DictionaryToJsonDatabinder struct {
	DictVarName string
	Output      string
}

// RemoveEntryDatabinder Define the RemoveEntry parameters
type RemoveEntryDatabinder struct {
	DictVarName string
	Key         string
	KeyVarName  string
	KeyIsVar    bool
}
