# read-dependencies-go
Tooling for parsing application manifest files to extract a list of dependencies

Currently only supports NuGet for C#, by default it will walk the directory tree from the current working directory and parse any csproj files it finds for PackageReference entries, but it can be directed to a specific file, any dependencies it reads are written to a file in CSV format.  This intermediate format can then be used as the source for some companion utilities which will do things such as creating reports or checking and reporting any known security vulnerabilities.

## Usage
A pre-built Docker image is available publicly on GitHub Container Registry, which you can run as so:
```
docker run -it -v `pwd`:/workspace -w /workspace ghcr.io/streeva/read-packages [parameters]
```
### Arguments
```bash
Usage of ./read-dependencies:
  -d string
        Specify directory (default ".")
  -e string
        Specify ecosystem (default "nuget")
  -f string
        Specify filename
  -o string
        Specify output file name (default "./dependencies.txt")
```
## Build
Clone the repo
```
git clone git@github.com:streeva/read-dependencies-go.git

cd read-dependencies-go
```
Build the application
```
go build
```
Run directly
```
./read-dependencies
```
Or build the Docker image
```
docker build . -t read-dependencies
```

## Output file format
The output dependencies file is CSV with the following fields:  
```
<Source manifest file name>,<Package Management Ecosystem>,<Package Name>,<Package Version>
```
E.g.
```
streeva.csproj,NuGet,Microsoft.CodeAnalysis.CSharp,3.7.0
```
