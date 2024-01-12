package proxy

import (
	"database/sql"
	"errors"
)

type Repo struct {
	DB *sql.DB
}

func (it *Repo) GetByHost(host string) (*Proxy, error) {
	var p Proxy
	query := `SELECT id, ip, protocol, domain, port, login, password, location 
			FROM proxy
			WHERE domain = ? AND alive = 1 AND active = 1 LIMIT 1`
	err := it.DB.QueryRow(query, host).
		Scan(&p.ID, &p.IP, &p.Protocol, &p.Domain, &p.Port, &p.Login, &p.Password, &p.Location)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}
