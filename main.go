package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const EXIT_SUCCESS = 0
const EXIT_FAILURE = 1

type Ecosystem interface {
	IsManifestFile(filename string) bool
	ReadDependencies(manifestFile string) ([]Dependency, error)
}

type Dependency struct {
	ManifestFilename	string
	Ecosystem					string
	Name							string
	Version						string
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(EXIT_FAILURE)
	}
}

func main() {
	var dirname string
	var filename string
	var ecosystem string
	var ofile string

	flag.StringVar(&dirname, "d", ".", "Specify directory")
	flag.StringVar(&filename, "f", "", "Specify filename")
	flag.StringVar(&ecosystem, "e", "nuget", "Specify ecosystem")
	flag.StringVar(&ofile, "o", "./dependencies.txt", "Specify output file name")
	flag.Parse()

	if ecosystem != "nuget" {
		fmt.Println("unsupported package ecosystem")
		os.Exit(3)
	}

	var eco NuGet
	manifestfiles, err := FindManifestFiles(filename, dirname, eco)
	check(err)
	if len (manifestfiles) <= 0 {
		fmt.Println("no manifest files found")
		os.Exit(EXIT_SUCCESS)
	}

	dependencies, err := ReadDependencies(manifestfiles, eco)
	check(err)

	file, err := os.Create(ofile)
	for _, dependency := range dependencies {
		file.WriteString(fmt.Sprintln(fmt.Sprintf("%s,%s,%s,%s", dependency.ManifestFilename, dependency.Ecosystem, dependency.Name, dependency.Version)))
	}
}

func FindManifestFiles(filename string, directory string, ecosystem Ecosystem) ([]string, error) {
	var projectfiles []string
	if len(filename) > 0 {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return make([]string, 0), err
		}

		projectfiles = append(projectfiles, filename)
	} else {
		err := filepath.Walk(directory,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
						return err
				}
				
				if ecosystem.IsManifestFile(path) {
					projectfiles = append(projectfiles, path)
					fmt.Println("Located manifest file " + path)
				}

				return nil
			})

		if err != nil {
			return make([]string, 0), err
		}
	}

	return projectfiles, nil;
}

func ReadDependencies(manifestFiles []string, ecosystem Ecosystem) ([]Dependency, error) {
	var dependencies []Dependency
	for _, manifestFile := range manifestFiles {
		manifestDependencies, err := ecosystem.ReadDependencies(manifestFile)
		if err != nil {
			return dependencies, nil
		}

		dependencies = append(dependencies, manifestDependencies...)
	}

	return dependencies, nil
}