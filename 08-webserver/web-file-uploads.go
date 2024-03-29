package main

import (
	"io"
	"log"
	"net/http"
	"io/ioutil"
	"html/template"
	"os"
)

const (
	UPLOAD_DIR = "/home/wwm/go-path/temp-file"
)

func main() {
    http.HandleFunc("/", listHandler)
    http.HandleFunc("/view", viewHandler)
    http.HandleFunc("/upload", uploadHandler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}

/* *
 * 直接输出图片
 */
func viewHandler(w http.ResponseWriter, r *http.Request) {
    imageId := r.FormValue("id")
    imagePath := UPLOAD_DIR + "/" + imageId
    if exists := isExists(imagePath);!exists {
        http.NotFound(w, r)
        return
    }
    // - 1. 设置Header ..
    w.Header().Set("Content-Type", "image")
    // - 2. 输出文件 
    http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    return os.IsExist(err)
}

// - 文件上传页面
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, err := template.ParseFiles("upload.html")
        if err != nil {
            http.Error(w, err.Error(),http.StatusInternalServerError)
            return
        }
        t.Execute(w, nil)
        return
    }
    if r.Method == "POST" {
        f, h, err := r.FormFile("image")
        if err != nil {
            http.Error(w, err.Error(),
            http.StatusInternalServerError)
            return
        }
        filename := h.Filename
        defer f.Close()
        t, err := os.Create(UPLOAD_DIR + "/" + filename)
        if err != nil {
            http.Error(w, err.Error(),
            http.StatusInternalServerError)
            return
        }
        defer t.Close()
        if _, err := io.Copy(t, f); err != nil {
            http.Error(w, err.Error(),
            http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/view?id="+filename,
        http.StatusFound)
    }
}

// - 文件展示页面
func listHandler(w http.ResponseWriter, r *http.Request) {
    fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
    if err != nil {
        http.Error(w, err.Error(),
        http.StatusInternalServerError)
        return
    }
    locals := make(map[string]interface{})
    images := []string{}
    for _, fileInfo := range fileInfoArr {
        images = append(images, fileInfo.Name())
    }
    locals["images"] = images
    t, err := template.ParseFiles("list.html")
    if err != nil {
        http.Error(w, err.Error(),
        http.StatusInternalServerError)
        return
    }
    t.Execute(w, locals)
}
