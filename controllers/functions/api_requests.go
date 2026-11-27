package functions

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"msys_api_gateway/conf/api"
	"msys_api_gateway/structs/requests"
	"msys_api_gateway/structs/responses"

	gctrace "github.com/msys/go-common/trace"
)

// log returns the context-aware, structured logger for the gateway service.
func log(ctx context.Context) *slog.Logger { return gctrace.Logger(ctx, "gateway") }

func ExistingNumber(ctx context.Context, clientid string, number string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/existing-number/"+number,
		api.GET)
	client := api.Client{
		Request: request,
		Type_:   "get",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func NameInquiryRequest(ctx context.Context, clientid string, number string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/name-inquiry/"+number,
		api.GET)
	client := api.Client{
		Request: request,
		Type_:   "get",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func ValidatePin(ctx context.Context, clientid string, number string, password string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/verify-pin",
		api.POST)
	request.Params["number"] = number
	request.Params["password"] = password
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func ResetPin(ctx context.Context, clientid string, number string, password string, newpassword string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/reset-pin",
		api.POST)
	request.Params["number"] = number
	request.Params["oldPassword"] = password
	request.Params["newPassword"] = newpassword
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func ListAccounts(ctx context.Context, clientid string, id string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/list-accounts",
		api.POST)
	request.Params["id"] = id
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func ListCustAccounts(ctx context.Context, clientid string, number string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/list-cust-accounts",
		api.POST)
	request.Params["number"] = number
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func ListCustAccountsV2(ctx context.Context, clientid string, number string) (resp responses.AccountsApiDataResponse) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/list-cust-accounts",
		api.POST)
	request.Params["number"] = number
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return resp
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return resp
	}

	json.Unmarshal(read, &resp)
	return resp
}

func CreditAccount(ctx context.Context, clientid string, accountNumber string, amount string, reference string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/credit-account",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	request.Params["amount"] = amount
	request.Params["reference"] = reference

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func CreditAccountV2(ctx context.Context, clientid string, accountNumber string, amount string, reference string, channel string) (resp responses.CreditAccountApiDataResponse) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/credit-account",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	request.Params["amount"] = amount
	request.Params["reference"] = reference
	request.Params["channel"] = channel

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return resp
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return resp
	}

	json.Unmarshal(read, &resp)
	return resp
}

func DebitAccountV2(ctx context.Context, clientid string, accountNumber string, amount string, reference string, channel string) (resp responses.CreditAccountApiDataResponse) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/debit-account",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	request.Params["amount"] = amount
	request.Params["reference"] = reference
	request.Params["channel"] = channel

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return resp
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return resp
	}

	json.Unmarshal(read, &resp)
	return resp
}

func CreditSharesAccountV2(ctx context.Context, clientid string, accountNumber string, amount string, reference string, channel string) (resp responses.CreditAccountApiDataResponse) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/credit-shares-account",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	request.Params["amount"] = amount
	request.Params["reference"] = reference
	request.Params["channel"] = channel

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return resp
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return resp
	}

	json.Unmarshal(read, &resp)
	return resp
}

func AccountBalance(ctx context.Context, clientid string, accountNumber string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/account-balance",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func AccountBalanceV2(ctx context.Context, clientid string, accountNumber string) (resp responses.AccountBalanceApiResponse2) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/account-balance",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return resp
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return resp
	}

	json.Unmarshal(read, &resp)
	return resp
}

func AccountStatementV2(ctx context.Context, clientid string, accountNumber string, fromDate string, toDate string) (resp responses.AccountStatementApiResponse) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/account-statement/"+accountNumber,
		api.GET)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return resp
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return resp
	}

	json.Unmarshal(read, &resp)
	return resp
}

func GetContactInfo(ctx context.Context, clientid string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/get-contact-info",
		api.POST)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func FieldDeposit(ctx context.Context, clientid string, accountNumber string, amount string, mobileNumber string, paymentMethod string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream",
		slog.String("client_id", clientid),
		slog.String("upstream", host),
		slog.String("payment_method", paymentMethod))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/field-deposit",
		api.POST)
	request.Params["accountNumber"] = accountNumber
	request.Params["amount"] = amount
	request.Params["mobileNumber"] = mobileNumber
	request.Params["paymentMethod"] = paymentMethod
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func RegisterCustomer(ctx context.Context, clientid string, firstName string, lastName string, gender string, mobileNumber string, source string, chargeAmount string, transactionRef string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/register-customer",
		api.POST)
	request.Params["firstName"] = firstName
	request.Params["lastName"] = lastName
	request.Params["gender"] = gender
	request.Params["mobileNumber"] = mobileNumber
	request.Params["source"] = source
	request.Params["chargeAmount"] = chargeAmount
	request.Params["txnRef"] = transactionRef
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func CreateFieldAccount(ctx context.Context, clientid string, firstName string, lastName string, gender string, mobileNumber string, agentMobileNumber string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/create-field-account",
		api.POST)
	request.Params["firstName"] = firstName
	request.Params["lastName"] = lastName
	request.Params["gender"] = gender
	request.Params["mobileNumber"] = mobileNumber
	request.Params["agentMobileNumber"] = agentMobileNumber
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func VerifyCustomer(ctx context.Context, clientid string, username string, firstName string, lastName string, gender string, mobileNumber string, email string, dob string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

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
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func FetchVerifiedCustomers(ctx context.Context, clientid string) (resp responses.ApprovedCustomersApiResponse) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/fetch-approved-accounts",
		api.GET)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return resp
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return resp
	}

	json.Unmarshal(read, &resp)
	return resp
}

func ActivateApprovedCustomer(ctx context.Context, clientid string, username string, mobileNumber string) (resp interface{}) {
	host := HostMapping(clientid)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientid), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientid+"/activate-verified-customer",
		api.POST)
	request.Params["mobileNumber"] = mobileNumber
	request.Params["username"] = username
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	data := map[string]interface{}{}
	json.Unmarshal(read, &data)
	return data
}

func ListAccountLoans(ctx context.Context, clientId string, accountNumber string) (resp responses.ListLoansApiDataResponse) {
	host := HostMapping(clientId)
	log(ctx).Info("forwarding request to upstream", slog.String("client_id", clientId), slog.String("upstream", host))

	request := api.NewRequest(
		host,
		"/api/"+clientId+"/list-account-loans/"+accountNumber,
		api.GET)
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return resp
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return resp
	}

	json.Unmarshal(read, &resp)
	return resp
}

func LoanRepayment(ctx context.Context, req requests.LoanRepaymentApiRequest) (resp responses.RepayLoanApiDataResponse) {
	host := HostMapping(req.ClientId)
	log(ctx).Info("forwarding loan repayment to upstream",
		slog.String("client_id", req.ClientId),
		slog.String("upstream", host),
		slog.String("loan_id", req.LoanId),
		slog.String("account_number", req.AccountNumber))

	request := api.NewRequest(
		host,
		"/api/"+req.ClientId+"/loan-repayment",
		api.POST)
	request.Params["mobileNumber"] = req.MobileNumber
	request.Params["loanId"] = req.LoanId
	request.Params["amount"] = req.Amount
	request.Params["accountNumber"] = req.AccountNumber
	request.Params["paymentMode"] = req.PaymentMode

	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest(ctx)
	if err != nil {
		log(ctx).Error("upstream request failed", slog.String("upstream", host), slog.String("error", err.Error()))
		return resp
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		return resp
	}

	json.Unmarshal(read, &resp)
	return resp
}
