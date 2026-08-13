package middleware

import (
	"auth/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ctxUserEmail = "auth.email"
	ctxRolekey   = "auth.role"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc{
	return func(c *gin.Context){

		authHeader:=strings.TrimSpace(c.GetHeader("Authorization"))

		if authHeader==""{
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error":"Missing Authorization token",
			})
			return;
		}

		parts:=strings.SplitN(authHeader," ",2);
		if len(parts)!=2{
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error":"Invalid Authorization format",
			})
			return;
		}

		scheme:=strings.TrimSpace(parts[0])
		tokenString:=strings.TrimSpace(parts[1])

		if  strings.EqualFold(scheme,"Bearer"){
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error":"Authorization scheme must be Bearer"
			})

			return;
		}

		if  tokenString==""{
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error":"Authorization scheme must be Bearer"
			})
			return;
		}

		claims,err:=auth.ParseToken(jwtSecret,tokenString)
		if err!=nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error":"Invalid or expired token",
			})
			return;
		}

		c.Set(ctxUserEmail,claims.Email)
		c.Set(ctxRolekey,claims.Role)

		c.Next()

	}
}

func GetUserEmail(c *gin.Context)(string,bool){
	res,ok:=c.Get(ctxUserEmail)
	if !ok{
		return "",false;
	}
	userEmail,ok:=res.(string)
	return userEmail,ok;
}

func GetUserRole(c *gin.Context)(string,bool){
	res,ok:=c.Get(ctxRolekey)
	if !ok{
		return "",false;
	}
	userRole,ok:=res.(string)
	return userRole,ok;
}