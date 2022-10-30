- Local dev is difficult; could use a tutorial on using .terraformrc...the whole local registsry namespace fileystem thing appears to be missing something I haven't been able to put my finger on yet...maybe a very explicit "for dummies" like thing would help

- I've had some difficulty figuring out the configure function, we only have the hashicups example and that's unfortunately not a close match for what I'm doing since I'm using a gRPC API

'''
| The plugin encountered an error, and failed to respond to the plugin.(*GRPCProvider).ApplyResourceChange call. The plugin logs may contain more details.
╵

Stack trace from the terraform-provider-gpio plugin:

panic: interface conversion: interface {} is nil, not string

goroutine 40 [running]:
github.com/andybaran/terraform-provider-gpio/internal/provider.resourcePWMCreate({0x103165e20?, 0x1400013a210?}, 0x0?, {0x10303fc20?, 0x140000889c0})
	/Users/andybaran/code/mentorship/terraform-provider-gpio/internal/provider/resource_gpio_pwm.go:49 +0x234
github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema.(*Resource).create(0x140001ca9a0, {0x103165e20, 0x1400013a210}, 0xd?, {0x10303fc20, 0x140000889c0})
	/Users/andybaran/go/pkg/mod/github.com/hashicorp/terraform-plugin-sdk/v2@v2.24.0/helper/schema/resource.go:707 +0xe8
github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema.(*Resource).Apply(0x140001ca9a0, {0x103165e20, 0x1400013a210}, 0x14000142410, 0x14000148180, {0x10303fc20, 0x140000889c0})
	/Users/andybaran/go/pkg/mod/github.com/hashicorp/terraform-plugin-sdk/v2@v2.24.0/helper/schema/resource.go:837 +0x86c
github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema.(*GRPCProviderServer).ApplyResourceChange(0x14000417e30, {0x103165e20?, 0x1400013a0c0?}, 0x14000138000)
	/Users/andybaran/go/pkg/mod/github.com/hashicorp/terraform-plugin-sdk/v2@v2.24.0/helper/schema/grpc_provider.go:1021 +0xb70
github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server.(*server).ApplyResourceChange(0x140002d0960, {0x103165e20?, 0x1400053a300?}, 0x140002200e0)
	/Users/andybaran/go/pkg/mod/github.com/hashicorp/terraform-plugin-go@v0.14.0/tfprotov5/tf5server/server.go:818 +0x3b8
github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5._Provider_ApplyResourceChange_Handler({0x103134660?, 0x140002d0960}, {0x103165e20, 0x1400053a300}, 0x14000220070, 0x0)
	/Users/andybaran/go/pkg/mod/github.com/hashicorp/terraform-plugin-go@v0.14.0/tfprotov5/internal/tfplugin5/tfplugin5_grpc.pb.go:385 +0x170
google.golang.org/grpc.(*Server).processUnaryRPC(0x140002cc1e0, {0x103169800, 0x14000003860}, 0x14000544120, 0x14000432030, 0x1035c2fe0, 0x0)
	/Users/andybaran/go/pkg/mod/google.golang.org/grpc@v1.50.1/server.go:1340 +0xb7c
google.golang.org/grpc.(*Server).handleStream(0x140002cc1e0, {0x103169800, 0x14000003860}, 0x14000544120, 0x0)
	/Users/andybaran/go/pkg/mod/google.golang.org/grpc@v1.50.1/server.go:1713 +0x82c
google.golang.org/grpc.(*Server).serveStreams.func1.2()
	/Users/andybaran/go/pkg/mod/google.golang.org/grpc@v1.50.1/server.go:965 +0x84
created by google.golang.org/grpc.(*Server).serveStreams.func1
	/Users/andybaran/go/pkg/mod/google.golang.org/grpc@v1.50.1/server.go:963 +0x290

Error: The terraform-provider-gpio plugin crashed!

This is always indicative of a bug within the plugin. It would be immensely
helpful if you could report the crash with the plugin's maintainers so that it
can be fixed. The output above should help diagnose the issue.
'''

- As a new go developer documentation like this is a little obtuse https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2@v2.24.0/diag?utm_source=gopls#Diagnostics ... are there provider devs out there like me who aren't neccesarily coming from a strong dev background but want to create or extend a provider to make their job easier with TF?

- Do we really only have one diag type?