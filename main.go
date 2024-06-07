package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Constants for generating short URLs
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const shortURLLength = 8

// URL struct represents a URL entry
type URL struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

// Initialize the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// generateShortURL generates a random short URL
func generateShortURL() string {
	b := make([]byte, shortURLLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// shorten_url handles shortening a URL
func shorten_url(c *gin.Context) {
	var url URL
	// Bind the JSON request to URL struct
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate short URL
	url.ShortURL = generateShortURL()
	// Prepare SQL statement to insert into database
	stmt, err := db.Prepare(`INSERT INTO urls(short_url,original_url) VALUES(?,?)`)
	if err != nil {
		log.Fatal(err)
	}

	// Execute SQL statement
	_, err = stmt.Exec(url.ShortURL, url.OriginalURL)
	if err != nil {
		log.Fatal(err)
	}

	// Set short URL in Redis cache
	setKey(url.ShortURL, url.OriginalURL, 24*time.Hour)

	// Respond with the shortened URL
	c.JSON(http.StatusOK, url)
}

// getOriginalURL resolves a short URL to its original URL
func getOriginalURL(c *gin.Context) {
	shortURL := c.Param("short_url")

	originalURL, err := getKey(shortURL)
	if err == redis.Nil {
		// If not found in cache, query from database
		row := db.QueryRow(`SELECT original_url FROM urls WHERE short_url = ?`, shortURL)
		err := row.Scan(&originalURL)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		// Set in cache
		setKey(shortURL, originalURL, 24*time.Hour)
	} else if err != nil {
		log.Fatal(err)
	}

	// Redirect to the original URL
	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func main() {
	router := gin.Default()

	// Initialize database and Redis
	initializeDb()
	initializeRedis()

	// Routes
	router.POST("/shorten", shorten_url)
	router.GET("/:short_url", getOriginalURL)

	// Start server
	fmt.Fprintln(gin.DefaultWriter, "Server is running on port 8080")
	router.Run(":8080")
}
