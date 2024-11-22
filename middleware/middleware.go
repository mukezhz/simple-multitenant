package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sachin-gautam/gin-api/constants"
	"github.com/sachin-gautam/gin-api/database"
)

type TenantMiddleware struct {
	db *database.Database
}

func NewTenantMiddleware(db *database.Database) *TenantMiddleware {
	return &TenantMiddleware{
		db,
	}
}

func (t *TenantMiddleware) ExtractTenantIDFromDomain() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		host := c.Request.Host
		domain := origin
		if domain == "" {
			domain = host
		} else if strings.Contains(domain, constants.HTTP) {
			domain = strings.Replace(domain, constants.HTTP, "", -1)
		} else if strings.Contains(domain, constants.HTTPS) {
			domain = strings.Replace(domain, constants.HTTPS, "", -1)
		}
		tenantID := t.db.FindTenantIDByDomain(domain)
		c.Set(constants.TenantID, tenantID)
		c.Next()
	}
}
