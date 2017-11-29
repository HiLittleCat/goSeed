package model

import (
	"common/conn"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	collectionName = "user"
)

type User struct {
	Id     string `bson:"_id"`
	Mobile string `bson:"mobile"`
	Name   string `bson:"name"`
	Logo   string `bson:"logo"`
}

func (user *User) Get() (err error) {
	conn.GetMgoPool(conn.MgoBosh).Exec(collectionName, func(c *mgo.Collection) {
		err = c.Find(bson.M{
			"_id": bson.ObjectIdHex(user.Id),
		}).Select(bson.M{
			"mobile": 1,
			"name":   1,
			"logo":   1,
		}).One(user)
	})
	return err
}
