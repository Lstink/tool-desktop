package message

type Data struct {
	Cmd    int
	Remark string
	List   []Item
}

type Item struct {
	Key   string
	Value string
}
