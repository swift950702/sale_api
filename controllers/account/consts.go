package account

type Rps struct {
	Errcode int         `json:"errcode"`
	Data    interface{} `json:"data,omitempty"`
}
