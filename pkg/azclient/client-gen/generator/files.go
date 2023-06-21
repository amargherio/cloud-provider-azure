// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

package generator

import (
	"bytes"
	"fmt"
	"io"

	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
)

func WriteToFile(ctx *genall.GenerationContext, root *loader.Package, fileName string, headerText string, importList map[string]map[string]struct{}, codeSnips *bytes.Buffer) error {

	outputFile, err := ctx.Open(root, fileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	return DumpToWriter(outputFile, headerText, importList, root.Name, codeSnips)
}

func DumpToWriter(writer io.Writer, headerText string, importList map[string]map[string]struct{}, pkgName string, codeSnips *bytes.Buffer) error {
	outContent := new(bytes.Buffer)

	outContent.WriteString(headerText)
	outContent.WriteRune('\n')
	outContent.WriteString("// Code generated by client-gen. DO NOT EDIT.\n")
	outContent.WriteString(fmt.Sprintf("package %s\n", pkgName))

	if len(importList) > 0 {
		if err := ImportTemplate.Execute(outContent, importList); err != nil {
			return err
		}
	}

	if _, err := io.Copy(outContent, codeSnips); err != nil {
		return err
	}
	rawBytes := outContent.Bytes()
	counter, err := writer.Write(rawBytes)
	if err != nil {
		return err
	}
	if counter < len(rawBytes) {
		return io.ErrShortWrite
	}
	return nil
}