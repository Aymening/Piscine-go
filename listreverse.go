package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}

	tmp := l.Head
	cur := l.Head
	prv := l.Head
	prv = nil

	for cur != nil {
		next := cur.Next
		cur.Next = prv
		prv = cur
		cur = next
	}
	l.Head = prv
	l.Tail = tmp
	l.Tail.Next = nil
}
