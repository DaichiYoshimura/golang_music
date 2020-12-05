package tonality

type Tonality struct {
	tonal string
	key   string
}

func (t *Tonality) Tonal() string {
	return t.tonal
}

func (t *Tonality) SetTonal(tonal string) {
	t.tonal = tonal
}

func (t *Tonality) Key() string {
	return t.key
}

func (t *Tonality) SetKey(key string) {
	t.key = key
}