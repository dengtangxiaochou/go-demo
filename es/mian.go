package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	//初始化连接得到一个Clinet
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")
	p1 := student{
		Name:    "Rion1",
		Age:     22,
		Married: false,
	}
	//链式操作

	put1, err := client.Index().
		Index("student").
		Type("go").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}