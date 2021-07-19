package repositories

import (
	"context"
	"log"
	"server-with-firebase-go/entities"

	"cloud.google.com/go/firestore"
)

type PostRepository interface {
	Save(post *entities.Post) (*entities.Post, error)
	//FindAll() ([]*entities.Post, error)
}

type repository struct{}

const (
	projectId      = "server-with-go"
	collectionName = "posts"
)

func NewPostRepository() PostRepository {
	return &repository{}
}

func (*repository) Save(post *entities.Post) (*entities.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Falló la creación del cliente de Firestore: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"id":    post.Id,
		"title": post.Title,
		"text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Falló la creación del post: %v", err)
		return nil, err
	}
	return post, nil
}

//func (*repository) FindAll() ([]*entities.Post, error)  {

//}
