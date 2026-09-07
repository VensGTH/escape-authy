package model

type ProtonBase struct {
	Version int         	`json:"version"`
	Entries []ProtonEntry 	`json:"entries"`
}

type ProtonEntry struct {
	Id   		string    		`json:"id"`
	Content   	ProtonContent   `json:"content"`
	Note   		*string    		`json:"note"`
}

type ProtonContent struct {
	Uri 		string `json:"uri"`
	EntryType   string `json:"entry_type"`
	Name   		string `json:"name"`
}