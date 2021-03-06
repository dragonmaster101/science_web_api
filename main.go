package main

import (
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/dragonmaster101/science_web_api/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/*
	PATH CONSTANTS FOR EACH OPERATIONS THAT THE API PROVIDES
*/

const userOpSearchPath = "/users/search/:userid"    // GET
const userOpCreatePath = "/users/create"            // POST
const userOpUpdatePath = "users/update/:userid"     // POST
const userOpAuthenticatePath = "users/authenticate" // POST

const postOpGetAllPath = "/posts"           // GET
const postOpCreatePath = "post/create"      // POST
const postOpGetbyidPath = "post/id/:postid" // GET

/*
	END END END END OF PATH CONSTANTS
*/

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://google.com"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	db := database.Setup("school-exhibition-firebase-adminsdk-lubef-117af78def.json",
		"https://school-exhibition-default-rtdb.asia-southeast1.firebasedatabase.app/")

	router.GET(userOpSearchPath, func(c *gin.Context) {
		userID := c.Param("userid")

		account, fetchErr := db.GetUserInfo(userID)
		if fetchErr != nil {
			errString := fmt.Sprintf("an error occured while fetching account from database , err : %v", fetchErr)
			c.String(http.StatusBadRequest, errString)
			return
		}

		jsonData, jsonConversionError := json.Marshal(account)
		if jsonConversionError != nil {
			c.String(http.StatusBadRequest, "internal error converting account data into json , err : %v", jsonConversionError)
			return
		}

		c.Data(http.StatusOK, gin.MIMEJSON, jsonData)
	})

	/*router.GET Methods for the database.Post type*/

	router.GET(postOpGetAllPath, func(c *gin.Context) {
		posts, fetchErr := db.GetPostsInfo()
		if fetchErr != nil {
			c.String(http.StatusBadRequest, "error fetching the posts , err : %v", fetchErr)
			return
		}

		jsonData, jsonConversionError := json.Marshal(database.PostArray{Posts: posts})
		if jsonConversionError != nil {
			c.String(http.StatusBadRequest, "internal error converting account data into json , err : %v", jsonConversionError)
			return
		}

		c.Data(http.StatusOK, gin.MIMEJSON, jsonData)
	})

	router.GET(postOpGetbyidPath, func(c *gin.Context) {
		postID := c.Param("postid")

		post, getPostErr := db.GetPostInfo(postID)
		if getPostErr != nil {
			c.String(http.StatusBadRequest, "error in retrieving the post from the database (make sure the post exists) , err : %v", getPostErr)
		}

		jsonData, jsonConversionError := json.Marshal(post)
		if jsonConversionError != nil {
			c.String(http.StatusBadRequest, "internal error converting account data into json , err : %v", jsonConversionError)
			return
		}

		c.Data(http.StatusOK, gin.MIMEJSON, jsonData)
	})

	// -----------------------------------------------------------------------------------------------------

	router.POST(userOpCreatePath, func(c *gin.Context) {
		account := database.Account{}

		decodeErr := json.NewDecoder(c.Request.Body).Decode(&account)
		if decodeErr != nil {
			c.String(http.StatusBadRequest,
				"error decoding the request body make sure you followed the proper format , err : %v", decodeErr)
			return
		}

		postErr := db.PostUserInfo(&account)
		if postErr != nil {
			c.String(http.StatusBadRequest, "error posting the account details to the database , errr : %v", postErr)
			return
		}

		c.String(http.StatusOK, "Successfully created the account")
	})

	router.POST(userOpUpdatePath, func(c *gin.Context) {

		userID := c.Param("userid")
		account := &database.NilAccount{}

		updateErr := db.UpdateUserInfo(userID, account)
		if updateErr != nil {
			c.String(http.StatusBadRequest, "error updating the account make sure you followed the proper format , err : %v", updateErr)
			return
		}

		c.String(http.StatusOK, "Successfully updated the account with the given userID")
	})

	router.POST(userOpAuthenticatePath, func(c *gin.Context) {
		form := database.AuthenticatorForm{}

		decodeErr := json.NewDecoder(c.Request.Body).Decode(&form)
		if decodeErr != nil {
			c.String(http.StatusBadRequest,
				"error decoding the request body make sure you followed the proper format , err : %v", decodeErr)
			return
		}

		result, authenticationErr := db.AuthenticateUserInfo(&form)
		if authenticationErr != nil {
			c.String(http.StatusBadRequest, "error getting authentication results from the database , err : %v", authenticationErr)
		}

		c.String(http.StatusOK, "%v", result)
	})

	// ---------------------------------------------------------------------------------------

	/*router.POST methods for the Post data type*/

	router.POST(postOpCreatePath, func(c *gin.Context) {
		var post database.Post

		decodeErr := json.NewDecoder(c.Request.Body).Decode(&post)
		if decodeErr != nil {
			c.String(http.StatusBadRequest,
				"error decoding the request body make sure you followed the proper format , err : %v", decodeErr)
			return
		}

		createErr := db.CreatePostInfo(&post)
		if createErr != nil {
			c.String(http.StatusBadRequest, "error creating the post make sure you followed the proper format , err : %v", createErr)
			return
		}

		c.String(http.StatusOK, "Successfully created the post")
	})

	router.Run()
}
