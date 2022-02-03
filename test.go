package main 

import (
	"log"

	"github.com/dragonmaster101/science_web_api/database"
)

func main(){
	log.Println("Hello World!");
	db := database.Setup("school-exhibition-firebase-adminsdk-lubef-117af78def.json","https://school-exhibition-default-rtdb.asia-southeast1.firebasedatabase.app/")

	var err error

	post := database.Post{
		Title: "A A post that should come second",
		Author: "Omer Ali Malik",
		Date : "3 February , 2022",
		Url : "nothing",
		Description: "just testing",
	};
	err = db.CreatePostInfo(&post);
	if err != nil {
		log.Fatal(err);
	}

	results , err := db.GetPostsInfo()
	if err != nil {
		log.Fatal(err);
	}

	log.Printf("results : \n");
	for _ , result := range results {
		log.Printf("Post : %v\n" , result);
	}
}