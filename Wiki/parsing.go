package main

import (
  //"strings"
)

//represents all nodes even though not all members are used
//for each type of node
//Go does not have classes and we need polymorphism
type Node struct {
  tag string
  value string
  parent *Node
  children []*Node
  emitHTML func( *Node ) string
}

func createNode( tag string, parent *Node ) *Node {
  return nil
}

//this notation allows us to use a dot notation to call the function
//i.e. n.addChild( c )
func ( n *Node ) addChild( c *Node ) {
  n.children = append(n.children, c)
}

//Emitters to be assigned at creation
func emitSection( c *Node ) string {
  var str string = "<section class='section'>"
  for _, value := range c.children {
    str += value.emitHTML( value )
  }
  str += "</section>"
  return str
}
func emitDiv( c *Node ) string {
  var str string = "<div class='container'><button class='ignore' onClick= 'this.parentNode.appendChild(CreateParagraph())'>Add Paragraph</button><button class='ignore' onClick= 'this.parentNode.appendChild(CreateList())' }>Add List</button>"
  for _, value := range c.children {
    str += value.emitHTML( value )
  }
  str += "</div>"
  return str
}
func emitUl( c *Node ) string {
  var str string = "<ul class='list'><button class='ignore' onClick= 'this.parentNode.appendChild(CreateItem())' }>Add Item</button>"
  for _, value := range c.children {
    str += value.emitHTML( value )
  }
  str += "</ul>"
  return str
}
func emitBody( c *Node ) string {
  var str string = ""
  for _, value := range c.children {
    str += value.emitHTML( value )
  }
  return str
}

func emitPara( t *Node ) string {
  return "<textarea class='paragraph'>" + t.value + "</textarea>"
}
func emitHn( t *Node ) string {
  return "<input type='text' class='heading' value='" + t.value + "'></input>"
}
func emitLi( t *Node ) string {
  return "<input type='text' class='item' value='" + t.value + "'></input>"
}

func wrapElement( emitted string ) string {
  return emitted
}


//returns the content between an opening and closing tag
// ex. <h1>blah blah</h1> returns blah blah
func getContent( tag string, str string ) string {
  return ""
}

//returns the text between <>
//Because all tags should be right next to one another
//we can assume all str just start with a tag
func readTag( str string ) string {
  return ""
}


// Create Intermediate Representation
// Generate a Node parent that contains the parsed text as Nodes
func createIR( str string ) *Node {
  return nil
}


// Given a string that has read only html tags, transform it into a string that has input tags
// Use all of the previously defined functions to achieve this
func TransformStr( str string ) string {
  return str
}