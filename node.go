package main

type Node struct {
	domain string
	ip string
	data map[string]interface{}
	flag bool
}
func (n *Node) put(key string, value interface{}) {
	if n.flag == true {
		n.data[key] = value
	}
	//n.data[key] = value
}
func (n *Node) remove(key string) {
	if n.flag == true {
		delete(n.data, key)
	}
	//delete(n.data, key)
}
func (n *Node) get(key string) bool {
	var ok bool = false
	if n.flag == true {
		_, ok = n.data[key]
	}
	return ok
	
	//_, ok := n.data[key]
	//return ok
}
