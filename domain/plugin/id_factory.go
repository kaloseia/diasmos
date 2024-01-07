package plugin

func NewID(key string) (ID, error) {
	id := ID{}
	parseErr := id.FromKey(key)
	if parseErr != nil {
		return ID{}, parseErr
	}
	return id, nil
}
