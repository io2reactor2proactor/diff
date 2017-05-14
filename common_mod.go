package main

import (
	"log"
)

type CommonHashCluster struct {
	nodes []Node
}
//普通hash
func NewCommonHashCluster(length int) *CommonHashCluster {
	return &CommonHashCluster{
		nodes: make([]Node, 0, length),
	}
}
func (nhc *CommonHashCluster) addNode(node Node) {
	nhc.nodes = append(nhc.nodes, node)
}
func (nhc *CommonHashCluster) removeNode(node Node) {
	for k, v := range nhc.nodes {
		if node.domain == v.domain || node.ip == v.ip {
			nhc.nodes[k].flag = false
		}
	}
}
func (nhc *CommonHashCluster) get(key string) *Node {
	var hash = int(bkdrHash(key))
	//log.Printf("hash:%v\n", hash)
	var idx = hash % len(nhc.nodes)
	//log.Printf("idx:%v\n", idx)
	if nhc.nodes[idx].flag == false {
		return &Node{}
	}
	return &nhc.nodes[idx]
}
func (nhc *CommonHashCluster) printNode() {
	for _, v := range nhc.nodes {
		if v.flag == false {
			continue
		}
		log.Printf("ip: %+v, size:%v\n", v.ip, len(v.data))
	}
}
func (nhc *CommonHashCluster) nodeSize() {
	log.Printf("node size: %+v\n", len(nhc.nodes))
}
func (nhc *CommonHashCluster) sortVirtualNode() {
}
