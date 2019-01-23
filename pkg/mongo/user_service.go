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

func NewWishService(mongoSession *MongoSession, config *config.MongoConfig) *WishService {
	collection := mongoSession.session.DB(config.DbName).C("Wish")
	collection.EnsureIndex(wishModelIndex())
	return &WishService{collection}
}

//CreateWish create a wish
func (ws *WishService) CreateWish(w *model.Wish) error {
	return ws.collection.Insert(&w)
}
