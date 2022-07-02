package _4_01_golden_file

// GenerateMusicQuery ...
func GenerateMusicQuery(keyword string) elastic.Query {
	query := elastic.NewFunctionScoreQuery()
	query.Query(
		elastic.NewMatchQuery("title", keyword),
	)

	query.Add(elastic.NewMatchPhraseQuery("artist", keyword), elastic.NewWeightFactorFunction(2))
	query.Add(elastic.NewMatchPhraseQuery("genre", keyword), elastic.NewWeightFactorFunction(10))

	query.ScoreMode("multiply")
	query.Boost(5)
	query.BoostMode("multiply")

	return query
}
