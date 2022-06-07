/*
 * fatfile
 * print all files under given path, sorted by size
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

const ONE_K = 1000
const ONE_MB = 1000000

func main() {

    var listFileElements = []FileElement{}
    var path string
    var totalFileCount int
    var totalSize int64

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

    filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        totalFileCount ++
        fileSize := info.Size()
        totalSize += fileSize

        fileElement := FileElement{path: path, size: fileSize}
        listFileElements = append(listFileElements, fileElement)

        return nil
    })

    sort.Slice(listFileElements, func(i, j int) bool {
        return listFileElements[i].size < listFileElements[j].size
    })

    for  _, e := range listFileElements {
        fmt.Printf("%0.2fM %s\n",  float64(e.size) / float64(ONE_MB), e.path)
    }
    fmt.Printf("total files: %d total size: %0.2fM\n",  totalFileCount, float64(totalSize) / float64(ONE_MB))

}
