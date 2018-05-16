package model

import (
	"github.com/HiLittleCat/conn"
	"github.com/HiLittleCat/core"
	"github.com/HiLittleCat/goSeed/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mgoName        = config.Default.MongoDB.DatebaseName
	collectionName = "user"
	errDBName      = "mongodb." + mgoName + "." + collectionName
)

func dbError(err error) error {
	if err != nil {
		return (&core.DBError{}).New(errDBName, err.Error())
	}
	return nil
}

type User struct {
	core.Model
	ID     string `bson:"_id"`
	Mobile string `bson:"mobile"`
	Name   string `bson:"name"`
	Logo   string `bson:"logo"`
}

func (user *User) Create() (err error) {
	conn.GetMgoPool(mgoName).Exec(collectionName, func(c *mgo.Collection) {
		err = c.Insert(bson.M{
			"mobile": user.Mobile,
			"name":   user.Name,
			"logo":   user.Logo,
		})
	})
	return dbError(err)
}

func (user *User) GetByID() (err error) {
	conn.GetMgoPool(mgoName).Exec(collectionName, func(c *mgo.Collection) {
		err = c.Find(bson.M{
			"_id": bson.ObjectIdHex(user.ID),
		}).Select(bson.M{
			"mobile": 1,
			"name":   1,
			"logo":   1,
		}).One(user)
	})
	return dbError(err)
}

func (user *User) GetCountByID() (count int, err error) {
	conn.GetMgoPool(mgoName).Exec(collectionName, func(c *mgo.Collection) {
		count, err = c.Find(bson.M{
			"_id": bson.ObjectIdHex(user.ID),
		}).Count()
	})
	return count, dbError(err)
}

type UserList struct {
	core.Model
	Page      int
	PageCount int
	List      []User
}

func (list *UserList) GetPage() (err error) {
	conn.GetMgoPool(mgoName).Exec(collectionName, func(c *mgo.Collection) {
		err = c.Find(nil).Select(bson.M{
			"mobile": 1,
			"name":   1,
			"logo":   1,
		}).Skip((list.Page - 1) * list.PageCount).Limit(list.PageCount).All(&list.List)
	})
	return dbError(err)
}
