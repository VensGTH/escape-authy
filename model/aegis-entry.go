package model

type AegisBase struct {
	Version int         `json:"version"`
	Header  AegisHeader `json:"header"`
	Db      AegisDB     `json:"db"`
}

type AegisHeader struct {
	Slots  *string `json:"slots"`
	Params *string `json:"params"`
}

type AegisDB struct {
	Version int          `json:"version"`
	Entries []AegisEntry `json:"entries"`
}

type AegisEntry struct {
	Type   string    `json:"type"`
	UUID   string    `json:"uuid"`
	Name   string    `json:"name"`
	Issuer string    `json:"issuer"`
	Note   string    `json:"note"`
	Icon   *string   `json:"icon"` // Nullable
	Info   AegisInfo `json:"info"`
}

type AegisInfo struct {
	Secret string `json:"secret"`
	Algo   string `json:"algo"`
	Digits int    `json:"digits"`
	Period int    `json:"period"`
}
