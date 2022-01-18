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
	nlpTagsRowsWithPlaceHolder = strings.Join(stringx.Remove(nlpTagsFieldNames, "`doc_id_sent_index`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheNlpTagsDocIdSentIndexPrefix     = "cache:nlpTags:docIdSentIndex:"
	cacheNlpTagsDocIdSentenceIndexPrefix = "cache:nlpTags:docId:sentenceIndex:"
	cacheNlpTagsDocIdPrefix              = "cache:nlpTags:docId:"
)

type (
	NlpTagsModel interface {
		Insert(data *NlpTags) (sql.Result, error)
		FindOne(docIdSentIndex string) (*NlpTags, error)
		FindOneByDocIdSentenceIndex(docId string, sentenceIndex int64) (*NlpTags, error)
		FindMoreByDocId(docId string) ([]NlpTags, error)
		Update(data *NlpTags) error
		Delete(docIdSentIndex string) error
	}

	defaultNlpTagsModel struct {
		sqlc.CachedConn
		table string
	}

	NlpTags struct {
		DocIdSentIndex string `db:"doc_id_sent_index"`
		DocId          string `db:"doc_id"`
		SentenceIndex  int64  `db:"sentence_index"`
		SentenceText   string `db:"sentence_text"`
		Tokens         string `db:"tokens"`
		Lemmas         string `db:"lemmas"`
		PosTags        string `db:"pos_tags"`
		NerTags        string `db:"ner_tags"`
		DocOffsets     string `db:"doc_offsets"`
		DepTypes       string `db:"dep_types"`
		DepTokens      string `db:"dep_tokens"`
	}
)

func NewNlpTagsModel(conn sqlx.SqlConn, c cache.CacheConf) NlpTagsModel {
	return &defaultNlpTagsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`nlpTags`",
	}
}

func (m *defaultNlpTagsModel) Insert(data *NlpTags) (sql.Result, error) {
	nlpTagsDocIdSentIndexKey := fmt.Sprintf("%s%v", cacheNlpTagsDocIdSentIndexPrefix, data.DocIdSentIndex)
	nlpTagsDocIdSentenceIndexKey := fmt.Sprintf("%s%v:%v", cacheNlpTagsDocIdSentenceIndexPrefix, data.DocId, data.SentenceIndex)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, nlpTagsRowsExpectAutoSet)
		return conn.Exec(query, data.DocIdSentIndex, data.DocId, data.SentenceIndex, data.SentenceText, data.Tokens, data.Lemmas, data.PosTags, data.NerTags, data.DocOffsets, data.DepTypes, data.DepTokens)
	}, nlpTagsDocIdSentIndexKey, nlpTagsDocIdSentenceIndexKey)
	return ret, err
}

func (m *defaultNlpTagsModel) FindOne(docIdSentIndex string) (*NlpTags, error) {
	nlpTagsDocIdSentIndexKey := fmt.Sprintf("%s%v", cacheNlpTagsDocIdSentIndexPrefix, docIdSentIndex)
	var resp NlpTags
	err := m.QueryRow(&resp, nlpTagsDocIdSentIndexKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `doc_id_sent_index` = ? limit 1", nlpTagsRows, m.table)
		return conn.QueryRow(v, query, docIdSentIndex)
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

func (m *defaultNlpTagsModel) FindOneByDocIdSentenceIndex(docId string, sentenceIndex int64) (*NlpTags, error) {
	nlpTagsDocIdSentenceIndexKey := fmt.Sprintf("%s%v:%v", cacheNlpTagsDocIdSentenceIndexPrefix, docId, sentenceIndex)
	var resp NlpTags
	err := m.QueryRowIndex(&resp, nlpTagsDocIdSentenceIndexKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `doc_id` = ? and `sentence_index` = ? limit 1", nlpTagsRows, m.table)
		if err := conn.QueryRow(&resp, query, docId, sentenceIndex); err != nil {
			return nil, err
		}
		return resp.DocIdSentIndex, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultNlpTagsModel) FindMoreByDocId(docId string) ([]NlpTags, error) {
	// 性能有待测试；当数据量大的时候，估计这里会查询很慢。
	// TODO: 确认该API
	var resp []NlpTags
	query := fmt.Sprintf("select %s from %s where `doc_id` = ?", nlpTagsRows, m.table)
	err := m.QueryRowsNoCache(&resp, query, docId)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultNlpTagsModel) Update(data *NlpTags) error {
	nlpTagsDocIdSentIndexKey := fmt.Sprintf("%s%v", cacheNlpTagsDocIdSentIndexPrefix, data.DocIdSentIndex)
	nlpTagsDocIdSentenceIndexKey := fmt.Sprintf("%s%v:%v", cacheNlpTagsDocIdSentenceIndexPrefix, data.DocId, data.SentenceIndex)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `doc_id_sent_index` = ?", m.table, nlpTagsRowsWithPlaceHolder)
		return conn.Exec(query, data.DocId, data.SentenceIndex, data.SentenceText, data.Tokens, data.Lemmas, data.PosTags, data.NerTags, data.DocOffsets, data.DepTypes, data.DepTokens, data.DocIdSentIndex)
	}, nlpTagsDocIdSentIndexKey, nlpTagsDocIdSentenceIndexKey)
	return err
}

func (m *defaultNlpTagsModel) Delete(docIdSentIndex string) error {
	data, err := m.FindOne(docIdSentIndex)
	if err != nil {
		return err
	}

	nlpTagsDocIdSentIndexKey := fmt.Sprintf("%s%v", cacheNlpTagsDocIdSentIndexPrefix, docIdSentIndex)
	nlpTagsDocIdSentenceIndexKey := fmt.Sprintf("%s%v:%v", cacheNlpTagsDocIdSentenceIndexPrefix, data.DocId, data.SentenceIndex)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `doc_id_sent_index` = ?", m.table)
		return conn.Exec(query, docIdSentIndex)
	}, nlpTagsDocIdSentIndexKey, nlpTagsDocIdSentenceIndexKey)
	return err
}

func (m *defaultNlpTagsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheNlpTagsDocIdSentIndexPrefix, primary)
}

func (m *defaultNlpTagsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `doc_id_sent_index` = ? limit 1", nlpTagsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
