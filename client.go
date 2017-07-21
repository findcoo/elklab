package main

import (
	"context"
	"log"
	"reflect"

	elastic "gopkg.in/olivere/elastic.v5"
)

// TestDoc 테스트 도큐먼트 구조체
type TestDoc struct {
	First  string
	Second string
}

// ES elastic search 클라이언트 구조체
type ES struct {
	client *elastic.Client
	ctx    context.Context
}

// NewES elastic search v5 이상의 클라이언트 생성
func NewES() *ES {
	conn, err := elastic.NewClient()
	if err != nil {
		log.Fatal()
	}

	es := &ES{
		client: conn,
		ctx:    context.Background(),
	}

	return es
}

// NewTestIndex 테스트 인덱스 생성
func (e *ES) NewTestIndex() {
	_, err := e.client.CreateIndex("test").Do(e.ctx)
	if err != nil {
		log.Panic(err)
	}
}

// SetTestDoc 테스트 도큐먼트 생성
func (e *ES) SetTestDoc() {
	testDoc := TestDoc{First: "Hello", Second: "world"}

	_, err := e.client.Index().Index("testDoc1").
		Type("test").
		Id("1").
		BodyJson(testDoc).
		Refresh(true).
		Do(e.ctx)
	if err != nil {
		log.Panic(err)
	}
}

// QueryTest 테스트 쿼리
func (e *ES) QueryTest(params) {
	query := elastic.NewTermQuery("first", "Hello")
	result, err := e.client.Search().
		Index(testDoc1).
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

func main() {
	es := NewES()
}
