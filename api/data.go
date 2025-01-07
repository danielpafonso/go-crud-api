package api

type Data struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func LoadData() []Data {
	return []Data{
		{
			ID:    1,
			Value: "hello mom!",
		},
		{
			ID:    2,
			Value: "Lorem ipsum dolor sit amet",
		},
	}
}

func LoadMapData() map[int]Data {

	data := LoadData()

	mapData := make(map[int]Data)

	for _, elm := range data {
		mapData[elm.ID] = elm
	}

	return mapData
}
