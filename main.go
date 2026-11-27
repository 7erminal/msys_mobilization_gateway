package main

import (
	"net"
	"os"
	"strings"

	"msys_api_gateway/gatewaylog"
	"msys_api_gateway/middleware"
	gctrace "github.com/msys/go-common/trace"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "msys_api_gateway/routers"
)

func main() {
	// Single-format structured logging: everything (new gateway logs AND the legacy
	// beego `logs.*` calls) flows through slog as JSON on stdout.
	gctrace.Configure(getenv("GATEWAY_LOG_LEVEL", "info"), getenv("GATEWAY_LOG_FORMAT", "json"))

	// Route legacy beego logs through the slog JSON adapter (no call-site changes).
	logs.Register("slogjson", &gatewaylog.SlogJSONAdapter{})
	logs.SetLogger("slogjson", "")
	logs.SetLevel(logs.LevelDebug) // slog performs final level filtering

	sqlConn, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		logs.Error("%s", err)
	}
	orm.RegisterDataBase("default", "mysql", sqlConn)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	orm.Debug = true

	// Tracing middleware owns panic recovery and access logging. Disable Beego's
	// built-in recover/access-log so we don't double-handle or duplicate logs.
	beego.BConfig.RecoverPanic = false
	beego.BConfig.Log.AccessLogs = false

	cfg := middleware.Config{
		ServiceName:       "gateway",
		TrustedProxies:    parseCIDRs(os.Getenv("TRUSTED_PROXIES")),
		LBRequestIDHeader: os.Getenv("LB_REQUEST_ID_HEADER"),
		UpstreamErrorMode: getenv("UPSTREAM_ERROR_MODE", "legacy"),
	}

	// Wrap the Beego HTTP handler with the tracing middleware.
	beego.BeeApp.Handlers = middleware.Middleware(beego.BeeApp.Handlers, cfg)

	beego.Run()
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

// parseCIDRs parses a comma-separated list of IPs/CIDRs into trusted proxy ranges.
// An empty string yields an empty slice (directly internet-facing behavior).
func parseCIDRs(s string) []net.IPNet {
	var out []net.IPNet
	if s == "" {
		return out
	}
	for _, p := range strings.Split(s, ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		var ipnet *net.IPNet
		if strings.Contains(p, "/") {
			_, n, err := net.ParseCIDR(p)
			if err != nil {
				continue
			}
			ipnet = n
		} else {
			ip := net.ParseIP(p)
			if ip == nil {
				continue
			}
			if ip4 := ip.To4(); ip4 != nil {
				ipnet = &net.IPNet{IP: ip4, Mask: net.CIDRMask(32, 32)}
			} else {
				ipnet = &net.IPNet{IP: ip, Mask: net.CIDRMask(128, 128)}
			}
		}
		out = append(out, *ipnet)
	}
	return out
}
