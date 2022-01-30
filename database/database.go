package database

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"hash/fnv"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go/v4"
	db "firebase.google.com/go/v4/db"
)

/*
	?This function sets up the firebase app and the firebase database and also authenticates the instance
	?This function returns an Instance type for manipulating the database
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
	*The Actual Account type which will be stored in the database
*/
type Account struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (a *Account) Init(email , name , username , password string){
	a.Email = email;
	a.Name = name;
	a.Username = username;
	a.Password = HashAsString(password);
}

func (a *Account) Key() string {
	return a.Username;
}

func (a *Account) GetPath() string {
	return "users/" + a.Key();
}

/*
	*An account type for handling nil fields
*/
type NilAccount struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Password *string `json:"password"`
	Username *string `json:"username"`
}

/*
	*Hashes a string and returns the hash as an unsigned 32 bit integer
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
	*Returns hashed string as another string
*/
func HashAsString(s string) string {
	h := Hash(s);
	return strconv.Itoa(int(h));
}

/*
	*A Type which contains all the methods required to manipulate the firebase database instance
*/
type Instance struct {
	database *db.Client
	ctx      context.Context
}

const usersPath = "users/";

/*
	*Acts as a constructor for the Instance type
*/
func (i *Instance) Init(ctx context.Context, database *db.Client) {
	i.ctx = ctx
	i.database = database
}

/*
	?Adds a new entry in the usersPath of the firebases database with the given account details
*/
func (i *Instance) PostUserInfo(acc *Account) (err error){
	err = i.database.NewRef(usersPath + acc.Key()).Set(i.ctx , *acc);

	return err;
} 

/*
	?Gets user information with the given userId from the firebase database
*/
func (i *Instance) GetUserInfo(userId string) (acc *Account , err error) {
	acc = &Account{};
	err = i.database.NewRef(usersPath + userId).Get(i.ctx , acc);

	return acc , err;
}

/*
	?Updates the User information given the userId i.e email or username. The caller has to fill in the fields to update and the remaining fields should be left nil in the struct

	
	E.g 
	{
		username: "Username"
	} This struct is a valid parameter
*/
func (i *Instance) UpdateUserInfo(userId string , newAcc *NilAccount) (error) {
	
	// *reads the account to be updated and updates it with the non-nil fields of new account and writes it again
	// *this way all the previous account details are not jeopardised and only the updated fields have any chance of corruption
	prevAcc , readErr := i.GetUserInfo(userId);

	if readErr != nil {
		return readErr;
	}

	switch newAcc {
	case nil : 
		return fmt.Errorf("err account with updated information is nil");
	default:
	}

	switch newAcc.Email {
	case nil:
		// prevAcc.Email = prevAcc.Email;
	default:
		prevAcc.Email = *newAcc.Email;
	}

	switch newAcc.Name {
	case nil:
		// prevAcc.Name = prevAcc.Name;
	default:
		prevAcc.Name = *newAcc.Name;
	}

	switch newAcc.Username {
	case nil:
		// prevAcc.Email = prevAcc.Email;
	default:
		prevAcc.Username = *newAcc.Username;
	}

	switch newAcc.Password {
	case nil:
		// nothing
	default:
		prevAcc.Password = *newAcc.Password;
	}


	i.database.NewRef("users/" + prevAcc.Key()).Set(i.ctx , prevAcc);

	return nil;
}

/*
	?This DataType acts as the form for the AuthenticateUserInfo method
*/
type AuthenticatorForm struct {
	UserId string	 `json:"userId"` 
	Password string  `json:"password"`
	Hashed bool 	 `json:"hashed"`
}

/*
	*Validates the given credentials in the firebase database
*/
func (i *Instance) AuthenticateUserInfo(form *AuthenticatorForm) (bool , error){
	trueAccount , getAccountErr := i.GetUserInfo(form.UserId);
	if getAccountErr != nil {
		return false , fmt.Errorf("this Account does not exist. Recieved this Error when retrieving account : %v" , getAccountErr);
	}


	switch form.Hashed {
	case false:
		var pwd string = form.Password;
		pwd = HashAsString(pwd);
		if pwd != trueAccount.Password {
			return false , fmt.Errorf("passwords do not match. you sent password = %v , got password = %v" , pwd , trueAccount.Password);
		} else {
			return true , nil;
		}
	case true:
		if form.Password != trueAccount.Password {
			return false , fmt.Errorf("passwords do not match. you sent password = %v , got password = %v" , form.Password , trueAccount.Password);
		} else {
			return true , nil;
		}

	default:
		return false, fmt.Errorf("you have not specified if the password is hashed or not using the hashed field");
	}
}