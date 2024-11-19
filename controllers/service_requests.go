package controllers

import (
	"encoding/json"
	"msys_api_gateway/controllers/functions"

	"msys_api_gateway/structs/requests"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Api requests
type Service_requestsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Service_requestsController) URLMapping() {
	c.Mapping("NameInquiry", c.NameInquiry)
	c.Mapping("VerifyPin", c.VerifyPin)
	c.Mapping("ResetPin", c.ResetPin)
	c.Mapping("ListAccounts", c.ListAccounts)
	c.Mapping("ListCustAccounts", c.ListCustAccounts)
	c.Mapping("CreditAccount", c.CreditAccount)
	c.Mapping("AccountBalance", c.AccountBalance)
	c.Mapping("GetAccountInfo", c.GetAccountInfo)
	c.Mapping("FieldDeposit", c.FieldDeposit)
	c.Mapping("RegisterCustomer", c.RegisterCustomer)
	c.Mapping("CreateFieldAccount", c.CreateFieldAccount)
	c.Mapping("ExistingNumber", c.ExistingNumber)
}

// Name Inquiry ...
// @Title NameInquiry
// @Description name inquiry with number
// @Param	number		path 	string	true		"The key for staticblock"
// @Param	clientId		header	true		"header for requests"
// @Success 200 {object} models.Service_requests
// @Failure 403 :number is empty
// @router /name-inquiry/:number [get]
func (c *Service_requestsController) NameInquiry() {
	clientId := c.Ctx.Input.Header("clientId")
	number := c.Ctx.Input.Param(":number")
	logs.Debug("Client id is ", clientId)
	logs.Debug("Number is ", number)

	resp := functions.NameInquiryRequest(clientId, number)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}

// Verify Pin ...
// @Title VerifyPin
// @Description verify pin
// @Param	body		body 	requests.ValidatePin	true		"body for verify pin"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /verify-pin [post]
func (c *Service_requestsController) VerifyPin() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.ValidatePin
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.ValidatePin(clientId, v.Number, v.Password)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}

// Reset Pin ...
// @Title ResetPin
// @Description Reset pin
// @Param	body		body 	requests.ResetPin	true		"body for Reset pin"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /reset-pin [post]
func (c *Service_requestsController) ResetPin() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.ResetPin
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.ResetPin(clientId, v.Number, v.OldPassword, v.NewPassword)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}

// List Accounts ...
// @Title ListAccounts
// @Description list first 3 accounts
// @Param	body		body 	requests.Id	true		"body for listing of accounts"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /list-accounts [post]
func (c *Service_requestsController) ListAccounts() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.Id
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.ListAccounts(clientId, v.Id)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}

// List Cust Accounts ...
// @Title ListCustAccounts
// @Description List Customer Accounts
// @Param	body		body 	requests.Number	true		"body for listing of customer accounts"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /list-cust-accounts [post]
func (c *Service_requestsController) ListCustAccounts() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.Number
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.ListCustAccounts(clientId, v.Number)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}

// Credit Account ...
// @Title CreditAccount
// @Description credit account
// @Param	body		body 	requests.CreditAccountRequest	true		"body for crediting of account"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /credit-account [post]
func (c *Service_requestsController) CreditAccount() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.CreditAccountRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	logs.Debug("Account number is ", v.AccountNumber)
	logs.Debug("Amount is ", v.Amount)

	resp := functions.CreditAccount(clientId, v.AccountNumber, v.Amount)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}

// Account Balance ...
// @Title AccountBalance
// @Description account balance
// @Param	body		body 	requests.AccountNumber	true		"body for account balance"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /account-balance [post]
func (c *Service_requestsController) AccountBalance() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.AccountNumber
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.AccountBalance(clientId, v.AccountNumber)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}

// Get Account Info ...
// @Title GetAccountInfo
// @Description get account info
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /get-contact-info [post]
func (c *Service_requestsController) GetAccountInfo() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	resp := functions.GetContactInfo(clientId)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}

// Field Deposit ...
// @Title FieldDeposit
// @Description verify pin
// @Param	body		body 	requests.FieldDepositRequest	true		"body for field deposit"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /field-deposit [post]
func (c *Service_requestsController) FieldDeposit() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.FieldDepositRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.FieldDeposit(clientId, v.AccountNumber, v.Amount, v.MobileNumber)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}

// Register Customer ...
// @Title RegisterCustomer
// @Description Register customer
// @Param	body		body 	requests.RegisterCustomer	true		"body for registering customers"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /register-customer [post]
func (c *Service_requestsController) RegisterCustomer() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.RegisterCustomer
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.RegisterCustomer(clientId, v.FirstName, v.LastName, v.Gender, v.MobileNumber)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}

// Create Field Account ...
// @Title CreateFieldAccount
// @Description create field account
// @Param	body		body 	requests.CreateFieldAccount	true		"body for creating field deposit"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /create-field-account [post]
func (c *Service_requestsController) CreateFieldAccount() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.CreateFieldAccount
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.CreateFieldAccount(clientId, v.FirstName, v.LastName, v.Gender, v.MobileNumber, v.AgentMobileNumber)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}

// Existing number ...
// @Title ExistingNumber
// @Description Check if a number exists
// @Param	number		path 	string	true		"The key for staticblock"
// @Param	clientId		header	true		"header for requests"
// @Success 200 {object} models.Service_requests
// @Failure 403 :number is empty
// @router /existing-number/:number [get]
func (c *Service_requestsController) ExistingNumber() {
	clientId := c.Ctx.Input.Header("clientId")
	number := c.Ctx.Input.Param(":number")
	logs.Debug("Client id is ", clientId)
	logs.Debug("Number is ", number)

	resp := functions.ExistingNumber(clientId, number)

	logs.Debug("Response is ", resp)

	c.Data["json"] = resp

	c.ServeJSON()
}
