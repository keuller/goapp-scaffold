package anime

type AnimeService interface {
	GetAll() []Anime
}

type animeService struct {
	repo AnimeRepository
}

func NewAnimeService() AnimeService {
	repo := NewAnimeRepo()
	return animeService{
		repo: repo,
	}
}

func (as animeService) GetAll() []Anime {
	list := []Anime{
		{ID: "MI5M8WHG34", Name: "Solo Leveling", Picture: "", Complete: false},
		{ID: "AC37A37EHG", Name: "Dragon Ball Z", Picture: "", Complete: true},
		{ID: "DW6IN4217B", Name: "Demon Slayer", Picture: "", Complete: false},
		{ID: "52FE8DE390", Name: "Ultraman", Picture: "", Complete: true},
	}
	return list
}
