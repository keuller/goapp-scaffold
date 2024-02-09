package anime

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type AnimeRepository interface {
	FetchAnimes() []Anime
}

type animeRepo struct {
	conn *sqlx.DB
}

func NewAnimeRepo() AnimeRepository {
	db, err := sqlx.Connect("sqlite3", "anime.db")
	if err != nil {
		slog.Error("fail to connect to DB", "error", err)
		return nil
	}
	return animeRepo{conn: db}
}

func (ar animeRepo) FetchAnimes() []Anime {
	var list []Anime
	ar.conn.Select(&list, "SELECT id, name, picture, complete FROM animes ORDER BY name LIMIT 100")
	return list
}
