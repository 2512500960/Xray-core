package api

import (
	reverseService "github.com/xtls/xray-core/app/reverse/command"
	"github.com/xtls/xray-core/main/commands/base"
)

var cmdDelReversePortal = &base.Command{
	CustomFlags: true,
	UsageLine:   "{{.Exec}} api delportal [--server=127.0.0.1:8080] -tag tag",
	Short:       "Add Reverse Portal",
	Long: `
Del Reverse Portal from Xray.
Arguments:
	-s, -server 
		The API server address. Default 127.0.0.1:8080
	-t, -timeout
		Timeout seconds to call API. Default 3
	-tag 
		tag for the portal
Example:
	{{.Exec}} {{.LongName}} --server=127.0.0.1:8080 -tag portal_tag1 
`,
	Run: executeDelReversePortal,
}

func executeDelReversePortal(cmd *base.Command, args []string) {
	setSharedFlags(cmd)
	tag := cmd.Flag.String("tag", "", "")
	cmd.Flag.Parse(args)

	conn, ctx, close := dialAPIServer()
	defer close()

	client := reverseService.NewReverseServiceClient(conn)
	r := &reverseService.RemovePortalRequest{Tag: *tag}
	resp, err := client.RemovePortal(ctx, r)
	if err != nil {
		base.Fatalf("failed to del portal: %s", err)
	}
	showJSONResponse(resp)
}
