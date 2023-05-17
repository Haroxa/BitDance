package Repository

import (
	"bufio"
	"encoding/json"
	"os"
	"v1/Common"
	"v1/Util"
)

func StoreOne(fileName string, data interface{}) error {
	// 打开文件     // https://blog.csdn.net/nssddwbzd/article/details/123263756
	open, err := os.OpenFile(Common.DataPath+fileName, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		Util.Logger.Error("open file one err :" + err.Error())
		return err
	}

	// 创建写入器
	writer := bufio.NewWriter(open)
	text, err := json.Marshal(data)
	if err != nil {
		Util.Logger.Error("store one json Marshal err :" + err.Error())
		return err
	}
	// 添加换行符
	text = append(text, '\n')
	// 写入
	if _, err := writer.Write(text); err != nil {
		Util.Logger.Error("store one write file err :" + err.Error())
		return err
	}
	writer.Flush() // 必须要刷新缓存区才能成功写入 //https://haicoder.net/golang/golang-bufio.html
	open.Close()
	return nil
}
