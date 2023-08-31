package api

import (
	reverseService "github.com/xtls/xray-core/app/reverse/command"
	"github.com/xtls/xray-core/main/commands/base"
)

var cmdGetReversePortals = &base.Command{
	CustomFlags: true,
	UsageLine:   "{{.Exec}} api portals [--server=127.0.0.1:8080]",
	Short:       "Get Reverse Portals",
	Long: `
Get Reverse Portals from Xray.
Arguments:
	-s, -server 
		The API server address. Default 127.0.0.1:8080
	-t, -timeout
		Timeout seconds to call API. Default 3
Example:
	{{.Exec}} {{.LongName}} --server=127.0.0.1:8080
`,
	Run: executeGetReversePortals,
}

func executeGetReversePortals(cmd *base.Command, args []string) {
	setSharedFlags(cmd)
	cmd.Flag.Parse(args)

	conn, ctx, close := dialAPIServer()
	defer close()

	client := reverseService.NewReverseServiceClient(conn)
	r := &reverseService.GetPortalsRequest{}
	resp, err := client.GetPortals(ctx, r)
	if err != nil {
		base.Fatalf("failed to get portals: %s", err)
	}
	showJSONResponse(resp)
}
