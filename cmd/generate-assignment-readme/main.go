package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// iterate on assignment folder
	for _, subFolder := range []string{"01-intermediate-dsa", "02-advanced-dsa"} {

		var modList strings.Builder
		modList.WriteString("# " + moduleTitle(subFolder) + "\n\n")

		// iterate on module folders
		modFolders, err := os.ReadDir("000-assignments/" + subFolder)
		if err != nil {
			panic(err)
		}

		for _, modFolder := range modFolders {
			if !modFolder.IsDir() || modFolder.Name() == ".pdf" {
				continue
			}

			modList.WriteString(fmt.Sprintf("- [%s](./%s/)", moduleTitle(modFolder.Name()), modFolder.Name()))

			// iterate problem folders
			probFolders, err := os.ReadDir("000-assignments/" + subFolder + "/" + modFolder.Name())
			if err != nil {
				panic(err)
			}

			for _, probFolder := range probFolders {
				if !probFolder.IsDir() {
					continue
				}
				modList.WriteString(fmt.Sprintf("\n\t- [%s](./%s/%s/)", probelmTitle(probFolder.Name()), modFolder.Name(), probFolder.Name()))
			}

			modList.WriteString("\n")
		}

		os.WriteFile("000-assignments/"+subFolder+"/README.md", []byte(modList.String()), 0644)
	}
}

// replace - with space and title case
func moduleTitle(s string) string {
	return strings.Title(
		strings.Join(
			strings.Split(s, "-")[1:],
			" ",
		),
	)
}

// replace space with - and title case
func probelmTitle(s string) string {
	return strings.TrimSpace(
		strings.Title(
			strings.ReplaceAll(s, "-", " "),
		),
	)
}
