package mock

type DataSet interface {
	Find(id uint8, name string) *Data
	Add(data *Data) *dataSet
	Delete(id uint8, name string) *dataSet
}

type dataSet struct {
	dataSet []*Data
}

type Data struct {
	id      uint8
	name    string
	address string
}

func InitData() DataSet {
	return &dataSet{
		dataSet: []*Data{
			&Data{1, "taro", "tokyo"},
			&Data{2, "jiro", "osaka"},
			&Data{3, "saburo", "okinawa"},
			&Data{4, "shiro", "hokkaido"},
			&Data{5, "goro", "kyoto"},
		},
	}
}

func NewData() DataSet {
	return &dataSet{
		dataSet: make([]*Data, 0),
	}
}

func (d *dataSet) Find(id uint8, name string) *Data {
	for _, v := range d.dataSet {
		if v.id == id && v.name == name {
			return v
		}
	}
	return nil
}

func (d *dataSet) Add(data *Data) *dataSet {
	d.dataSet = append(d.dataSet, data)
	return d
}

func (d *dataSet) Delete(id uint8, name string) *dataSet {
	res := make([]*Data, 0)
	for _, v := range d.dataSet {
		if v.id != id || v.name != name {
			res = append(res, v)
		}
	}
	return &dataSet{dataSet: res}
}
