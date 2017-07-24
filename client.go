package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	elastic "gopkg.in/olivere/elastic.v5"
)

// TestDoc 테스트 도큐먼트 구조체
type TestDoc struct {
	First  string
	Second string
}

const indexSetting = `
{
	"settings": {
		"index": {
			"number_of_shards": 3,
			"number_of_replicas": 0
		}
	}
}
`

// ES elastic search 클라이언트 구조체
type ES struct {
	client *elastic.Client
	ctx    context.Context
}

// NewES elastic search v5 이상의 클라이언트 생성
func NewES() *ES {
	conn, err := elastic.NewClient(
		elastic.SetBasicAuth("elastic", "changeme"),
		elastic.SetSniff(false),
	)

	if err != nil {
		log.Fatal(err)
	}

	es := &ES{
		client: conn,
		ctx:    context.Background(),
	}

	return es
}

// NewTestIndex 테스트 인덱스 생성
func (e *ES) NewTestIndex() {
	exists, err := e.client.IndexExists("test").Do(e.ctx)
	if err != nil {
		log.Panic(err)
	}

	if !exists {
		_, err := e.client.CreateIndex("test").BodyString(indexSetting).Do(e.ctx)
		if err != nil {
			log.Panic(err)
		}
	}
}

// SetTestDoc 테스트 도큐먼트 생성
func (e *ES) SetTestDoc() {
	testDoc := TestDoc{First: "Hello", Second: "world"}

	_, err := e.client.Index().Index("test").
		Type("test").
		Id("1").
		BodyJson(testDoc).
		Refresh("true").
		Do(e.ctx)
	if err != nil {
		log.Panic(err)
	}
}

// QueryTest 테스트 쿼리
func (e *ES) QueryTest() {
	query := elastic.NewTermQuery("first", "Hello")
	result, err := e.client.Search().
		Index("test").
		Query(query).
		Sort("first", true).
		From(0).Size(10).
		Pretty(true).
		Do(e.ctx)
	if err != nil {
		log.Panic(err)
	}

	var testDoc TestDoc
	for _, item := range result.Each(reflect.TypeOf(testDoc)) {
		if t, ok := item.(TestDoc); ok {
			log.Printf("search result: %s", t.First)
		}
	}
}

// Ping elasticsearch 서버의 현재 상태를 검사한다.
func (e *ES) Ping() {
	info, code, err := e.client.Ping("http://127.0.0.1:9200").Do(e.ctx)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

func main() {
	e := NewES()
	e.NewTestIndex()
	e.SetTestDoc()
	e.QueryTest()
}
