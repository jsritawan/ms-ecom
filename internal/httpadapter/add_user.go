package httpadapter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsritawan/ms-ecom/ecomapi"
)

func (a *Adapter) AddUser(c *gin.Context) {
	var req ecomapi.SignUpUserRequest
	if err := c.Bind(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("Could not bind request"))
	}

	resp, err := a.userService.AddNewUser(c.Request.Context(), &req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, resp)
}
