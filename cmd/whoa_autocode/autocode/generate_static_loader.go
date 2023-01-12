package autocode

import "unicode"

func (g *Generator) globalDBName(name string) string {
	gname := []rune("g_" + name + "DB")
	gname[2] = unicode.ToLower(gname[2])
	return string(gname)
}

func (g *Generator) generateStaticDBHeader() error {
	file, err := g.NewPrinter("src/db/ClientDatabases.hpp")
	if err != nil {
		return err
	}

	file.Printf("#ifndef DB_CLIENT_DATABASES_HPP\n")
	file.Printf("#define DB_CLIENT_DATABASES_HPP\n")
	file.Printf("\n")

	localimports := []string{
		"db/ClientDefs.hpp",
	}

	for _, localimport := range localimports {
		file.Printf("#include \"%s\"\n", localimport)
	}

	file.Printf("\n")

	for _, target := range g.layouts {
		file.Printf("extern WowClientDB<%sRec> %s;\n", target.Definition.Name, g.globalDBName(target.Definition.Name))
	}

	file.Printf("\n")

	file.Printf("void StaticDBLoadAll();\n")

	file.Printf("\n")

	file.Printf("#endif")

	return file.Close()
}

func (g *Generator) generateStaticDBLoader() error {
	file, err := g.NewPrinter("src/db/ClientDatabases.cpp")
	if err != nil {
		return err
	}

	localimports := []string{
		"db/ClientDatabases.hpp",
		// "db/ClientDefs.hpp",
	}

	for _, localimport := range localimports {
		file.Printf("#include \"%s\"\n", localimport)
	}

	file.Printf("\n")

	file.Printf("void StaticDBLoadAll() {\n")

	for _, target := range g.layouts {
		file.Printf("\t%s.Load();\n", g.globalDBName(target.Definition.Name))
	}

	file.Printf("}\n")
	file.Printf("\n")
	return nil
}

func (g *Generator) generateStaticLoader() error {
	if err := g.generateStaticDBHeader(); err != nil {
		return err
	}

	if err := g.generateStaticDBLoader(); err != nil {
		return err
	}

	return nil
}
