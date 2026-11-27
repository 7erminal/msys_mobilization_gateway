package api

import (
	"context"
	"net/http"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/beego/beego/v2/core/logs"

	gctrace "github.com/msys/go-common/trace"
)

type Client struct {
	Request *Request
	Type_   string
}

// SendRequest performs the upstream HTTP call. It:
//   - forwards the three validated trace IDs (from ctx) on the upstream request,
//   - records upstream service/status/duration into the per-request State (for the
//     access log and outcome classification),
//   - classifies transport/5xx errors.
// The trace IDs and State are carried in ctx (see github.com/msys/go-common/trace).
func (c *Client) SendRequest(ctx context.Context) (*http.Response, error) {
	logs.Debug("Path received is " + c.Request.BaseURL + c.Request.Path)
	logs.Debug("Method is ", c.Request.Method.String())
	// TODO: BaseURLとPathのvalidation

	// Forward the validated client-supplied IDs to the upstream, unchanged.
	ids := gctrace.FromContext(ctx)
	gctrace.ApplyToRequest(c.Request.HeaderField, ids)

	beegoRequest := httplib.NewBeegoRequest(
		c.Request.BaseURL+c.Request.Path,
		c.Request.Method.String())
	for key, value := range c.Request.HeaderField {
		beegoRequest.Header(key, value)
	}
	if c.Type_ == "params" {
		for key, value := range c.Request.Params {
			beegoRequest.Param(key, value)
		}
	} else if c.Type_ == "body" {
		beegoRequest.JSONBody(c.Request.Params)
	} else {
	}

	start := time.Now()
	res, err := beegoRequest.Response()
	dur := time.Since(start)

	if st := gctrace.FromState(ctx); st != nil {
		st.UpstreamService = c.Request.BaseURL
		st.UpstreamDuration = dur
		if err != nil {
			st.Outcome, st.ErrorType = gctrace.ClassifyError(err)
		} else if res != nil {
			st.UpstreamStatus = res.StatusCode
			if res.StatusCode >= 500 {
				st.Outcome = gctrace.OutcomeUpstream5xx
				st.ErrorType = "http_5xx"
			}
		}
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}
