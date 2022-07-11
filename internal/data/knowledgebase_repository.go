package data

import (
	"ecommerce-bot/knowledgebase"
)

type KnowledgeBaseRepository struct {
	Data *Data
}

func (kb *KnowledgeBaseRepository) GetAll() ([]knowledgebase.KnowledgeBase, error) {
	q := `SELECT id, topic , answers  FROM knowledgebase ;`

	rows, err := kb.Data.DB.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var knowledgeBases []knowledgebase.KnowledgeBase
	for rows.Next() {
		var k knowledgebase.KnowledgeBase
		rows.Scan(&k.ID, &k.Topic, &k.Answers)
		knowledgeBases = append(knowledgeBases, k)
	}

	return knowledgeBases, nil

}
