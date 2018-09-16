package standings

// GetFromDB fetches standings from DB
func GetFromDB(league, venue string, season int) ([]Standings, error) {
	stnds := []Standings{}
	var err error

	switch venue {
	case "all":
		stnds, err = allFromDB(league, season)
	case "home":
		stnds, err = homeFromDB(league, season)
	case "away":
		stnds, err = awayFromDB(league, season)
	default:
		stnds, err = allFromDB(league, season)
	}

	// stnds, err := allFromDB(league, venue, season)

	if err != nil {
		return nil, err
	}
	return stnds, err
}
