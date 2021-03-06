package mongo

import (
	config "../config"
	model "../model"
	mgo "gopkg.in/mgo.v2"
)

//WishService which of type WishService interface
type WishService struct {
	collection *mgo.Collection
}

//NewWishService create and initiate a WishService
func NewWishService(mongoSession *MongoSession) *WishService {
	config := config.GetMongoConfig()

	collection := mongoSession.session.DB(config.DbName).C("Wish")
	collection.EnsureIndex(wishModelIndex())
	return &WishService{collection}
}

//CreateWish create a wish
func (ws *WishService) CreateWish(w *model.Wish) error {
	return ws.collection.Insert(&w)
}
