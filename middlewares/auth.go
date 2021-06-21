package mwauth

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware : to verify all authorized operations
func AuthMiddleware(ginContext *gin.Context) {
	firebaseAuth := ginContext.MustGet("firebaseAuth").(*auth.Client)
	authorizationToken := ginContext.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
		ginContext.Abort()
		return
	}
	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		ginContext.Abort()
		return
	}
	ginContext.Set("UUID", token.UID)
	ginContext.Next()
}
