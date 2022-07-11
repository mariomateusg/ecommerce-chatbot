// AgentBot.
//
// API that exposes the functionality of the agent. It has 2 methods. Greeting, used when the conversation begins, and reply that receives the query and returns the answer https://osmandi.now.sh/post/simpsongo
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0
//
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

package bot

import (
	"ecommerce-bot/internal/data"
	"ecommerce-bot/knowledgebase"
	"ecommerce-bot/samplequestion"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
)

type AgentCase struct {
	Bot
	name            string
	title           string
	thumbnailPath   string
	knowledgeBase   map[string]string
	knowledgeCorpus []string
	sampleQuestions []string
	Repository      knowledgebase.Repository
	RepositorySQ    samplequestion.Repository
}

func NewAgentCase() *AgentCase {
	// connection to the database.
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}
	agentCase := &AgentCase{name: "Bot", title: "Ecommer Bot", thumbnailPath: "/static/images/chat/Case.png", Repository: &data.KnowledgeBaseRepository{Data: d}, RepositorySQ: &data.SampleQuestionRepository{Data: d}}
	agentCase.initializeIntelligence()
	return agentCase
}

func (a *AgentCase) Name() string {
	return a.name
}

func (a *AgentCase) Title() string {
	return a.title
}

func (a *AgentCase) ThumbnailPath() string {
	return a.thumbnailPath
}

func (a *AgentCase) SetName(name string) {
	a.name = name
}

func (a *AgentCase) SetTitle(title string) {
	a.title = title
}

func (a *AgentCase) SetThumbnailPath(thumbnailPath string) {
	a.thumbnailPath = thumbnailPath
}

func randomNumber(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func (a *AgentCase) Greetings() string {

	sampleQuestionIndex := randomNumber(0, len(a.sampleQuestions))
	greeting := "Hi there! I'm Case. You can ask me questions such as " + a.sampleQuestions[sampleQuestionIndex] + "\""
	return greeting

}

func (a *AgentCase) Reply(query string) string {

	var result string

	vectoriser := nlp.NewCountVectoriser()
	transformer := nlp.NewTfidfTransformer()

	reducer := nlp.NewTruncatedSVD(4)

	matrix, _ := vectoriser.FitTransform(a.knowledgeCorpus...)
	matrix, _ = transformer.FitTransform(matrix)
	lsi, _ := reducer.FitTransform(matrix)

	matrix, _ = vectoriser.Transform(query)
	matrix, _ = transformer.Transform(matrix)
	queryVector, _ := reducer.Transform(matrix)

	highestSimilarity := -1.0
	var matched int
	_, docs := lsi.Dims()
	for i := 0; i < docs; i++ {
		similarity := pairwise.CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), lsi.(mat.ColViewer).ColView(i))
		if similarity > highestSimilarity {
			matched = i
			highestSimilarity = similarity
		}
	}

	if highestSimilarity == -1 {
		result = "I don't know the answer to that one."
	} else {
		result = a.knowledgeBase[a.knowledgeCorpus[matched]]
	}

	return result

}

func (a *AgentCase) initializeIntelligence() {

	knowledgeBases, err := a.Repository.GetAll()
	a.knowledgeBase = make(map[string]string)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(knowledgeBases); i++ {
		//fmt.Println(knowledgeBases[i])
		a.knowledgeBase[knowledgeBases[i].Topic] = knowledgeBases[i].Answers
	}

	fmt.Println(a.knowledgeBase)

	a.knowledgeCorpus = make([]string, 1)
	for k, _ := range a.knowledgeBase {
		a.knowledgeCorpus = append(a.knowledgeCorpus, k)
	}

	samplequestions, err := a.RepositorySQ.GetAll()

	a.sampleQuestions = make([]string, 1)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(samplequestions); i++ {
		a.sampleQuestions = append(a.sampleQuestions, samplequestions[i].Question)

	}

	fmt.Println("Samples ", a.sampleQuestions)

	//a.sampleQuestions = []string{"What is isomorphic go?", "What are the benefits of this technology?", "Does isomorphic go offer anything react-like?", "How can I recompile code instantly?", "How can I perform routing in my web app?", "Where can I get starter code?", "Where can I find a talk on this topic?"}

}
