package libmpv

/*
#include <mpv/client.h>

struct mpv_node* GetNodeFromList(struct mpv_node* list,int i){
	return &list[i];
}

char* GetString(char** strings,int i){
	return strings[i];
}

*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"log"
	"unsafe"
)

type Node struct {
	value  interface{}
	format Format
}

func NewNode(value interface{}, format Format) *Node {
	n := &Node{}
	n.value = value
	n.format = format
	return n
}

func (n *Node) GetVal() interface{} {
	return n.value
}

func (n *Node) GetCNode() *C.mpv_node {
	ptr := data2Ptr(n.format, n.value)
	if ptr == nil {
		return nil
	}
	writer := new(bytes.Buffer)
	err := binary.Write(writer, binary.LittleEndian, uint64(uintptr(ptr)))
	if err != nil {
		log.Println("Error write bin", err)
		return nil
	}
	buf := writer.Bytes()
	node := &C.mpv_node{}
	for i := 0; i < len(buf) && i < 8; i++ {
		node.u[i] = buf[i]
	}
	node.format = C.mpv_format(n.format)
	return node
}

func GetNode(node *C.mpv_node) (*Node, error) {
	n := &Node{}
	var err error
	n.value, err = GetValue(node)
	n.format = Format(node.format)
	return n, err
}

func FreeMpvNode(cnode *C.mpv_node) {
	C.mpv_free_node_contents(cnode)
}

func GetValue(node *C.mpv_node) (interface{}, error) {
	format := Format(node.format)
	buf := bytes.NewReader(C.GoBytes(unsafe.Pointer(&node.u[0]), 8))
	var ptr uint64
	err := binary.Read(buf, binary.LittleEndian, &ptr)
	if err != nil {
		log.Println("Error binary read", err)
		return nil, err
	}
	switch format {
	case FORMAT_STRING:
		{
			var ret *C.char
			ret = (*C.char)((unsafe.Pointer)(uintptr(ptr)))
			return C.GoString(ret), nil
		}
	case FORMAT_FLAG:
		{
			var ret bool
			ret = *(*C.int)((unsafe.Pointer)(uintptr(ptr))) != 0
			return ret, nil
		}
	case FORMAT_INT64:
		{
			var ret C.int64_t
			ret = *(*C.int64_t)((unsafe.Pointer)(uintptr(ptr)))
			return int64(ret), nil
		}
	case FORMAT_DOUBLE:
		{
			var ret C.double
			ret = *(*C.double)((unsafe.Pointer)(uintptr(ptr)))
			return float64(ret), nil
		}
	case FORMAT_NODE:
		{
			var ret C.mpv_node
			ret = *(*C.mpv_node)((unsafe.Pointer)(uintptr(ptr)))
			return GetNode(&ret)
		}
	case FORMAT_NODE_ARRAY:
		{
			var ret C.mpv_node_list
			ret = *(*C.mpv_node_list)((unsafe.Pointer)(uintptr(ptr)))
			return GetNodeList(&ret)
		}
	case FORMAT_NODE_MAP:
		{
			var ret C.mpv_node_list
			ret = *(*C.mpv_node_list)((unsafe.Pointer)(uintptr(ptr)))
			return GetNodeMap(ret)
		}
	case FORMAT_BYTE_ARRAY:
		{
			var ret C.mpv_byte_array
			ret = *(*C.mpv_byte_array)((unsafe.Pointer)(uintptr(ptr)))
			return C.GoBytes(ret.data, C.int(ret.size)), nil
		}
	default:
		{
			return nil, nil
		}
	}
}

func GetNodeList(clist *C.mpv_node_list) ([]*Node, error) {
	nodes := make([]*Node, clist.num)
	var err error
	var n C.int
	for n = 0; n < clist.num; n++ {
		nodes[n], err = GetNode(C.GetNodeFromList(clist.values, n))
		if err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func GetCNodeList(list []*Node) *C.mpv_node_list {
	/*if len(list) <= 0 {
		return nil
	}
	cnlist := &C.mpv_node_list{}
	cnlist.num = C.int(len(list))
	carr := unsafe.NewArray(C.mpv_node, len(list))
	for i, n := range list {
		carr[i] = *n.GetCNode()
	}*/
	return nil
}

func GetNodeMap(cmap C.mpv_node_list) (map[string]*Node, error) {
	nodes, err := GetNodeList(&cmap)
	if err != nil {
		return nil, err
	}
	var mapnode map[string]*Node
	for i, n := range nodes {
		mapnode[C.GoString(C.GetString(cmap.keys, C.int(i)))] = n
	}
	return mapnode, nil
}
