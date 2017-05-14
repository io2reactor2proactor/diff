package main

import (
	"fmt"
	"log"
	"sort"
)

const SPILT string = "#"
//一致性hash
type Item struct {
	vir uint32
	rel *Node
	flag bool
}

type ConsistentHashCluster struct {
	nodes []Node
	replicas []Item
	virtual int
}
func NewConsistentHashCluster(length, vir int) *ConsistentHashCluster {
	return &ConsistentHashCluster{
		nodes: make([]Node, 0, length),
		replicas: make([]Item, 0, length * vir),
		virtual: vir,
	}
}
func (nhc *ConsistentHashCluster) addNode(node Node) {
	var hash uint32
	nhc.nodes = append(nhc.nodes, node)
	for index := 0; index < nhc.virtual; index++ {
		hash = bkdrHash(node.ip + SPILT + fmt.Sprintf("%v", index))
		nhc.replicas = append(nhc.replicas, Item{
			vir: hash,
			rel: &node,
			flag: true,
		})
	}
}
func (nhc *ConsistentHashCluster) removeNode(node Node) {
	for k, v := range nhc.nodes {
		if node.domain == v.domain || node.ip == v.ip {
			nhc.nodes[k].flag = false
			break
		}
	}
	
	for k, v := range nhc.replicas {
		if node.domain == v.rel.domain || node.ip == v.rel.ip {
			nhc.replicas[k].flag = false
		}
	}
}
func (nhc *ConsistentHashCluster) get(key string) *Node {
	var hash = bkdrHash(key)
	
	var pos int
	for k, v := range nhc.replicas {
		if v.flag == false {
			continue
		}
	
		if hash > v.vir {
			continue
		}

		pos = k
		break
	}
	//log.Printf("calc hash:%v, virtual node:%+v\n", hash, nhc.replicas[pos].vir)

	return nhc.replicas[pos].rel
}
func (nhc *ConsistentHashCluster) printNode() {
	for _, v := range nhc.nodes {
		if v.flag == false {
			continue
		}
		log.Printf("ip: %+v, size:%v\n", v.ip, len(v.data))
	}
}
func (nhc *ConsistentHashCluster) nodeSize() {
	var nodeTotal int = 0
	for _, v := range nhc.nodes {
		if v.flag == false {
			continue
		}
		nodeTotal = nodeTotal + 1
//		log.Printf("<%+v, %+v, %+v>\n", v.domain, v.ip, v.flag)
	}

	var virTotal int = 0
	for _, v := range nhc.replicas {
		if v.flag == false {
			continue
		}
		virTotal = virTotal + 1
//		log.Printf("<%+v: %+v, %+v, %+v>\n", v.vir, v.rel.domain, v.rel.ip, v.rel.flag)
	}
	log.Printf("real node size: %v, virtual node size: %v\n", nodeTotal, virTotal)
}
func (nhc *ConsistentHashCluster) sortVirtualNode() {
	sort.Sort(Key(nhc.replicas))
}
