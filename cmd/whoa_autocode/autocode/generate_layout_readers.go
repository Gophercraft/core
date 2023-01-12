package autocode

import (
	"fmt"

	"github.com/Gophercraft/core/format/dbc"
	"github.com/Gophercraft/core/format/dbc/dbd"
)

func (g *Generator) generateLayoutReader(target *layoutTarget) error {
	var anyStrings bool

	locSize, err := dbc.LocStringSize(g.Build)
	if err != nil {
		return err
	}

	file, err := g.NewPrinter(fmt.Sprintf("src/db/record/%s.cpp", target.Definition.Name))
	if err != nil {
		return err
	}

	file.Printf("\n")
	file.Printf("#include \"db/ClientDefs.hpp\"\n")
	file.Printf("\n")

	file.Printf("bool %sRec::Read(SFile* f, const char* stringBuffer) {\n", target.Definition.Name)

	// Create string index arrays
	for _, columnLayout := range target.Layout.Columns {
		columnDef := target.Definition.Column(columnLayout.Name)
		fieldCount := columnLayout.ArraySize
		if fieldCount == -1 {
			fieldCount = 1
		}
		if columnDef.Type == dbd.String || columnDef.Type == dbd.LocString {
			if columnDef.Type == dbd.LocString {
				fieldCount *= locSize
			}
			file.Printf("\tuint32_t temp%sIndices[%d];\n", columnDef.Name, fieldCount)
		}
	}

	file.Printf("\tif (\n")

	// Read integer values & string indices
	for i, columnLayout := range target.Layout.Columns {
		columnDef := target.Definition.Column(columnLayout.Name)
		elementCount := columnLayout.ArraySize
		if elementCount == -1 {
			elementCount = 1
		}

		if columnDef.Type == dbd.String || columnDef.Type == dbd.LocString {
			anyStrings = true
			if columnDef.Type == dbd.LocString {
				elementCount = locSize * elementCount
				if columnLayout.ArraySize > 0 {
					panic("no 2D array yet")
				}
			}
			// Todo: support 2D array (if such a thing is used in DBC)
			for f := 0; f < elementCount; f++ {
				file.Printf("\t\t")

				file.Printf("SFile::Read(f, &temp%sIndices[%d], sizeof(uint32_t), nullptr, nullptr, nullptr) == 0", columnLayout.Name, f)

				if !(f+1 == elementCount && i+1 == len(target.Layout.Columns)) {
					file.Printf(" ||")
				}
				file.Printf("\n")
			}
		} else {
			isArray := columnLayout.ArraySize > -1

			for f := 0; f < elementCount; f++ {
				file.Printf("\t\t")

				if isArray {
					file.Printf("SFile::Read(f, &m_%s[%d], sizeof(m_%s[0]), nullptr, nullptr, nullptr) == 0", columnLayout.Name, f, columnLayout.Name)
				} else {
					file.Printf("SFile::Read(f, &m_%s, sizeof(m_%s), nullptr, nullptr, nullptr) == 0", columnLayout.Name, columnLayout.Name)
				}

				if !(f+1 == elementCount && i+1 == len(target.Layout.Columns)) {
					file.Printf(" ||")
				}

				file.Printf("\n")
			}
		}
	}

	file.Printf("\t)\n\t{\n")
	file.Printf("\t\tConsoleWrite(\"Error reading %s\", WARNING_COLOR);\n", target.Definition.Name)
	file.Printf("\t\treturn false;\n")
	file.Printf("\t}\n")

	// optimization
	if !anyStrings {
		// file.Printf("\telse\n\t{\n\t\tresult = true;\n")
		goto done
	}

	file.Printf("\tif (stringBuffer) {\n")
	// file.Printf("\t\tresult = true;\n")

	// Assign values from string buffer (if present)
	for _, columnLayout := range target.Layout.Columns {
		columnDef := target.Definition.Column(columnLayout.Name)
		if columnDef.Type == dbd.String || columnDef.Type == dbd.LocString {
			elementCount := columnLayout.ArraySize
			if elementCount == -1 {
				elementCount = 1
			}

			if columnDef.Type == dbd.LocString {
				elementCount = locSize * elementCount
			}
			isArray := columnLayout.ArraySize > 0 || columnDef.Type == dbd.LocString

			for f := 0; f < elementCount; f++ {
				if isArray {
					file.Printf("\t\tm_%s[%d] = &stringBuffer[temp%sIndices[%d]];\n", columnLayout.Name, f, columnLayout.Name, f)
				} else {
					file.Printf("\t\tm_%s = &stringBuffer[temp%sIndices[0]];\n", columnLayout.Name, columnLayout.Name)
				}
			}
		}
	}

	// file.PRint

	file.Printf("\t\treturn true;\n")
	file.Printf("\t}\n")

	// \telse\n\t{\n")
	// file.Printf("\t\tresult = true;\n")

	// Assign empty string values if stringBuffer is not present.
	for _, columnLayout := range target.Layout.Columns {
		columnDef := target.Definition.Column(columnLayout.Name)
		if columnDef.Type == dbd.String || columnDef.Type == dbd.LocString {
			elementCount := columnLayout.ArraySize
			if elementCount == 0 {
				elementCount = 1
			}

			if columnDef.Type == dbd.LocString {
				elementCount = locSize * elementCount
			}
			isArray := columnLayout.ArraySize > 0 || columnDef.Type == dbd.LocString

			for f := 0; f < elementCount; f++ {
				if isArray {
					file.Printf("\t\tm_%s[%d] = \"\";\n", columnLayout.Name, f)
				} else {
					file.Printf("\t\tm_%s = \"\";\n", columnLayout.Name)
				}
			}
		}
	}

done:
	file.Printf("\treturn true;\n")
	file.Printf("}\n")

	// Add filename getter
	file.Printf("\n")

	file.Printf("const char* %sRec::GetFilename() {\n", target.Definition.Name)
	file.Printf("\treturn \"DBFilesClient\\\\%s.dbc\";\n", target.Definition.Name)
	file.Printf("}\n")

	return file.Close()
}

func (g *Generator) generateLayoutReaders() error {
	for _, target := range g.layouts {
		if err := g.generateLayoutReader(target); err != nil {
			return err
		}
	}

	return nil
}
