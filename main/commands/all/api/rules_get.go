package api

import (
	routerService "github.com/xtls/xray-core/app/router/command"
	"github.com/xtls/xray-core/main/commands/base"
)

var cmdGetRouterRules = &base.Command{
	CustomFlags: true,
	UsageLine:   "{{.Exec}} api routerules [--server=127.0.0.1:8080]",
	Short:       "Get Router rules",
	Long: `
Get Router rule from Xray.
Arguments:
	-s, -server 
		The API server address. Default 127.0.0.1:8080
	-t, -timeout
		Timeout seconds to call API. Default 3
Example:
	{{.Exec}} {{.LongName}} --server=127.0.0.1:8080
`,
	Run: executeGetRouterRules,
}

func executeGetRouterRules(cmd *base.Command, args []string) {
	setSharedFlags(cmd)
	cmd.Flag.Parse(args)

	conn, ctx, close := dialAPIServer()
	defer close()

	client := routerService.NewRoutingServiceClient(conn)
	r := &routerService.GetRoutingRulesRequest{}
	resp, err := client.GetRules(ctx, r)
	if err != nil {
		base.Fatalf("failed to get rule: %s", err)
	}
	showJSONResponse(resp)
}
