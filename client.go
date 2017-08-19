package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	elastic "gopkg.in/olivere/elastic.v5"
)

// TestDoc 테스트 도큐먼트 구조체
type TestDoc struct {
	First  string
	Second string
}

var indexSetting = readMapping()

// ES elastic search 클라이언트 구조체
type ES struct {
	client *elastic.Client
	ctx    context.Context
}

type TestKoreanDoc struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Message  string    `json:"message"`
	Address  string    `json:"address"`
	Location []float64 `json:"location"`
	Phone    string    `json:"phone"`
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
	doc := TestKoreanDoc{
		ID:       1,
		Message:  "한글 korean 테스트",
		Name:     "정의성",
		Address:  "서울특별시 한글동 1 번지",
		Location: []float64{0.1, 0.22},
		Phone:    "000-0000-0000",
	}

	_, err := e.client.Index().Index("test").
		Type("test").
		Id("1").
		BodyJson(doc).
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

func readMapping() string {
	buff, err := ioutil.ReadFile("./mapping.json")
	if err != nil {
		log.Fatal(err)
	}
	return string(buff)
}

func main() {
	e := NewES()
	e.NewTestIndex()
	e.SetTestDoc()
	e.QueryTest()
}
