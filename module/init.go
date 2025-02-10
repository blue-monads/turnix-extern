package module

import (
	"github.com/blue-monads/turnix/backend/registry"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
)

func init() {
	registry.Register("externtest", New)
}

func New(opt xproject.BuilderOption) (*xproject.Defination, error) {

	def := &xproject.Defination{
		Name:                "Externtest",
		Slug:                "externtest",
		Info:                "This is externl test",
		Icon:                "book-open",
		NewFormSchemaFields: []xproject.PTypeField{},
		Version:             "1.0.0",
		Perminssions:        []string{"read", "write", "read/write"},
		OnAPIMount: func(rg *gin.RouterGroup) {
			rg.GET("/test", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Hello World",
				})
			})
		},
	}

	return def, nil
}
