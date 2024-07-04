package piscine

func ListPushFront(l *List, data interface{}) {
	if l.Head == nil {
		newNode := &NodeL{Data: data}

		l.Head = newNode
		return
	}
	newNode := &NodeL{Data: data, Next: l.Head}
	l.Head = newNode
}
