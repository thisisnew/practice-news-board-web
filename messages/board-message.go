package messages

const (
	DAY_WRITE_LIMIT = 5
)

type Board struct {
	BoardId      string `json:"boardId"`
	BoardName    string `json:"boardName"`
	BoardExplain string `json:"boardExplain"`
	BoardNo      string `json:"boardNo"`
	BoardState   uint   `json:"boardState"`
}

type BoardList struct {
	Items   []Board
	IsLimit bool
	News    []News
	User    User
}

type BoardDetail struct {
	Board    Board
	Comments []Comment
}

type News struct {
	By          string
	Descendants uint
	Id          string
	Kids        []uint
	Parts       []uint
	Score       uint
	Time        uint
	Title       string
	Type        string
}
