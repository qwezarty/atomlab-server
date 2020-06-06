package models

type Titanic struct {
	Sex              float64 `json:"sex,string" form:"sex"`
	Age              float64 `json:"age,string" form:"age"`
	NSiblingsSpouses float64 `json:"n_siblings_spouses,string" form:"n_siblings_spouses"`
	Parch            float64 `json:"parch,string" form:"parch"`
	Fare             float64 `json:"fare,string" form:"fare"`
	Class            float64 `json:"class,string" form:"class"`
	Deck             float64 `json:"deck,string" form:"deck"`
	EmbarkTown       float64 `json:"embark_town,string" form:"embark_town"`
	Alone            float64 `json:"alone,string" form:"alone"`
}

func Titanic2Slice(t *Titanic) []float64 {
	// we have to keep the order correctly
	ret := make([]float64, 0)
	ret = append(ret, t.Sex)
	ret = append(ret, t.Age)
	ret = append(ret, t.NSiblingsSpouses)
	ret = append(ret, t.Parch)
	ret = append(ret, t.Fare)
	ret = append(ret, t.Class)
	ret = append(ret, t.Deck)
	ret = append(ret, t.EmbarkTown)
	ret = append(ret, t.Alone)

	return ret
}
