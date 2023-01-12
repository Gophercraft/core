package autocode

import (
	"fmt"
	"sort"

	"github.com/Gophercraft/core/format/dbc"
	"github.com/Gophercraft/core/format/dbc/dbd"
	"github.com/Gophercraft/core/format/dbc/dbdefs"
)

type layoutTarget struct {
	Definition *dbd.Definition
	Layout     *dbd.Layout
}

func (g *Generator) findLayoutTargets() error {
	// fmt.Println(len(dbdefs.All))

	for _, def := range dbdefs.All {
		for l := range def.Layouts {
			layout := &def.Layouts[l]
			foundBuild := false
			for _, b := range layout.VerifiedBuilds {
				if b == g.Build {
					foundBuild = true
					break
				}
			}
			//
			if !foundBuild {
				for _, br := range layout.BuildRanges {
					if br.Contains(g.Build) {
						g.layouts = append(g.layouts, &layoutTarget{
							Definition: def,
							Layout:     layout,
						})
						foundBuild = true
						break
					}
				}
				if foundBuild {
					break
				}
			}

			if foundBuild {
				g.layouts = append(g.layouts, &layoutTarget{
					Definition: def,
					Layout:     layout,
				})
				break
			}
		}
	}

	sort.Slice(g.layouts, func(i, j int) bool {
		return g.layouts[i].Definition.Name <= g.layouts[j].Definition.Name
	})

	return nil
}

func (g *Generator) writeLayout(file *Printer, target *layoutTarget) error {
	locSize, err := dbc.LocStringSize(g.Build)
	if err != nil {
		return err
	}

	numColumns := 0
	rowSize := 0
	indexIsID := target.Definition.Column("ID") == nil

	// Calculate number of columns and size of row

	for _, column := range target.Layout.Columns {
		columnDef := target.Definition.Column(column.Name)

		numColumnElements := 1

		if column.ArraySize > 0 {
			numColumnElements = column.ArraySize
		}

		columnSize := 0

		switch columnDef.Type {
		case dbd.LocString:

		case dbd.Int, dbd.Uint:
			columnSize = (column.Bits / 8) * numColumnElements
		case dbd.Float:
			columnSize = (column.Bits / 8) * numColumnElements
		case dbd.String:
			columnSize = numColumnElements * 4
		default:
			panic("unknown type")
		}

		rowSize += columnSize

		if columnDef.Type == dbd.LocString {
			numColumns += locSize
			continue
		}

		numColumns += numColumnElements
	}

	// Begin writing record definition

	file.Printf("struct %sRec {\n", target.Definition.Name)

	file.Printf("\tstatic constexpr uint32_t NumColumns = %d;\n", numColumns)
	file.Printf("\tstatic constexpr uint32_t RowSize = %d;\n", rowSize)
	file.Printf("\tstatic constexpr bool IndexIsID = %t;\n", indexIsID)

	file.Printf("\n")

	for _, column := range target.Layout.Columns {
		var (
			arraySuffix string
			cppType     string
			memberName  string
			columnDef   = target.Definition.Column(column.Name)
		)

		memberName = column.Name

		switch columnDef.Type {
		case dbd.LocString:
			cppType = "const char*"
			arraySuffix = fmt.Sprintf("[%d]", locSize)
		case dbd.String:
			cppType = "const char*"
		case dbd.Uint:
			cppType = fmt.Sprintf("uint%d_t", column.Bits)
		case dbd.Int:
			cppType = fmt.Sprintf("int%d_t", column.Bits)
		case dbd.Float:
			switch column.Bits {
			case 32:
				cppType = "float"
			case 64:
				cppType = "double"
			default:
				panic(column.Bits)
			}
		}

		if column.ArraySize > -1 {
			arraySuffix = fmt.Sprintf("[%d]", column.ArraySize) + arraySuffix
		}

		file.Printf("\t%s m_%s%s;\n", cppType, memberName, arraySuffix)
	}

	file.Printf("\n")
	file.Printf("\tconst char * GetFilename();\n")
	file.Printf("\tbool Read(SFile* f, const char* stringBuffer);\n")
	file.Printf("};\n\n")

	return nil
}

func (g *Generator) writeLayoutFile() error {
	layoutFile, err := g.NewPrinter("src/db/ClientDefs.hpp")
	if err != nil {
		return err
	}

	layoutFile.Printf("#ifndef DB_CLIENT_DEFS_HPP\n")
	layoutFile.Printf("#define DB_CLIENT_DEFS_HPP\n")
	layoutFile.Printf("\n")

	var localimports = []string{
		"client/Console.hpp",
		"util/SFile.hpp",
	}

	for _, localimport := range localimports {
		layoutFile.Printf("#include \"%s\"\n", localimport)
	}

	layoutFile.Printf("\n")

	for _, layout := range g.layouts {
		if err := g.writeLayout(layoutFile, layout); err != nil {
			return err
		}
	}

	layoutFile.Printf("\n")
	layoutFile.Printf("#endif")

	return layoutFile.Close()
}

func (g *Generator) generateLayouts() error {
	if err := g.findLayoutTargets(); err != nil {
		return err
	}

	if err := g.writeLayoutFile(); err != nil {
		return err
	}

	return nil
}
