package services

import (
	"context"
	"errors"
	db "redemption/models/db"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateProduct(name string, manufacturer string, createdBy primitive.ObjectID) (*db.Product, error) {

	product := db.NewProduct(name, manufacturer, createdBy)
	err := mgm.Coll(product).Create(product)
	if err != nil {
		return nil, errors.New("cannot create new product")
	}

	return product, nil
}

func FindProductById(productId primitive.ObjectID) (*db.Product, error) {
	product := &db.Product{}
	err := mgm.Coll(product).FindByID(productId, product)
	if err != nil {
		return nil, errors.New("cannot find user")
	}

	return product, nil
}

func GetAllProducts(createdBy primitive.ObjectID, limit int64, offset int64) ([]*db.Product, int64, error) {
	var products []*db.Product

	opts := &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
	}

	total, err := mgm.Coll(&db.Product{}).CountDocuments(context.Background(), bson.M{"created_by": createdBy})
	if err != nil {
		return nil, -1, errors.New("failed to count products: " + err.Error())
	}

	err = mgm.Coll(&db.Product{}).SimpleFind(&products, bson.M{"created_by": createdBy}, opts)
	if err != nil {
		return nil, -1, errors.New("cannot find products")
	}

	return products, total, nil
}

func DeleteAllProducts(createdBy primitive.ObjectID) error {

	_, err := mgm.Coll(&db.Product{}).DeleteMany(context.Background(), bson.M{"created_by": createdBy})
	if err != nil {
		return errors.New("failed to delete products: " + err.Error())
	}

	return nil
}
