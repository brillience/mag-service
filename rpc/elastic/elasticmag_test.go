package elastic_test

import (
	"github.com/elastic/go-elasticsearch/v7"
	"mag_service/rpc/elastic"
	"testing"
)

func TestCreateDocument(t *testing.T) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		t.Error(err.Error())
	}
	magEs, err := elastic.NewMagEs(client)
	if err != nil {
		t.Error(err.Error())
	}
	err = magEs.CreateDocument(elastic.Abstract{Id: "test", Content: "test"})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestDeleteDocument(t *testing.T) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		t.Error(err.Error())
	}
	magEs, err := elastic.NewMagEs(client)
	if err != nil {
		t.Error(err.Error())
	}
	err = magEs.DeleteDocument("test")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestUpdateDocument(t *testing.T) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		t.Error(err.Error())
	}
	magEs, err := elastic.NewMagEs(client)
	err = magEs.UpdateDocument(elastic.Abstract{Id: "test", Content: "test_update"})
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetDocumentById(t *testing.T) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		t.Error(err.Error())
	}
	magEs, err := elastic.NewMagEs(client)
	abstract, err := magEs.GetDocumentById("test")
	if err != nil {
		t.Error(err.Error())
	}
	if abstract.Id != "test" || abstract.Content != "test_update" {
		t.Error("逻辑错误！")
	}
}

func TestSearchDocumentsByKeyWord(t *testing.T) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		t.Error(err.Error())
	}
	magEs, err := elastic.NewMagEs(client)
	abstracts, err := magEs.SearchDocumentsByKeyWord("stroma*")
	if err != nil {
		t.Error(err.Error())
	}
	if len(abstracts) == 0 {
		t.Error("得到空结果！")
	}
}
