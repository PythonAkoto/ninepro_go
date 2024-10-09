package handler

import (
	"net/http"

	"github.com/PythonAkoto/ninepro_go/pkg/mail"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "nine_pro_base.html", nil)
}

func Contact(c *gin.Context) {
	userName := c.PostForm("user_name")
	userEmail := c.PostForm("user_email")
	userPhone := c.PostForm("user_phone")
	userMessage := c.PostForm("user_message")

	// Send email
	err := mail.SendEmail(userName, userEmail, userPhone, userMessage)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "nine_pro_base.html", gin.H{
			"error": "Failed to send message: " + err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "nine_pro_base.html", gin.H{
		"success": "Message sent successfully!",
	})
}
