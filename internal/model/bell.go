package model

type BellInfo struct {
	ID           int    `json:"id"`
	Status       string `json:"-"`
	CallDate     string `json:"call_date"`
	OperatorFio  string `json:"operator_fio"`
	ClientPhone  string `json:"client_phone"`
	ContactAudio string `json:"contact_audio"`
	StatusBell   string `json:"status_bell"`
}
