
package user
import (
    "fmt"
    "log"
    "lib/redis"
    "encoding/json"

    "crypto/md5"
    "io"
    "strconv"
    "time"
    "config"
    "lib/mysql"
)

type registData struct{
    Errcode int
    Msg string
    UserId string
}

type loginData struct{
    Errcode int
    Msg string
    Token string
    GateIp string
    GatePort string
}

const(
    HTTP_OK int = iota
    HTTP_USER_REGIST_FAIL // 注册失败
    HTTP_USER_REGISTED   //  玩家已经注册
    HTTP_USER_REGIST_OK  //  注册成功
    HTTP_USER_REGIST_EMPTY


    HTTP_USER_LOGIN_NOUSER      //  玩家不存在
    HTTP_USER_LOGIN_OK   // 玩家登录成功
    HTTP_USER_LOGIN_PWDERR   //
    HTTP_USER_LOGIN_FAILED   //

    HTTP_REATED_MAC // 重复的硬件码

    HTTP_USER_LOGIN_REPEATED_OK   // 玩家重复登录成功
)

var noticeMsg map[int]string

func init(){
    noticeMsg = make(map[int]string)
    noticeMsg[HTTP_OK] = "OK"
    noticeMsg[HTTP_USER_REGIST_FAIL] = "Regist Failed"
    noticeMsg[HTTP_USER_REGISTED] = "Regist Existed User"
    noticeMsg[HTTP_USER_REGIST_OK] = "Regist OK"
    noticeMsg[HTTP_USER_REGIST_EMPTY] = "Regist Param Error"
    noticeMsg[HTTP_USER_LOGIN_NOUSER] = "Login User Not Existed"
    noticeMsg[HTTP_USER_LOGIN_OK] = "Login OK"
    noticeMsg[HTTP_USER_LOGIN_PWDERR] = "Login Error Password"
    noticeMsg[HTTP_USER_LOGIN_FAILED] = "Login Failed"
    noticeMsg[HTTP_REATED_MAC] = "Mac Error"
    noticeMsg[HTTP_USER_LOGIN_REPEATED_OK] = "Login OK Repeated"
}

func notice(v int) string{
    if msg, ok := noticeMsg[v];ok{
        return msg
    }
    return "Empty"
    
}

func getToken() string {
    crutime := time.Now().UnixNano()
    h := md5.New()
    io.WriteString(h, strconv.FormatInt(crutime, 10))
    token := fmt.Sprintf("%x", h.Sum(nil))
    return token
}

var userIdBase int = 0

func createUserId() int{
    if userIdBase == 0{
        if value, err := redis.HGet("numid", "base");err == nil{
            v, _ := strconv.Atoi(value)
            userIdBase = v
        } else {
            if err := redis.HSet("numid", "base", "100000"); err == nil{
                userIdBase = 100000
                return userIdBase
            } else  {
                return 0
            }
        }
    }

    if err := redis.HIncrBy("numid", "base", 1);err == nil{
        userIdBase++
    }
    return userIdBase
}

func getTimeGlag() string{
    now := time.Now() //获取当前时间
    // log.Printf("current time:%v\n", now)
    year := now.Year()     //年
    month := now.Month()   //月
    day := now.Day()       //日
    hour := now.Hour()     //小时
    minute := now.Minute() //分钟
    second := now.Second() //秒
    return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func Regist(name, pwd, mac string) []byte {
    reponse := &registData{}

    if pwd == "" || name == "" || mac == "" {
        reponse.Errcode = HTTP_USER_REGIST_EMPTY
        reponse.Msg = notice(reponse.Errcode)
        goto OVER
    }

    if value, err := redis.HGet("mac", mac); err == nil{
        reponse.Errcode = HTTP_USER_REGISTED
        reponse.Msg = notice(reponse.Errcode)
        reponse.UserId = value
    } else {
        // redis 不命中，从sql中查询
        sql := fmt.Sprintf("Select numid, registTime from user where mac=\"%s\" limit 1", mac)

        rows, _ := mysql.Select(sql)
        // log.Println(ret, err)
        defer rows.Close()
        var numid int
        var registTime string
        for rows.Next() {
            rows.Scan(&numid, &registTime)
        }
        if numid != 0 {

            if err := redis.HSet("mac", mac, strconv.Itoa(numid));err == nil{
                redis.HSet("userids", strconv.Itoa(numid), registTime)
                reponse.Errcode = HTTP_USER_REGISTED
                reponse.Msg = notice(reponse.Errcode)
                reponse.UserId = strconv.Itoa(numid)
            } else {
                // log.Println("11111111", err)
                reponse.Errcode = HTTP_USER_REGIST_FAIL
                reponse.Msg = notice(reponse.Errcode)
            }
        } else {
            timeFlag :=  getTimeGlag()
            numid := createUserId()
            sql := fmt.Sprintf("Insert into user(numid, mac, name, pwd, registTime) values(%d, \"%s\", \"%s\", \"%s\", \"%s\")", numid, mac, name, pwd, timeFlag)
            _, err := mysql.Insert(sql)
            if err != nil{
                log.Println(err)
            } else {
                if err := redis.HSet("mac", mac, strconv.Itoa(numid));err == nil{
                    redis.HSet("userids", strconv.Itoa(numid), timeFlag)
                    reponse.Errcode = HTTP_USER_REGIST_OK
                    reponse.Msg = notice(reponse.Errcode)
                    reponse.UserId = strconv.Itoa(numid)
                } else {
                    // log.Println("11111112")
                    reponse.Errcode = HTTP_USER_REGIST_FAIL
                    reponse.Msg = notice(reponse.Errcode)
                }
            }
        }
    }

OVER:
    jsonData, _ := json.Marshal(reponse)
    return jsonData
}

func Login(numid, pwd string) []byte {
    // testSql(name, pwd)
    reponse := &loginData{}

    if _, err := redis.HGet("userids", numid); (err != nil && !redis.IsNil(err)) {
        // Login查询失败应该去数据库查询一次并更新redisKey
        // log.Println(err, value)
        reponse.Errcode = HTTP_USER_LOGIN_NOUSER
        reponse.Msg = notice(reponse.Errcode)
    } else {
        sql := fmt.Sprintf("Select pwd from user where numid=\"%s\" and pwd=\"%s\" limit 1", numid, pwd)
        rows, err := mysql.Select(sql)
        defer rows.Close()

        if err != nil || !rows.Next() {
            // log.Println(err, rows)
            reponse.Errcode = HTTP_USER_LOGIN_PWDERR
            reponse.Msg = notice(reponse.Errcode)
        } else {
            t, _ := redis.Get(numid + ":token")
            gateIp, _ := config.GetConfigString("gateIp")
            gatePort, _ := config.GetConfigString("gatePort")
            if t != "" {
                reponse.Errcode = HTTP_USER_LOGIN_REPEATED_OK
                reponse.Token = t
                reponse.Msg = notice(reponse.Errcode)
                reponse.GateIp = gateIp
                reponse.GatePort = gatePort

            } else {
                token := getToken()
                // token过期时间为3分钟
                if err := redis.SetAndExpire(numid + ":token", token, time.Millisecond * 1000 * 180);err == nil{
                    reponse.Errcode = HTTP_USER_LOGIN_OK
                    reponse.Token = token
                    reponse.Msg = notice(reponse.Errcode)
                    reponse.GateIp = gateIp
                    reponse.GatePort = gatePort
                }

            }
            
        }
    }
    jsonData, _ := json.Marshal(reponse)
    return jsonData
}

