package libmpv

/*
#include <mpv/client.h>
#include <stdlib.h>
#cgo LDFLAGS: -lmpv

char** makeCharArray1(int size) {
    return calloc(sizeof(char*), size);
}
void setArrayString1(char** a, int i, char* s) {
    a[i] = s;
}

*/
import "C"

import (
	"unsafe"
)

type Mpv struct {
	handle              *C.mpv_handle
	wakeup_callbackVar  interface{}
	wakeup_callbackFunc func(d interface{})
}

func Create() *Mpv {
	ctx := C.mpv_create()
	if ctx == nil {
		return nil
	}
	return &Mpv{ctx, nil, nil}
}

func (m *Mpv) ClientName() string {
	return C.GoString(C.mpv_client_name(m.handle))
}

func (m *Mpv) Initialize() error {
	return NewError(C.mpv_initialize(m.handle))
}

func (m *Mpv) TerminateDestroy() {
	C.mpv_terminate_destroy(m.handle)
}

func (m *Mpv) CreateClient(name string) *Mpv {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cmpv := C.mpv_create_client(m.handle, cname)
	if cmpv != nil {
		return &Mpv{cmpv, nil, nil}
	}
	return nil
}

func (m *Mpv) LoadConfigFile(fileName string) error {
	cfn := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfn))
	return NewError(C.mpv_load_config_file(m.handle, cfn))
}

func (m *Mpv) GetTimeUS() int64 {
	return int64(C.mpv_get_time_us(m.handle))
}

func (m *Mpv) SetOption(name string, format Format, data interface{}) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ptr := data2Ptr(format, data)
	return NewError(C.mpv_set_option(m.handle, cname, C.mpv_format(format), ptr))
}

func (m *Mpv) SetOptionString(name, data string) error {
	cname := C.CString(name)
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cname))
	defer C.free(unsafe.Pointer(cdata))
	return NewError(C.mpv_set_option_string(m.handle, cname, cdata))
}

func (m *Mpv) Command(command []string) error {
	cArray := C.makeCharArray1(C.int(len(command) + 1))
	if cArray == nil {
		panic("got NULL from calloc")
	}
	defer C.free(unsafe.Pointer(cArray))

	for i, s := range command {
		cStr := C.CString(s)
		C.setArrayString1(cArray, C.int(i), cStr)
		defer C.free(unsafe.Pointer(cStr))
	}

	return NewError(C.mpv_command(m.handle, cArray))
}

func (m *Mpv) CommandNode(command []string) int {
	//int mpv_command_node(mpv_handle *ctx, mpv_node *args, mpv_node *result);
	//TODO
	panic("Not supported command")
	return -1
}

func (m *Mpv) CommandString(command string) error {
	ccmd := C.CString(command)
	defer C.free(unsafe.Pointer(ccmd))
	return NewError(C.mpv_command_string(m.handle, ccmd))
}

func (m *Mpv) CommandAsync(replyUserdata uint64, command []string) error {
	cArray := C.makeCharArray1(C.int(len(command) + 1))
	if cArray == nil {
		panic("got NULL from calloc")
	}
	defer C.free(unsafe.Pointer(cArray))

	for i, s := range command {
		cStr := C.CString(s)
		C.setArrayString1(cArray, C.int(i), cStr)
		defer C.free(unsafe.Pointer(cStr))
	}

	return NewError(C.mpv_command_async(m.handle, C.uint64_t(replyUserdata), cArray))
}

func (m *Mpv) CommandNodeAsync(command []string) int {
	//int mpv_command_node_async(mpv_handle *ctx, uint64_t reply_userdata, mpv_node *args);
	//TODO
	panic("Not supported command")
	return -1
}

func (m *Mpv) SetProperty(name string, format Format, data interface{}) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ptr := data2Ptr(format, data)
	return NewError(C.mpv_set_property(m.handle, cname, C.mpv_format(format), ptr))
}

func (m *Mpv) SetPropertyString(name, data string) error {
	cname := C.CString(name)
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cname))
	defer C.free(unsafe.Pointer(cdata))
	return NewError(C.mpv_set_property_string(m.handle, cname, cdata))
}

func (m *Mpv) SetPropertyAsync(name string, replyUserdata uint64, format Format, data interface{}) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	ptr := data2Ptr(format, data)
	return NewError(C.mpv_set_property_async(m.handle, C.uint64_t(replyUserdata), cname, C.mpv_format(format), ptr))
}

func (m *Mpv) GetProperty(name string, format Format) (interface{}, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	switch format {
	case FORMAT_STRING, FORMAT_OSD_STRING:
		{
			var cval *C.char
			err := NewError(C.mpv_get_property(m.handle, cname, C.mpv_format(format), unsafe.Pointer(&cval)))
			if err != nil {
				return nil, err
			}
			defer C.mpv_free(unsafe.Pointer(cval))
			return C.GoString(cval), nil
		}
	case FORMAT_INT64:
		{
			var cval C.int64_t
			err := NewError(C.mpv_get_property(m.handle, cname, C.mpv_format(format), unsafe.Pointer(&cval)))
			if err != nil {
				return nil, err
			}
			return int64(cval), nil
		}
	case FORMAT_DOUBLE:
		{
			var cval C.double
			err := NewError(C.mpv_get_property(m.handle, cname, C.mpv_format(format), unsafe.Pointer(&cval)))
			if err != nil {
				return nil, err
			}
			return float64(cval), nil
		}
	case FORMAT_FLAG:
		{
			var cval C.int
			err := NewError(C.mpv_get_property(m.handle, cname, C.mpv_format(format), unsafe.Pointer(&cval)))
			if err != nil {
				return nil, err
			}
			return cval == 1, nil
		}
	case FORMAT_NONE:
		{
			err := NewError(C.mpv_get_property(m.handle, cname, C.mpv_format(format), nil))
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
	case FORMAT_NODE:
		{
			var cval C.mpv_node
			err := NewError(C.mpv_get_property(m.handle, cname, C.mpv_format(format), unsafe.Pointer(&cval)))
			if err != nil {
				return nil, err
			}
			return GetNode(&cval)
		}
	case FORMAT_NODE_ARRAY:
		{
			var cval C.mpv_node_list
			err := NewError(C.mpv_get_property(m.handle, cname, C.mpv_format(format), unsafe.Pointer(&cval)))
			if err != nil {
				return nil, err
			}
			return GetNodeList(&cval)
		}
	case FORMAT_NODE_MAP:
		{
			var cval C.mpv_node_list
			err := NewError(C.mpv_get_property(m.handle, cname, C.mpv_format(format), unsafe.Pointer(&cval)))
			if err != nil {
				return nil, err
			}
			return GetNodeMap(cval)
		}
	default:
		panic("Not supported format")
	}
}

func (m *Mpv) GetPropertyString(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	cstr := C.mpv_get_property_string(m.handle, cname)
	if cstr != nil {
		str := C.GoString(cstr)
		C.mpv_free(unsafe.Pointer(cstr))
		return str
	}

	return ""
}

func (m *Mpv) GetPropertyOsdString(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cstr := C.mpv_get_property_osd_string(m.handle, cname)
	if cstr != nil {
		str := C.GoString(cstr)
		C.mpv_free(unsafe.Pointer(cstr))
		return str
	}

	return ""
}

func (m *Mpv) GetPropertyAsync(name string, replyUserdata uint64, format Format) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return NewError(C.mpv_get_property_async(m.handle, C.uint64_t(replyUserdata), cname, C.mpv_format(format)))
}

func (m *Mpv) ObserveProperty(replyUserdata uint64, name string, format Format) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewError(C.mpv_observe_property(m.handle, C.uint64_t(replyUserdata), cname, C.mpv_format(format)))
}

func (m *Mpv) UnObserveProperty(registeredReplyUserdata uint64) error {
	return NewError(C.mpv_unobserve_property(m.handle, C.uint64_t(registeredReplyUserdata)))
}

func (m *Mpv) RequestEvent(event EventId, enable bool) error {
	var en C.int = 0
	if enable {
		en = 1
	}
	return NewError(C.mpv_request_event(m.handle, C.mpv_event_id(event), en))
}

func (m *Mpv) RequestLogMessages(minLevel string) error {
	clevel := C.CString(minLevel)
	defer C.free(unsafe.Pointer(clevel))
	return NewError(C.mpv_request_log_messages(m.handle, clevel))
}

func (m *Mpv) WaitEvent(timeout float32) *Event {
	var cevent *C.mpv_event
	cevent = C.mpv_wait_event(m.handle, C.double(timeout))
	if cevent == nil {
		return nil
	}

	e := &Event{}

	e.Event_Id = EventId(cevent.event_id)
	e.Reply_Userdata = uint64(cevent.reply_userdata)
	e.Error = NewError(cevent.error)
	if e.Event_Id == EVENT_END_FILE {
		var eef *C.mpv_event_end_file = (*C.struct_mpv_event_end_file)(cevent.data)
		efr := EventEndFile{}
		efr.Reason = EndFileReason(eef.reason)
		efr.ErrCode = Error(eef.error)
		e.Data = &efr
	} else if e.Event_Id == EVENT_PROPERTY_CHANGE {
		rawData := (*C.mpv_event_property)(cevent.data)
		mep := EventProperty{
			Name:   C.GoString(rawData.name),
			Format: Format(rawData.format),
			Data:   rawData.data,
		}
		e.Data = &mep
	} else {
		e.Data = cevent.data
	}
	return e
}

func (m *Mpv) Wakeup() {
	C.mpv_wakeup(m.handle)
}

func (m *Mpv) GetWakeupPipe() int {
	return int(C.mpv_get_wakeup_pipe(m.handle))
}

func (m *Mpv) WaitAsyncRequests() {
	C.mpv_wait_async_requests(m.handle)
}

func data2Ptr(format Format, data interface{}) unsafe.Pointer {
	var ptr unsafe.Pointer = nil
	switch format {
	case FORMAT_STRING, FORMAT_OSD_STRING:
		{
			ptr = unsafe.Pointer(&[]byte(data.(string))[0])
		}
	case FORMAT_INT64:
		{
			i, ok := data.(int64)
			if !ok {
				i = int64(data.(int))
			}
			val := C.int64_t(i)
			ptr = unsafe.Pointer(&val)
		}
	case FORMAT_DOUBLE:
		{
			val := C.double(data.(float64))
			ptr = unsafe.Pointer(&val)
		}
	case FORMAT_FLAG:
		{
			val := C.int(0)
			if data.(bool) {
				val = 1
			}
			ptr = unsafe.Pointer(&val)
		}
	case FORMAT_NONE:
		{
			return nil
		}

	case FORMAT_NODE:
		{
			val := (data.(*Node))
			cnode := val.GetCNode()
			ptr = unsafe.Pointer(cnode)
		}

	case FORMAT_NODE_ARRAY, FORMAT_NODE_MAP:
		{
			return nil
		}
	}
	return ptr
}

type Event struct {
	Event_Id       EventId
	Error          error
	Reply_Userdata uint64
	Data           interface{}
}

type EventProperty struct {
	Name   string
	Format Format
	Data   interface{}
}

type EventEndFile struct {
	Reason  EndFileReason
	ErrCode Error
}
