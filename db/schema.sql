CREATE TABLE urls (
    id TEXT PRIMARY KEY,
    long_url TEXT NOT NULL,
    short_url TEXT UNIQUE NOT NULL,
    created_at TEXT NOT NULL
)