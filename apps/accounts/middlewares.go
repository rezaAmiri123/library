package accounts

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/rezaAmiri123/library/conf"
	"net/http"
	"strings"
)

func UpdateContextUser(ctx *gin.Context, uid uint) {
	var u User
	if uid != 0 {
		db := conf.GetDB()
		db.First(&u, uid)
	}
	ctx.Set("user", u)
	ctx.Set("userId", uid)
}

//Strip 'TOKEN ' prefix from token
func StripBearerPrefixFromTokenString(token string) (string, error) {
	if len(token) > 5 && strings.ToUpper(token[0:6]) == "TOKEN " {
		return token[6:], nil
	}
	return token, nil
}

//Extract token from Authorization header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	StripBearerPrefixFromTokenString,
}

//Extractor for OAuth2 access tokens. Looks in 'Authorization'
// header then 'access_token' argument for a token
var Auth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func GetByteSecretKey(token *jwt.Token) (interface{}, error) {
	bsp := []byte(conf.SecretKey)
	return bsp, nil
}

// You can custom middlewares yourself as the doc: https://github.com/gin-gonic/gin#custom-middleware
//  r.Use(AuthMiddleware(true))
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		UpdateContextUser(ctx, 0)
		token, err := request.ParseFromRequest(ctx.Request, Auth2Extractor, GetByteSecretKey)
		if err != nil {
			if auto401 {
				ctx.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			uid := uint(claims["id"].(float64))
			UpdateContextUser(ctx, uid)
		}
	}
}

