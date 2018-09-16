package standings

// GetFromDB fetches standings from DB
func GetFromDB(league, venue string, season int, conference string) ([]Standings, error) {
	stnds := []Standings{}
	var err error

	switch venue {
	case "all":
		stnds, err = allFromDB(league, season, conference)
	case "home":
		stnds, err = homeFromDB(league, season, conference)
	case "away":
		stnds, err = awayFromDB(league, season, conference)
	default:
		stnds, err = allFromDB(league, season, conference)
	}

	if err != nil {
		return nil, err
	}
	return stnds, err
}
