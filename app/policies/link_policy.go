package policies

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/link"
)

func CanModifyLink(c *gin.Context, linkModel link.Link) bool {
	return true
}

// func CanViewLink(c *gin.Context, linkModel link.Link) bool {}
// func CanCreateLink(c *gin.Context, linkModel link.Link) bool {}
// func CanUpdateLink(c *gin.Context, linkModel link.Link) bool {}
// func CanDeleteLink(c *gin.Context, linkModel link.Link) bool {}
