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
    // log.Println(1)
    if instanceConn == nil {
        // log.Println(2)
        lock.Lock()
        // log.Println(3)
        if instanceConn == nil {
            // log.Println(4)
            addr := config.GetString("mysqlUser") + ":" +
                    config.GetString("mysqlPwd") + "@tcp(" +
                    config.GetString("mysqlIp") + ":" +
                    config.GetString("mysqlPort") + ")/" +
                    config.GetString("mysqlDb")
            // log.Println(5, addr)

            db, err := sql.Open("mysql", addr)
            // log.Println(6, err)
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