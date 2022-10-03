package response

import "github.com/gin-gonic/gin"

type APIResponse struct {
	StatusCode int         `json:"-"`
	Result     bool        `json:"result"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func (r *APIResponse) Send(c *gin.Context) {
	c.JSON(r.StatusCode, r)
	return
}
