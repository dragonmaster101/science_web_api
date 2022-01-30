package database

import (
	"context"
	"log"
	"strconv"

	"hash/fnv"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go/v4"
	db "firebase.google.com/go/v4/db"
)

/*
	This function sets up the firebase app and the firebase database and also authenticates the instance
	This function returns an Instance type for manipulating the database
*/
func Setup() *Instance {
	opt := option.WithCredentialsFile("codexia-dc073-firebase-adminsdk-v40vf-bd80ee0a87.json")

	ctx := context.Background()

	config := &firebase.Config{
		DatabaseURL: "https://codexia-dc073-default-rtdb.asia-southeast1.firebasedatabase.app/",
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
	}

	instance := Instance{}
	instance.Init(ctx, client)

	return &instance
}

/*
	The Actual Account type which will be stored in the database
*/
type Account struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Username string `json:"username"`
}

/*
	An account type for handling nil fields
*/
type NilAccount struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Password *string `json:"password"`
	Username *string `json:"username"`
}

/*
	Hashes a string and returns the hash as an unsigned 32 bit integer
*/
func Hash(s string) uint32{
	
	hasher := fnv.New32a();
	_ , err := hasher.Write([]byte(s));
	if err != nil {
		log.Fatal(err);
	}
	return hasher.Sum32();
}

/*
	Returns hashed string as another string
*/
func HashAsString(s string) string {
	h := Hash(s);
	return strconv.Itoa(int(h));
}

/*
	A Type which contains all the methods required to manipulate the firebase database instance
*/
type Instance struct {
	database *db.Client
	ctx      context.Context
}

/*
	Acts as a constructor for the Instance type
*/
func (i *Instance) Init(ctx context.Context, database *db.Client) {
	i.ctx = ctx
	i.database = database
}