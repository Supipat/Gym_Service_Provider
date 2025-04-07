package handlers

import (
	"gymservice/connectdb"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("Authorization")
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API Key required"})
			c.Abort()
			return
		}

		var affiliatorID int
		var affiliatorFname, affiliatorLname string

		err := connectdb.DB.QueryRow(`
			SELECT affiliator_id, affiliator_fname, affiliator_lname 
			FROM affiliators 
			WHERE api_key = $1`, apiKey).
			Scan(&affiliatorID, &affiliatorFname, &affiliatorLname)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
			c.Abort()
			return
		}

		c.Set("affiliator_id", affiliatorID)
		c.Set("affiliator_fname", affiliatorFname)
		c.Set("affiliator_lname", affiliatorLname)
		c.Next()
	}
}


func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		affiliatorID, exists := c.Get("affiliator_id")
	
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		affiliatorFname, _ := c.Get("affiliator_fname")
		affiliatorLname, _ := c.Get("affiliator_lname")

		var pathParams string
		for _, param := range c.Params {
    		pathParams += param.Key + "=" + param.Value + "&"
		}

		_, err := connectdb.DB.Exec("INSERT INTO log_api_requests (affiliator_id, affiliator_fname, affiliator_lname, endpoint, path_parameters, query_parameters, method) VALUES ($1, $2, $3, $4, $5, $6, $7)",
    	affiliatorID, affiliatorFname, affiliatorLname, c.Request.URL.Path, pathParams, c.Request.URL.RawQuery, c.Request.Method)

		if err != nil {
			log.Printf("Failed to log request: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not log request"})
			c.Abort()
			return
		}

		c.Next()
	}
}