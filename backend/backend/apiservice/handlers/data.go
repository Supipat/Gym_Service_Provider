package handlers

import (
	"database/sql"
	"fmt"
	"gymservice/connectdb"
	"gymservice/models"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllGym(c *gin.Context, params models.QueryParams) {
	query := `
        SELECT * FROM gyms
        WHERE 1 = 1
    `
    args := []interface{}{}
    paramIndex := 1

    if params.Search != "" {
        query += fmt.Sprintf(`
            AND (gym_name ILIKE $%d OR gym_location ILIKE $%d OR gym_type ILIKE $%d)
        `, paramIndex, paramIndex+1, paramIndex+2)
		args = append(args, "%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%")
		paramIndex += 3
    }

    if params.GymID != 0 {
        query += fmt.Sprintf(" AND g.gym_id = $%d", paramIndex)
		args = append(args, params.GymID)
		paramIndex++
    }

    sortFields := map[string]string{
        "name": "gym_name",
        "id": "gym_id",
    }

    orderDirections := map[string]string{
        "asc": "ASC",
        "desc": "DESC",
    }

    sortColumn, ok := sortFields[params.Sort]
    if !ok {
        sortColumn = "gym_id"
    }

    orderDirection, ok := orderDirections[strings.ToLower(params.Order)]
    if !ok {
        orderDirection = "ASC"
    }

    query += fmt.Sprintf(" ORDER BY %s %s , gym_id ASC", sortColumn, orderDirection)

	rows, err := connectdb.DB.Query(query, args...)
	if err != nil {
		log.Printf("Failed to fetch data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch data"})
		return
	}
	defer rows.Close()

	gyms := make([]models.Gym, 0)
	for rows.Next() {
		var gym models.Gym
		if err := rows.Scan(
			&gym.GymId, &gym.GymName, &gym.GymType, &gym.GymOfficeHours, &gym.GymLocation, &gym.GymContactNumber, &gym.GymEmail); err != nil {
			log.Printf("Failed to scan row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data scan error"})
			return
		}
		gyms = append(gyms, gym)
	}
	c.JSON(http.StatusOK, gyms)
}

func GetGymId(c *gin.Context, gymId int) {
    query := `
        SELECT * FROM gyms WHERE gym_id = $1
    `

    row := connectdb.DB.QueryRow(query, gymId)

    var gym models.Gym
    if err := row.Scan(
        &gym.GymId, &gym.GymName, &gym.GymType, &gym.GymOfficeHours, &gym.GymLocation, &gym.GymContactNumber, &gym.GymEmail); err != nil {
        if err == sql.ErrNoRows {
            log.Printf("Gym with id %d not found", gymId)
            c.JSON(http.StatusNotFound, gin.H{"error": "Gym not found"})
        } else {
            log.Printf("Failed to scan row: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Data scan error"})
        }
        return
    }
    c.JSON(http.StatusOK, gym)   
}

func GetGymClass(c *gin.Context, gymId int, params models.QueryParams) {
    query := `
        SELECT gc.class_id, gc.class_name, gc.description, gc.class_type, gc.class_duration, gc.class_schedule, gc.class_price, gc.class_level, gc.class_url, gc.gym_id
        FROM gym_classes gc
        WHERE 1 = 1 AND gc.gym_id = $1
    `
    args := []interface{}{gymId}
    paramIndex := 2

    if params.Search != "" {
        query += fmt.Sprintf(`
            AND (gc.class_name ILIKE $%d OR gc.description ILIKE $%d OR gc.class_type ILIKE $%d)
        `, paramIndex, paramIndex+1, paramIndex+2)
        args = append(args, "%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%")
        paramIndex += 3
    }

    sortFields := map[string]string{
        "name": "gc.class_name",
        "price": "gc.class_price",
        "level": "gc.class_level",
        "id": "gc.class_id",
    }

    orderDirections := map[string]string{
        "asc": "ASC",
        "desc": "DESC",
    }

    sortColumn, ok := sortFields[params.Sort]
    if !ok {
        sortColumn = "gc.class_id"
    }

    orderDirection, ok := orderDirections[strings.ToLower(params.Order)]
    if !ok {
        orderDirection = "ASC"
    }

    query += fmt.Sprintf(" ORDER BY %s %s", sortColumn, orderDirection)

    rows, err := connectdb.DB.Query(query, args...)
    if err != nil {
        log.Printf("Failed to fetch data: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch data"})
        return
    }
    defer rows.Close()

    gymClasses := make([]models.GymClass, 0)
    for rows.Next() {
        var class models.GymClass
        if err := rows.Scan(
            &class.ClassId, &class.ClassName, &class.Description, &class.ClassType,
            &class.ClassDuration, &class.ClassSchedule, &class.ClassPrice, &class.ClassLevel, &class.ClassUrl, &class.GymId); err != nil {
            log.Printf("Failed to scan row: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Data scan error"})
            return
        }
        gymClasses = append(gymClasses, class)
    }
    c.JSON(http.StatusOK, gymClasses)

}

func GetGymClassId(c *gin.Context, gymId int, classId int) {
    query := `
        SELECT gc.class_id, gc.class_name, gc.description, gc.class_type, gc.class_duration, gc.class_schedule, gc.class_price, gc.class_level, gc.class_url
        FROM gym_classes gc
        WHERE 1 = 1 AND gc.gym_id = $1 AND gc.class_id = $2
    `
    row := connectdb.DB.QueryRow(query, gymId, classId)

    var class models.GymClass
    if err := row.Scan(
        &class.ClassId, &class.ClassName, &class.Description, &class.ClassType, &class.ClassDuration, &class.ClassSchedule, &class.ClassPrice, &class.ClassLevel, &class.ClassUrl); err != nil {
        if err == sql.ErrNoRows {
            log.Printf("Class with id %d not found", classId)
            c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
        } else {
            log.Printf("Failed to scan row: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Data scan error"})
        }
        return
    }
    c.JSON(http.StatusOK, class)   

}

func CreateGym(c *gin.Context) {
    var gym models.Gym

	if err := c.ShouldBindJSON(&gym); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	query := `
        INSERT INTO gyms (gym_name, gym_type, gym_office_hours, gym_location, gym_contact_number, gym_email)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING gym_id, gym_name, gym_type, gym_office_hours, gym_location, gym_contact_number, gym_email
    `
	err := connectdb.DB.QueryRow(
		query,
		gym.GymName,
		gym.GymType,
		gym.GymOfficeHours,
		gym.GymLocation,
		gym.GymContactNumber,
		gym.GymEmail,
	).Scan(
		&gym.GymId,
		&gym.GymName,
		&gym.GymType,
		&gym.GymOfficeHours,
		&gym.GymLocation,
		&gym.GymContactNumber,
		&gym.GymEmail,
	)

	if err != nil {
		log.Printf("Failed to insert gym: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create gym"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Gym created successfully",
		"gym":     gym,
	})
}

func CreateGymClass(c *gin.Context) {
    var gymclass models.GymClass

	if err := c.ShouldBindJSON(&gymclass); err != nil {
		log.Printf("Failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	query := `
        INSERT INTO gym_classes (class_name, description, class_type, class_duration, class_schedule, class_price, class_level, class_url, gym_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING class_id, class_name, description, class_type, class_duration, class_schedule, class_price, class_level, class_url, gym_id
    `
	err := connectdb.DB.QueryRow(
		query,
		gymclass.ClassName,
		gymclass.Description,
		gymclass.ClassType,
		gymclass.ClassDuration,
		gymclass.ClassSchedule,
		gymclass.ClassPrice,
		gymclass.ClassLevel,
        gymclass.ClassUrl,
		gymclass.GymId,
	).Scan(
		&gymclass.ClassId,
		&gymclass.ClassName,
		&gymclass.Description,
		&gymclass.ClassType,
		&gymclass.ClassDuration,
		&gymclass.ClassSchedule,
		&gymclass.ClassPrice,
		&gymclass.ClassLevel,
        &gymclass.ClassUrl,
		&gymclass.GymId,
	)

	if err != nil {
		log.Printf("Failed to insert gym class: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create class"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Class created successfully",
		"gymclass": gymclass,
	})
}

func DeleteGym(c *gin.Context) {
	Param := c.Param("gymId")
	gymId, err := strconv.Atoi(Param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		return
	}

	query := `DELETE FROM gyms WHERE gym_id = $1`

	_, err = connectdb.DB.Exec(query, gymId)
	if err != nil {
		log.Printf("Failed to delete gym: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete gym"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Gym deleted successfully"})
}

func DeleteGymClass(c *gin.Context) {
	Param := c.Param("classId")
	classId, err := strconv.Atoi(Param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	query := `DELETE FROM gym_classes WHERE class_id = $1`
	_, err = connectdb.DB.Exec(query, classId)
	if err != nil {
		log.Printf("Failed to delete gym class: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete gym class"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gym class deleted successfully"})
}
