package main

import (
  "fmt" //for formatting messages to the console
  "net/http" //for web service
  "log" //logging errors
  "errors" //creating new errors
  "os" //reading and writing files
  "html/template" //for generating the page HTML
  "strconv" //converting status codes to string
)

//Structure to represent each wiki page
type WikiPage struct {
  DocName string //public
  path string
  statusCode int
  Content []byte //io libraries use []byte
  HTMLContent template.HTML //public
}

var FileNotFound *WikiPage // 404
var ErrorOccur *WikiPage // 500 status

//returns the WikiPage associated with the provided document name
//If the DocName does not coorspond to a valid file, return the 404 WikiPage
func getWikiPage( DocName string ) (*WikiPage, error) {
  var filePath string = "WebComponents/Documents/" +  DocName + ".csoc"
  body, err := os.ReadFile(filePath)
  if (err != nil) { //nil here means nothing wrong
    return FileNotFound, err
  }
  return &WikiPage{DocName: DocName, path: filePath, statusCode: 200, Content: body}, err
}

//returns the WikiPage for error files
//these exist in a different directory to not conflict with regular file names
func getErrorPage( statusCode int ) (*WikiPage, error) {
  var filePath string = "WebComponents/Errors/" + strconv.Itoa(statusCode) + ".csoc"
  body, err := os.ReadFile(filePath)
  if (err != nil) { //nil here means nothing wrong
    return ErrorOccur, err
  }
  return &WikiPage{DocName: strconv.Itoa(statusCode), path: filePath, statusCode: statusCode, Content: body}, err
}

//returns a WikiPage* associated with the provided file path
// if the provided path does not exist, FileNotFound is returned instead
// ***this will become defunct with getWikiPage creation***
func getFile( filePath string ) (*WikiPage, error) {
  body, err := os.ReadFile(filePath)
  if ( err != nil ) {
    return FileNotFound, err
  }
  return &WikiPage{ DocName: "temp", path: filePath, statusCode: 200, Content: body}, err
}

//writes the Content of the page to the file structure
// errors.New("Page is empty")
func writeWikiPage( page *WikiPage ) error {
  if ( page == nil ) {
    return errors.New("Page is nil")
  }

  return os.WriteFile(page.path, page.Content, 0666)
}

//you can leave this alone
//simply grabs the file specified by the request
//needed for just general files needed by the web browser
func getFileContents(writer http.ResponseWriter, request *http.Request) {

  //fmt.Printf("%s\n", request.URL.Path)
  //fmt.Printf("%s\n", request)

  var path string = request.URL.Path[1:]
  http.ServeFile( writer, request, path)
}

//viewing a wiki page
func viewHandler(writer http.ResponseWriter, request *http.Request) {
  var path string = request.URL.Path[ len("/view/"): ]
  page, _ := getWikiPage(path)
  //writer.WriteHeader( page.statusCode )
  //http.ServeFile( writer, request, page.path)

  //check if user gave a bad Path (remember the field statusCode)
  // if it is a bad path then template error.html
    page.HTMLContent = template.HTML( page.Content )

if ( page.statusCode != 200 ) {
  parsedTemplate, _ := template.ParseFiles("WebComponents/Templates/error.html")
  err := parsedTemplate.Execute( writer, page )

  if err != nil {
    log.Println("something bad happend")
  }
} else {
  parsedTemplate, _ := template.ParseFiles("WebComponents/Templates/view.html")
  err := parsedTemplate.Execute( writer, page )

  if err != nil {
    log.Println("something bad happend")
  }
}





}


//editing a wiki page
func editHandler(writer http.ResponseWriter, request *http.Request) {
  getFileContents( writer, request )
}

//creating a new wiki page
func newHandler(writer http.ResponseWriter, request *http.Request) {
  getFileContents( writer, request )
}

//saving a wiki page
func saveHandler( writer http.ResponseWriter, request *http.Request ) {
  //redirect if trying to use anything other than a POST
  return
}


//home screen
// home screen should show a bulleted list of all available to view wiki pages
func homeHandler( writer http.ResponseWriter, request *http.Request ) {
  //TODO
  //a way to get all the files in the Documents folder

  //get a file name from a file descriptor, removing the .csoc file extension

  //use the template home.html to display the page

  //remember, setup some form of error handling to make your life easier

  return
}

func main() {
  //Initialize error files
  FileNotFound, _ = getErrorPage( 404 )
  ErrorOccur, _ = getErrorPage( 500 )


  fmt.Printf("Running...\n")
  //TODO:
  // Create handlers
  http.HandleFunc( "/", getFileContents )
  http.HandleFunc( "/view/", viewHandler )
  // Create listen and serve to start server

  log.Fatal(http.ListenAndServe(":8000", nil) )


}
