package dict

//Dict :
type Dict map[uint]string
type rDict map[string]uint

const (
	eIndex = 0
	eValue = ""
)

// IndexOf :
func (d Dict) IndexOf(value string) uint {
	r := d.reverse()
	if v,ok:= r[value];ok{
		return v
	}
	return eIndex
}

// ValueOf :
func (d Dict) ValueOf(index uint) string {
	if v,ok := d[index];ok{
		return v
	}
	return eValue
}

// Indexes :
func (d *Dict) Indexes() []uint {
	r := []uint{}
	for i := range *d {
		r = append(r,i)
	}
	return r
}

// Values : 
func (d *Dict) Values() []string {
	r := []string{}
	for _, v := range *d {
		r = append(r,v)
	}
	return r
}

func (d *Dict) reverse() map[string]uint {
	r := rDict{}
	for i, v := range *d {
		r[v] = i
	}
	return r
}
