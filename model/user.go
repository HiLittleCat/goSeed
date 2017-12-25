package model

import (
	"github.com/HiLittleCat/goSeed/conn"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	collectionName = "user"
)

type User struct {
	ID     string `bson:"_id"`
	Mobile string `bson:"mobile"`
	Name   string `bson:"name"`
	Logo   string `bson:"logo"`
}

func (user *User) Create() (err error) {
	conn.GetMgoPool(conn.MgoBosh).Exec(collectionName, func(c *mgo.Collection) {
		err = c.Insert(bson.M{
			"mobile": user.Mobile,
			"name":   user.Name,
			"logo":   user.Logo,
		})
	})
	return err
}

func (user *User) Get() (err error) {
	conn.GetMgoPool(conn.MgoBosh).Exec(collectionName, func(c *mgo.Collection) {
		err = c.Find(bson.M{
			"_id": bson.ObjectIdHex(user.ID),
		}).Select(bson.M{
			"mobile": 1,
			"name":   1,
			"logo":   1,
		}).One(user)
	})
	return err
}

type UserList struct {
	Page      int
	PageCount int
	List      []User
}

func (list *UserList) GetPage() (err error) {
	conn.GetMgoPool(conn.MgoBosh).Exec(collectionName, func(c *mgo.Collection) {
		err = c.Find(nil).Select(bson.M{
			"mobile": 1,
			"name":   1,
			"logo":   1,
		}).Skip((list.Page - 1) * list.PageCount).Limit(list.PageCount).All(&list.List)
	})
	return err
}
