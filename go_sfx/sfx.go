package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
)

type ZipFilePath struct {
    FullPath string
    ZipName string
}

func packageSfx(targetFilePath string, prepareSfxFilePath string,
    dir0 string, entryCmd string) error {
    var err error

    sfxConfig := fmt.Sprintf(`
;The comment below contains SFX script commands\r
Setup={}\r
TempMode\r
Silent=1\r
`, entryCmd)

    files := make([]ZipFilePath,0)

    walkfn := func(path string, info os.FileInfo, err error) error {
        if err !=nil {
            return nil
        }
        zipName,err := filepath.Rel(path,dir0)
        if err != nil {
            zipName = path
        }
        files = append(files, ZipFilePath{path,zipName})
        return nil
    }

    if err = filepath.Walk(dir0, walkfn); err != nil {
        return err
    }


    fw,err := os.Create(targetFilePath)
    if err != nil {
        return err
    }
    defer fw.Close()
    fr,err := os.Open(prepareSfxFilePath)
    if err != nil {
        return err
    }
    defer fr.Close()

    _,_ = io.Copy(fw,fr)


    zip.Create()

}

func main(){


    fmt.Printf("main exit")
}