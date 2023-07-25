package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {

	// g := generator.New()
	// data, err := io.ReadAll(g)
	// if err != nil {
	// 	g.Error(err, "reading input")
	// }

	// if err := proto.Unmarshal(data, g.Request); err != nil {
	// 	g.Error(err, "parsing input proto")
	// }
	// if len(g.Request.FileToGenerate) == 0 {
	// 	g.Fail("no file to generate")
	// }
	// g.CommandLineParameters(g.Request.GetParameter())
	// g.WrapTypes()
	// g.SetPackageNames()
	// g.BuildTypeNameMap()
	// g.GenerateAllFiles()

	// data, err = proto.Marshal(g.Response)
	// if err != nil {
	// 	g.Error(err, "failed to marshal output proto")
	// }
	// _, err = os.Stdout.Write(data)
	// if err != nil {
	// 	g.Error(err, "failed to write output proto")
	// }
	// parse the input
	var request pluginpb.CodeGeneratorRequest
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	err = proto.Unmarshal(data, &request)
	if err != nil {
		panic(err)
	}
	// crate a new codeGeneratorResponse.

	var response pluginpb.CodeGeneratorResponse

	opts := protogen.Options{}

	g, err := opts.New(&request)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range g.Files {
		for _, m := range f.Messages {
			fmt.Printf("Generating code for message %v\n", m.GoIdent)
			filename := fmt.Sprintf("%s.pb.go", m.GoIdent.GoName)
			output := g.NewGeneratedFile(filename, f.GoImportPath)
			output.P("package", f.GoPackageName)
			output.P("type", m.GoIdent.GoName, "struct{\n")
			for _, field := range m.Fields {
				output.P(field.GoName, " ", field.Desc.Kind(), "\n")
			}
			output.P("}\n")
		}
	}

	response.File = append(response.File, &pluginpb.CodeGeneratorResponse_File{
		Name: proto.String(fmt.Sprintf("%s.pb.go",
			g.Request.ProtoReflect().Descriptor().Name())),

		Content:           proto.String(g.Request.FileToGenerate[0]),
		GeneratedCodeInfo: &descriptorpb.GeneratedCodeInfo{},
	})

	data, err = proto.Marshal(&response)
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}
}

// build .
//protoc --proto_path=. --go_out=. --go_opt=paths=source_relative --go_opt=Mproto/tour.proto=github.com/your/repo/tour proto/tour.proto
/*
--proto_path=.：指定 .proto 文件的搜索路径为当前目录。
--go_out=.：指定生成的 Go 代码的输出目录为当前目录。
--go_opt=paths=source_relative：指定使用相对路径来生成 Go 代码文件。
--go_opt=Mproto/tour.proto=github.com/your/repo/tour：指定 proto/tour.proto 文件的 Go 导入路径为 github.com/your/repo/tour，这里假设你的 Git 仓库地址为 github.com/your/repo。
proto/tour.proto：指定要生成 Go 代码的 .proto 文件路径为 proto/tour.proto。
*/
