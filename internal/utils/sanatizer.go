package utils

import (
	"bytes"
	"encoding/json"
	"html"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func sanitizeInput(input string) string {
	sanitized := html.EscapeString(input)
	sanitized = strings.ReplaceAll(sanitized, "'", "''")

	return sanitized
}

func SanitizeMiddleware(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		contentType := c.GetHeader("Content-Type")

		if strings.Contains(contentType, "application/json") {
			// JSON payload
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
				return
			}

			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			var data map[string]interface{}
			if err := json.Unmarshal(body, &data); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
				return
			}

			for key, value := range data {
				if strValue, ok := value.(string); ok {
					data[key] = sanitizeInput(strValue)
				}
			}

			sanitizedBody, err := json.Marshal(data)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Sanitization error"})
				return
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(sanitizedBody))
		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			// Form data
			if err := c.Request.ParseForm(); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
				return
			}

			for key, values := range c.Request.PostForm {
				for i, value := range values {
					c.Request.PostForm[key][i] = sanitizeInput(value)
				}
			}
		}
	}

	c.Next()
}
