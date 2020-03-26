package wsserver

import (
	"log"
    // "net"
	"net/http"
	"github.com/gorilla/websocket"
    ag "libnet/agent"
    "libnet/common"
    "errors"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:    4096,
    WriteBufferSize:   4096,
    EnableCompression: true,
    // 允许跨域
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type LinkInfo struct{
    cache []byte
    conn * websocket.Conn
    agent * ag.Agent
}

type Linker interface{
    append([]byte)
    callback(*common.NetPackHead,[]byte)
}

func (l * LinkInfo) append(b []byte){
    if len(b) <= 0{
        return
    }
    l.cache = append(l.cache, b...)
    for {
        if len(l.cache) <= common.PackHead {
            return
        }
        head := common.ChangeByteSliceToHead(l.cache[0:common.PackHead])
        if int(head.PackLen) <= len(l.cache){
            l.callback(head,l.cache[common.PackHead:head.PackLen])
            l.cache = l.cache[head.PackLen:]
        } else {
            break
        }
    }
}

func (l * LinkInfo) callback(head * common.NetPackHead, data [] byte){
    l.agent.Callback(head, data)
}

type ClassServer struct {
    count int
    linkMap map[*websocket.Conn] *LinkInfo
}

type Servicer interface{
    Init()
    GetConnectCnt() int
    RegistProto(interface {}) bool
    Start(string) bool
    handleSession(*websocket.Conn)
    addLink(*websocket.Conn) (bool,error)
    removeLink(*websocket.Conn)
    getLink(*websocket.Conn) *LinkInfo
}

func (c * ClassServer) Init(){
    c.linkMap = make(map[*websocket.Conn] *LinkInfo)
}

func (c *ClassServer) GetConnectCnt() int {
    return c.count
}

func (c *ClassServer) RegistProto(interface {}) bool {
    return true
}

func (c * ClassServer) addLink(conn *websocket.Conn) (bool, error){
    if _, ok := c.linkMap[conn]; ok {
        return false, errors.New("Exist")
    } else {
        c.linkMap[conn] = &LinkInfo {
            conn : conn,
            cache : make([]byte, 0),
            agent : &ag.Agent{Wsconn : conn, Nettype : "ws"},
        }
        c.count += 1
        return true, nil
    }
}

func (c * ClassServer) removeLink(conn *websocket.Conn) {
    if l, ok := c.linkMap[conn]; ok {
        l.agent.Offline()
        delete(c.linkMap, conn)
        c.count -= 1
    }
}

func (c * ClassServer) getLink(conn *websocket.Conn) *LinkInfo{
    link, ok := c.linkMap[conn]
    if ok {
        return link
    }
    return nil
}

func (c *ClassServer) Start(address string) bool {
    data := func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
            if err != nil {
            log.Print("upgrade:", err)
            return
        }
        if ok,_ := c.addLink(conn);ok{
            go c.handleSession(conn)
        }
    }
    log.Println("listen: " + address)
    http.HandleFunc("/data", data)
    log.Fatal(http.ListenAndServe(address, nil))
    return true
}

func (c *ClassServer) handleSession(conn *websocket.Conn) {
    log.Println("Session started:", conn.RemoteAddr())
    for {
        _, message, err := conn.ReadMessage()
        if err == nil {
            linkinfo := c.getLink(conn)
            if linkinfo != nil {
                linkinfo.append(message)
            } else {
                log.Println("linkinfo nil")
            }
            // conn.WriteMessage(websocket.BinaryMessage, message)
        } else {
            c.removeLink(conn)
            log.Println("Session closed", err)
            conn.Close()
            break
        }
    }
}