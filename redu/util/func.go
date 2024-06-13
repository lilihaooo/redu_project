package util

import (
	"fmt"
	"os"
	"path/filepath"
	"redu/global"
	"time"
)

func GetOffset(currentPage, pageSize int) int {
	offset := (currentPage - 1) * pageSize
	if offset < 0 {
		return 0
	}
	return offset
}

// TimeParse layout "2006-01-02 15:04:05"; "2006-01-02";
func TimeParse(timeStr string, layout string) (t time.Time) {
	// 解析时间
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		global.Logrus.Error(err)
		return time.Time{}
	}
	return t
}

// TimeToString layout "2006-01-02 15:04:05"; "2006-01-02"
func TimeToString(t time.Time, layout string) string {
	return t.Format(layout)
}

// 目录相关

// FindProjectRoot 默认go.mod同级目录为项目目录
func FindProjectRoot() (string, error) {
	// 从当前工作目录开始搜索
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// 递归地向上搜索go.mod文件
	for {
		// 检查当前目录是否有go.mod文件
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			// 找到go.mod文件，返回当前目录作为项目根目录
			return dir, nil
		}
		// 如果没有找到go.mod文件，尝试向上移动到父目录
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// 到达了文件系统的根目录，但没有找到go.mod文件
			return "", fmt.Errorf("could not find project root (go.mod file)")
		}
		dir = parentDir
	}
}
