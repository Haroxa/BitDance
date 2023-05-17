package Model

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"` // 字段名必须正确
	UserID    string `json:"user_id"`
}
