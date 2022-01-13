package mag

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

type Abstract struct {
	Id      string `json:"docid"`
	Content string `json:"abstract"`
}

func (a *Abstract) String() string {
	return fmt.Sprintf("Abstract: id->{%s}, content->{%s};", a.Id, a.Content)
}

func LoadCsv(csvPath string) chan Abstract {
	//创建一个大小为128的带缓冲通道
	session := make(chan Abstract, 128)
	if path.Ext(csvPath) != ".csv" {
		log.Fatalln("You should use .csv format file!")
	}
	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatalln("open csv file err: ", err.Error())
	}
	reader := csv.NewReader(file)
	//采用生成器模式
	go func() {
		//流式读入
		for {
			line, err := reader.Read()
			if err == io.EOF {
				log.Printf("file: %s read over! \n", csvPath)
				break
			}
			if err != nil {
				log.Fatalln("csv read error ：", err.Error())
			}
			session <- Abstract{Id: line[0], Content: line[1]}
		}
		close(session)
	}()
	return session
}
