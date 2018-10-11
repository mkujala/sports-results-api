package nhl

import "sports-results/standings"

type yearlyStnds struct {
	list map[string][]standings.Standings
}

// wins, loses, otwins, otloses, win%reg, win%all, ot%, gf, ga, gf/gp, ga/gp, pts, pts/gp
type nhlAverages struct {
	Season      int
	Description string
	Wins,
	Loses,
	Otwins,
	Otloses,
	WinRegP,
	WinAllP,
	OtP,
	GF,
	GA,
	GFgp,
	GAgp,
	Points,
	PTSgp float64
}
