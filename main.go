package main

import (
	"fmt"
	"os"

	"listFiles/cli"
	"listFiles/style"
)

func main() {
	rows := [][]string{}
	var filePermissions string

	args := cli.Get()
	colours := style.CreateColours()
	permissions := style.PermissionsStyles{Use: args.Colour, Colours: colours}
	permissionsMap := permissions.CreateMapPermissions()

	entries, _ := os.ReadDir(".")
	for _, entry := range entries {
		fileInfo, _ := entry.Info()
		size := float64(fileInfo.Size()) / 1024.0 / 1024.0
		filePermissions = fileInfo.Mode().Perm().String()

		// colour world
		var colorPermissionsString string
		for _, str := range filePermissions {
			literal, perm := permissionsMap[string(str)]
			if perm {
				colorPermissionsString += literal
			} else { // else delimiter
				colorPermissionsString += string(str)
			}
		}
		if fileInfo.IsDir() {
			colorPermissionsString = permissions.DirStyle() + colorPermissionsString[1:]
		}
		rows = append(rows, []string{entry.Name(), fmt.Sprintf("%f MB", size), colorPermissionsString, fileInfo.ModTime().String()})
	}
	table := style.CreateTable([]string{"File Name", "File Size", "File Permissions", "File Modified"}, rows...)
	fmt.Println(table)
}
