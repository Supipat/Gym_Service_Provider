package handlers

import (
	"gymservice/connectdb"
	"gymservice/models"
	"log"
	"net/http"

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

func LogClick(c *gin.Context) {
	var click models.ClickLog
	if err := c.ShouldBindJSON(&click); err != nil {
		log.Printf("error %v:", err)
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}

	query := `
        INSERT INTO click_logs (class_url, class_id, class_type, click_timestamp)
        VALUES ($1, $2, $3, $4)
        RETURNING click_id, class_url, class_id, class_type, click_timestamp`

	var result models.ClickLog
	err := connectdb.DB.QueryRow(query,
		click.ClickTarget,
		click.ClickTargetId,
		click.ClickTargetType,
		click.ClickTimestamp,
	).Scan(
		&result.ClickId,
		&result.ClickTarget,
		&result.ClickTargetId,
		&result.ClickTargetType,
		&result.ClickTimestamp,
	)

	if err != nil {
		log.Println("Error inserting click log:", err)
		c.JSON(500, gin.H{"error": "Failed to log click"})
		return
	}

	c.JSON(200, gin.H{"message": "Click logged successfully", "click_id": result.ClickId})
}

func GetClickLogs(c *gin.Context) {

	query := `
        SELECT c.click_id, c.class_url, c.class_id, c.class_type, c.click_timestamp
        FROM click_logs c
        ORDER BY c.click_timestamp DESC`

	rows, err := connectdb.DB.Query(query)
	if err != nil {
		log.Println("Error querying click logs:", err)
		c.JSON(500, gin.H{"error": "Failed to retrieve click logs"})
		return
	}
	defer rows.Close()

	var clickLogs []models.ClickLog
	for rows.Next() {
		var clickLog models.ClickLog
		err := rows.Scan(&clickLog.ClickId, &clickLog.ClickTarget, &clickLog.ClickTargetId, &clickLog.ClickTargetType, &clickLog.ClickTimestamp)
		if err != nil {
			log.Println("Error scanning row:", err)
			c.JSON(500, gin.H{"error": "Failed to retrieve click log data"})
			return
		}
		clickLogs = append(clickLogs, clickLog)
	}

	if len(clickLogs) == 0 {
		c.JSON(404, gin.H{"message": "No click logs found"})
		return
	}

	c.JSON(200, gin.H{"click_logs": clickLogs})
}
