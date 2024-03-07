package plugin

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"google.golang.org/protobuf/types/descriptorpb"
)

func init() {
	generator.RegisterPlugin(new(netRpcPlugin))
}

type netRpcPlugin struct {
	*generator.Generator
}

func (n *netRpcPlugin) Name() string {
	return "netRpc"
}

func (n *netRpcPlugin) Init(g *generator.Generator) {
	n.Generator = g
}

func (n *netRpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		n.genServiceCode(svc)
	}
}

func (n *netRpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		n.genImportCode(file)
	}
}

func (n *netRpcPlugin) genServiceCode(svc *descriptorpb.ServiceDescriptorProto) {
	n.P("// TODO: service code , Name = " + svc.GetName())
}

func (n *netRpcPlugin) genImportCode(file *generator.FileDescriptor) {
	n.P("// TODO: import code")
}

var _ generator.Plugin = (*netRpcPlugin)(nil)
