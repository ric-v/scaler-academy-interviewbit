package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	out, _ := exec.Command("tree", "-L", "2", "000-assignments").CombinedOutput()
	input := string(out)

	replacer := strings.NewReplacer(
		"├", "",
		"│", "",
		"─", "",
		"└", "",
		" ", "",
		"\t", "",
	)

	var markdown, moduleName string

	for _, line := range strings.Split(input, "\n") {

		line = replacer.Replace(line)

		if strings.Contains(line, "000-assignments") || line == "" || strings.Contains(line, "directories") {
			continue
		}

		var fileName string
		for _, word := range strings.Split(line, "-") {
			fileName += word + "-"
		}
		fileName = strings.TrimSpace(strings.TrimSuffix(fileName, "-"))

		if strings.Contains(fileName, "intermediate-dsa") || strings.Contains(fileName, "advanced-dsa") {
			markdown += fmt.Sprintf("- [%s](./000-assignments/%s/)\n", moduleTitle(fileName), fileName)
			moduleName = fileName
		} else if !strings.Contains(fileName, ".") {
			markdown += fmt.Sprintf("\t- [%s](./000-assignments/%s/%s/)\n", probelmTitle(fileName), moduleName, fileName)
		}
	}
	fmt.Println(markdown)

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
