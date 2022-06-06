/*
 * fatfile
 * print all files larger than 1M, sorted by size
 *
 * author: Chrisotpher Minson
 *
 */

package main

import  (
    "fmt"
    "os"
    "sort"
    "path/filepath"
)

type FileElement struct {
    path  string
    size  int64
}

const ONE_MB = 1000000

func main() {

    var listFileElements = []FileElement{}
    var path string

    switch len(os.Args) {
    case 1:
        cwd, err :=  os.Getwd()
        if err != nil {
            fmt.Println("Exit:  failed to find current directory")
            os.Exit(0)
        }
        path = cwd
    case 2:
        path = os.Args[1]
    default:
        fmt.Println("incorrect arguments:  fatfile <path>")
        os.Exit(0)
    }
    fmt.Printf("Scanning: %s\n", path)

    err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Println(err)
            return err
        }

        fileSize := info.Size()
        if fileSize > ONE_MB {

            fileElement := FileElement{path: path, size: fileSize}
            listFileElements = append(listFileElements, fileElement)
        }

        return nil
    })
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }

    sort.Slice(listFileElements, func(i, j int) bool {
        return listFileElements[i].size < listFileElements[j].size
    })

    for  _, e := range listFileElements {
        fmt.Printf("%dM %s\n",  e.size / ONE_MB, e.path)
    }
}

