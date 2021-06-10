package models

type Board struct {
	BoardId      uint   `json:"boardId"`
	BoardName    string `json:"boardName"`
	BoardExplain string `json:"boardExplain"`
	BoardNo      string `json:"boardNo"`
	BoardState   uint   `json:"boardState"`
}

func (Board) Table() string {
	return "board"
}