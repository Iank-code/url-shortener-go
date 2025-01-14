package model

type URL struct {
	ID        string `db:"id"`
    LongURL   string `db:"long_url"`
    ShortURL  string `db:"short_url"`
    CreatedAt string `db:"created_at"`
}