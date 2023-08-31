package api

import (
	"encoding/json"
	routerService "github.com/xtls/xray-core/app/router/command"
	"github.com/xtls/xray-core/infra/conf"

	"github.com/xtls/xray-core/main/commands/base"
)

var cmdAddRouterRule = &base.Command{
	CustomFlags: true,
	UsageLine:   "{{.Exec}} api addrouterule [--server=127.0.0.1:8080] -rule jsonstr",
	Short:       "Add Router rule",
	Long: `
Add Router rule to Xray.
Arguments:
	-s, -server 
		The API server address. Default 127.0.0.1:8080
	-t, -timeout
		Timeout seconds to call API. Default 3
	-r, -rule
		json string of rule 
Example:
	{{.Exec}} {{.LongName}} --server=127.0.0.1:8080 -rule '{"tag":"rule_no1","type":"field","inboundTag":["tunnel_3389"],"outbound":"portal"}'
`,
	Run: executeAddRouterRule,
}

func executeAddRouterRule(cmd *base.Command, args []string) {
	setSharedFlags(cmd)
	cmd.Flag.Parse(args)
	config_json := cmd.Flag.String("rule", "", "")
	conn, ctx, close := dialAPIServer()
	defer close()
	var rule, _ = conf.ParseRule(json.RawMessage(*config_json))
	client := routerService.NewRoutingServiceClient(conn)
	r := &routerService.AddRoutingRuleRequest{RoutingRule: rule}

	resp, err := client.AddRule(ctx, r)
	if err != nil {
		base.Fatalf("failed to get rule: %s", err)
	}
	showJSONResponse(resp)
}
