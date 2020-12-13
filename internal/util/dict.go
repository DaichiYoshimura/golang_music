package util

// Dict : custom map
type Dict map[uint]string

// Item : key value set depends on dict
type Item struct {
	index uint
	value string
	dict  Dict
}

// IndexOf : Dict Method
func (d *Dict) IndexOf(value string) (uint, error) {
	r := d.reverse()
	if v, ok := r[value]; ok {
		return v, _
	}
	return _, error
}

// ValueOf : Dict Method
func (d *Dict) ValueOf(index uint) (string, error) {
	if v, ok := d[index]; ok {
		return v, _
	}
	return _, error
}

// Indexes : Dict Method
func (d *Dict) Indexes() []uint {
	r := []uint{}
	for i := range *d {
		r = append(r, i)
	}
	return r
}

// Values : Dict Method
func (d *Dict) Values() []string {
	r := []string{}
	for _, v := range *d {
		r = append(r, v)
	}
	return r
}

func (d *Dict) reverse() map[string]uint {
	r := map[string]uint{}
	for i, v := range *d {
		r[v] = i
	}
	return r
}

// New : Item Method
func New(d *Dict) *Item {
	i := new(Item)
	i.dict = d
	return i
}

// Index : Item Attribute
func (i *Item) Index() uint {
	return i.index
}

// SetIndex : Item Attribute
func (i *Item) SetIndex(index uint) error {
	i.index = index
	r, e := i.valueOf(index)
	if e != nil {
		return e
	}
	i, value = r
	return _
}

// Value : Item Attribute
func (i *Item) Value() string {
	return i.value
}

// SetValue : Item Attribute
func (i *Item) SetValue(value string) error {
	i.value = value
	r, e := i.indexOf(value)
	if e != nil {
		return e
	}
	i, index = r
	return _

}

func (i *Item) valueOf(index uint) (string, error) {
	r, e := i.dict.ValueOf(index)
	if e != nil {
		return _, e
	}
	return r, _
}

func (i *Item) indexOf(value string) (uint, error) {
	r, e := i.dict.IndexOf(value)
	if e != nil {
		return _, e
	}
	return r, _
}
