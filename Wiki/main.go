package main

import (
  "fmt" //for formatting messages to the console
  // "net/http" //for web service
  // "log" //logging errors
  // "errors" //creating new errors
  // "os" //reading and writing files
  // "html/template" //for generating the page HTML
  // "strconv" //converting status codes to string
)

//Structure to represent each wiki page
type WikiPage struct {
  DocName string //public
  path string
  statusCode int
  Content []byte //io libraries use []byte
  HTMLContent template.HTML //public
}

var FileNotFound *WikiPage
var ErrorOccur *WikiPage

//returns the WikiPage associated with the provided document name
//If the DocName does not coorspond to a valid file, return the 404 WikiPage
func getWikiPage( DocName string ) (*WikiPage, error) {
  return nil, nil
}

//returns the WikiPage for error files
//these exist in a different directory to not conflict with regular file names
func getErrorPage( statusCode int ) (*WikiPage, error) {
  return nil, nil
}

//
func getFile( filePath string ) (*WikiPage, error) {
  return nil, nil
}

//writes the Content of the page to the file structure
func writeWikiPage( page *WikiPage ) error {
  return nil
}

//you can leave this alone
//simply grabs the file specified by the request
//needed for just general files needed by the web browser
func getFileContents(writer http.ResponseWriter, request *http.Request) {
  var path string = request.URL.Path[1:]
  http.ServeFile( writer, request, path)
}

//viewing a wiki page
func viewHandler(writer http.ResponseWriter, request *http.Request) {
  getFileContents( writer, request )
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
func homeHandler( writer http.ResponseWriter, request *http.Request ) {
  return
}

func main() {
  //Initialize error files
  FileNotFound, _ = getErrorPage( 404 )
  ErrorOccur, _ = getErrorPage( 500 )


  fmt.Printf("Running...\n")

  //TODO:
  // Create handlers

  // Create listen and serve to start server


}
