package service

type SecondConversionEvent interface {
	Create()
	Update()
	Get()
	GetList()
}

type RepoConversionEvent struct {
	Repo SecondConversionEvent
}

func NewConversionEventService(repo SecondConversionEvent) *RepoConversionEvent {
	return &RepoConversionEvent{
		Repo: repo,
	}
}
