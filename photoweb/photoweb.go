package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"io"
	"os"
	"html/template" 
	"fmt"
)

const (
	UPLOAD_DIR = "./uploads"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// io.WriteString(w, "<form method=\"POST\" action=\"/upload\" "+
		// " enctype=\"multipart/form-data\">"+
		// "Choose an image to upload: <input name=\"image\" type=\"file\" />"+
		// "<input type=\"submit\" value=\"Upload\" />"+
		// "</form>")
		t, err := template.ParseFiles("upload.html")
		if err != nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}
	fmt.Println("post", r.Method)
	if r.Method == "POST" {
		fmt.Println("post11111")
		f, h, err := r.FormFile("image")
		if err != nil {
			fmt.Println("post err", err.Error())
			http.Error(w, err.Error(),
			http.StatusInternalServerError)
			return
		}
		fmt.Println("post name", h.Filename)

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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath);!exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{})(err error){
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		return
	}
	err = t.Execute(w, locals)
	return
}

func listHandler(w http.ResponseWriter, r *http.Request){
	// var files []os.FileInfo
	fileInfoArr, err := ioutil.ReadDir("./uploads")
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
	if err = renderHtml(w, "list", locals); err != nil {
		http.Error(w, err.Error(),
		http.StatusInternalServerError)
	   }
	// var listHtml string
	// for _, fileInfo := range fileInfoArr {
	// 	imgid := string(fileInfo.Name)
	// 	str := "<li><a href=\"/view?id=" + imgid
	// 	listHtml += str
	// }
	// io.WriteString(w, "<ol>"+listHtml+"</ol>")
}

func main() {
	http.HandleFunc("/", listHandler) 
	http.HandleFunc("/view", viewHandler) 
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}


