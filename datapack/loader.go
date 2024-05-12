package datapack

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

type Loader struct {
	packs []*Pack
}

// Load a
func Load(directory string) (loader *Loader, err error) {
	var (
		fi os.FileInfo
		de []os.DirEntry
	)
	fi, err = os.Stat(directory)
	if err != nil {
		return
	}

	if !fi.IsDir() {
		err = fmt.Errorf("datapack: Load() can only load directories containing datapacks")
		return
	}

	loader = new(Loader)

	de, err = os.ReadDir(directory)
	if err != nil {
		return
	}

	var pack_files []string

	for _, directory_entry := range de {
		pack_files = append(pack_files, directory_entry.Name())
	}

	sort.Strings(pack_files)

	for _, directory_entry := range pack_files {
		datapack_path := filepath.Join(directory, directory_entry)

		var pack *Pack
		pack, err = Open(datapack_path)
		if err != nil {
			return
		}

		loader.packs = append(loader.packs, pack)
	}

	err = loader.resolve_load_order()

	return
}

func (loader *Loader) check_dependency(dependant *Pack, pack_dependency *PackDependency) (err error) {
	for _, pack := range loader.packs {
		if pack.info.ID == pack_dependency.ID {
			if pack.info.Version < pack_dependency.MinimumVersion {
				err = fmt.Errorf("datapack: (*Loader).check_dependency: datapack %s version %d is too old to be depended on by %s", pack_dependency.ID, pack.info.Version, dependant.info.ID)
				return
			}

			dependant.dependencies = append(dependant.dependencies, pack)
			return
		}
	}

	err = fmt.Errorf("datapack: (*Loader).check_dependency: datapack %s depended on by %s is missing", pack_dependency.ID, dependant.info.ID)
	return
}

func find_cycle_in_dependency(path []string, dependant *Pack, dependency *Pack) (err error) {
	if dependency == dependant {
		err = fmt.Errorf("dependency cycle in path %s", strings.Join(path, " 🡒 "))
		return
	}

	dependant.transitive_dependencies = append(dependant.transitive_dependencies, dependency)

	for _, dep_dependency := range dependency.dependencies {
		dependency_path := append(path, dep_dependency.info.ID)
		if err = find_cycle_in_dependency(dependency_path, dependant, dep_dependency); err != nil {
			return
		}
	}

	return
}

func find_cycle(dependant *Pack) (err error) {
	path := []string{dependant.info.ID}

	for _, dependency := range dependant.dependencies {
		dependency_path := append(path, dependency.info.ID)
		if err = find_cycle_in_dependency(dependency_path, dependant, dependency); err != nil {
			return
		}
	}

	return
}

// resolves the load order of items within the Loader.
// returns error in the case of a cyclic dependency
// Items
func (loader *Loader) resolve_load_order() (err error) {
	// Check presence
	for _, pack := range loader.packs {
		for _, dependency := range pack.info.Dependencies {
			err = loader.check_dependency(pack, &dependency)
			if err != nil {
				return
			}
		}
	}

	// Check for dependency cycle
	for _, pack := range loader.packs {
		err = find_cycle(pack)
		if err != nil {
			return
		}
	}

	// Change load order according to dependency
	sort.Slice(loader.packs, func(i, j int) bool {
		pack_i := loader.packs[i]
		pack_j := loader.packs[j]

		if slices.Contains(pack_j.transitive_dependencies, pack_i) {
			return true
		}

		if slices.Contains(pack_i.transitive_dependencies, pack_j) {
			return false
		}

		// base pack is less than non-base pack
		if pack_i.info.Base && !pack_j.info.Base {
			return true
		}

		// non-base pack is greater than base pack
		if !pack_i.info.Base && pack_j.info.Base {
			return false
		}

		// use normal sorting
		return pack_i.info.ID < pack_j.info.Name
	})

	return
}
