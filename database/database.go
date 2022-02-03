package database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"strconv"

	"hash/fnv"

	"google.golang.org/api/option"

	"github.com/dragonmaster101/science_web_api/algorithms"

	firebase "firebase.google.com/go/v4"
	db "firebase.google.com/go/v4/db"
)

/*Setup : This function sets up the firebase app and the firebase database and also authenticates the instance
This function returns an Instance type for manipulating the database
*/
func Setup(credentialFile string, url string) *Instance {
	opt := option.WithCredentialsFile(credentialFile)

	ctx := context.Background()

	config := &firebase.Config{
		DatabaseURL: url,
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

/*Account : The Actual Account type which will be stored in the database
 */
type Account struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Username string `json:"username"`
}

/*Init : The constructor for the Account data type */
func (a *Account) Init(email, name, username, password string) {
	a.Email = email
	a.Name = name
	a.Username = username
	a.Password = HashAsString(password)
}

/*Key : returns the userId for the given Account */
func (a *Account) Key() string {
	return a.Username
}

/*GetPath : returns the path of the account in the firebase database */
func (a *Account) GetPath() string {
	return "users/" + a.Key()
}

/*NilAccount : An account type for handling nil fields
 */
type NilAccount struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Password *string `json:"password"`
	Username *string `json:"username"`
}

/*Hash : Hashes a string and returns the hash as an unsigned 32 bit integer
 */
func Hash(s string) uint32 {

	hasher := fnv.New32a()
	_, err := hasher.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}
	return hasher.Sum32()
}

/*HashAsString : Returns hashed string as another string
 */
func HashAsString(s string) string {
	h := Hash(s)
	return strconv.Itoa(int(h))
}

/*IsHash : Returns the boolean indicator of whether or not the given string is a hash or not in the context of passwords*/
func IsHash(str string) bool {
	_, err := strconv.Atoi(str)

	return err == nil
}

/*Instance : A Type which contains all the methods required to manipulate the firebase database instance
 */
type Instance struct {
	database *db.Client
	ctx      context.Context
}

const usersPath = "users/"

/*Init : Acts as a constructor for the Instance type */
func (i *Instance) Init(ctx context.Context, database *db.Client) {
	i.ctx = ctx
	i.database = database
}

/*PostUserInfo : Adds a new entry in the usersPath of the firebases database with the given account details */
func (i *Instance) PostUserInfo(acc *Account) (err error) {
	if !IsHash(acc.Password) {
		acc.Password = HashAsString(acc.Password)
	}
	err = i.database.NewRef(usersPath+acc.Key()).Set(i.ctx, *acc)

	return err
}

/*GetUserInfo : Gets user information with the given userId from the firebase database */
func (i *Instance) GetUserInfo(userID string) (acc *Account, err error) {
	acc = &Account{}
	err = i.database.NewRef(usersPath+userID).Get(i.ctx, acc)

	return acc, err
}

/*UpdateUserInfo :

Updates the User information given the userId i.e email or username. The caller has to fill in the fields to update and the remaining fields should be left nil in the struct


E.g
{
	username: "Username"
} This struct is a valid argument to the function
*/
func (i *Instance) UpdateUserInfo(userID string, newAcc *NilAccount) error {

	// *reads the account to be updated and updates it with the non-nil fields of new account and writes it again
	// *this way all the previous account details are not jeopardised and only the updated fields have any chance of corruption
	prevAcc, readErr := i.GetUserInfo(userID)

	if readErr != nil {
		return readErr
	}

	switch newAcc {
	case nil:
		return fmt.Errorf("err account with updated information is nil")
	default:
	}

	switch newAcc.Email {
	case nil:
		// prevAcc.Email = prevAcc.Email;
	default:
		prevAcc.Email = *newAcc.Email
	}

	switch newAcc.Name {
	case nil:
		// prevAcc.Name = prevAcc.Name;
	default:
		prevAcc.Name = *newAcc.Name
	}

	switch newAcc.Username {
	case nil:
		// prevAcc.Email = prevAcc.Email;
	default:
		prevAcc.Username = *newAcc.Username
	}

	switch newAcc.Password {
	case nil:
		// nothing
	default:
		prevAcc.Password = *newAcc.Password
	}

	i.database.NewRef("users/"+prevAcc.Key()).Set(i.ctx, prevAcc)

	return nil
}

/*AuthenticatorForm : This DataType acts as the form for the AuthenticateUserInfo method */
type AuthenticatorForm struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
	Hashed   bool   `json:"hashed"`
}

/*AuthenticateUserInfo : Validates the given credentials in the firebase database */
func (i *Instance) AuthenticateUserInfo(form *AuthenticatorForm) (bool, error) {
	trueAccount, getAccountErr := i.GetUserInfo(form.UserID)
	if getAccountErr != nil {
		return false, fmt.Errorf("this Account does not exist. Recieved this Error when retrieving account : %v", getAccountErr)
	}

	switch form.Hashed {
	case false:
		var pwd string = form.Password
		pwd = HashAsString(pwd)
		if pwd != trueAccount.Password {
			return false, fmt.Errorf("passwords do not match. you sent password = %v , got password = %v", pwd, trueAccount.Password)
		}
		return true, nil
	case true:
		if form.Password != trueAccount.Password {
			return false, fmt.Errorf("passwords do not match. you sent password = %v , got password = %v", form.Password, trueAccount.Password)
		}
		return true, nil

	default:
		return false, fmt.Errorf("you have not specified if the password is hashed or not using the hashed field")
	}
}


/*
<Post 
            title="Quantum Pacman"
            author="Omer Ali Malik"
            date="29 February 2020"
            url="https://www.youtube.com/watch?v=LMagNcngvcU&ab_channel=JavaScriptMastery"
            description="This is a game that uses quantum circuits and qauntum game theory to simulate pacman in an entangled state instance. Winners and losers are just predicted"
            card={false}
        />
*/

type Post struct {
	Title  		string `json:"title"`
	Author 		string `json:"author"`
	Date   		string `json:"date"`
	Url    		string `json:"url"`
	Description string `json:"description"`
}

type NilPost struct {
	Title  		*string `json:"title"`
	Author 		*string `json:"author"`
	Date   		*string `json:"date"`
	Url    		*string `json:"url"`
	Description *string `json:"description"`
}

const postsPath = "posts";

func (i *Instance) CreatePostInfo(post *Post) error {
	switch post {
	case nil:
		return fmt.Errorf("post paramater is nil");
	default:
	}	
	
	ref := i.database.NewRef("");
	postsRef := ref.Child(postsPath);
	newPostRef , postRefErr := postsRef.Push(i.ctx , nil);

	if postRefErr != nil {
		return fmt.Errorf("unable to get new reference for this post , err : %v" , postRefErr);
	}

	newPostRef.Set(i.ctx , post);
	
	return nil;
}

func (i *Instance) GetPostsInfo() ([]Post , error) {
	ref := i.database.NewRef("");
	postsRef := ref.Child(postsPath);

	var posts []Post;

	titleResults , titleQueryErr := postsRef.OrderByChild("title").GetOrdered(i.ctx);
	if titleQueryErr != nil {
		return nil , fmt.Errorf("error querying posts by titles from firebase database , err : %v" , titleQueryErr);
	}

	for _ , title := range titleResults {
		var post Post 
		if err := title.Unmarshal(&post); err != nil {
			return nil , fmt.Errorf("error decoding post data from title query results , err : %v" , err);
		}
		posts = append(posts, post);
	}

	return posts , nil;
}

func (i *Instance) GetPostInfo(postID string) (*Post , error) {
	ref := i.database.NewRef(postsPath + "/" + postID);
	post := Post{};
	err := ref.Get(i.ctx , &post);
	if err != nil {
		return nil , fmt.Errorf("error trying to retrieve post by this post id , err : %v" , err);
	}

	return &post , nil;
}



/*---------------------------- SEARCH POSTS METHODS START ----------------------------*/


func (i *Instance) SearchPostsTitle(query string) ([]Post , error) {
	query = strings.ToLower(query);

	ref := i.database.NewRef("");
	postsRef := ref.Child(postsPath);

	var titles []string;
	var posts []Post;

	titleResults , titleQueryErr := postsRef.OrderByChild("title").GetOrdered(i.ctx);
	if titleQueryErr != nil {
		return nil , fmt.Errorf("error querying posts by titles from firebase database , err : %v" , titleQueryErr);
	}

	for _ , title := range titleResults {
		var post Post 
		if err := title.Unmarshal(&post); err != nil {
			return nil , fmt.Errorf("error decoding Post data from title query Results , err : %v" , err);
		}
		titles = append(titles , strings.ToLower(post.Title));
		posts = append(posts, post);
	}

	indices , ok := algorithms.SearchStrings(titles , query);
	if !ok {
		return nil , nil;
	}

	var filteredPosts []Post;

	for _ , index := range indices {
		filteredPosts = append(filteredPosts , posts[index]);
	}

	return filteredPosts , nil;
}

func (i *Instance) SearchPostsAuthor(query string) ([]Post , error) {
	query = strings.ToLower(query);

	ref := i.database.NewRef("");
	postsRef := ref.Child(postsPath);

	var authors []string;
	var posts []Post;

	authorResults , authorQueryErr := postsRef.OrderByChild("title").GetOrdered(i.ctx);
	if authorQueryErr != nil {
		return nil , fmt.Errorf("error querying posts by Authors from firebase database , err : %v" , authorQueryErr);
	}

	for _ , author := range authorResults {
		var post Post 
		if err := author.Unmarshal(&post); err != nil {
			return nil , fmt.Errorf("error decoding Post data from Author query Results , err : %v" , err);
		}
		authors = append(authors , strings.ToLower(post.Author));
		posts = append(posts, post);
	}

	indices , ok := algorithms.SearchStrings(authors , query);
	if !ok {
		return nil , nil;
	}

	var filteredPosts []Post;

	for _ , index := range indices {
		filteredPosts = append(filteredPosts , posts[index]);
	}

	return filteredPosts , nil;
}

func PostSame(postA *Post , postB *Post) bool {
	return postA.Title == postB.Title;
}

func PostIsDuplicate(posts []Post , postCmp *Post) bool {
	for _ , post := range posts {
		return PostSame(&post , postCmp);
	} 
	return true;
}

func (i *Instance) SearchPosts(query string) ([]Post , error) {
	titlePosts , titleQueryErr := i.SearchPostsTitle(query);
	if titleQueryErr != nil {
		return nil , titleQueryErr;
	}

	authorPosts , authorQueryErr := i.SearchPostsAuthor(query);
	if authorQueryErr != nil {
		return nil , authorQueryErr;
	}

	var posts []Post;
	if len(titlePosts) != 0 {
		posts = titlePosts;
		for _ , authorPost := range authorPosts {
			if !PostIsDuplicate(posts , &authorPost) {
				posts = append(posts, authorPost);
			}
		}
	} else {
		if len(authorPosts) == 0 {
			return nil , nil;
		} else {
			posts = authorPosts;
		}
	}

	return posts , nil;
}

/*-------------------------------------- SEARCH POSTS METHODS END -------------------------------------------*/

func (i *Instance) UpdatePostInfo(postID string , updatedPost *NilPost) error{
	oldPost , oldPostErr := i.GetPostInfo(postID);
	if oldPostErr != nil {
		return fmt.Errorf("error in post update , error fetching post (maybe doesn't exist) , err : %v" , oldPostErr);
	}

	switch updatedPost {
	case nil:
		return nil;
	}


	switch updatedPost.Author {
	case nil:
		break;
	default:
		oldPost.Author = *updatedPost.Author;
	}

	switch updatedPost.Date {
	case nil:
		break;
	default:
		oldPost.Date = *updatedPost.Date;
	}

	switch updatedPost.Description {
	case nil:
		break;
	default:
		oldPost.Description = *updatedPost.Description;
	}

	switch updatedPost.Title {
	case nil:
		break;
	default:
		oldPost.Title = *updatedPost.Title;
	}

	switch updatedPost.Url {
	case nil:
		break;
	default:
		oldPost.Url = *updatedPost.Url;
	}

	ref := i.database.NewRef(postsPath + "/" + postID);
	err := ref.Set(i.ctx , oldPost);
	if err != nil {
		return fmt.Errorf("error trying to retrieve post by this post id , err : %v" , err);
	}

	return nil;
}