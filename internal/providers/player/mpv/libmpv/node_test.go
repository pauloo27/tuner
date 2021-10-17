package libmpv

import (
	"log"
	"testing"
)

func TestNode(t *testing.T) {
	{
		val := "123456789"
		format := FORMAT_STRING
		node := NewNode(val, format)
		cnode := node.GetCNode()
		node, err := GetNode(cnode)
		log.Println("CNode:", cnode, "Node val:", node.GetVal().(string), "Error:", err)
		if cnode == nil {
			log.Fatal("Fail node convert string")
		}
	}
	{
		val := int64(0x7FFFFFFFFFFFFFFF)
		format := FORMAT_INT64
		node := NewNode(val, format)
		cnode := node.GetCNode()
		node, err := GetNode(cnode)
		log.Println("CNode:", cnode, "Node val:", node.GetVal(), "Error:", err)
		if cnode == nil {
			log.Fatal("Fail node convert int64")
		}
	}
	{
		val := bool(true)
		format := FORMAT_FLAG
		node := NewNode(val, format)
		cnode := node.GetCNode()
		node, err := GetNode(cnode)
		log.Println("CNode:", cnode, "Node val:", node.GetVal(), "Error:", err)
		if cnode == nil {
			log.Fatal("Fail node convert int64")
		}
	}

	{
		val := float64(1.7976931348623157E+308)
		format := FORMAT_DOUBLE
		node := NewNode(val, format)
		cnode := node.GetCNode()
		node, err := GetNode(cnode)
		log.Println("CNode:", cnode, "Node val:", node.GetVal(), "Error:", err)
		if cnode == nil {
			log.Fatal("Fail node convert int64")
		}
	}

	{
		val := NewNode(123, FORMAT_INT64)
		format := FORMAT_NODE
		node := NewNode(val, format)
		cnode := node.GetCNode()
		node, err := GetNode(cnode)
		log.Println("CNode:", cnode, "Node val:", node.GetVal(), "Error:", err)
		if cnode == nil {
			log.Fatal("Fail node convert int64")
		}
	}
}
