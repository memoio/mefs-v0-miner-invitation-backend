package sqlite

import (
)

type Subscribe struct{
	Id string
	Name string
	Email string
	Role string
	Status string
}

func NewSubscribe() Subscribe {
	return Subscribe{}
}
func (this Subscribe) Insert(data Subscribe) {
	db.Exec("insert into subscribe(id,name,email,role,status) values(?,?,?,?,?)", data.Id,data.Name,data.Email,data.Role,data.Status)
}
func (this Subscribe) UpdateByEmail(data Subscribe) {
	db.Exec("update subscribe set status = ? where email = ?",data.Status,data.Email)
}
func (this Subscribe) UpdateByEmailAndRole(data Subscribe) {
	db.Exec("update subscribe set status = ? where email = ? and role = ?",data.Status,data.Email,data.Role)
}
func (this Subscribe) SelectByEmail(data Subscribe) Subscribe {
	var res Subscribe
	db.Raw("select id,name,email,role,status from subscribe where email = ?", data.Email).Scan(&res)
	return res
}
func (this Subscribe) SelectByEmailAndRole(data Subscribe) Subscribe {
	var res Subscribe
	db.Raw("select id,name,email,role,status from subscribe where email = ? and role = ?", data.Email,data.Role).Scan(&res)
	return res
}
func (this Subscribe) SelectByStatus(data Subscribe) Subscribe {
	var res Subscribe
	db.Raw("select id,name,email,role,status from subscribe where status = ?", data.Status).Scan(&res)
	return res
}