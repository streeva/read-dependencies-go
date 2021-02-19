package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type NuGet struct {
}

type Project struct {
	ItemGroups []ItemGroup `xml:"ItemGroup"`
}

type ItemGroup struct {
	PackageReferences []PackageReference `xml:"PackageReference"`
}

type PackageReference struct {
	Name		string	`xml:"Include,attr"`
	Version	string	`xml:"Version,attr"`
}

func (n NuGet) IsManifestFile(filename string) bool {
	if filepath.Ext(filename) == ".csproj" {
		return true
	}

	return false
}

func (n NuGet) ReadDependencies(manifestFile string) ([]Dependency, error) {
	var dependencies []Dependency
		fmt.Println("Opening " + manifestFile)
		file, err := os.Open(manifestFile)
		if err != nil {
			return make([]Dependency, 0), err
		}

		byteValue, err := ioutil.ReadAll(file)
		if err != nil {
			return make([]Dependency, 0), err
		}

		var project Project
		xml.Unmarshal(byteValue, &project)
		for _, itemGroup := range project.ItemGroups {
			for _, packageReference := range itemGroup.PackageReferences {
				dependencies = append(dependencies, Dependency { 
					ManifestFilename:	filepath.Base(manifestFile),
					Ecosystem:				"NuGet",
					Name:							packageReference.Name,
					Version:					packageReference.Version,
				})
			}
		}

	return dependencies, nil
}