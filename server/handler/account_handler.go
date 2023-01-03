package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/model"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/repository"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/response"
	"github.com/pankaj-katyare-wiz/airway-cargo-shipping-tracking/server/utils"

	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
)

type AccountHandler struct {
	DB      *sqlx.DB
	queries *repository.Queries
}

func NewAccountHandler(DB *sqlx.DB) *AccountHandler {
	return &AccountHandler{
		DB:      DB,
		queries: repository.New(DB),
	}
}

func (handler AccountHandler) CreateAccount(context *gin.Context) {

	var account model.Account

	if err := context.ShouldBind(&account); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		fmt.Println("error", err)
		return
	}

	_, err := handler.queries.CreateAccountDetails(context, repository.CreateAccountDetailsParams{
		ID:          uuid.New().String(),
		Name:        sql.NullString{String: account.Name, Valid: true},
		Email:       account.Email,
		CompanyName: sql.NullString{String: account.CompanyName, Valid: true},
		Mobile:      sql.NullString{String: account.Mobile, Valid: true},
		Roles:       sql.NullString{String: account.Roles, Valid: true},
		City:        sql.NullString{String: account.City, Valid: true},
		Password:    utils.GetHash(account.Password),
	})

	if err != nil {
		fmt.Println("error", err)
		response.ErrorResponse(context, http.StatusBadRequest, "Error in inserting account")
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Customer Created successfully",
	})
}

func (handler AccountHandler) Login(context *gin.Context) {

	var login model.LoginRequest

	if err := context.ShouldBind(&login); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		fmt.Println("error", err)
		return
	}

	if login.Email == "" {
		response.ErrorResponse(context, http.StatusNotFound, "ID not specified")
		return
	}
	if login.Password == "" {
		response.ErrorResponse(context, http.StatusNotFound, "password not specified")
		return
	}

	state, err := handler.queries.Login(context, repository.LoginParams{
		Email:    login.Email,
		Password: utils.GetHash(login.Password),
	})

	if err != nil {
		fmt.Println("error", err)
		response.ErrorResponse(context, http.StatusNotFound, "Invalid credentials")
		return
	}

	token, err := jwt.Sign(jwt.HS256, []byte(os.Getenv("SECRET")), model.TokenData{
		CustomerID: state.ID,
		Role:       state.Roles.String,
	}, jwt.MaxAge(2*time.Hour))
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		response.ErrorResponse(context, http.StatusInternalServerError, "Error generating token")
		return
	}

	context.Header("auth-bearer", string(token))
	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "login successfully",
	})
}

func (handler AccountHandler) GetAccountByID(context *gin.Context) {

	id := context.Request.URL.Query().Get("id")
	if id == "" {
		response.ErrorResponse(context, http.StatusNotFound, "ID not specified")
		return
	}

	state, err := handler.queries.GetAccountDetails(context, id)

	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, err.Error())
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "account Data",
		"data":    state,
	})
}

func (handler AccountHandler) UpdateAccountDetails(context *gin.Context) {

	var account model.Account

	if err := context.ShouldBind(&account); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, "Required fields are empty")
		return
	}

	err := handler.queries.UpdateAccountDetails(context, repository.UpdateAccountDetailsParams{
		ID:          account.Id,
		Name:        sql.NullString{String: account.Name, Valid: true},
		CompanyName: sql.NullString{String: account.CompanyName, Valid: true},
		Email:       account.Email,
		Mobile:      sql.NullString{String: account.Mobile, Valid: true},
		Roles:       sql.NullString{String: account.Roles, Valid: true},
		City:        sql.NullString{String: account.City, Valid: true},
		Password:    utils.GetHash(account.Password),
	})

	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, "Error on Update account details")
		fmt.Println("error", err)
		return
	}
	// TODO return, nothing to update
	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Updated suceessfully",
	})
}

func (handler AccountHandler) GetAllAccount(context *gin.Context) {

	quotes, err := handler.queries.ListAccountDetails(context)

	if err != nil {
		response.SuccessResponse(context, map[string]interface{}{
			"code":    "success",
			"message": "Error int get all account",
			"error":   err,
		})
		return
	}

	response.SuccessResponse(context, map[string]interface{}{
		"code":    "success",
		"message": "Fetched all account list",
		"data":    quotes,
	})
}

func AuthorizationMiddleware(c *gin.Context) {
	bearer := c.Request.Header.Get("Authorization")
	if bearer == "" {
		fmt.Println("Authorization token not found in request")
		c.Status(http.StatusUnauthorized)
		return
	}

	token := strings.Split(bearer, " ")[1]
	verifiedToken, err := jwt.Verify(jwt.HS256, []byte(os.Getenv("SECRET")), []byte(token))
	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusUnauthorized)
		return
	}

	var claims model.TokenData

	if err = verifiedToken.Claims(&claims); err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusUnauthorized)
		return
	}

	c.Set("claims", claims)

	c.Next()
}
