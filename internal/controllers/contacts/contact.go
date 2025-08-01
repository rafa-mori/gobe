// Package contacts provides the ContactController for handling contact form submissions.
package contacts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	ci "github.com/rafa-mori/gobe/internal/interfaces"
	t "github.com/rafa-mori/gobe/internal/types"
	gl "github.com/rafa-mori/gobe/logger"
)

type SMTPConfig struct {
	Host string
	Port string
	User string
	Pass string
}

type ContactController struct {
	queue      chan ci.ContactForm
	properties map[string]any
	APIWrapper *t.APIWrapper[ci.ContactForm]
}

func NewContactController(properties map[string]any) *ContactController {
	return &ContactController{
		queue:      make(chan ci.ContactForm, 100),
		properties: properties,
		APIWrapper: t.NewApiWrapper[ci.ContactForm](),
	}
}

func (c *ContactController) HandleContact(ctx *gin.Context) {
	var form t.ContactForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error processing data"})
		gl.Log("debug", fmt.Sprintf("Error processing data: %v", err.Error()))
		return
	}

	envT := c.properties["env"].(*t.Property[ci.IEnvironment])
	env := envT.GetValue()
	secretToken := env.Getenv("SECRET_TOKEN")

	if form.Token != secretToken {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid token"})
		gl.Log("warn", fmt.Sprintf("Invalid token: %s", form.Token))
		return
	}

	if err := sendEmailWithRetry(c, form, 2); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
		gl.Log("debug", fmt.Sprintf("Error sending email: %v", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message sent successfully!"})
	gl.Log("success", "Message sent successfully!")
}

func (c *ContactController) GetContact(ctx *gin.Context) {
	var form t.ContactForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error processing data"})
		gl.Log("debug", fmt.Sprintf("Error processing data: %v", err.Error()))
		return
	}

	envT := c.properties["env"].(*t.Property[ci.IEnvironment])
	env := envT.GetValue()
	secretToken := env.Getenv("SECRET_TOKEN")

	if form.Token != secretToken {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid token"})
		gl.Log("warn", fmt.Sprintf("Invalid token: %s", form.Token))
		return
	}

	if err := sendEmailWithRetry(c, form, 2); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
		gl.Log("debug", fmt.Sprintf("Error sending email: %v", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message sent successfully!"})
	gl.Log("success", "Message sent successfully!")
}

func (c *ContactController) PostContact(ctx *gin.Context) {
	var form t.ContactForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error processing data"})
		gl.Log("debug", fmt.Sprintf("Error processing data: %v", err.Error()))
		return
	}

	envT := c.properties["env"].(*t.Property[ci.IEnvironment])
	env := envT.GetValue()
	secretToken := env.Getenv("SECRET_TOKEN")

	if form.Token != secretToken {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid token"})
		gl.Log("warn", fmt.Sprintf("Invalid token: %s", form.Token))
		return
	}

	if err := sendEmailWithRetry(c, form, 2); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
		gl.Log("debug", fmt.Sprintf("Error sending email: %v", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message sent successfully!"})
	gl.Log("success", "Message sent successfully!")
}

func (c *ContactController) GetContactForm(ctx *gin.Context) {
	var form t.ContactForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error processing data"})
		gl.Log("debug", fmt.Sprintf("Error processing data: %v", err.Error()))
		return
	}

	envT := c.properties["env"].(*t.Property[ci.IEnvironment])
	env := envT.GetValue()
	secretToken := env.Getenv("SECRET_TOKEN")

	if form.Token != secretToken {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid token"})
		gl.Log("warn", fmt.Sprintf("Invalid token: %s", form.Token))
		return
	}

	if err := sendEmailWithRetry(c, form, 2); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
		gl.Log("debug", fmt.Sprintf("Error sending email: %v", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message sent successfully!"})
	gl.Log("success", "Message sent successfully!")
}

func (c *ContactController) GetContactFormByID(ctx *gin.Context) {
	var form t.ContactForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error processing data"})
		gl.Log("debug", fmt.Sprintf("Error processing data: %v", err.Error()))
		return
	}

	envT := c.properties["env"].(*t.Property[ci.IEnvironment])
	env := envT.GetValue()
	secretToken := env.Getenv("SECRET_TOKEN")

	if form.Token != secretToken {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid token"})
		gl.Log("warn", fmt.Sprintf("Invalid token: %s", form.Token))
		return
	}

	if err := sendEmailWithRetry(c, form, 2); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending email"})
		gl.Log("debug", fmt.Sprintf("Error sending email: %v", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message sent successfully!"})
	gl.Log("success", "Message sent successfully!")
}
