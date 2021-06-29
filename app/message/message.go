package message

type Message struct {
	Identifier string
	Content    interface{}
}

func New(identifer string, content interface{}) *Message {
	return &Message{
		Identifier: identifer,
		Content:    content,
	}
}
