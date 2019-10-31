package sqlite

import (
)

type Dict struct{
	Key,Value,Desc string
}

func NewDict() Dict {
	return Dict{}
}
func (this Dict) Get(key string) string {
	var res Dict
	db.Raw("select key,value,desc from dict where key = ?", key).Scan(&res)
	return res.Value
}