package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"gymservice/connectdb"
	"gymservice/models"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func generateAPIKey() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func RegisterClient(c *gin.Context) {
    var client models.Client
    if err := c.ShouldBindJSON(&client); err != nil {
        log.Printf("error %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    client.APIKey = generateAPIKey()

    var affiliatorID int
    err := connectdb.DB.QueryRow(`
        INSERT INTO affiliators (
            affiliator_fname, affiliator_lname, affiliator_email, affiliator_phone, api_key
        ) VALUES ($1, $2, $3, $4, $5)
        RETURNING affiliator_id
    `, client.AffiliatorFirstName, client.AffiliatorLastName, client.AffiliatorEmail, client.AffiliatorPhone, client.APIKey).Scan(&affiliatorID)

    if err != nil {
        log.Printf("error %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register affiliator"})
        return
    }

    for _, website := range client.AffiliatorWebsite {
        _, err := connectdb.DB.Exec(`
            INSERT INTO referrer_websites (affiliator_id, website_url)
            VALUES ($1, $2)
        `, affiliatorID, website)

        if err != nil {
            log.Printf("error %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register website"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "message":  "Client registered",
        "api_key":  client.APIKey,
        "websites": client.AffiliatorWebsite,
    })
}

func GetClient(c *gin.Context) {
	query := `
    SELECT aff.affiliator_id, aff.affiliator_fname, aff.affiliator_lname, aff.affiliator_email, aff.affiliator_phone, 
	aff.api_key, string_agg(rw.website_url, ',') AS website_urls
    FROM affiliators aff
    JOIN referrer_websites rw ON aff.affiliator_id = rw.affiliator_id
    GROUP BY aff.affiliator_id
`
	rows, err := connectdb.DB.Query(query)
	if err != nil {
		log.Printf("Failed to fetch data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch data"})
		return
	}
	defer rows.Close()

	clients := make([]models.Client, 0)
	for rows.Next() {
		var client models.Client
    var websiteUrls string
    if err := rows.Scan(&client.AffiliatorId, &client.AffiliatorFirstName, &client.AffiliatorLastName, &client.AffiliatorEmail, &client.AffiliatorPhone, &client.APIKey, &websiteUrls); err != nil {
        log.Printf("Failed to scan row: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Data scan error"})
        return
    }

    client.AffiliatorWebsite = strings.Split(websiteUrls, ",")
    clients = append(clients, client)
	}
	c.JSON(http.StatusOK, clients)
}

func GetLog(c *gin.Context) {
	rows, err := connectdb.DB.Query("SELECT * FROM log_api_requests")
	if err != nil {
		log.Printf("Failed to fetch data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch dara"})
		return
	}
	defer rows.Close()

	var log []models.LogApiRequest
	for rows.Next() {
		var logapi models.LogApiRequest
		if err := rows.Scan(&logapi.ResuestID, &logapi.AffiliatorId, &logapi.AffiliatorFname, &logapi.AffiliatorLname, &logapi.RequestTimestamp, &logapi.Endpoint, &logapi.PathParameters, &logapi.QueryParameters, &logapi.Method); err != nil {
			//log.Printf("Failed to scan row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data scan error"})
		}
		log = append(log, logapi)
	}
	c.JSON(http.StatusOK, log)
}

func DeleteAffiliator(c *gin.Context) {
	Param := c.Param("affiliatorId")
	affiliatorId, err := strconv.Atoi(Param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid affiliator ID"})
		return
	}

	query := `DELETE FROM affiliators WHERE affiliator_id = $1`
	_, err = connectdb.DB.Exec(query, affiliatorId)
	if err != nil {
		log.Printf("Failed to delete affiliator: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete affiliator"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Affiliator deleted successfully"})
}

func CreateLogin(c *gin.Context) {
	var input struct {
        AffiliatorFirstName string `json:"affiliator_fname"`
        AffiliatorLastName  string `json:"affiliator_lname"`
        AffiliatorEmail     string `json:"affiliator_email"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        log.Printf("error %v:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var client models.Client
    query := `SELECT affiliator_id, affiliator_fname, affiliator_lname, affiliator_email, affiliator_phone, api_key FROM affiliators WHERE affiliator_email = $1 AND affiliator_fname = $2 AND affiliator_lname = $3`
    row := connectdb.DB.QueryRow(query, input.AffiliatorEmail, input.AffiliatorFirstName, input.AffiliatorLastName)
    err := row.Scan(&client.AffiliatorId, &client.AffiliatorFirstName, &client.AffiliatorLastName, &client.AffiliatorEmail, &client.AffiliatorPhone, &client.APIKey)
    
    if err != nil {
        if err == sql.ErrNoRows {
			log.Printf("Failed to authorized %v:", err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        } else {
			log.Printf("%v:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        }
        return
    }
	c.JSON(http.StatusOK, input)

}

func GetAffiliatorByEmail(c *gin.Context) {
    email := c.Param("email")

    var affiliator models.Client
    err := connectdb.DB.QueryRow(`
        SELECT affiliator_id, affiliator_fname, affiliator_lname, affiliator_email, affiliator_phone, api_key
        FROM affiliators
        WHERE affiliator_email = $1
    `, email).Scan(&affiliator.AffiliatorId, &affiliator.AffiliatorFirstName, &affiliator.AffiliatorLastName, &affiliator.AffiliatorEmail, &affiliator.AffiliatorPhone, &affiliator.APIKey)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Affiliator not found"})
        return
    }

    rows, err := connectdb.DB.Query(`
        SELECT website_url FROM referrer_websites WHERE affiliator_id = $1
    `, affiliator.AffiliatorId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch websites"})
        return
    }
    defer rows.Close()

    var websites []string
    for rows.Next() {
        var websiteUrls string
        if err := rows.Scan(&websiteUrls); err == nil {
            websites = append(websites, websiteUrls)
        }
    }

    affiliator.AffiliatorWebsite = websites

    c.JSON(http.StatusOK, affiliator)
}
