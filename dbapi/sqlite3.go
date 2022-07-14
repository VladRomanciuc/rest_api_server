package dbapi

import (
	"database/sql"
	"log"
	"os"

	"github.com/VladRomanciuc/Go-classes/api/models"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct{}

//DB client
func NewSQLiteDb() models.DbOps {
	//remove previous version
	os.Remove("./posts.db")
	//Acces the DB, handling error
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
	}
	//Close connection after operation
	defer db.Close()
	//Create table statement
	createTable := `
	create table posts (id integer not null primary key, title text, txt text);
	delete from posts;
	`
	//Execute the create table statement and handling the error
	_, err = db.Exec(createTable)
	if err != nil {
		log.Printf("%q: %s\n", err, createTable)
	}
	return &sqlite{}
}

//Add post function
func (*sqlite) AddPost(post *models.Post) (*models.Post, error) {
	//Access the DB and handling error
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Access the db cli and handle the error
	prep, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Prepare the statement to add a post, handling the error
	addpost, err := prep.Prepare("insert into posts(id, title, txt) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Close connection after operation
	defer addpost.Close()
	//Execute add post cmd, handling the error
	_, err = addpost.Exec(post.Id, post.Title, post.Text)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Commit changes
	prep.Commit()
	return post, nil
}
//Get all post function from db
func (*sqlite) GetAll() ([]models.Post, error) {
	//Acces the DB and handling error
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Prepare the query to get all posts
	entry, err := db.Query("select id, title, txt from posts")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Close connection after operation
	defer entry.Close()
	//Declare a slice of type post
	var posts []models.Post
	//looping every entry and scan for data
	for entry.Next() {
		var id string
		var title string
		var text string
		err = entry.Scan(&id, &title, &text)
		//hendling the error
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		//declare type of data and map with post structure
		post := models.Post{
			Id:    id,
			Title: title,
			Text:  text,
		}
		//add each entry to the slice
		posts = append(posts, post)
	}
	//Error checker and handler
	err = entry.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//return the slice
	return posts, nil
}

func (*sqlite) DeleteById(id string) (*models.Post, error) {
	//Acces the DB and handling error
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Access the db cli and handle the error
	prep, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Prepare the statement to delete the post, handling the error
	delete, err := prep.Prepare("delete from posts where id =?")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Close connection after operation
	defer delete.Close()
	//Execute the statement and handle the error
	_, err = delete.Exec(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Commit changes
	prep.Commit()
	return nil, nil
}


func (*sqlite) GetById(id string) (*models.Post, error) {
	//Acces the DB and handling error
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//Querry a specific entry by id
	row := db.QueryRow("select id, title, txt from posts where id = ?", id)
	//Declare a variable of type post and map with post structure
	var post models.Post
	if row != nil {
		var id string
		var title string
		var text string
		result := row.Scan(&id, &title, &text)
		//Error handler
		if result != nil {
			return nil, err
		} else {
			//return the post
			post = models.Post{
				Id:    id,
				Title: title,
				Text:  text,
			}
		}
	}
	//return the post
	return &post, nil
}