package mysql

import(
	"sync"
	"config"
	"log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var lock sync.Mutex
var instanceConn * sql.DB

func instance() *  sql.DB{
	if instanceConn == nil {
		lock.Lock()
		if instanceConn == nil {
            mysqlUser, err := config.GetConfigString("mysqlUser")
            if err != nil{ panic(err) }
            mysqlPwd, err := config.GetConfigString("mysqlPwd")
            if err != nil{ panic(err) }
            mysqlIp, err := config.GetConfigString("mysqlIp")
            if err != nil{ panic(err) }
            mysqlPort, err := config.GetConfigString("mysqlPort")
            if err != nil{ panic(err) }
            mysqlDb, err := config.GetConfigString("mysqlDb")
            if err != nil{ panic(err) }
            addr := mysqlUser + ":" + mysqlPwd + "@tcp(" + mysqlIp + ":" + mysqlPort + ")/" + mysqlDb

			db, err := sql.Open("mysql", addr)
    		if err == nil{
    			instanceConn = db
    		} else {
                log.Println(err)
            }
		}
		lock.Unlock()
	}
    if instanceConn == nil{
        panic("mysql init failed!")
    }
	return instanceConn
}

func Select(str string) (*sql.Rows, error) {
    log.Println(str)
	rows, err := instance().Query(str);
    if err != nil {
        // log.Fatal(err)
        return nil, err
    }
    return rows, nil
    // defer rows.Close()
    // for rows.Next() {
    //     err := rows.Scan(&id, &name)
    //     if err != nil {
    //         log.Fatal(err)
    //     }
    //     log.Println(id, name)
    // }
    // err = rows.Err()
    // if err != nil {
    //     log.Fatal(err)
    // }
}

func Insert(str string) (int, error) {
	res, err:= instance().Exec(str)
    if err != nil {
    	return 0, err
        // log.Fatal(err)
    }
    count, _ := res.RowsAffected()
    return int(count), nil
    // if err != nil {
    //     log.Printf("res.RowsAffected() returned error: %s", err.Error())
    // }
    // log.Printf("expected 1 affected row, got %d", count)
}