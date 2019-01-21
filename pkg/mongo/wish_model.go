package mongo

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson" 
)

//WishModel type
type WishModel struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Title string
	Story string
}

func wishModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"title"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func (w *WishModel) createWish(db *mgo.Collection) error {
	return errors.New("Not implemented")
}
