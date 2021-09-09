package iterator

type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() interface{}
}

type ArrayInt []int

type ArrayListIterator struct {
	arrayList ArrayInt
	index     int
}

func (a ArrayInt) Iterator() Iterator {
	return &ArrayListIterator{arrayList: a, index: 0}
}

func (l *ArrayListIterator) Next() {
	l.index++
}
func (l *ArrayListIterator) HasNext() bool {
	return l.index < len(l.arrayList)-1
}

func (l *ArrayListIterator) CurrentItem() interface{} {
	return l.arrayList[l.index]
}
