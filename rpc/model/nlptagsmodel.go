package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	nlpTagsFieldNames          = builder.RawFieldNames(&NlpTags{})
	nlpTagsRows                = strings.Join(nlpTagsFieldNames, ",")
	nlpTagsRowsExpectAutoSet   = strings.Join(stringx.Remove(nlpTagsFieldNames, "`create_time`", "`update_time`"), ",")
	nlpTagsRowsWithPlaceHolder = strings.Join(stringx.Remove(nlpTagsFieldNames, "`doc_id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheNlpTagsDocIdPrefix = "cache:nlpTags:docId:"
)

type (
	NlpTagsModel interface {
		Insert(data *NlpTags) (sql.Result, error)
		FindOne(docId string) (*NlpTags, error)
		Update(data *NlpTags) error
		Delete(docId string) error
	}

	defaultNlpTagsModel struct {
		sqlc.CachedConn
		table string
	}

	NlpTags struct {
		DocId   string `db:"doc_id"`
		NlpTags string `db:"nlp_tags"` // json字符串
	}
)

func NewNlpTagsModel(conn sqlx.SqlConn, c cache.CacheConf) NlpTagsModel {
	return &defaultNlpTagsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`nlpTags`",
	}
}

func (m *defaultNlpTagsModel) Insert(data *NlpTags) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, nlpTagsRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.DocId, data.NlpTags)

	return ret, err
}

func (m *defaultNlpTagsModel) FindOne(docId string) (*NlpTags, error) {
	nlpTagsDocIdKey := fmt.Sprintf("%s%v", cacheNlpTagsDocIdPrefix, docId)
	var resp NlpTags
	err := m.QueryRow(&resp, nlpTagsDocIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `doc_id` = ? limit 1", nlpTagsRows, m.table)
		return conn.QueryRow(v, query, docId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNlpTagsModel) Update(data *NlpTags) error {
	nlpTagsDocIdKey := fmt.Sprintf("%s%v", cacheNlpTagsDocIdPrefix, data.DocId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `doc_id` = ?", m.table, nlpTagsRowsWithPlaceHolder)
		return conn.Exec(query, data.NlpTags, data.DocId)
	}, nlpTagsDocIdKey)
	return err
}

func (m *defaultNlpTagsModel) Delete(docId string) error {

	nlpTagsDocIdKey := fmt.Sprintf("%s%v", cacheNlpTagsDocIdPrefix, docId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `doc_id` = ?", m.table)
		return conn.Exec(query, docId)
	}, nlpTagsDocIdKey)
	return err
}

func (m *defaultNlpTagsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheNlpTagsDocIdPrefix, primary)
}

func (m *defaultNlpTagsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `doc_id` = ? limit 1", nlpTagsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
