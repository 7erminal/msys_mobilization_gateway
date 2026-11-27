package controllers

import (
	"context"
	"encoding/json"
	"log/slog"
	"msys_api_gateway/controllers/functions"
	"msys_api_gateway/structs/requests"
	"msys_api_gateway/structs/responses"
	beego "github.com/beego/beego/v2/server/web"

	gctrace "github.com/msys/go-common/trace"
)

// Api requests
type Service_requestsController struct {
	beego.Controller
}

// reqCtx returns the request context carrying the validated trace IDs and State.
func (c *Service_requestsController) reqCtx() context.Context { return c.Ctx.Request.Context() }

// log returns the context-aware, structured logger for the gateway service.
func (c *Service_requestsController) log() *slog.Logger {
	return gctrace.Logger(c.reqCtx(), "gateway")
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
	c.Mapping("VerifyCustomer", c.VerifyCustomer)
	c.Mapping("FetchApprovedCustomers", c.FetchApprovedCustomers)
	c.Mapping("ActivateVerifiedCustomers", c.ActivateVerifiedCustomers)
	c.Mapping("ListCustAccountsV2", c.ListCustAccountsV2)
	c.Mapping("GetAccountStatment", c.GetAccountStatment)
	c.Mapping("RegisterCustomerV2", c.RegisterCustomerV2)
	c.Mapping("AccountBalanceV2", c.AccountBalanceV2)
	c.Mapping("CreditAccount2", c.CreditAccount2)
	c.Mapping("LoanRepayment", c.LoanRepayment)
	c.Mapping("CloseAccount", c.CloseAccount)
}

// Name Inquiry ...
// @Title NameInquiry
// @Description name inquiry with number
// @Param	number		path 	string	true		"The key for staticblock"
// @Param	clientId	header	true		"header for requests"
// @Success 200 {object} models.Service_requests
// @Failure 403 :number is empty
// @router /name-inquiry/:number [get]
func (c *Service_requestsController) NameInquiry() {
	clientId := c.Ctx.Input.Header("clientId")
	number := c.Ctx.Input.Param(":number")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "NameInquiry"))

	resp := functions.NameInquiryRequest(c.reqCtx(), clientId, number)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Verify Pin ...
// @Title VerifyPin
// @Description verify pin
// @Param	body		body 	requests.ValidatePin	true		"body for verify pin"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /verify-pin [post]
func (c *Service_requestsController) VerifyPin() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "VerifyPin"))

	var v requests.ValidatePin
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.ValidatePin(c.reqCtx(), clientId, v.Number, v.Password)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Reset Pin ...
// @Title ResetPin
// @Description Reset pin
// @Param	body		body 	requests.ResetPin	true		"body for Reset pin"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /reset-pin [post]
func (c *Service_requestsController) ResetPin() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "ResetPin"))

	var v requests.ResetPin
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.ResetPin(c.reqCtx(), clientId, v.Number, v.Password, v.NewPassword)
	c.Data["json"] = resp
	c.ServeJSON()
}

// List Accounts ...
// @Title ListAccounts
// @Description list first 3 accounts
// @Param	body		body 	requests.Id	true		"body for listing of accounts"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /list-accounts [post]
func (c *Service_requestsController) ListAccounts() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "ListAccounts"))

	var v requests.Id
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.ListAccounts(c.reqCtx(), clientId, v.Id)
	c.Data["json"] = resp
	c.ServeJSON()
}

// List Cust Accounts ...
// @Title ListCustAccounts
// @Description List Customer Accounts
// @Param	body		body 	requests.Number	true		"body for listing of customer accounts"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /list-cust-accounts [post]
func (c *Service_requestsController) ListCustAccounts() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "ListCustAccounts"))

	var v requests.Number
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.ListCustAccounts(c.reqCtx(), clientId, v.Number)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Credit Account ...
// @Title CreditAccount
// @Description credit account
// @Param	body		body 	requests.CreditAccountRequest	true		"body for crediting of account"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /credit-account [post]
func (c *Service_requestsController) CreditAccount() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "CreditAccount"))

	var v requests.CreditAccountRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.CreditAccount(c.reqCtx(), clientId, v.AccountNumber, v.Amount, v.Reference)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Account Balance ...
// @Title AccountBalance
// @Description account balance
// @Param	body		body 	requests.AccountNumber	true		"body for account balance"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /account-balance [post]
func (c *Service_requestsController) AccountBalance() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "AccountBalance"))

	var v requests.AccountNumber
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.AccountBalance(c.reqCtx(), clientId, v.AccountNumber)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Account Balance V2 ...
// @Title AccountBalanceV2
// @Description account balance version 2
// @Param	body		body 	requests.AccountNumber	true		"body for account balance"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/account-balance [post]
func (c *Service_requestsController) AccountBalanceV2() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "AccountBalanceV2"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	var v requests.AccountNumber
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.AccountBalanceV2(c.reqCtx(), clientId, v.AccountNumber)

	var response responses.AccountBalanceResponse

	if resp.Data.StatusCode == 200 {
		c.log().Info("successfully fetched account balance", slog.String("client_id", clientId))
		var accountBal responses.AccountBalanceData
		if resp.Data.Result != nil {
			accountBal = responses.AccountBalanceData(*resp.Data.Result)
		} else {
			accountBal = responses.AccountBalanceData{}
		}
		response = responses.AccountBalanceResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     &accountBal,
			Client:     resp.Data.Client,
			RequestID:  requestID,
		}
	} else {
		c.log().Error("error fetching account balance", slog.String("client_id", clientId), slog.Int("upstream_status", resp.Data.StatusCode))
		response = responses.AccountBalanceResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     nil,
			Client:     resp.Data.Client,
			RequestID:  requestID,
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// Get Account Info ...
// @Title GetAccountInfo
// @Description get account info
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /get-contact-info [post]
func (c *Service_requestsController) GetAccountInfo() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "GetAccountInfo"))

	resp := functions.GetContactInfo(c.reqCtx(), clientId)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Field Deposit ...
// @Title FieldDeposit
// @Description verify pin
// @Param	body		body 	requests.FieldDepositRequest	true		"body for field deposit"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /field-deposit [post]
func (c *Service_requestsController) FieldDeposit() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "FieldDeposit"))

	var v requests.FieldDepositRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	paymentMethod := "CASH"
	if v.PaymentMethod != "" {
		paymentMethod = v.PaymentMethod
	}
	c.log().Info("field-deposit request", slog.String("client_id", clientId), slog.String("payment_method", paymentMethod))

	resp := functions.FieldDeposit(c.reqCtx(), clientId, v.AccountNumber, v.Amount, v.MobileNumber, paymentMethod)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Register Customer ...
// @Title RegisterCustomer
// @Description Register customer
// @Param	body		body 	requests.RegisterCustomer	true		"body for registering customers"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /register-customer [post]
func (c *Service_requestsController) RegisterCustomer() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "RegisterCustomer"))

	var v requests.RegisterCustomer
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	source := "USSD"
	chargeAmount := "0"
	txnRef := "USSD" + v.MobileNumber

	resp := functions.RegisterCustomer(c.reqCtx(), clientId, v.FirstName, v.LastName, v.Gender, v.MobileNumber, source, chargeAmount, txnRef)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Register Customer App ...
// @Title RegisterCustomerApp
// @Description Register customer App
// @Param	body		body 	requests.RegisterCustomer	true		"body for registering customers"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /2/register-customer [post]
func (c *Service_requestsController) RegisterCustomerV2() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "RegisterCustomerV2"))

	var v requests.RegisterCustomer2
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.RegisterCustomer(c.reqCtx(), clientId, v.FirstName, v.LastName, v.Gender, v.MobileNumber, v.Source, v.ChargeAmount, v.TransactionRef)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Create Field Account ...
// @Title CreateFieldAccount
// @Description create field account
// @Param	body		body 	requests.CreateFieldAccount	true		"body for creating field deposit"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /create-field-account [post]
func (c *Service_requestsController) CreateFieldAccount() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "CreateFieldAccount"))

	var v requests.CreateFieldAccount
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.CreateFieldAccount(c.reqCtx(), clientId, v.FirstName, v.LastName, v.Gender, v.MobileNumber, v.AgentMobileNumber)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Existing number ...
// @Title ExistingNumber
// @Description Check if a number exists
// @Param	number		path 	string	true		"The key for staticblock"
// @Param	clientId	header	true		"header for requests"
// @Success 200 {object} models.Service_requests
// @Failure 403 :number is empty
// @router /existing-number/:number [get]
func (c *Service_requestsController) ExistingNumber() {
	clientId := c.Ctx.Input.Header("clientId")
	number := c.Ctx.Input.Param(":number")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "ExistingNumber"))

	resp := functions.ExistingNumber(c.reqCtx(), clientId, number)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Verify Customer ...
// @Title VerifyCustomer
// @Description Register customer
// @Param	body		body 	requests.VerifyCustomer	true		"body for verifying customers"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /verify-customer [post]
func (c *Service_requestsController) VerifyCustomer() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "VerifyCustomer"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	var v requests.VerifyCustomer
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.VerifyCustomer(c.reqCtx(), clientId, v.Username, v.FirstName, v.LastName, v.Gender, v.MobileNumber, v.Email, v.Dob)

	var response responses.VerifyCustomerResponse
	if resp == nil {
		c.log().Error("error verifying customer", slog.String("client_id", clientId))
		response = responses.VerifyCustomerResponse{StatusCode: 500, StatusDesc: "Failed to verify customer", Result: "", RequestID: requestID}
	} else {
		if dataMap, ok := resp.(map[string]interface{}); ok {
			if data, ok := dataMap["data"].(map[string]interface{}); ok {
				result := data["Result"]
				if result == nil {
					result = ""
				}
				response = responses.VerifyCustomerResponse{
					StatusCode: int(data["StatusCode"].(float64)),
					StatusDesc: data["StatusDesc"].(string),
					Result:     result.(string),
					Client:     data["Client"].(string),
					RequestID:  requestID,
				}
			} else {
				c.log().Error("error extracting 'data' from response map", slog.String("client_id", clientId))
				response = responses.VerifyCustomerResponse{StatusCode: 500, StatusDesc: "Failed to verify customer", Result: "", RequestID: requestID}
			}
		} else {
			c.log().Error("error asserting response type", slog.String("client_id", clientId))
			response = responses.VerifyCustomerResponse{StatusCode: 500, StatusDesc: "Failed to verify customer", Result: "", RequestID: requestID}
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// Fetch Verified Customer ...
// @Title FetchApprovedCustomers
// @Description Fetch approved customers
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /fetch-approved-customers [get]
func (c *Service_requestsController) FetchApprovedCustomers() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "FetchApprovedCustomers"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	resp := functions.FetchVerifiedCustomers(c.reqCtx(), clientId)

	var response responses.ApprovedCustomersResponse

	if resp.Data.StatusCode == 200 {
		c.log().Info("successfully fetched approved customers", slog.String("client_id", clientId))
		customerAccounts := []responses.CustomerAccountsData{}
		for _, account := range *resp.Data.Result {
			customerAccounts = append(customerAccounts, responses.CustomerAccountsData(account))
		}
		response = responses.ApprovedCustomersResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     customerAccounts,
			RequestID:  requestID,
		}
	} else {
		c.log().Error("error fetching approved customers", slog.String("client_id", clientId), slog.Int("upstream_status", resp.Data.StatusCode))
		response = responses.ApprovedCustomersResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     nil,
			RequestID:  requestID,
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// Activate Verified Customer ...
// @Title ActivateVerifiedCustomers
// @Description Register customer
// @Param	body		body 	requests.ActivateCustomer	true		"body for verifying customers"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /activate-verified-customers [post]
func (c *Service_requestsController) ActivateVerifiedCustomers() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "ActivateVerifiedCustomers"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	var v requests.ActivateCustomer
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.ActivateApprovedCustomer(c.reqCtx(), clientId, v.Username, v.MobileNumber)

	var response responses.ActivateCustomerResponse
	if resp == nil {
		c.log().Error("error activating verified customer", slog.String("client_id", clientId))
		response = responses.ActivateCustomerResponse{StatusCode: 500, StatusDesc: "Failed to activate verified customer", Result: "", RequestID: requestID}
	} else {
		if dataMap, ok := resp.(map[string]interface{}); ok {
			if data, ok := dataMap["data"].(map[string]interface{}); ok {
				response = responses.ActivateCustomerResponse{
					StatusCode: int(data["StatusCode"].(float64)),
					StatusDesc: data["StatusDesc"].(string),
					Result:     "SUCCESS",
					RequestID:  requestID,
				}
			} else {
				c.log().Error("error extracting 'data' from response map", slog.String("client_id", clientId))
				response = responses.ActivateCustomerResponse{StatusCode: 500, StatusDesc: "Failed to activate verified customer", Result: "FAILED", RequestID: requestID}
			}
		} else {
			c.log().Error("error asserting response type", slog.String("client_id", clientId))
			response = responses.ActivateCustomerResponse{StatusCode: 500, StatusDesc: "Failed to activate verified customer", Result: "FAILED", RequestID: requestID}
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// List Cust Accounts version 2 ...
// @Title ListCustAccountsV2
// @Description List Customer Accounts version 2
// @Param	body		body 	requests.Number	true		"body for listing of customer accounts"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/list-cust-accounts [post]
func (c *Service_requestsController) ListCustAccountsV2() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "ListCustAccountsV2"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	var v requests.Number
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.ListCustAccountsV2(c.reqCtx(), clientId, v.Number)

	response := responses.AccountsResponse{}

	if resp.Data.StatusCode == 200 {
		if resp.Data.Result != nil {
			c.log().Info("successfully fetched customer accounts", slog.String("client_id", clientId))
			accounts := []responses.AccountsApiData{}
			for _, account := range *resp.Data.Result {
				accountData := responses.AccountsApiData{
					AccountNumber: account.AccountNumber,
					Product:       account.Product,
				}
				accounts = append(accounts, accountData)
			}
			response = responses.AccountsResponse{
				StatusCode: resp.Data.StatusCode,
				StatusDesc: resp.Data.StatusDesc,
				Result:     &accounts,
				RequestID:  requestID,
			}
		} else {
			c.log().Info("no accounts found for the customer", slog.String("client_id", clientId))
			response = responses.AccountsResponse{
				StatusCode: resp.Data.StatusCode,
				StatusDesc: "No accounts found for the customer",
				Result:     &[]responses.AccountsApiData{},
				RequestID:  requestID,
			}
		}
	} else {
		c.log().Error("error fetching customer accounts", slog.String("client_id", clientId), slog.Int("upstream_status", resp.Data.StatusCode))
		response = responses.AccountsResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     nil,
			RequestID:  requestID,
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// Credit Account V2 ...
// @Title CreditAccountV2
// @Description credit account version 2
// @Param	body		body 	requests.CreditAccountRequest	true		"body for crediting of account"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/credit-account [post]
func (c *Service_requestsController) CreditAccount2() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "CreditAccount2"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	var v requests.CreditAccountRequestV2
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := responses.CreditAccountApiDataResponse{}
	if v.DepositType == "DEPOSIT" || v.DepositType == "NORMAL" {
		resp = functions.CreditAccountV2(c.reqCtx(), clientId, v.AccountNumber, v.Amount, v.Reference, v.Channel)
	} else if v.DepositType == "SHARES" {
		resp = functions.CreditSharesAccountV2(c.reqCtx(), clientId, v.AccountNumber, v.Amount, v.Reference, v.Channel)
	} else {
		c.log().Error("unsupported deposit type", slog.String("client_id", clientId), slog.String("deposit_type", v.DepositType))
		resp.Data.StatusCode = 400
		resp.Data.StatusDesc = "Unsupported deposit type"
	}

	var response responses.CreditAccountResponse

	if resp.Data.StatusCode == 200 {
		c.log().Info("successfully credited account", slog.String("client_id", clientId))
		response = responses.CreditAccountResponse{
			StatusCode: 200,
			StatusDesc: "Account credited successfully",
			Result:     resp.Data.Result,
			RequestID:  requestID,
		}
	} else {
		c.log().Error("error crediting account", slog.String("client_id", clientId), slog.Int("upstream_status", resp.Data.StatusCode))
		response = responses.CreditAccountResponse{
			StatusCode: 500,
			StatusDesc: "Failed to credit account",
			Result:     "",
			RequestID:  requestID,
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// Credit Shares Account V2 ...
// @Title CreditSharesAccountV2
// @Description credit account version 2
// @Param	body		body 	requests.CreditAccountRequest	true		"body for crediting of account"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/credit-shares-account [post]
func (c *Service_requestsController) CreditSharesAccount2() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "CreditSharesAccount2"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	var v requests.CreditAccountRequestV2
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.CreditSharesAccountV2(c.reqCtx(), clientId, v.AccountNumber, v.Amount, v.Reference, v.Channel)

	var response responses.CreditAccountResponse

	if resp.Data.StatusCode == 200 {
		c.log().Info("successfully credited account", slog.String("client_id", clientId))
		response = responses.CreditAccountResponse{
			StatusCode: 200,
			StatusDesc: "Account credited successfully",
			Result:     resp.Data.Result,
			RequestID:  requestID,
		}
	} else {
		c.log().Error("error crediting account", slog.String("client_id", clientId), slog.Int("upstream_status", resp.Data.StatusCode))
		response = responses.CreditAccountResponse{
			StatusCode: 500,
			StatusDesc: "Failed to credit account",
			Result:     "",
			RequestID:  requestID,
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// Debit Account V2 ...
// @Title DebitAccountV2
// @Description Debit account version 2
// @Param	body		body 	requests.DebitAccountRequestV2	true		"body for crediting of account"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/debit-account [post]
func (c *Service_requestsController) DebitAccount2() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "DebitAccount2"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	var v requests.DebitAccountRequestV2
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.DebitAccountV2(c.reqCtx(), clientId, v.AccountNumber, v.Amount, v.Reference, v.Channel)

	var response responses.DebitAccountResponse

	if resp.Data.StatusCode == 200 {
		c.log().Info("successfully debited account", slog.String("client_id", clientId))
		response = responses.DebitAccountResponse{
			StatusCode: 200,
			StatusDesc: "Account debited successfully",
			Result:     resp.Data.Result,
			RequestID:  requestID,
		}
	} else {
		c.log().Error("error debiting account", slog.String("client_id", clientId), slog.Int("upstream_status", resp.Data.StatusCode))
		response = responses.DebitAccountResponse{
			StatusCode: 500,
			StatusDesc: "Failed to debit account",
			Result:     "",
			RequestID:  requestID,
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// Get Account Statement ...
// @Title GetAccountStatment
// @Description Get account statement
// @Param	body		body 	requests.AccountStatementRequestV2	true		"body for crediting of account"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/account-statement [post]
func (c *Service_requestsController) GetAccountStatment() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "GetAccountStatment"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	var v requests.AccountStatementRequestV2
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.AccountStatementV2(c.reqCtx(), clientId, v.AccountNumber, v.FromDate, v.ToDate)

	var response responses.AccountStatementResponse

	if resp.Data.StatusCode == 200 {
		c.log().Info("successfully fetched account statement", slog.String("client_id", clientId))
		accountStatements := []responses.AccountStatementData{}
		if resp.Data.Result != nil {
			for _, statement := range *resp.Data.Result {
				accountStatement := responses.AccountStatementData{
					TransactionDate:        statement.TransactionDate,
					TransactionDescription: statement.TransactionDescription,
					Debit:                  statement.Debit,
					Credit:                 statement.Credit,
				}
				accountStatements = append(accountStatements, accountStatement)
			}
		}
		response = responses.AccountStatementResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     &accountStatements,
			RequestID:  requestID,
		}
	} else {
		c.log().Error("error fetching account statement", slog.String("client_id", clientId), slog.Int("upstream_status", resp.Data.StatusCode))
		response = responses.AccountStatementResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     nil,
			RequestID:  requestID,
		}
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// Get Account Loans ...
// @Title GetAccountLoans
// @Description Get account loans
// @Param	body		body 	requests.AccountLoansRequest	true		"body for crediting of account"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /list-account-loans [post]
func (c *Service_requestsController) ListAccountLoans() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "ListAccountLoans"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	var v requests.AccountLoansRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	status := 400
	statusMessage := "Error retrieving account loans"

	resp := functions.ListAccountLoans(c.reqCtx(), clientId, v.AccountNumber)

	var response responses.ListLoansResponse

	var result []responses.LoanData

	if resp.Data.StatusCode == 200 {
		c.log().Info("successfully fetched account loans", slog.String("client_id", clientId))
		status = 200
		statusMessage = "Successfully fetched account loans"
		if resp.Data.Result != nil {
			result = *resp.Data.Result
		} else {
			result = []responses.LoanData{}
			c.log().Info("no loans found for the account", slog.String("client_id", clientId))
			statusMessage = "No loans found for the account"
		}
	} else {
		status = resp.Data.StatusCode
		c.log().Error("error fetching account loans", slog.String("client_id", clientId), slog.Int("upstream_status", resp.Data.StatusCode))
		statusMessage = resp.Data.StatusDesc
	}

	response = responses.ListLoansResponse{
		StatusCode: status,
		StatusDesc: statusMessage,
		Result:     &result,
		RequestID:  requestID,
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// Loan Repayment ...
// @Title LoanRepayment
// @Description Repay loan
// @Param	body		body 	requests.LoanRepaymentRequest	true		"body for crediting of account"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /loan-repayment [post]
func (c *Service_requestsController) LoanRepayment() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "LoanRepayment"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	var v requests.LoanRepaymentRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	status := false
	statusMessage := "Error repaying loan"
	result := "Repayment failed"

	req := requests.LoanRepaymentApiRequest{
		AccountNumber: v.AccountNumber,
		Amount:        v.Amount,
		MobileNumber:  v.MobileNumber,
		LoanId:        v.LoanId,
		ClientId:      clientId,
	}

	resp := functions.LoanRepayment(c.reqCtx(), req)

	var response responses.RepayLoanResponse

	if resp.Data.StatusCode == 200 {
		c.log().Info("successfully repaid loan", slog.String("client_id", clientId))
		status = true
		statusMessage = "Successfully repaid loan"
		result = resp.Data.Result
	} else {
		c.log().Error("error repaying loan", slog.String("client_id", clientId), slog.Int("upstream_status", resp.Data.StatusCode))
		statusMessage = resp.Data.StatusDesc
	}

	response = responses.RepayLoanResponse{
		StatusCode: status,
		StatusDesc: statusMessage,
		Result:     result,
		RequestID:  requestID,
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// Close Account ...
// @Title CloseAccount
// @Description Close Account
// @Param	body		body 	requests.LoanRepaymentRequest	true		"body for crediting of account"
// @Param	clientId	header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/close-account [post]
func (c *Service_requestsController) CloseAccount() {
	clientId := c.Ctx.Input.Header("clientId")
	c.log().Debug("request received", slog.String("client_id", clientId), slog.String("action", "CloseAccount"))
	requestID := gctrace.FromContext(c.reqCtx()).RequestID

	var v requests.LoanRepaymentRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	status := false
	statusMessage := "Error retrieving account loans"
	result := "Repayment failed"

	req := requests.LoanRepaymentApiRequest{
		AccountNumber: v.AccountNumber,
		Amount:        v.Amount,
		MobileNumber:  v.MobileNumber,
		LoanId:        v.LoanId,
		ClientId:      clientId,
		PaymentMode:   v.PaymentMode,
	}

	resp := functions.LoanRepayment(c.reqCtx(), req)

	var response responses.RepayLoanResponse

	if resp.Data.StatusCode == 200 {
		c.log().Info("successfully closed account", slog.String("client_id", clientId))
		status = true
		statusMessage = "Successfully fetched account loans"
		result = resp.Data.Result
	} else {
		c.log().Error("error closing account", slog.String("client_id", clientId), slog.Int("upstream_status", resp.Data.StatusCode))
		statusMessage = resp.Data.StatusDesc
	}

	response = responses.RepayLoanResponse{
		StatusCode: status,
		StatusDesc: statusMessage,
		Result:     result,
		RequestID:  requestID,
	}

	c.Data["json"] = response
	c.ServeJSON()
}
