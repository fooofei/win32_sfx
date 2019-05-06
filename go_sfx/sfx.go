package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"
)
type ZipFilePath struct {
    FullPath string
    ZipName string
    FileInfo os.FileInfo
}


func appendEntryToZip(f *ZipFilePath, zw *zip.Writer) error {

    fih,err := zip.FileInfoHeader(f.FileInfo)
    if err != nil {
        return nil
    }
    fih.Name = f.ZipName
    fih.Method = zip.Store

    fw,err := zw.CreateHeader(fih)
    if err != nil {
        return err
    }
    if f.FileInfo.IsDir() {
        return nil
    }
    fh,err := os.Open(f.FullPath)
    if err != nil {
        return err
    }
    defer fh.Close()

    if _,err = io.Copy(fw,fh); err != nil {
        return err
    }
    return nil
}

func appendEntrysToZip(files []*ZipFilePath, zw * zip.Writer) error {
    var err error
    for idx, _ := range files {
        if err = appendEntryToZip(files[idx],zw); err != nil {
            return err
        }
    }
    return nil
}

// walk the files in dir0, not include directory
func packageSfx(targetFilePath string, prepareSfxFilePath string,
    dir0 string, entryCmd string) error {

    var err error

    files := make([]*ZipFilePath,0)

    walkfn := func(path0 string, info os.FileInfo, err error) error {
        if err !=nil {
            return nil
        }
        if path0 == dir0 {
            return nil
        }
        zipName,err := filepath.Rel(dir0, path0)
        if err != nil {
            zipName = path0
        }
        if info.IsDir() && ! strings.HasSuffix(zipName, "/") {
            zipName += "/"
        }
        files = append(files, &ZipFilePath{path0,zipName,info})
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
    if _,err = io.Copy(fw,fr); err != nil {
        return err
    }

    zw := zip.NewWriter(fw)
    defer zw.Close()
    if err = appendEntrysToZip(files,zw); err != nil {
        return err
    }
    cmt := ";The comment below contains SFX script commands\r\n" +
fmt.Sprintf("Setup=%v\r\n", entryCmd) +
"TempMode\r\n" +
"Silent=1\r\n" ;

    err = zw.SetComment(cmt)
    return err
}

func main(){


    fmt.Printf("main exit")
}
