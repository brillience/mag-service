package mag

type GetReponse struct {
	Index       string   `json:"_index"`
	Type        string   `json:"_type"`
	Id          string   `json:"_id"`
	Version     int      `json:"_version"`
	SeqNo       int      `json:"_seq_no"`
	PrimaryTerm int      `json:"_primary_term"`
	Found       bool     `json:"found"`
	Source      Abstract `json:"_source"`
}
