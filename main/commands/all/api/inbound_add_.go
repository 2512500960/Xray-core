package api

import (
	"fmt"
	"strings"

	handlerService "github.com/xtls/xray-core/app/proxyman/command"
	"github.com/xtls/xray-core/infra/conf"
	"github.com/xtls/xray-core/infra/conf/serial"
	"github.com/xtls/xray-core/main/commands/base"
)

var cmdAddInbound = &base.Command{
	CustomFlags: true,
	UsageLine:   "{{.Exec}} api adis [--server=127.0.0.1:8080] -rule '[jsonconfig]'",
	Short:       "Add inbound from string",
	Long: `
Add inbound to Xray.
Arguments:
	-s, -server 
		The API server address. Default 127.0.0.1:8080
	-t, -timeout
		Timeout seconds to call API. Default 3
	-r, -rule
		json string of inbound 
Example:
    {{.Exec}} {{.LongName}} --server=127.0.0.1:8080 -rule '{"inbounds":[{"listen":"0.0.0.0","tag":"25125br_ptlin","port":10080,"protocol":"dokodemo-door","settings":{"address":"127.0.0.1","network":"tcp"}}]}'
`,
	Run: executeAddInbound,
}

func executeAddInbound(cmd *base.Command, args []string) {
	setSharedFlags(cmd)
	config_json := cmd.Flag.String("rule", "", "")

	cmd.Flag.Parse(args)

	ins := make([]conf.InboundDetourConfig, 0)
	if len(*config_json) > 0 {
		conf, err := serial.DecodeJSONConfig(strings.NewReader(*config_json))
		if err != nil {
			base.Fatalf("failed to decode %s: %s", *config_json, err)
		}
		ins = append(ins, conf.InboundConfigs...)
	}
	if len(ins) == 0 {
		base.Fatalf("no valid inbound found")
	}

	conn, ctx, close := dialAPIServer()
	defer close()

	client := handlerService.NewHandlerServiceClient(conn)
	for _, in := range ins {
		fmt.Println("adding:", in.Tag)
		i, err := in.Build()
		if err != nil {
			base.Fatalf("failed to build conf: %s", err)
		}
		r := &handlerService.AddInboundRequest{
			Inbound: i,
		}
		resp, err := client.AddInbound(ctx, r)
		if err != nil {
			base.Fatalf("failed to add inbound: %s", err)
		}
		showJSONResponse(resp)
	}
}
