// Copyright © 2020 Hedzr Yeh.

package gen

import (
	"bytes"
	"fmt"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/log/dir"
	"go/ast"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

func pr(msg string, args ...interface{}) {
	fmt.Printf("%v\n", fmt.Sprintf(msg, args...))
}

func prflag(name string) {
	fmt.Printf("%v => %q\n", name, cmdr.GetStringR("application."+name))
}

func genApp(cmd *cmdr.Command, args []string) (err error) {
	prflag("module-name")
	prflag("name")
	prflag("package")
	prflag("processing-filename")
	prflag("processing-file-lineno")

	goFile := cmdr.GetStringR("application.processing-filename")
	goFilePath := path.Join(dir.GetCurrentDir(), goFile)
	goPackage := cmdr.GetStringR("application.package")
	goModName := goPackage
	if goModName == "main" {
		goModName = trimGopath(path.Dir(goFilePath))
	}

	var pkgInfo *build.Package

	pr("goFilePath: %q", goFilePath)
	pr("goPackage: %q", goPackage)
	pr("goModName: %q", goModName)
	pr("generating and target to %q ...", goModName)
	pkgInfo, err = build.ImportDir(path.Dir(goFilePath), 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("pkgInfo: %v", pkgInfo)

	parseAst(goModName, goPackage, pkgInfo)
	return
}

func parseAst(modName, pkgName string, pkgInfo *build.Package) {
	fset := token.NewFileSet()
	for _, file := range pkgInfo.GoFiles {
		f, err := parser.ParseFile(fset, file, nil, 0)
		if err != nil {
			log.Fatal(err)
		}
		typ := ""
		ast.Inspect(f, func(n ast.Node) bool {
			decl, ok := n.(*ast.GenDecl)
			// 只需要const
			if !ok || decl.Tok != token.CONST {
				return true
			}
			for _, spec := range decl.Specs {
				vspec := spec.(*ast.ValueSpec)
				if vspec.Type == nil && len(vspec.Values) > 0 {
					typ = ""
					continue
				}

				if vspec.Type != nil {
					ident, ok := vspec.Type.(*ast.Ident)
					if !ok {
						continue
					}
					typ = ident.Name
				}
				pr("typ = %q | vspec = %v / %v", typ, vspec.Names, vspec.Values)

				////typ是否是需处理的类型
				//consts, ok := typesMap[typ]
				//if !ok {
				//	continue
				//}
				////将所有const变量名保存
				//for _, n := range vspec.Names {
				//	consts = append(consts, n.Name)
				//}
				//typesMap[typ] = consts
			}
			return true
		})
	}
	return
}

func genString(modName, pkgName string, pkgInfo *build.Package, types map[string][]string) []byte {
	const strTmp = `
	package {{.pkg}}
	import "fmt"
	
	{{range $typ,$consts :=.types}}
	func (c {{$typ}}) String() string{
		switch c { {{range $consts}}
			case {{.}}:return "{{.}}"{{end}}
		}
		return fmt.Sprintf("Status(%d)", c)	
	}
	{{end}}
	`

	//pkgName := os.Getenv("GOPACKAGE")
	if pkgName == "" {
		pkgName = pkgInfo.Name
	}
	data := map[string]interface{}{
		"pkg":   pkgName,
		"types": types,
	}

	t, err := template.New("").Parse(strTmp)
	if err != nil {
		log.Fatal(err)
	}
	buff := bytes.NewBufferString("")
	err = t.Execute(buff, data)
	if err != nil {
		log.Fatal(err)
	}
	//格式化
	src, err := format.Source(buff.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	return src
}

func trimGopath(s string) string {
	gopath := os.Getenv("GOPATH")
	// pr("gopath = %q", gopath)
	for _, p := range strings.Split(gopath, ":") {
		if strings.HasPrefix(s, p) {
			v := s[len(p)+1:]
			if strings.HasPrefix(v, "src/") {
				v = v[4:]
			}
			return v
		}
	}

	if strings.Contains(s, "src/github.com") {
		idx := strings.Index(s, "src/github.com")
		return s[idx+4:]
	}

	return s
}
