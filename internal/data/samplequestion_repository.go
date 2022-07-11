package data

import "ecommerce-bot/samplequestion"

type SampleQuestionRepository struct {
	Data *Data
}

func (sr *SampleQuestionRepository) GetAll() ([]samplequestion.SampleQuestion, error) {

	q := `SELECT id, question  FROM samplequestion ;`

	rows, err := sr.Data.DB.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var samplequestions []samplequestion.SampleQuestion

	for rows.Next() {
		var s samplequestion.SampleQuestion
		rows.Scan(&s.ID, &s.Question)
		samplequestions = append(samplequestions, s)
	}

	return samplequestions, nil

}
