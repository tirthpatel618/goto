package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

const configFileName = ".goto.json"

type Bookmarks map[string]string

func getConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, configFileName)
}

func loadBookmarks() Bookmarks {
	path := getConfigPath()
	file, err := os.ReadFile(path)
	if err != nil {
		return Bookmarks{}
	}
	var bm Bookmarks
	json.Unmarshal(file, &bm)
	return bm
}

func saveBookmarks(bm Bookmarks) {
	data, _ := json.MarshalIndent(bm, "", "  ")
	os.WriteFile(getConfigPath(), data, 0644)
}

func Execute() {
	args := os.Args[1:]
	bookmarks := loadBookmarks()

	if len(args) == 0 {
		fmt.Println("Usage: goto [add|rm|edit|list|name]")
		return
	}

	switch args[0] {
	case "add":
		if len(args) != 3 {
			fmt.Println("Usage: goto add <name> <path>")
			return
		}
		bookmarks[args[1]] = args[2]
		saveBookmarks(bookmarks)
		fmt.Printf("Added %s → %s\n", args[1], args[2])

	case "rm":
		if len(args) != 2 {
			fmt.Println("Usage: goto rm <name>")
			return
		}
		delete(bookmarks, args[1])
		saveBookmarks(bookmarks)
		fmt.Printf("Removed bookmark: %s\n", args[1])

	case "list":
		if len(bookmarks) == 0 {
			fmt.Println("No bookmarks saved.")
			return
		}
		names := make([]string, 0, len(bookmarks))
		for name := range bookmarks {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Printf("%s → %s\n", name, bookmarks[name])
		}

	case "edit":
		if len(args) != 3 {
			fmt.Println("Usage: goto edit <name> <new_path>")
			return
		}
		name, newPath := args[1], args[2]
		if _, ok := bookmarks[name]; !ok {
			fmt.Printf("Bookmark '%s' does not exist.\n", name)
			return
		}
		bookmarks[name] = newPath
		saveBookmarks(bookmarks)
		fmt.Printf("Updated %s → %s\n", name, newPath)
	

	default:
		dest, ok := bookmarks[args[0]]
		if !ok {
			fmt.Printf("No such bookmark: %s\n", args[0])
			os.Exit(1)
		}
		fmt.Print(dest)
	}
}
