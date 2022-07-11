package knowledgebase

type KnowledgeBase struct {
	ID      uint   `json:"id,omitempty"`
	Topic   string `json:"topic,omitempty"`
	Answers string `json:"answers,omitempty"`
}
