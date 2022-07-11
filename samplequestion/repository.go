package samplequestion

type Repository interface {
	GetAll() ([]SampleQuestion, error)
}
