package km

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"time"
)

func ParseTime(str string) time.Time {
	res, err := time.ParseInLocation("2006-01-02 15:04:05.000", str, time.Local)
	if err != nil {
		return time.Time{}
	}
	return res
}

func formatTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05.000")
}

// LogStat 统计文件
func LogStat(t string, start time.Time, end time.Time) map[string]interface{} {

	file := GetLogFile(t, formatTime(start)[0:10])
	reader := bufio.NewReader(file)
	// 循环读取文件的内容
	stat := map[string]int{}
	total := 0
	for {
		// 读到一个换行符就结束
		line, err := reader.ReadString('\n')
		if "" == line || err == io.EOF { // io.EOF表示文件的末尾
			break
		}
		if len(line) < 24 {
			continue
		}
		t := ParseTime(line[0:23])
		if t.After(start) && t.Before(end) {
			key := line[24 : len(line)-1]
			value := stat[key]
			stat[key] = value + 1
			total++
		}

	}
	result := map[string]interface{}{
		"stat":  stat,
		"total": total,
	}
	return result
}

// GetLogFile 获取日志文件
func GetLogFile(t string, day string) *os.File {
	usr, err := user.Current()
	fileDir := fmt.Sprintf("%s\\.kmgo", usr.HomeDir)
	_, err = os.Stat(fileDir)
	if os.IsNotExist(err) {
		fmt.Println(err.Error())
		os.MkdirAll(fileDir, os.ModePerm)
	}

	filePath := fmt.Sprintf("%s\\%s%s.log", fileDir, t, day)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) // O_CREATE 能创建文件
	if err != nil {
		fmt.Println(err.Error())
	}
	return file
}

func LogToFile(t string, str string) {

	now := formatTime(time.Now())
	file := GetLogFile(t, now[0:10])
	// 及时关闭文件句柄
	defer file.Close()
	// 准备写入的内容
	// 写入时，使用带缓冲方式的 bufio.NewWriter(w io.Writer) *Writer
	writer := bufio.NewWriter(file)
	content := fmt.Sprintf("%s\u0001%s\n", now, str)
	fmt.Print(content)
	if _, err := writer.WriteString(content); err != nil {
		fmt.Printf("%s", err.Error())
	}

	// 因为 writer 是带缓存的，所以需要 Flush 方法将缓冲中的数据真正写入到文件中
	_ = writer.Flush()
}
