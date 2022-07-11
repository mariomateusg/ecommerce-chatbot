package knowledgebase

type Repository interface {
	GetAll() ([]KnowledgeBase, error)
}
