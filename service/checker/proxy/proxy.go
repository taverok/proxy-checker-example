package proxy

type Proxy struct {
	ID       int    `db:"id"`
	IP       string `db:"ip"`
	Protocol string `db:"protocol"`
	Domain   string `db:"domain"`
	Port     int    `db:"port"`
	Login    string `db:"login"`
	Password string `db:"password"`
	Location string `db:"location"`
}
