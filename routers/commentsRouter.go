package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Client_servicesController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Client_servicesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Client_servicesController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Client_servicesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Client_servicesController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Client_servicesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Client_servicesController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Client_servicesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Client_servicesController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Client_servicesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:ClientsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:ClientsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:ClientsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:ClientsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:ClientsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:ClientsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:ClientsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:ClientsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:ClientsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:ClientsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "AccountBalance",
            Router: `/account-balance`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "ActivateVerifiedCustomers",
            Router: `/activate-verified-customers`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "CreateFieldAccount",
            Router: `/create-field-account`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "CreditAccount",
            Router: `/credit-account`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "ExistingNumber",
            Router: `/existing-number/:number`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "FetchApprovedCustomers",
            Router: `/fetch-approved-customers`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "FieldDeposit",
            Router: `/field-deposit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "GetAccountInfo",
            Router: `/get-contact-info`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "ListAccounts",
            Router: `/list-accounts`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "ListCustAccounts",
            Router: `/list-cust-accounts`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "NameInquiry",
            Router: `/name-inquiry/:number`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "RegisterCustomer",
            Router: `/register-customer`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "ResetPin",
            Router: `/reset-pin`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "VerifyCustomer",
            Router: `/verify-customer`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:Service_requestsController"],
        beego.ControllerComments{
            Method: "VerifyPin",
            Router: `/verify-pin`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:ServicesController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:ServicesController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:ServicesController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:ServicesController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["msys_api_gateway/controllers:ServicesController"] = append(beego.GlobalControllerRouter["msys_api_gateway/controllers:ServicesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
