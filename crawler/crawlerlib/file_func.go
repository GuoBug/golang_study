package crawlerlib

import (
	"os"
	"io"
	"fmt"
)


func CheckFileIsExist(fileName string) (bool) {
    var exist = true;
    if _, err := os.Stat(fileName); os.IsNotExist(err) {
        exist = false; 
    }
    return exist;
}

func CreateFile(fileName string) {
    fp, err := os.Create(fileName)
    if err != nil {
        panic(err)
    }
    defer fp.Close()
}

func RecordInfo(fileName string ,record string) {

//有则追加
    fmt.Printf("\n*************True****************\n")
    fp, err := os.OpenFile(fileName, os.O_APPEND, 0666)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    n, writeErr := io.WriteString(fp, record)
    if writeErr != nil {
        panic (writeErr)
    }
    defer fp.Close()
    fmt.Printf("本次输入 **%d** 个字节,%s", n,record)

    fp.Close()
}
