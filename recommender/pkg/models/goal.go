package models

type Goal struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Completed bool   `json:"completed"`
	Content   string `json:"content"`
	Extended  string `json:"extended"`
}

func (g *Goal) SetExtended(extended string) {
	g.Extended = extended
}
