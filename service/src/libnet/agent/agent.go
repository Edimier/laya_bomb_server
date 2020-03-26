package agent

import(
	"github.com/golang/protobuf/proto"
	"libnet/common"
	"libnet/protocol"
	"log"
	"net"
	"reflect"
	"proto/client"
	"proto/rpcdatabase"
	"github.com/gorilla/websocket"
	// "proto/rpcredis"
)

type Agenter interface{
	Offline()
	Callback(head * common.NetPackHead, data [] byte)
	SendData(data []byte) error
	// DecodeProtoBuff( * common.NetPackHead, [] byte)
}

type Agent struct{
	Status common.NetStatusType
	Tcpconn net.Conn
	Nettype string
	Wsconn * websocket.Conn
}

func (a * Agent) Offline(){
	a.Status = common.NET_STATUS_OFFLINE
}

func (a * Agent) Callback(head * common.NetPackHead, data [] byte){
	if decodeData := protocol.Decode(head, data); decodeData != nil{
		log.Println("Server Recieve,",common.HeadMsg(head),",msg =",decodeData)
		// return
		switch decodeData.(type){
			case string:
			case []byte:
			case map[string]interface {}: // json

			// case * client.Login:

			case * client.ReqRank:
				if op, ok := decodeData.(*client.ReqRank); ok {
					respRank := & client.RespRank{
			            Id : op.Id,
			        }

			        respRank.Ranks = append(respRank.Ranks, &client.RankInfo{
			            Name : "小明",
			            Prop : &client.Prop{
			            	Propid : 1,
				            Proptype : 1,
				            Expire : 1,
				            Count : 100,
			            },
			        })

			        respRank.Ranks = append(respRank.Ranks, &client.RankInfo{
			            Name : "小强",
			            Prop : &client.Prop{
			            	Propid : 1,
				            Proptype : 1,
				            Expire : 1,
				            Count : 50,
			            },
			        })
	        		data, _ := proto.Marshal(respRank)
	        		a.SendData(protocol.EncodeProtoBuffer("client.RespRank", data))
				} else {
					log.Println("client.ReqRank error")
				}
			case * client.ReqPlayerInfo:
				if op, ok := decodeData.(*client.ReqPlayerInfo); ok{
					var reply = &rpcdatabase.RespPlayerInfo{}
				    var param = &rpcdatabase.ReqPlayerInfo{
				    	Id : op.Id,
				    }
				    err := rpcDataBase("DataBase.ReqPlayerInfo", &param, &reply)

				    if err != nil {
				        log.Println("DataBase RPC failed", err)
				    } else {
				    	respPlayerInfo := & client.RespPlayerInfo{
				            Id : op.Id,
				            Name : reply.Name,
				            Rich : reply.Rich,
				            Sex : reply.Sex,
				            Phone : reply.Phone,
				        }
				        for _, b := range reply.Bag {
				        	respPlayerInfo.Bag = append(respPlayerInfo.Bag, &client.Prop{
					            Propid : b.Propid,
					            Proptype : b.Proptype,
					            Expire : b.Expire,
					            Count : b.Count,
					        })
				        }
				        data, _ := proto.Marshal(respPlayerInfo)
				        a.SendData(protocol.EncodeProtoBuffer("client.RespPlayerInfo", data))
				    }
				} else {
					log.Println("client.ReqPlayerInfo error")
				}
			default:
				typeOfData := reflect.TypeOf(decodeData)
				log.Println("Not find case ", "Name:", typeOfData.Name(), "Kind:", typeOfData.Kind())
		}

	} else {
		log.Println("Decode proto buff failed")
	}
}

func (a * Agent) SendData(data []byte) error {
	// log.Println("Server sendData", data)
	if a.Nettype == "tcp"{
		_, err := a.Tcpconn.Write(data)
		return err
	} else {
		err := a.Wsconn.WriteMessage(websocket.BinaryMessage, data)
        return err
	}
	
}