package protocol

import(
    "github.com/golang/protobuf/proto"
    "libnet/common"
    "log"
    "strconv"
    "reflect"
    "encoding/json"
    "proto/protomp"
)

func getPropMap(key string) (string, bool){
    return protomp.GetProto(key)
}

func Decode(head * common.NetPackHead, data [] byte) interface{}{
	switch head.ProtoType {
        case common.NET_PROTO_STRING:
            return string(data)
        case common.NET_PROTO_JSON:
            var retJsonData interface{}
            json.Unmarshal(data, &retJsonData)
            if retJsonData != nil{
                return retJsonData
            } else {
               log.Println("json decode failed")
            }
        case common.NET_PROTO_PROTOBUFF:
            // return decodeProtoBuffer(head.ProtoIdx, data)
            str := strconv.Itoa(int(head.ProtoIdx))
            if protoName, ok := getPropMap(str);ok{
                mt := proto.MessageType(protoName)
                if mt != nil {
                    info := reflect.New(mt.Elem())
                    if err := proto.Unmarshal(data, info.Interface().(proto.Message)); err != nil {
                        log.Println("proto.Unmarshal failed", protoName)
                    } else {
                        return info.Interface()
                    }
                }
            } else {
                log.Println("Not found ProtoIdx proto from protoMap", head.ProtoIdx)
            }
            return nil
        case common.NET_PROTO_BYTE:
            return data
        default:
            log.Println("Can not found proto type", head.ProtoType)
    }
    return nil
}

func decodeProtoBuffer(idx int16, data []byte) interface{} {
	str := strconv.Itoa(int(idx))

	if protoName, ok := getPropMap(str);ok{
		mt := proto.MessageType(protoName)
		if mt != nil {
			info := reflect.New(mt.Elem())
			if err := proto.Unmarshal(data, info.Interface().(proto.Message)); err != nil {
				log.Println("proto.Unmarshal failed", protoName)
			} else {
				return info.Interface()
			}
		}
	} else {
		log.Println("Not found idx proto from protoMap", idx)
	}
	return nil
}

/********************************************************************************/
/********************************************************************************/
func encode(head * common.NetPackHead, data []byte) []byte{
    datalen := len(data) + common.PackHead
    if datalen > common.PackDataLen || datalen <= common.PackHead {
        log.Println("Data length error", datalen)
        return nil
    }
    head.PackLen = int32(datalen)
    // log.Println("head = ", head)
    headBytes := common.ChangeHeadToByteSlice(head)
    retData := make([]byte, 0)
    retData = append(retData, headBytes...)
    retData = append(retData, data...)
    return retData[0:]
}

func EncodeString(str string) [] byte{
    return encode(&common.NetPackHead{ProtoType : common.NET_PROTO_STRING}, []byte(str))
}

func EncodeJson(data []byte) [] byte{
    return encode(&common.NetPackHead{ProtoType : common.NET_PROTO_JSON}, data)
}

func EncodeBytes(data []byte) [] byte{
    return encode(&common.NetPackHead{ProtoType : common.NET_PROTO_BYTE}, data)
}

func EncodeProtoBuffer(protoName string, data []byte) []byte{
    if idx, ok := getPropMap(protoName); ok {
        if tmp, err := strconv.Atoi(idx); err == nil {
            return encode(&common.NetPackHead{ProtoType : common.NET_PROTO_PROTOBUFF, ProtoIdx : int16(tmp)}, data)
        }        
    }
    panic("EncodeProtoBuffer failed,protoName=" + protoName)
    return nil
}

func GetPropName(idx int16) string {
    str := strconv.Itoa(int(idx))
    if protoName, ok := getPropMap(str);ok{
        return protoName
    }
    return ""
}

/********************************************************************************/
/********************************************************************************/