package common

import (
	"unsafe"
	"log"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type Error_t int
const (
    NET_OK 			Error_t = iota
    NET_WriteErr
    NET_ReadErr
    NET_ConnectError
)

type ProtoType = int8
const (
	NET_PROTO_DEFAULT ProtoType = iota
	NET_PROTO_STRING
	NET_PROTO_JSON
	NET_PROTO_PROTOBUFF
	NET_PROTO_BYTE
)

type NetStatusType = int
const (
	NET_STATUS_OK NetStatusType = iota
	NET_STATUS_OFFLINE
)

/***********************************************************/
/***********************************************************/

type NetResult struct {
	Id 		int
	Errcode Error_t
	Errmsg  string
}

const (
	PackHead int = 8
	PackDataLen int = 65535 // 虽然一次可以发很多，但限制一次发包最多65535个字节
)

// 注意变量字节顺序
type NetPackHead struct{
	PackLen   int32   // 包长度，包括头
	ProtoIdx  int16   // 根据协议类型做的附加字段
	ProtoType int8    // 协议类型，见ProtoType
	Extend    int8    // 预留字段
}

type sliceHeader struct {
    addr uintptr
    len  int
    cap  int
}

func ChangeHeadToByteSlice(t * NetPackHead) []byte {
	if t != nil {
		s := &sliceHeader{
	        addr: uintptr(unsafe.Pointer(t)),
	        len:  int(unsafe.Sizeof(*t)),
	        cap:  int(unsafe.Sizeof(*t)),
	    }
	    return *(*[]byte)(unsafe.Pointer(s))
	}
	return nil
}

func ChangeByteSliceToHead(data []byte) *NetPackHead {
	return *(**NetPackHead)(unsafe.Pointer(&data))	
}

func ReadJsonFile(filename string) map[string] string {
	data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Println("json file read error : %s", err)
        return nil
    }
    
    var jsonData interface{}
    err = json.Unmarshal(data, &jsonData)
    if err == nil {
    	tmp, ok := jsonData.(map[string]interface{})
    	if ok {
    		retData := make(map[string]string)
        	for k1, v1 := range tmp {
            	switch v2 := v1.(type) {
	            	case string:
	            		_, exist := retData[k1]
	            		if exist{
	            			log.Println(k1,"Have exist,cover it")
	            		}
	            		retData[k1] = v2
	            	default:
	            		log.Println(k1, "类型未知")
	            }
        	}
        	if len(retData) > 0{
        		return retData
        	}
    	}
    }
    return nil
}

func GetProtoType(t ProtoType) string {
	switch t{
	case NET_PROTO_STRING:
		return "string"
	case NET_PROTO_JSON:
		return "json"
	case NET_PROTO_PROTOBUFF:
		return "protoBuff"
	case NET_PROTO_BYTE:
		return "bytes"
	default:
		return "default"
	}

}
func HeadMsg(head * NetPackHead) string{
	if head != nil{
		return "PackLen=" + strconv.Itoa(int(head.PackLen)) +
		 		",ProtoIdx=" + strconv.Itoa(int(head.ProtoIdx)) +
		 		",ProtoType=" + GetProtoType(head.ProtoType) +
		 		",Extend=" + strconv.Itoa(int(head.Extend))
	} 
	return ""
}