package repository

import (
	"context"
	"log"

	"github.com/Juliannars1/invoice-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceRepository struct {
	collection *mongo.Collection
}

func NewInvoiceRepository(client *mongo.Client, dbName, collectionName string) *InvoiceRepository {
	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	return &InvoiceRepository{
		collection: collection,
	}
}

func (repo *InvoiceRepository) CreateInvoice(ctx context.Context, invoice interface{}) error {
	_, err := repo.collection.InsertOne(ctx, invoice)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
func (repo *InvoiceRepository) GetInvoice(ctx context.Context, Number string) (*models.Invoice, error) {
    var invoice models.Invoice 
    err := repo.collection.FindOne(ctx, bson.M{"number": Number}).Decode(&invoice)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    return &invoice, nil 
}
func (repo *InvoiceRepository) GetAllInvoices(ctx context.Context) (interface{}, error) {
	cur, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cur.Close(context.Background())
	var invoices []*models.Invoice

	for cur.Next(ctx) {
		var invoice models.Invoice
		if err := cur.Decode(&invoice); err != nil {
			log.Println(err)
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return invoices, nil
	
}
func (repo *InvoiceRepository) UpdateInvoice(ctx context.Context, number string, invoice *map[string]interface{}) error {
	update := bson.M{"$set": *invoice}
	_, err := repo.collection.UpdateOne(ctx, bson.M{"number": number}, update)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repo *InvoiceRepository) DeleteInvoice(ctx context.Context, number string) error {
	_, err := repo.collection.DeleteOne(ctx, bson.M{"number": number})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
