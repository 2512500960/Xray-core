package api

import (
	routerService "github.com/xtls/xray-core/app/router/command"
	"github.com/xtls/xray-core/main/commands/base"
)

var cmdDelRouterRule = &base.Command{
	CustomFlags: true,
	UsageLine:   "{{.Exec}} api delrouterule [--server=127.0.0.1:8080] -tag rule_tag",
	Short:       "Del Router rule",
	Long: `
Del Router rule from Xray.
Arguments:
	-s, -server 
		The API server address. Default 127.0.0.1:8080
	-t, -timeout
		Timeout seconds to call API. Default 3
	-tag
		tag of the rule
Example:
	{{.Exec}} {{.LongName}} --server=127.0.0.1:8080
`,
	Run: executeDelRouterRules,
}

func executeDelRouterRules(cmd *base.Command, args []string) {
	setSharedFlags(cmd)
	tag := cmd.Flag.String("tag", "", "")
	cmd.Flag.Parse(args)

	conn, ctx, close := dialAPIServer()
	defer close()
	client := routerService.NewRoutingServiceClient(conn)
	r := &routerService.RemoveRoutingRuleRequest{Tag: *tag}
	resp, err := client.RemoveRule(ctx, r)
	if err != nil {
		base.Fatalf("failed to del rule: %s", err)
	}
	showJSONResponse(resp)
}
