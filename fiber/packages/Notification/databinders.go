package notification

// CreateDatabinder Define the NotificationCreate parameters
type CreateDatabinder struct {
	Title          string
	TitleVarName   string
	TitleIsVar     bool
	Message        string
	MessageVarName string
	MessageIsVar   bool
}
