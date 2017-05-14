package main

import (
	"fmt"
	"log"
	"time"
	"flag"
)

type Cluster interface {
	addNode(n Node) //新增节点
	removeNode(n Node) //删除节点
	get(key string) *Node //获取节点
	printNode() //遍历节点信息
	sortVirtualNode() //虚拟节点排序
	nodeSize() //节点信息
}

var (
	kind *string = flag.String("k", "", "db or cache storage kind, value is common or consistent")
)

const (
	VIRT_COUNT int = 1024
	NODE_COUNT int = 10
	DATA_COUNT int = 1000000
	PRE_KEY string = "test"
)

func main() {
	flag.PrintDefaults()
	flag.Parse()
	
	var c Cluster
	switch *kind {
	case "common":
		c = NewCommonHashCluster(NODE_COUNT)
	case "consistent":
		c = NewConsistentHashCluster(NODE_COUNT, VIRT_COUNT)
	default:
		return
	}

	//var c Cluster = NewConsistentHashCluster(NODE_COUNT, VIRT_COUNT)
	//var c Cluster = NewCommonHashCluster(NODE_COUNT)
	c.addNode(Node{
		domain: "c1.yywang.info", 
		ip: "192.168.0.1", 
		data: make(map[string]interface{}),
		flag: true,
	})
	c.addNode(Node{
		domain: "c2.yywang.info", 
		ip: "192.168.0.2", 
		data: make(map[string]interface{}),
		flag: true,
	})
	c.addNode(Node{
		domain: "c3.yywang.info", 
		ip: "192.168.0.3",
		data: make(map[string]interface{}),
		flag: true,
	})
	c.addNode(Node{
		domain: "c4.yywang.info", 
		ip: "192.168.0.4", 
		data: make(map[string]interface{}),
		flag: true,
	})
	c.sortVirtualNode() //虚拟节点升序排序
	c.nodeSize()
	fmt.Println()

	var val string
	var start = time.Now()
	for index := 0; index < DATA_COUNT; index++ {
		val = PRE_KEY + fmt.Sprintf("%v", index)
		var node = c.get(val)
		node.put(val, "temp data")
	}
	log.Printf("take time: %+v\n", time.Now().Sub(start))

	//增加一个节点
	c.addNode(Node{
		domain: "c5.yywang.info", 
		ip: "192.168.0.5", 
		data: make(map[string]interface{}),
		flag: true,
	})
	
	//删除一个节点
	c.removeNode(Node{
		domain: "c4.yywang.info",
		ip: "192.168.0.4",
		flag: false,
	})
	
	c.sortVirtualNode() //虚拟节点升序排序
	c.nodeSize()
	fmt.Println()
	
	fmt.Println()
	log.Printf("数据分布情况:\n")
	c.printNode()
	
	//缓存命中率		
	var hitCount int = 0
	for index := 0; index < DATA_COUNT; index++ {
		val = PRE_KEY + fmt.Sprintf("%v", index)
		if c.get(val).get(val) {
			hitCount = hitCount + 1
		}
	}
	log.Printf("hitCount: %v, totalCount: %v\n", hitCount, DATA_COUNT)
	log.Printf("缓存命中率：%v\n", float64(hitCount) / float64(DATA_COUNT))
}
