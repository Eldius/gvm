package installer

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Eldius/go-version-manager/config"
	"github.com/Eldius/go-version-manager/versions"
)

/*
Install installs go versions
*/
func Install(version string) {
	v := versions.FilterVersion(version, versions.ListAvailableVersions())
	if v == nil {
		fmt.Println("Version not found...")
		return
	}
	f, err := v.Download()
	if err != nil {
		fmt.Printf("Failed to download file from '%s'\n", v.LinuxAmd64)
		log.Printf("Failed to download file from '%s'\n", v.LinuxAmd64)
		log.Panic(err.Error())
	}
	unpack(f, v)
	//os.Setenv("PATH", fmt.Sprintf("%s:%s", filepath.Join(config.GetWorkspaceDir(), "bin"), os.Getenv("PATH")))
	//fmt.Println(os.Getenv("PATH"))
	fmt.Println(f)
}

func unpack(file string, v *versions.GoVersion) {
	installDir := filepath.Join(config.GetVersionsDir(), v.Name)
	fmt.Printf("Version will be installed at %s\n", installDir)
	if _, err := os.Stat(installDir); err == nil {
		fmt.Println("Install dir already exists.\nWe will try to remove it to make a fresh install...")
		if err := os.RemoveAll(installDir); err != nil {
			fmt.Println("Failed to remove old directory...")
			log.Println("Failed to remove old directory...")
			log.Panic(err.Error())
		}
	}
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tarReader := tar.NewReader(gzf)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		name := header.Name

		fmt.Println("-", name, header.Typeflag)
		dest := filepath.Join(installDir, name)
		switch header.Typeflag {
		case tar.TypeDir:
			fmt.Printf("  -> Creating folder %s\n", dest)
			if err := os.MkdirAll(dest, os.ModePerm); err != nil {
				fmt.Printf("Failed to create folder %s\n", dest)
				log.Printf("Failed to create folder %s\n", dest)
				log.Panicln(err.Error())
			}
		case tar.TypeReg:
			fmt.Printf("  -> Creating file %s\n", dest)
			f, err := os.OpenFile(dest, os.O_CREATE|os.O_RDWR, os.ModePerm)
			if err != nil {
				fmt.Printf("Failed to create file %s\n", dest)
				log.Printf("Failed to create file %s\n", dest)
				log.Panicln(err.Error())
			}
			_, err = io.Copy(f, tarReader)
			fmt.Printf("  -> Writing content to file %s\n", dest)
			if err != nil {
				fmt.Printf("Failed to write to file file %s\n", dest)
				log.Printf("Failed to write to file %s\n", dest)
				log.Panicln(err.Error())
			}
		default:
			fmt.Printf("%s : %c %s %s\n",
				"Yikes! Unable to figure out type",
				header.Typeflag,
				"in file",
				name,
			)
		}
	}
}
