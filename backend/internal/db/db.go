package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type DB struct{ Conn *sql.DB }

func New(url string) (*DB, error) {
    if url == "" {
        url = "file:perimeter.db?_foreign_keys=1"
    }
    conn, err := sql.Open("sqlite3", url)
    if err != nil {
        return nil, err
    }
    schema := `CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY AUTOINCREMENT, site TEXT, type TEXT, payload TEXT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP);`
    if _, err := conn.Exec(schema); err != nil {
        return nil, err
    }
    return &DB{Conn: conn}, nil
}

func (db *DB) Close() error { return db.Conn.Close() }

func (db *DB) InsertEvent(site, etype, payload string) error {
    _, err := db.Conn.Exec("INSERT INTO events(site,type,payload) VALUES(?,?,?)", site, etype, payload)
    return err
}

func (db *DB) ListEvents(limit int) ([]map[string]interface{}, error) {
    rows, err := db.Conn.Query("SELECT id,site,type,payload,created_at FROM events ORDER BY created_at DESC LIMIT ?", limit)
    if err != nil { return nil, err }
    defer rows.Close()
    var out []map[string]interface{}
    for rows.Next() {
        var id int
        var site, t, payload, created string
        rows.Scan(&id, &site, &t, &payload, &created)
        out = append(out, map[string]interface{}{"id": id, "site": site, "type": t, "payload": payload, "created_at": created})
    }
    return out, nil
}
