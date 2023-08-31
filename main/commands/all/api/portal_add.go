package api

import (
	"github.com/xtls/xray-core/app/reverse"
	reverseService "github.com/xtls/xray-core/app/reverse/command"
	"github.com/xtls/xray-core/main/commands/base"
)

var cmdAddReversePortal = &base.Command{
	CustomFlags: true,
	UsageLine:   "{{.Exec}} api addportal [--server=127.0.0.1:8080] -tag tag -domain domainname",
	Short:       "Add Reverse Portal",
	Long: `
Add Reverse Portal to Xray.
Arguments:
	-s, -server 
		The API server address. Default 127.0.0.1:8080
	-t, -timeout
		Timeout seconds to call API. Default 3
	-tag 
		tag for the portal
	-domain
		domain for the portal
Example:
	{{.Exec}} {{.LongName}} --server=127.0.0.1:8080 -tag portal_tag1 -domain reversetrest.xray.com
`,
	Run: executeAddReversePortal,
}

func executeAddReversePortal(cmd *base.Command, args []string) {
	setSharedFlags(cmd)
	cmd.Flag.Parse(args)

	conn, ctx, close := dialAPIServer()
	defer close()

	tag := cmd.Flag.String("tag", "", "")
	domain := cmd.Flag.String("domain", "", "")
	client := reverseService.NewReverseServiceClient(conn)
	r := &reverseService.AddPortalRequest{Config: &reverse.PortalConfig{Tag: *tag, Domain: *domain}}
	resp, err := client.AddPortal(ctx, r)
	if err != nil {
		base.Fatalf("failed to add portal: %s", err)
	}
	showJSONResponse(resp)
}
