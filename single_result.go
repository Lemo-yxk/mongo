/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-28 15:32
**/

package longo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type SingleResult struct {
	singleResult *mongo.SingleResult
}

func (sg *SingleResult) Do(result interface{}) error {
	return sg.singleResult.Decode(result)
}
