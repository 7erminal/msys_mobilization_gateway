package controllers

import (
	"encoding/json"
	"msys_api_gateway/controllers/functions"

	"msys_api_gateway/structs/requests"
	"msys_api_gateway/structs/responses"

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
	c.Mapping("VerifyCustomer", c.VerifyCustomer)
	c.Mapping("FetchApprovedCustomers", c.FetchApprovedCustomers)
	c.Mapping("ActivateVerifiedCustomers", c.ActivateVerifiedCustomers)
	c.Mapping("ListCustAccountsV2", c.ListCustAccountsV2)
	c.Mapping("GetAccountStatment", c.GetAccountStatment)
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

	logs.Debug("ID is ", v.Id)

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

	logs.Debug("Number is ", v.Number)

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

	// logs.Debug("Request::: ", c.Ctx.Input.RequestBody)
	logs.Debug("Credit Account:::: Account number:: ", v.AccountNumber, " Amount:: ", v.Amount, " Reference:: ", v.Reference)

	resp := functions.CreditAccount(clientId, v.AccountNumber, v.Amount, v.Reference)

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

// Account Balance V2 ...
// @Title AccountBalanceV2
// @Description account balance version 2
// @Param	body		body 	requests.AccountNumber	true		"body for account balance"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/account-balance [post]
func (c *Service_requestsController) AccountBalanceV2() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.AccountNumber
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.AccountBalanceV2(&c.Controller, clientId, v.AccountNumber)

	logs.Debug("Response is ", resp)

	var response responses.AccountBalanceResponse

	if resp.Data.StatusCode == 200 {
		logs.Info("Successfully fetched account balance")
		var accountBal responses.AccountBalanceData

		if resp.Data.Result != nil {
			for _, account := range *resp.Data.Result {
				accountBal = responses.AccountBalanceData(account)
			}
		} else {
			accountBal = responses.AccountBalanceData{}
		}

		response = responses.AccountBalanceResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     &accountBal,
			Client:     resp.Data.Client,
		}
	} else {
		logs.Error("Error fetching account balance")
		response = responses.AccountBalanceResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     nil,
			Client:     resp.Data.Client,
		}
	}

	c.Data["json"] = response

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

// Verify Customer ...
// @Title VerifyCustomer
// @Description Register customer
// @Param	body		body 	requests.VerifyCustomer	true		"body for verifying customers"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /verify-customer [post]
func (c *Service_requestsController) VerifyCustomer() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.VerifyCustomer
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	reqBody, err := json.Marshal(v)
	if err != nil {
		logs.Error("Error marshalling request body: %v", err)
	} else {
		logs.Debug("VerifyCustomer request: %s", string(reqBody))
	}

	resp := functions.VerifyCustomer(clientId, v.Username, v.FirstName, v.LastName, v.Gender, v.MobileNumber, v.Email, v.Dob)

	logs.Debug("Response is ", resp)
	var response responses.VerifyCustomerResponse
	if resp == nil {
		logs.Error("Error verifying customer")
		response = responses.VerifyCustomerResponse{StatusCode: 500, StatusDesc: "Failed to verify customer", Result: ""}
	} else {
		// Type assert resp to the expected struct type
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
				}
			} else {
				logs.Error("Error extracting 'data' from response map")
				response = responses.VerifyCustomerResponse{StatusCode: 500, StatusDesc: "Failed to verify customer", Result: ""}
			}
		} else {
			logs.Error("Error asserting response type")
			response = responses.VerifyCustomerResponse{StatusCode: 500, StatusDesc: "Failed to verify customer", Result: ""}
		}
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// Fetch Verified Customer ...
// @Title FetchApprovedCustomers
// @Description Fetch approved customers
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /fetch-approved-customers [get]
func (c *Service_requestsController) FetchApprovedCustomers() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	resp := functions.FetchVerifiedCustomers(&c.Controller, clientId)

	logs.Debug("Response is ", resp)

	var response responses.ApprovedCustomersResponse

	if resp.Data.StatusCode == 200 {
		logs.Info("Successfully fetched approved customers")
		customerAccounts := []responses.CustomerAccountsData{}
		for _, account := range *resp.Data.Result {
			customerAccounts = append(customerAccounts, responses.CustomerAccountsData(account))
		}
		response = responses.ApprovedCustomersResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     customerAccounts,
		}
	} else {
		logs.Error("Error fetching approved customers")
		response = responses.ApprovedCustomersResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     nil,
		}
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// Activate Verified Customer ...
// @Title ActivateVerifiedCustomers
// @Description Register customer
// @Param	body		body 	requests.ActivateCustomer	true		"body for verifying customers"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /activate-verified-customers [post]
func (c *Service_requestsController) ActivateVerifiedCustomers() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.ActivateCustomer
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	resp := functions.ActivateApprovedCustomer(clientId, v.Username, v.MobileNumber)

	logs.Debug("Response is ", resp)

	var response responses.ActivateCustomerResponse
	if resp == nil {
		logs.Error("Error activating verified customer")
		response = responses.ActivateCustomerResponse{StatusCode: 500, StatusDesc: "Failed to activate verified customer", Result: ""}
	} else {
		// Type assert resp to the expected struct type
		if dataMap, ok := resp.(map[string]interface{}); ok {
			if data, ok := dataMap["data"].(map[string]interface{}); ok {

				response = responses.ActivateCustomerResponse{
					StatusCode: int(data["StatusCode"].(float64)),
					StatusDesc: data["StatusDesc"].(string),
					Result:     "SUCCESS",
				}
			} else {
				logs.Error("Error extracting 'data' from response map")
				response = responses.ActivateCustomerResponse{StatusCode: 500, StatusDesc: "Failed to activate verified customer", Result: "FAILED"}
			}
		} else {
			logs.Error("Error asserting response type")
			response = responses.ActivateCustomerResponse{StatusCode: 500, StatusDesc: "Failed to activate verified customer", Result: "FAILED"}
		}
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// List Cust Accounts version 2 ...
// @Title ListCustAccountsV2
// @Description List Customer Accounts version 2
// @Param	body		body 	requests.Number	true		"body for listing of customer accounts"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/list-cust-accounts [post]
func (c *Service_requestsController) ListCustAccountsV2() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.Number
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	logs.Debug("Number is ", v.Number)

	resp := functions.ListCustAccountsV2(&c.Controller, clientId, v.Number)

	logs.Debug("Response is ", resp)

	response := responses.AccountsResponse{}

	if resp.Data.StatusCode == 200 {
		if resp.Data.Result != nil {
			logs.Info("Successfully fetched customer accounts")
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
			}
		} else {
			logs.Info("No accounts found for the customer")
			response = responses.AccountsResponse{
				StatusCode: resp.Data.StatusCode,
				StatusDesc: "No accounts found for the customer",
				Result:     &[]responses.AccountsApiData{},
			}
		}
	} else {
		logs.Error("Error fetching customer accounts")
		response = responses.AccountsResponse{
			StatusCode: resp.Data.StatusCode,
			StatusDesc: resp.Data.StatusDesc,
			Result:     nil,
		}
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// Credit Account V2 ...
// @Title CreditAccountV2
// @Description credit account version 2
// @Param	body		body 	requests.CreditAccountRequest	true		"body for crediting of account"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/credit-account [post]
func (c *Service_requestsController) CreditAccount2() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.CreditAccountRequestV2
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	reqBody, err := json.Marshal(v)
	if err != nil {
		logs.Error("Error marshalling request body: %v", err)
	} else {
		logs.Debug("CreditSharesAccount2 request: %s", string(reqBody))
	}

	// logs.Debug("Request::: ", c.Ctx.Input.RequestBody)
	logs.Debug("Credit Account:::: Account number:: ", v.AccountNumber, " Amount:: ", v.Amount, " Reference:: ", v.Reference)

	resp := functions.CreditAccountV2(&c.Controller, clientId, v.AccountNumber, v.Amount, v.Reference, v.Channel)

	logs.Debug("Response is ", resp)

	var response responses.CreditAccountResponse

	if resp.Data.StatusCode == 200 {
		logs.Info("Successfully credited account")
		response = responses.CreditAccountResponse{
			StatusCode: 200,
			StatusDesc: "Account credited successfully",
			Result:     resp.Data.Result,
		}
	} else {
		logs.Error("Error crediting account")
		response = responses.CreditAccountResponse{
			StatusCode: 500,
			StatusDesc: "Failed to credit account",
			Result:     "",
		}
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// Credit Shares Account V2 ...
// @Title CreditSharesAccountV2
// @Description credit account version 2
// @Param	body		body 	requests.CreditAccountRequest	true		"body for crediting of account"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/credit-shares-account [post]
func (c *Service_requestsController) CreditSharesAccount2() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.CreditAccountRequestV2
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	// logs.Debug("Request::: ", c.Ctx.Input.RequestBody)
	reqBody, err := json.Marshal(v)
	if err != nil {
		logs.Error("Error marshalling request body: %v", err)
	} else {
		logs.Debug("CreditSharesAccount2 request: %s", string(reqBody))
	}
	logs.Debug("Credit Shares Account:::: Account number:: ", v.AccountNumber, " Amount:: ", v.Amount, " Reference:: ", v.Reference)

	resp := functions.CreditSharesAccountV2(&c.Controller, clientId, v.AccountNumber, v.Amount, v.Reference, v.Channel)

	logs.Debug("Response is ", resp)

	var response responses.CreditAccountResponse

	if resp.Data.StatusCode == 200 {
		logs.Info("Successfully credited account")
		response = responses.CreditAccountResponse{
			StatusCode: 200,
			StatusDesc: "Account credited successfully",
			Result:     resp.Data.Result,
		}
	} else {
		logs.Error("Error crediting account")
		response = responses.CreditAccountResponse{
			StatusCode: 500,
			StatusDesc: "Failed to credit account",
			Result:     "",
		}
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// Debit Account V2 ...
// @Title DebitAccountV2
// @Description Debit account version 2
// @Param	body		body 	requests.DebitAccountRequestV2	true		"body for crediting of account"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/debit-account [post]
func (c *Service_requestsController) DebitAccount2() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.DebitAccountRequestV2
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	// logs.Debug("Request::: ", c.Ctx.Input.RequestBody)
	reqBody, err := json.Marshal(v)
	if err != nil {
		logs.Error("Error marshalling request body: %v", err)
	} else {
		logs.Debug("DebitAccountRequestV2 request: %s", string(reqBody))
	}
	logs.Debug("Credit Account:::: Account number:: ", v.AccountNumber, " Amount:: ", v.Amount, " Reference:: ", v.Reference)

	resp := functions.DebitAccountV2(&c.Controller, clientId, v.AccountNumber, v.Amount, v.Reference, v.Channel)

	logs.Debug("Response is ", resp)

	var response responses.DebitAccountResponse

	if resp.Data.StatusCode == 200 {
		logs.Info("Successfully debited account")
		response = responses.DebitAccountResponse{
			StatusCode: 200,
			StatusDesc: "Account debited successfully",
			Result:     resp.Data.Result,
		}
	} else {
		logs.Error("Error debiting account")
		response = responses.DebitAccountResponse{
			StatusCode: 500,
			StatusDesc: "Failed to debit account",
			Result:     "",
		}
	}

	c.Data["json"] = response

	c.ServeJSON()
}

// Get Account Statement ...
// @Title GetAccountStatment
// @Description Get account statement
// @Param	body		body 	requests.AccountStatementRequestV2	true		"body for crediting of account"
// @Param	clientId		header	true		"header for requests"
// @Success 201 {object} models.Service_requests
// @Failure 403 body is empty
// @router /v2/account-statement [post]
func (c *Service_requestsController) GetAccountStatment() {
	clientId := c.Ctx.Input.Header("clientId")
	logs.Debug("Client id is ", clientId)

	var v requests.AccountStatementRequestV2
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	// logs.Debug("Request::: ", c.Ctx.Input.RequestBody)
	reqBody, err := json.Marshal(v)
	if err != nil {
		logs.Error("Error marshalling request body: %v", err)
	} else {
		logs.Debug("GetAccountStatment request: %s", string(reqBody))
	}
	logs.Debug("Account Statement:::: Account number:: ", v.AccountNumber, " From:: ", v.FromDate, " To:: ", v.ToDate)

	resp := functions.AccountStatementV2(&c.Controller, clientId, v.AccountNumber, v.FromDate, v.ToDate)

	logs.Debug("Response is ", resp)

	var response responses.AccountStatementResponse

	if resp.StatusCode == 200 {
		logs.Info("Successfully fetched account statement")
		accountStatements := []responses.AccountStatementData{}
		if resp.Result != nil {
			for _, statement := range *resp.Result {
				accountStatement := responses.AccountStatementData{
					TransactionDate:   statement.TransactionDate,
					TransactionType:   statement.TransactionType,
					TransactionAmount: statement.TransactionAmount,
					Balance:           statement.Balance,
					Narration:         statement.Narration,
					Reference:         statement.Reference,
				}
				accountStatements = append(accountStatements, accountStatement)
			}
		}
		response = responses.AccountStatementResponse{
			StatusCode: resp.StatusCode,
			StatusDesc: resp.StatusDesc,
			Result:     &accountStatements,
		}
	} else {
		logs.Error("Error fetching account statement")
		response = responses.AccountStatementResponse{
			StatusCode: resp.StatusCode,
			StatusDesc: resp.StatusDesc,
			Result:     nil,
		}
	}

	c.Data["json"] = response

	c.ServeJSON()
}
