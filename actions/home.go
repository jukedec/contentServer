package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
  "fmt"
//   "bytes"
//   "strings"
  "io"
  "os"
"path/filepath"
//   "reflect"
  "os/exec"
  "log"
  "math/rand"
  "strconv"
  "strings"
//   "strings"
//   "io/ioutil"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
// 	return c.Render(http.StatusOK, r.HTML("index.html"))
  return c.Render(http.StatusOK, r.HTML("jukeDecIndex.html"))
}

func UploadIndex(c buffalo.Context) error {
  return c.Render(http.StatusOK, r.HTML("uploadIndex.html"))
}

func UploadPost(c buffalo.Context) error {
  
//   fmt.Println("UPLOADING now")

//   f, err := c.File("files")
//   if err != nil {
//     return err
//   }
//   fmt.Println(f)
  
//   for _, track := range f {
// 		fmt.Println("A track is:", track)
// 	}
  req := c.Request()
//   w := c.Response()
  
//       req.ParseMultipartForm(32 << 20) // limit your max input length!
//     var buf bytes.Buffer
//     // in your case file would be fileupload
//     file, header, err := req.FormFile("files")
//     if err != nil {
//         panic(err)
//     }
//     defer file.Close()
//     name := strings.Split(header.Filename, ".")
//     fmt.Printf("File name %s\n", name[0])
//     // Copy the file data to my buffer
//     io.Copy(&buf, file)
//     // do something with the contents...
//     // I normally have a struct defined and unmarshal into a struct, but this will
//     // work as an example
//     contents := buf.String()
//     fmt.Println(contents)
//     // I reset the buffer in case I want to use it again
//     // reduces memory allocations in more intense projects
//     buf.Reset()
//     // do something else
//     // etc write header
    
  //parse the multipart form in the request
		err := req.ParseMultipartForm(100000)
		if err != nil {
			return err // http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
		}

		//get a ref to the parsed multipart form
		m := req.MultipartForm

		//get the *fileheaders
		files := m.File["files"]
  dir := ""
  req_id := ""
  req_id = strconv.Itoa(rand.Intn(9999999))
//   req_id = c.Param("bandname")
  fmt.Println("GOT NAME?")
  
  fmt.Println(req_id)
  
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				return err // http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
			}
			//create destination file making sure the path is writeable.
      
       // DOESNT WORK FOR NOW
//       		var irid interface{}
//           irid = c.Session().Get("requestor_id")
//       req_id = fmt.Sprintf("%v", irid)
      
      

      
//       xType := reflect.TypeOf(c.Session())
//       fmt.Println(irid)
      
      dir = filepath.Join(".", "uploads/" + req_id)
        if err := os.MkdirAll(dir, 0755); err != nil {
          return err
        }
//       fmt.Println(dir)
      
      
			dst, err := os.Create(dir + "/" + files[i].Filename)
			defer dst.Close()
			if err != nil {
				return err //http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				return err // http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
			}

		}

    fmt.Println(dir)
  fmt.Println(req_id)
  
  commandDir := "/home/jukedec/go/src/github.com/jukedec/Hugo-Mp3-Preprocessor/bash.sh"
  mp3Dir := "/home/jukedec/go/src/github.com/jukedec/content_server/uploads/" + req_id
  siteDir := "/home/jukedec/go/myTemp/" + req_id
  serveDir := "/var/www/html/sites"
  command := commandDir
  
  args := []string{mp3Dir, siteDir, serveDir, req_id}
//   args := mp3Dir + " " + siteDir
  fmt.Println(command)
  fmt.Println(args)
  
//   outstuff, probablyError := exec.Command(command, args).Output()
// 	fmt.Printf("maybe an error: %v", probablyError)
  
//     out, commandError := exec.Command("date").Output()
//     if commandError != nil {
//         log.Fatal(commandError)
//     }
    
  
  
  cmd, commandError := exec.Command(command, args...).Output()
      if commandError != nil {
        log.Fatal(commandError)
    }
    
	fmt.Printf("Running command and waiting for it to finish...")
// 	processErr := cmd.Run()
// 	fmt.Printf("Command finished with error: %v", processErr)
  
  fmt.Printf("The OUTPUT: %s\n", cmd)
  
//   out, err := exec.Command("date").Output()
  
  
//   fmt.Println(name)
//   siteUrl := "http://play.jukedec.com/" + req_id + "/public"
  
  
//   fmt.Printf("RAN OUTPUT: %s\n", outstuff)
//     line := string(b)
//     if(strings.Contains(line, "theArtistName=") ){
      
//       url = line[14:len(line)]
//     }
  
  n:= string(cmd)
  res1 := strings.Index(n, "theArtistName=")
  fmt.Println("Result 1: ", res1) 
  inputFmt:= n[res1+15:len(n)-1]
  fmt.Println("Result 2: ", inputFmt)
  
  baseUrl := "http://play.jukedec.com/"
  url := inputFmt
//   cmd, e := exec.Run(command, nil, nil, exec.DevNull, exec.Pipe, exec.MergeWithStdout); e == nil {

//   }
  siteUrl := baseUrl + url
  c.Set("siteUrl", siteUrl)
  fmt.Println(siteUrl)
  
  return c.Render(http.StatusOK, r.HTML("uploadAction.html"))
}
// a.GET("/", jukedecHomeHandler)


// a.GET("/",(c buffalo.Context) error {
// 	return c.Render(http.StatusOK, r.HTML("jukedec_home_handler/jukedec_home_handler.html"))
// }