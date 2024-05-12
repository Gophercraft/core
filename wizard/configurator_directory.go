package wizard

var global_directory = ""

func Directory() string {
	if global_directory != "" {
		return global_directory
	}

	return get_directory()
}

func SetDirectory(directory string) {
	global_directory = directory
}
