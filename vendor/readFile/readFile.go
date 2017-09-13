package readFile

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type ReadFile struct{}

var (
	file     []interface{}
	dataPath = "data/algs4-data/"
)

// new readFile
func New() *ReadFile {
	return &ReadFile{}
}

// get data path
func (r *ReadFile) GetDataPath() (string, error) {
	filePath, err := os.Getwd()
	if err != nil {
		return filePath, err
	}
	end := strings.Index(filePath, "algs") + 5
	return filePath[:end] + dataPath, nil

}

// get bufio reader
func (r *ReadFile) read(path string) error {
	file = []interface{}{}
	inputFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	inputRead := bufio.NewReader(inputFile)
	for {
		line, err := inputRead.ReadString('\n')
		if err == io.EOF {
			return nil
		} else {
			file = append(file, line)
		}
	}
}

// read file into []int
func (r *ReadFile) ReadToInt(path string, args *[]int) error {
	err := r.read(path)
	if err != nil {
		return err
	}
	for _, v := range file {
		vStr := v.(string)
		vStr = strings.Replace(vStr, "\n", "", -1) // 去除换行
		vStr = strings.Replace(vStr, " ", "", -1)  // 去除空格
		lineInt, _ := strconv.Atoi(vStr)
		*args = append(*args, lineInt)
	}
	return nil
}

// read file into []string
func (r *ReadFile) ReadToString(path string, args *[]string) error {
	err := r.read(path)
	if err != nil {
		return err
	}
	for _, v := range file {
		vStr := v.(string)
		vStr = strings.Replace(vStr, "\n", "", -1) // 去除换行
		vStr = strings.Replace(vStr, " ", "", -1)  // 去除空格
		*args = append(*args, vStr)
	}
	return nil
}
