package server
import (
    "log"
    "net"
    // "bytes"
    "libnet/common"
    "errors"
    ag "libnet/agent"
)

type LinkInfo struct{
    cache []byte
    conn net.Conn
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
    linkMap map[net.Conn] *LinkInfo
}

type Servicer interface{
    Init()
    GetConnectCnt() int
    RegistProto(interface {}) bool
    Start(string) bool
    handleSession(net.Conn)
    addLink(net.Conn) (bool,error)
    removeLink(net.Conn)
    getLink(net.Conn) *LinkInfo
}

func (c * ClassServer) Init(){
    c.linkMap = make(map[net.Conn] *LinkInfo)
}

func (c *ClassServer) GetConnectCnt() int {
    return c.count
}

func (c *ClassServer) RegistProto(interface {}) bool {
    return true
}

func (c * ClassServer) addLink(conn net.Conn) (bool, error){
    if _, ok := c.linkMap[conn]; ok {
        return false, errors.New("Exist")
    } else {
        c.linkMap[conn] = &LinkInfo {
            conn : conn,
            cache : make([]byte, 0),
            agent : &ag.Agent{Tcpconn : conn, Nettype : "tcp"},
        }
        c.count += 1
        return true, nil
    }
}

func (c * ClassServer) removeLink(conn net.Conn) {
    if l, ok := c.linkMap[conn]; ok {
        l.agent.Offline()
        delete(c.linkMap, conn)
        c.count -= 1
    }
}

func (c * ClassServer) getLink(conn net.Conn) *LinkInfo{
    link, ok := c.linkMap[conn]
    if ok {
        return link
    }
    return nil
}

func (c *ClassServer) Start(address string) bool {
     l, err := net.Listen("tcp", address)
    // 如果侦听发生错误, 打印错误并退出
    if err != nil {
        log.Println(err.Error())
        // c.exitChan <- 1
        return false
    }
    // 打印侦听地址, 表示侦听成功
    log.Println("listen: " + address)
    // 延迟关闭侦听器
    defer l.Close()
    // 侦听循环
    for {
        // 新连接没有到来时, Accept是阻塞的
        conn, err := l.Accept() 

        // 发生任何的侦听错误, 打印错误并退出服务器
        if err != nil {
            log.Println(err.Error())
            continue
        }
        if ok, _ := c.addLink(conn); ok{
            go c.handleSession(conn)
        }
    }
    return true
}

func (c *ClassServer) handleSession(conn net.Conn) {
    log.Println("Session started:", conn.RemoteAddr())
    for {
        // result := bytes.NewBuffer(nil)
        var buf [common.PackDataLen] byte
        n, err := conn.Read(buf[0:])
        if err == nil{
            // result.Write(buf[0:n])
            // log.Println("recieve len=", n)

            linkinfo := c.getLink(conn)
            if linkinfo != nil {
                linkinfo.append(buf[0:n])
            } else {
                log.Println("linkinfo nil")
            }
            // _, err = conn.Write(buf[0:n])
        } else {
            c.removeLink(conn)
            log.Println("Session closed", err)
            conn.Close()
            break
        }
    }
}
