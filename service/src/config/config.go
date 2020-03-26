package config

import(
	"os"
	"github.com/yuin/gopher-lua"
	"sync"
    "log"
    "strconv"
    "errors"
    "fmt"
)

var lock sync.Mutex
var luaState = lua.NewState()
var flag string

func init() {
	
}

func checkConfig() {
	SetConfig("")
}

func SetConfig(str string) {
	if flag == "" {
        lock.Lock()
        if flag == "" {
            defaultList := []string{
            	str,
            	"../config/config.lua",
            	"config/config.lua",
            	"../../config/config.lua",
            }
            var foundFile string = ""
            for _, file := range defaultList {
            	_, err := os.Stat(file)
            	if err == nil {
            		foundFile = file
            		break
            	}
            }
            if foundFile == "" {
				panic("Config file Not Found!")
            }
			if err := luaState.DoFile(foundFile); err != nil {
				panic(err)
			}
			flag = foundFile
            log.Printf("Use config:%s\n", flag)
        }
        lock.Unlock()
    }
}

func GetConfigString(key string) (string,error) {
    checkConfig()
    lv := luaState.Get(-1)
    if tbl, ok := lv.(*lua.LTable); ok {
        value := tbl.RawGetString(key)
        return value.String(), nil
    }
    err := fmt.Sprintf("Not found %s from %s\n", key, flag)
    return "", errors.New(err)
}

func GetConfigInt(key string) (int,error) {
    checkConfig()
    lv := luaState.Get(-1)
    if tbl, ok := lv.(*lua.LTable); ok {
        value := tbl.RawGetString(key)
        num, _ := strconv.Atoi(value.String())
        // str := strconv.Itoa(num)
        return num, nil
    }
    err := fmt.Sprintf("Not found %s from %s\n", key, flag)
    return 0, errors.New(err)
}


