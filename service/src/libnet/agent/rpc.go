package agent

import(
	"sync"
	"config"
	"net/rpc"
)

var lock sync.Mutex
var instanceConn * rpc.Client	

func rpcDataBase(rpcCbkName string, param interface{}, reply interface{}) error {
	if instanceConn == nil {
		lock.Lock()
		if instanceConn == nil {
			if addr, err := config.GetConfigString("database");err == nil{
				if conn, err := rpc.Dial("tcp", addr); err == nil{
					instanceConn = conn
				}
			}
		}
		lock.Unlock()
	}
	if instanceConn == nil{
		panic("rpc database failed")
	}
	return instanceConn.Call(rpcCbkName, param, reply)
}