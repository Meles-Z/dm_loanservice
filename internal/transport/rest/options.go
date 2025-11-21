package rest

import (
	"context"
	"net/http"
	"dm_loanservice/drivers/goconf"

	httptransport "github.com/go-kit/kit/transport/http"

	Logger "dm_loanservice/drivers/logger/zap"
	"dm_loanservice/drivers/utils"
	ctxDM "dm_loanservice/drivers/utils/context"
)

var opts = []httptransport.ServerOption{
	httptransport.ServerBefore(serverBefore),
	httptransport.ServerAfter(serverAfter),
}

func serverBefore(ctx context.Context, req *http.Request) context.Context {
	reqId := req.Header.Get("X-Correlation-ID")
	if len(reqId) == 0 {
		reqId = utils.GenerateThreadId()
	}
	ctxSess := ctxDM.New(Logger.GetLogger()).
		SetXCorrelationID(reqId).
		SetAppName(goconf.Config().GetString("app.name")).
		SetAppVersion(goconf.Config().GetString("app.version")).
		SetPort(goconf.Config().GetInt("rest.port")).
		SetSrcIP(req.Host).
		SetURL(req.URL.Path).
		SetMethod(req.Method).
		SetHeader(req.Header)

	return context.WithValue(ctx, ctxDM.AppSession, ctxSess)
}

func serverAfter(ctx context.Context, w http.ResponseWriter) context.Context {
	w.Header().Set("Content-Type", "application/json")
	return ctx
}

func LoggerContext(req *http.Request) context.Context {
	return serverBefore(req.Context(), req)
}
