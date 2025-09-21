package functions

import (
	"bytes"
	"encoding/json"
	"io"
	"msys_api_gateway/conf/api"
	"msys_api_gateway/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func ExistingNumber(clientid string, number string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/existing-number/"+number,
		api.GET)
	// request.Params["UserId"] = strconv.Itoa(int(userid))
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "get",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	logs.Info("Raw response received is ", res)
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)

	return data
}

func NameInquiryRequest(clientid string, number string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/name-inquiry/"+number,
		api.GET)
	// request.Params["UserId"] = strconv.Itoa(int(userid))
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "get",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	logs.Info("Raw response received is ", res)
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)

	return data
}

func ValidatePin(clientid string, number string, password string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/verify-pin",
		// "/api/verify-pin",
		api.POST)
	request.Params["number"] = number
	request.Params["password"] = password
	// request.Params = {"password": strconv.Itoa(int(userid))}
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	logs.Info("Raw response received is ", res)
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}

func ResetPin(clientid string, number string, password string, newpassword string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/reset-pin",
		api.POST)
	request.Params["number"] = number
	request.Params["oldPassword"] = password
	request.Params["newPassword"] = newpassword
	// request.Params = {"password": strconv.Itoa(int(userid))}
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	logs.Info("Raw response received is ", res)
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}

func ListAccounts(clientid string, id string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/list-accounts",
		api.POST)
	request.Params["id"] = id
	// request.Params = {"password": strconv.Itoa(int(userid))}
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	logs.Info("Raw response received is ", res)
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}

func ListCustAccounts(clientid string, number string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/list-cust-accounts",
		api.POST)
	request.Params["number"] = number
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}

func ListCustAccountsV2(c *beego.Controller, clientid string, number string) (resp responses.AccountsApiDataResponse) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/list-cust-accounts",
		api.POST)
	request.Params["number"] = number
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	var data responses.AccountsApiDataResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func CreditAccount(clientid string, accountNumber string, amount string, reference string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/credit-account",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	request.Params["amount"] = amount
	request.Params["reference"] = reference

	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	logs.Info("Raw response received is ", res)
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}

func CreditAccountV2(c *beego.Controller, clientid string, accountNumber string, amount string, reference string, channel string) (resp responses.CreditAccountApiDataResponse) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/credit-account",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	request.Params["amount"] = amount
	request.Params["reference"] = reference
	request.Params["channel"] = channel

	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	var data responses.CreditAccountApiDataResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data
	// var data responses.NameInquiryResponse

	return data
}

func DebitAccountV2(c *beego.Controller, clientid string, accountNumber string, amount string, reference string, channel string) (resp responses.CreditAccountApiDataResponse) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/debit-account",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	request.Params["amount"] = amount
	request.Params["reference"] = reference
	request.Params["channel"] = channel

	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	var data responses.CreditAccountApiDataResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data
	// var data responses.NameInquiryResponse

	return data
}

func CreditSharesAccountV2(c *beego.Controller, clientid string, accountNumber string, amount string, reference string, channel string) (resp responses.CreditAccountApiDataResponse) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/credit-shares-account",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	request.Params["amount"] = amount
	request.Params["reference"] = reference
	request.Params["channel"] = channel

	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	var data responses.CreditAccountApiDataResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data
	// var data responses.NameInquiryResponse

	return data
}

func AccountBalance(clientid string, accountNumber string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/account-balance",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	logs.Info("Raw response received is ", res)
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}

func AccountBalanceV2(c *beego.Controller, clientid string, accountNumber string) (resp responses.AccountBalanceApiResponse) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/account-balance",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	var data responses.AccountBalanceApiResponse
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func GetContactInfo(clientid string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/get-contact-info",
		api.POST)

	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	logs.Info("Raw response received is ", res)
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}

func FieldDeposit(clientid string, accountNumber string, amount string, mobileNumber string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/field-deposit",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	request.Params["amount"] = amount
	request.Params["mobileNumber"] = mobileNumber
	// request.Params = {"password": strconv.Itoa(int(userid))}
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	logs.Info("Raw response received is ", res)
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}

func RegisterCustomer(clientid string, firstName string, lastName string, gender string, mobileNumber string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/register-customer",
		api.POST)
	request.Params["firstName"] = firstName
	request.Params["lastName"] = lastName
	request.Params["gender"] = gender
	request.Params["mobileNumber"] = mobileNumber
	// request.Params = {"password": strconv.Itoa(int(userid))}
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	logs.Info("Raw response received is ", res)
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}

func CreateFieldAccount(clientid string, firstName string, lastName string, gender string, mobileNumber string, agentMobileNumber string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/create-field-account",
		api.POST)
	request.Params["firstName"] = firstName
	request.Params["lastName"] = lastName
	request.Params["gender"] = gender
	request.Params["mobileNumber"] = mobileNumber
	request.Params["agentMobileNumber"] = agentMobileNumber
	// request.Params = {"password": strconv.Itoa(int(userid))}
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	logs.Info("Raw response received is ", res)
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}

func VerifyCustomer(clientid string, username string, firstName string, lastName string, gender string, mobileNumber string, email string, dob string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/verify-customer",
		api.POST)
	request.Params["firstName"] = firstName
	request.Params["lastName"] = lastName
	request.Params["gender"] = gender
	request.Params["mobileNumber"] = mobileNumber
	request.Params["username"] = username
	request.Params["email"] = email
	request.Params["dob"] = dob
	// request.Params = {"password": strconv.Itoa(int(userid))}
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}

func FetchVerifiedCustomers(c *beego.Controller, clientid string) (resp responses.ApprovedCustomersApiResponse) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/fetch-approved-accounts",
		api.GET)

	// request.Params = {"password": strconv.Itoa(int(userid))}
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	var data responses.ApprovedCustomersApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Response data is ", data)

	return data
}

func ActivateApprovedCustomer(clientid string, username string, mobileNumber string) (resp interface{}) {
	host := HostMapping(clientid)

	logs.Info("Sending client ID ", clientid)

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/activate-verified-customer",
		api.POST)

	request.Params["mobileNumber"] = mobileNumber
	request.Params["username"] = username

	// request.Params = {"password": strconv.Itoa(int(userid))}
	logs.Debug("Request to be sent is ", request)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	data := map[string]interface{}{}
	// var data responses.NameInquiryResponse
	json.Unmarshal(read, &data)

	return data
}
