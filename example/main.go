/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-26 19:10
**/

package main

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Lemo-yxk/longo"
)

func main() {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"

	mgo, _ := longo.NewClient().Connect(&longo.Config{Url: url})

	err := mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

	// Transaction can not create collection, so you have to create it before you run.
	// mgo.Transaction(func(sessionContext mongo.SessionContext) error {
	//
	// 	var err error
	//
	// 	_, err = mgo.DB("Test").C("test1").InsertOneWithSession(sessionContext, bson.M{"hello": "world"})
	//
	// 	_, err = mgo.DB("Test").C("test2").InsertOneWithSession(sessionContext, bson.M{"hello": "world"})
	//
	// 	return err
	// })
	var t = true
	a, err := mgo.DB("Test").C("test").Indexes().CreateOne(mongo.IndexModel{Keys: bson.M{"a": -1}, Options: &options.IndexOptions{Background: &t}})
	log.Println(a, err)
}
