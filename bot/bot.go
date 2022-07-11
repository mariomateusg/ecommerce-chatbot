package bot

type Bot interface {
	Greetings() string
	Reply(string) string
	Name() string
	Title() string
	ThumbnailPath() string
	SetName(string)
	SetTitle(string)
	SetThumbnailPath(string)
}
