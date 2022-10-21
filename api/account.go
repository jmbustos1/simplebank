package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jmbustos1/simplebank/db/sqlc"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"require"`
	Currency string `json:"currency" binding:"require,oneof=USD EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil { //client send invalid data
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
