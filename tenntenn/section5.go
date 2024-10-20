package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	// ルートディレクトリを取得
	path, _ := os.Getwd() 
	dst := filepath.Join(path, "log.txt")
	println(dst)

	df, err := os.Create(dst)
	if err != nil{
		fmt.Fprintln(os.Stderr, err.Error())
	}

	//標準エラー出力を設定
	os.Stderr = df
	// 遅延実行
	defer df.Close()
	


	fmt.Fprintln(os.Stderr, "エラー")   // 標準エラー出力に出力
	fmt.Fprintln(os.Stdout, "Hello") // 標準出力に出力

	catCommand("text.txt","text2.txt")
	
}

// インデックス対応
func catCommand(filepath ...string){
	var fileTexts [][]string
	for _, v := range filepath{
		fileTexts = append(fileTexts,  readfile(v))
	}
	dumptexts(fileTexts)
}

func dumptexts(filetexts [][]string){
	index := 1
	for _,file := range filetexts{
		for _,v := range file{
			fmt.Println(strconv.Itoa(index)+ ":" + v)
			index +=1 
		}
	}
}


func readfile(filepath string) ([]string, error){
	// 読み込み
	file, err := os.Open(filepath)
	if err != nil{
		fmt.Fprintln(os.Stderr, err.Error())
	}
	defer file.Close()

	var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return lines, nil
}
