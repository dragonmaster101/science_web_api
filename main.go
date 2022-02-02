package main

import (
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/dragonmaster101/acweb_api/database"
	"github.com/gin-gonic/gin"
)

/*
	PATH CONSTANTS FOR EACH OPERATIONS THAT THE API PROVIDES
*/

/*
	END END END END OF PATH CONSTANTS
*/

func main() {
	router := gin.Default()

	db := database.Setup("school-exhibition-firebase-adminsdk-lubef-117af78def.json",
		"https://school-exhibition-default-rtdb.asia-southeast1.firebasedatabase.app/")

	router.GET("/users/search/:userid", func(c *gin.Context) {
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

	router.POST("/users/create", func(c *gin.Context) {
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

	router.POST("users/update/:userid", func(c *gin.Context) {

		userID := c.Param("userid")
		account := &database.NilAccount{}

		updateErr := db.UpdateUserInfo(userID, account)
		if updateErr != nil {
			c.String(http.StatusBadRequest, "error updating the account make sure you followed the proper format , err : %v", updateErr)
			return
		}

		c.String(http.StatusOK, "Successfully updated the account with the given userID")
	})

	router.POST("users/authenticate", func(c *gin.Context) {
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

	router.Run()
}
