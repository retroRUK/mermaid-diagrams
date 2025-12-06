package java

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/retroRUK/mermaid-diagrams/src/models"
	"github.com/retroRUK/mermaid-diagrams/src/utilities"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/java"
)

func CreateClassDiagrams(rootPath string) error {
	var declarations []models.Declaration

	if err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasSuffix(d.Name(), ".java") {
			return nil
		}

		sourceCode, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Create a new parser
		parser := sitter.NewParser()
		parser.SetLanguage(java.GetLanguage())

		// Parse the source code
		tree, err := parser.ParseCtx(context.Background(), nil, sourceCode)
		if err != nil {
			log.Fatal(err)
		}
		defer tree.Close()

		// Get the root node
		root := tree.RootNode()

		var dec models.Declaration
		dec = findClassDeclarations(root, sourceCode)

		if dec.Name == "" {
			dec = findInterfaceDeclarations(root, sourceCode)
		}

		if dec.Name == "" {
			dec = findEnumDeclarations(root, sourceCode)
		}

		if dec.Type == "class" {
			dec.Implements = findImplementedInterfaces(root, sourceCode)
			dec.Properties = findProperties(root, sourceCode)
		}

		dec.Methods = findMethods(root, sourceCode)
		dec.Path = filepath.Dir(path)

		declarations = append(declarations, dec)

		return nil
	}); err != nil {
		return fmt.Errorf("Error walking directory: %v\n", err)
	}

	for _, d := range declarations {
		decs := getRelatedDeclarations(d, declarations)
		writeClassDiagram(decs, d.Path+"/"+d.Name+".md")
	}

	writeClassDiagram(declarations, rootPath+"/overall.md")

	return nil
}

func getRelatedDeclarations(d models.Declaration, allDeclarations []models.Declaration) []models.Declaration {
	related := []models.Declaration{d}
	seen := map[string]bool{d.Name: true}

	// Add extended class
	if d.Extends != "" && !seen[d.Extends] {
		for _, dd := range allDeclarations {
			if dd.Name == d.Extends {
				seen[dd.Name] = true
				related = append(related, dd)
				// Recursively add parent's dependencies
				parentRelated := getRelatedDeclarations(dd, allDeclarations)
				for _, pr := range parentRelated {
					if !seen[pr.Name] {
						seen[pr.Name] = true
						related = append(related, pr)
					}
				}
				break
			}
		}
	}

	// Add implemented interfaces
	for _, impl := range d.Implements {
		if !seen[impl] {
			for _, dd := range allDeclarations {
				if dd.Name == impl {
					seen[dd.Name] = true
					related = append(related, dd)
					break
				}
			}
		}
	}

	return related
}

func writeClassDiagram(declarations []models.Declaration, outputFile string) {
	readme := "```mermaid\nclassDiagram\n"
	for _, d := range declarations {
		if d.Extends != "" {
			readme += fmt.Sprintf("\t%s --|> %s\n", d.Name, d.Extends)
		}

		for _, i := range d.Implements {
			readme += fmt.Sprintf("\t%s ..|> %s\n", d.Name, i)
		}

		readme += fmt.Sprintf("\tclass %s {\n\t\t<<%s>>\n", d.Name, d.Type)

		for _, p := range d.Properties {
			readme += fmt.Sprintf("\t\t%s %s %s\n", getModifierSymbol(p.Modifier), p.Type, p.Name)
		}

		for _, m := range d.Methods {
			readme += fmt.Sprintf("\t\t%s %s %s(%s)\n", getModifierSymbol(m.Modifier), m.ReturnType, m.Name, strings.Join(m.Parameters, ", "))
		}

		readme += "\t}\n\n"
	}
	readme += "```"

	if err := os.WriteFile(outputFile, []byte(readme), 0644); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
}

func findClassDeclarations(node *sitter.Node, source []byte) models.Declaration {
	var declaration models.Declaration
	declaration.Type = "class"

	var walk func(*sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "class_declaration" {
			// Loop through ALL children to find both identifier and superclass
			for i := range int(n.ChildCount()) {
				child := n.Child(i)

				// Get class name (first identifier found)
				if child.Type() == "identifier" && declaration.Name == "" {
					declaration.Name = child.Content(source)
				}

				// Get superclass (extends)
				if child.Type() == "superclass" {
					for j := range int(child.ChildCount()) {
						superChild := child.Child(j)
						if superChild.Type() == "type_identifier" {
							declaration.Extends = superChild.Content(source)
							break
						}
					}
				}
			}
			// Don't continue walking once we've found a class_declaration
			return
		}

		for i := range int(n.ChildCount()) {
			walk(n.Child(i))
		}
	}

	walk(node)
	return declaration
}

func findInterfaceDeclarations(node *sitter.Node, source []byte) models.Declaration {
	var declaration models.Declaration
	declaration.Type = "interface"

	var walk func(*sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "interface_declaration" {
			// Loop through ALL children to find identifier and extends
			for i := range int(n.ChildCount()) {
				child := n.Child(i)

				// Get interface name
				if child.Type() == "identifier" && declaration.Name == "" {
					declaration.Name = child.Content(source)
				}

				// Get extends (interfaces can extend other interfaces)
				if child.Type() == "extends_interfaces" {
					for j := range int(child.ChildCount()) {
						extendsChild := child.Child(j)
						if extendsChild.Type() == "type_list" {
							// Interface can extend multiple interfaces
							for k := range int(extendsChild.ChildCount()) {
								typeChild := extendsChild.Child(k)
								if typeChild.Type() == "type_identifier" {
									declaration.Implements = append(declaration.Implements, typeChild.Content(source))
								}
							}
						}
					}
				}
			}
			return
		}

		for i := range int(n.ChildCount()) {
			walk(n.Child(i))
		}
	}

	walk(node)
	return declaration
}

func findEnumDeclarations(node *sitter.Node, source []byte) models.Declaration {
	var declaration models.Declaration

	var walk func(*sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "enum_declaration" {
			declaration.Type = "enum"

			for i := range int(n.ChildCount()) {
				child := n.Child(i)

				// Get enum name
				if child.Type() == "identifier" && declaration.Name == "" {
					declaration.Name = child.Content(source)
				}

				// Get enum constants
				if child.Type() == "enum_body" {
					for j := range int(child.ChildCount()) {
						bodyChild := child.Child(j)
						if bodyChild.Type() == "enum_constant" {
							// Get the constant name

							for k := range int(bodyChild.ChildCount()) {
								constChild := bodyChild.Child(k)
								if constChild.Type() == "identifier" {
									// Add enum constants as properties
									declaration.Properties = append(declaration.Properties, models.Property{
										Modifier: "",
										Type:     "",
										Name:     strings.TrimSpace(constChild.Content(source)),
									})
									break
								}
							}
						}
					}
				}

				// Enums can implement interfaces
				if child.Type() == "super_interfaces" {
					for j := range int(child.ChildCount()) {
						interfaceChild := child.Child(j)
						if interfaceChild.Type() == "type_list" {
							for k := range int(interfaceChild.ChildCount()) {
								typeChild := interfaceChild.Child(k)
								if typeChild.Type() == "type_identifier" {
									declaration.Implements = append(declaration.Implements, typeChild.Content(source))
								}
							}
						}
					}
				}
			}
			return
		}

		for i := range int(n.ChildCount()) {
			walk(n.Child(i))
		}
	}

	walk(node)
	log.Println(declaration.Properties)
	return declaration
}

func findImplementedInterfaces(node *sitter.Node, source []byte) []string {
	var interfaces []string

	var walk func(*sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "class_declaration" {
			for i := range int(n.ChildCount()) {
				child := n.Child(i)

				// Look for super_interfaces node
				if child.Type() == "super_interfaces" {
					for j := range int(child.ChildCount()) {
						interfaceChild := child.Child(j)
						if interfaceChild.Type() == "type_list" {
							for k := range int(interfaceChild.ChildCount()) {
								typeChild := interfaceChild.Child(k)
								if typeChild.Type() == "type_identifier" {
									interfaces = append(interfaces, typeChild.Content(source))
								}
							}
						}
					}
				}
			}
			return
		}

		for i := range int(n.ChildCount()) {
			walk(n.Child(i))
		}
	}

	walk(node)
	return interfaces
}

func findMethods(node *sitter.Node, source []byte) []models.Method {
	var methods []models.Method

	var walk func(*sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "method_declaration" {
			var modifier string
			var methodName string
			var returnType string
			var parameters []string

			for i := 0; i < int(n.ChildCount()); i++ {
				child := n.Child(i)

				// get modifier, public, private, etc
				if child.Type() == "modifiers" {
					modifier = child.Content(source)
				}

				// Get return type
				if child.Type() == "void_type" ||
					child.Type() == "type_identifier" ||
					child.Type() == "integral_type" ||
					child.Type() == "floating_point_type" ||
					child.Type() == "boolean_type" ||
					child.Type() == "generic_type" ||
					child.Type() == "array_type" {
					returnType = child.Content(source)
				}

				// Get method name
				if child.Type() == "identifier" {
					methodName = child.Content(source)
				}

				// Get parameters
				if child.Type() == "formal_parameters" {
					parameters = extractParameters(child, source)
				}
			}

			methods = append(methods, models.Method{
				Modifier:   modifier,
				ReturnType: utilities.ConvertMermaidString(returnType),
				Name:       methodName,
				Parameters: parameters,
			})
		}

		for i := range int(n.ChildCount()) {
			walk(n.Child(i))
		}
	}

	walk(node)
	return methods
}

func extractParameters(paramsNode *sitter.Node, source []byte) []string {
	var params []string

	for i := 0; i < int(paramsNode.ChildCount()); i++ {
		child := paramsNode.Child(i)

		if child.Type() == "formal_parameter" {
			var paramType string
			var paramName string

			for j := 0; j < int(child.ChildCount()); j++ {
				paramChild := child.Child(j)

				// Get parameter type
				if paramChild.Type() == "type_identifier" ||
					paramChild.Type() == "integral_type" ||
					paramChild.Type() == "floating_point_type" ||
					paramChild.Type() == "boolean_type" ||
					paramChild.Type() == "generic_type" ||
					paramChild.Type() == "array_type" {
					paramType = paramChild.Content(source)
				}

				// Get parameter name
				if paramChild.Type() == "identifier" {
					paramName = paramChild.Content(source)
				}
			}

			if paramType != "" && paramName != "" {
				params = append(params, fmt.Sprintf("%s %s",
					utilities.ConvertMermaidString(paramType),
					paramName))
			}
		}
	}

	return params
}

func findProperties(node *sitter.Node, source []byte) []models.Property {
	var properties []models.Property

	var walk func(*sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "field_declaration" {
			var fieldModifier string
			var fieldType string
			var fieldName string

			for i := range int(n.ChildCount()) {
				child := n.Child(i)

				if child.Type() == "modifiers" {
					fieldModifier = child.Content(source)
				}

				// Get field type - handle both simple and generic types
				if child.Type() == "type_identifier" ||
					child.Type() == "integral_type" || // int, long, byte, short, char
					child.Type() == "floating_point_type" || // float, double
					child.Type() == "boolean_type" || // boolean
					child.Type() == "generic_type" || // List<String>
					child.Type() == "array_type" { // String[]
					fieldType = child.Content(source)
				}

				// Get field name from variable_declarator
				if child.Type() == "variable_declarator" {
					for j := 0; j < int(child.ChildCount()); j++ {
						varChild := child.Child(j)
						if varChild.Type() == "identifier" {
							fieldName = varChild.Content(source)
							break
						}
					}
				}
			}

			properties = append(properties, models.Property{
				Modifier: fieldModifier,
				Type:     utilities.ConvertMermaidString(fieldType),
				Name:     fieldName,
			})
		}

		for i := range int(n.ChildCount()) {
			walk(n.Child(i))
		}
	}

	walk(node)
	return properties
}

func getModifierSymbol(modifier string) string {
	result := ""

	if strings.Contains(modifier, "public") {
		result += "+"
	}

	if strings.Contains(modifier, "private") {
		result += "-"
	}

	if strings.Contains(modifier, "protected") {
		result += "#"
	}

	if strings.Contains(modifier, "static") {
		result += "$"
	}

	if strings.Contains(modifier, "final") {
		result += "!"
	}

	if strings.Contains(modifier, "volatile") {
		result += "V"
	}

	if strings.Contains(modifier, "transient") {
		result += "T"
	}

	return result
}
