// Path manipulation
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("filepath package absolute")

	// if len(os.Args) != 2 { // ensure path is specified
	// 	fmt.Println("Please specify a path.")
	// 	return
	// }

	fmt.Println("OS Args", os.Args)

	fpath := "../go-docker-tutorial"
	// fpath := "~/Downloads/"

	fabs, err := filepath.Abs(fpath)

	if err != nil {
		fmt.Println("Error filepath Abs", err)
		return
	}
	fmt.Println("Absolute", fabs)

	fbase := filepath.Base(fpath)

	fmt.Println("Base", fbase)
	fmt.Println(filepath.Base("/foo/bar/baz.js"))
	fmt.Println(filepath.Base("/foo/bar/baz"))
	fmt.Println(filepath.Base("/foo/bar/baz/"))
	fmt.Println(filepath.Base("dev.txt"))
	fmt.Println(filepath.Base("../todo.txt"))
	fmt.Println(filepath.Base(".."))
	fmt.Println(filepath.Base("."))
	fmt.Println(filepath.Base("/"))
	fmt.Println(filepath.Base(""))

	fmt.Println("filepath Clean")
	fmt.Println(filepath.Clean("/.."))
	fmt.Println(filepath.Base("//var/www/html/go-system-programming/../go-learning"))
	fmt.Println(filepath.Base("//var/www/html/go-system-programming/.."))
	fmt.Println(filepath.Base("/var/www/html/go-system-programming"))

	fmt.Println("***filepath Dir***:")
	fmt.Println(filepath.Dir("/foo/bar/baz.js"))
	fmt.Println(filepath.Dir("/foo/bar/baz"))
	fmt.Println(filepath.Dir("/foo/bar/baz/"))
	fmt.Println(filepath.Dir("/dirty//path///"))
	fmt.Println(filepath.Dir("dev.txt"))
	fmt.Println(filepath.Dir("../todo.txt"))
	fmt.Println(filepath.Dir(".."))
	fmt.Println(filepath.Dir("."))
	fmt.Println(filepath.Dir("/"))
	fmt.Println(filepath.Dir(""))

	fmt.Println("***filepath EvalSymlinks***:")
	link, err := filepath.EvalSymlinks(fpath)
	if err != nil {
		fmt.Printf("Cannot read symbolic link '%s', error was: %s\n", fpath, err)
		return
	}
	fmt.Println("Link", link)

	fmt.Println("***filepath Ext***:")
	fmt.Printf("No dots: %q\n", filepath.Ext("index"))
	fmt.Printf("One dot: %q\n", filepath.Ext("index.js"))
	fmt.Printf("Two dots: %q\n", filepath.Ext("main.test.js"))

	fmt.Println("***filepath FromSlash***:")
	// Option 1
	examplePath1 := "dir" + string(os.PathSeparator) + "example"
	fmt.Println("PathSeparator: " + examplePath1)

	// Option 2
	examplePath2 := filepath.FromSlash("dir/example")
	fmt.Println("FromSlash: " + examplePath2)

	fmt.Println("***filepath Glob***:")
	// pattern := "*input*"

	// upperDirPattern := "../*go*"
	binDirPattern := "/usr/*bin"

	matches, err := filepath.Glob(binDirPattern)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(matches)

	fmt.Println("***filepath IsAbs***:")
	fmt.Println(filepath.IsAbs("/home/gopher"))
	fmt.Println(filepath.IsAbs(".bashrc"))
	fmt.Println(filepath.IsAbs(".."))
	fmt.Println(filepath.IsAbs("."))
	fmt.Println(filepath.IsAbs("/"))
	fmt.Println(filepath.IsAbs(""))

	fmt.Println("***filepath Join***:")
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("a/b", "c"))
	fmt.Println(filepath.Join("a/b", "/c"))

	fmt.Println(filepath.Join("a/b", "../../../xyz"))

	fmt.Println("***filepath Match***:")
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo"))
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo/bar"))
	fmt.Println(filepath.Match("/home/?opher", "/home/gopher"))
	fmt.Println(filepath.Match("/home/\\*", "/home/*"))

	fmt.Println("***filepath Rel***:")

	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}
	base := "/a"

	for _, p := range paths {
		rel, err := filepath.Rel(base, p)
		fmt.Printf("%q: %q - %v\n", p, rel, err)
	}

	fmt.Println("***filepath Split***:")
	pathsRel := []string{
		"/home/arnie/amelia.jpg",
		"/mnt/photos/",
		"rabbit.jpg",
		"/usr/local//go",
	}

	for _, p := range pathsRel {
		dir, file := filepath.Split(p)
		fmt.Printf("input: %q\n\tdir: %q\n\tfile: %q\n", p, dir, file)
	}

	fmt.Println("***filepath SplitList***:", filepath.SplitList("/a/b/c:/usr/bin"))

	fmt.Println("***filepath ToSlash***:")

	examplePath3 := filepath.ToSlash("dir/example")
	fmt.Println("ToSlash: " + examplePath3)

	fmt.Println("***filepath Walk***:")
	tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)
	os.Chdir(tmpDir)

	subDirToSkip := "skip"

	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}

	fmt.Println("***filepath Walk walkRoot***:")
	walkRoot()
}

func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}

	return tmpDir, nil
}
