package repositories

import (
	"context"
	"log"
	"server-with-firebase-go/entities"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type PostRepository interface {
	Save(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
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

func (*repository) FindAll() ([]entities.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Falló la creación del cliente de Firestore: %v", err)
		return nil, err
	}

	defer client.Close()
	var posts []entities.Post
	documentIterator := client.Collection(collectionName).Documents(ctx)
	defer documentIterator.Stop()
	for {
		doc, err := documentIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Falló la obtencion de la lista de posts: %v", err)
			return nil, err
		}
		data := doc.Data()
		post := entities.Post{
			Id:    data["id"].(int),
			Title: data["title"].(string),
			Text:  data["text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
