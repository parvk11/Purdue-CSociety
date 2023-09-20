var resultingTranslation;

function insertSection() {
  document.body.appendChild( CreateSection() )
}

function CreateSection( appendChildren = true ) {
  const section = document.createElement("section");
  if ( appendChildren ) {
    section.appendChild( CreateHeading() );
    section.appendChild( CreateContainer() );
  }

  section.className = "section"
  return section;
}

function CreateHeading( value = "") {
  const heading = document.createElement("input");
  heading.className = "heading";
  heading.placeholder = "Heading";
  heading.value = value;
  return heading;
}

function CreateContainer() {
  const cont = document.createElement("div");
  cont.className = "container";
  const makeP = document.createElement("button");
  makeP.className = "ignore";
  makeP.textContent = "Add Paragraph";
  makeP.onclick = function() {
    cont.appendChild( CreateParagraph() )
  }
  const makeList = document.createElement("button");
  makeList.className = "ignore";
  makeList.textContent = "Add List";
  makeList.onclick = function() {
    cont.appendChild( CreateList() )
  }

  cont.appendChild(makeP);
  cont.appendChild(makeList);

  return cont;
}

function CreateParagraph( value = "" ) {
  const p = document.createElement("textarea");
  p.className = "paragraph";
  p.rows = "5";
  p.cols = "50";
  p.value = value;

  return p;
}

function CreateList() {
  const list = document.createElement("ul");
  list.className = "list";
  const makeItem = document.createElement("button");
  makeItem.className = "ignore";
  makeItem.textContent = "Add Item";
  makeItem.onclick = function() {
    const item = document.createElement("input");
    item.className = "item";
    list.appendChild( item );
  }
  list.appendChild( makeItem )
  return list;
}

function CreateItem( value = "") {
  const item = document.createElement("input");
  item.className = "item";
  item.value = value;
  return item;
}

class Container {
  constructor() {
    this.children = [];
    this.tag = "";
  }

  addChild( child ) {
    this.children.push( child );
    return;
  }

  emitDOM() {
    var result = "";
    for (var i = 0; i < this.children.length; i++) {
      if ( this.children[i].className == "ignore") { continue; }
      result += this.children[i].emitDOM();
    }
    return result;
  }

  getTag() {
    return this.tag;
  }

  getParent() {
    return this.parent;
  }
}

class ContainerStr {
  constructor( parent ) {
    this.children = [];
    this.tag = "";
    this.parent = parent;
  }
  addChild( child ) {
    this.children.push( child );
    return;
  }

  emitDOM() {
    let body = document.body;
    for (let i = 1; i < this.children.length; i++) {
      body.appendChild( this.children[i].emitDOM() );
    }
    return body;
  }

  getTag() {
    return this.tag;
  }

  getParent() {
    return this.parent;
  }
}

class Text {
  constructor( value ) {
    this.value = value;
    this.tag = "p";
  }

  emitDOM() { return this.value; }

  getTag() {
    return this.tag;
  }
}

class TextStr {
  constructor( value ) {
    this.value = value;
    this.tag = "p";
  }

  emitDOM() { return this.value; }

  getTag() {
    return this.tag;
  }
}

class Heading extends Text {
  constructor( title, size ) {
    super( title );
    this.size = size;
    this.tag = "h";
  }

  emitDOM() {
    if ( this.size == 1) {
      const heading = "<h1>" + this.value + "</h1>";
      return heading;
    } else {
      const heading = "<h2>" + this.value + "</h2>";
      return heading;
    }
  }
}

class HeadingStr extends TextStr {
  constructor( title, size ) {
    super( title );
    this.size = size;
    this.tag = "h";
  }

  emitDOM() {
    const h = CreateHeading( this.value );
    return h;
  }
}

class Section extends Container {
  constructor( DOMNode ) {
    super();
    this.tag = "section";
    super.addChild( new Heading( DOMNode.children[0].value, 2 ) );
    super.addChild( new Div( this, DOMNode.children[1] ) );
  }

  emitDOM() {
    var section = "<section>" +
    this.children[0].emitDOM() +
    this.children[1].emitDOM() +
    "</section>";
    return section;
  }
}

class SectionStr extends ContainerStr {
  constructor( parent ) {
    super( parent );
    parent.addChild( this );
    this.tag = "section";
    //super.addChild( new Heading( DOMNode.children[0].value, 2 ) );
    //super.addChild( new Div( this, DOMNode.children[1] ) );
  }

  emitDOM() {
    const sec = CreateSection( false );
    console.log(this.children);
    for (let i = 0; i < this.children.length; i++ ) {
      sec.appendChild( this.children[i].emitDOM() );
    }
    return sec;
  }

  addChild( child ) {
    super.addChild( child );
  }
}

class Div extends Container {
  constructor( parent, DOMNode ) {
    super( parent );
    this.tag = "div";
    for ( var i = 0; i < DOMNode.children.length; i++ ) {
      switch ( DOMNode.children[i].className ) {
        case "paragraph":
          super.addChild( new Paragraph( DOMNode.children[i].value ) );
          break;
        case "list" :
          super.addChild( new List( this, DOMNode.children[i] ) );
          break;
        default:
          break;
      }
    }
  }

  emitDOM() {
    var div = "<div>";
    for (var i = 0; i < this.children.length; i++) {
      div += this.children[i].emitDOM();
    }
    div += "</div>";
    return div;
  }
}

class DivStr extends ContainerStr {
  constructor( parent ) {
    super( parent );
    parent.addChild( this );
    this.tag = "div";
  }

  emitDOM() {
    const div = CreateContainer();
    for (let i = 0; i < this.children.length; i++ ) {
      div.appendChild( this.children[i].emitDOM() );
    }
    return div;
  }
}

class Paragraph extends Text {
  constructor( text ) {
    super( text );
    this.tag = "p";
  }

  emitDOM() {
    const p = "<p>" + this.value + "</p>"
    return p;
  }
}
class ParagraphStr extends TextStr {
  constructor( text ) {
    super( text );
    this.tag = "p";
  }

  emitDOM() {
    return CreateParagraph( this.value );
  }
}
class List extends Container {
  constructor( parent, DOMNode ) {
    super( parent );
    this.tag = "ul";
    for ( var i = 0; i < DOMNode.children.length; i++ ) {
      if ( DOMNode.children[i].className == "item" ) {
          super.addChild( new ListItem( DOMNode.children[i].value ) );
      }
    }
  }

  emitDOM() {
      var list = "<ul>";
      for (var i = 0; i < this.children.length; i++) {
        list += this.children[i].emitDOM();
      }
      list += "</ul>";
      return list;
  }
}
class ListStr extends ContainerStr {
  constructor( parent ) {
    super( parent );
    parent.addChild( this );
    this.tag = "ul";
  }

  emitDOM() {
    let list = CreateList();
    for (let i = 0; i < this.children.length; i++ ) {
      list.appendChild( this.children[i].emitDOM() );
    }
    return list;
  }
}

class ListItem extends Text {
  constructor( text ) {
    super( text );
    this.tag = "li";
  }

  emitDOM() {
    const li = "<li>" + this.value + "</li>";
    return li;
  }
}

class ListItemStr extends TextStr {
  constructor( text ) {
    super( text );
    this.tag = "li";
  }

  emitDOM() {
    return CreateItem( this.value );
  }
}


//I'm going to assume good inputs bc I don't want to deal with errors rn
function TranslateDOM() {
  var root = new Container();

  //Start from the DOM Body
  const doc = document.body;

  for ( var i = 0; i < doc.children.length; i++ ) {
    //if ( doc.children[i].value == "" ) { continue; }
    switch ( doc.children[i].className ) {
      case "heading":
        root.addChild( new Heading( doc.children[i].value, 1 ) );
        break;
      case "section":
        root.addChild( new Section( doc.children[i] ) );
        break;
      default:
        break;
    }
  }

  //global since we can't save variables in the HTMX
  resultingTranslation = root;

  return root;
}


//ASSUMPTIONS
// We aren't using '<' or '>' in any text
// TODO: get rid of that assumption
function findNextTag( start, str ) {
  for ( let i = start; i < str.length; i++ ) {
    if ( str[i] == '<' ) { return i; }
  }
  return -1;
}

function getTag( start, str ) {
  let result = "";
  let i = start;
  while ( i < (str.length - 1 ) && str[i] != ">" ) {
    result += str[i];
    i++;
  }
  result = result.substring(1, result.length);

  return result;
}

function getText( start, str ) {
  let result = "";
  let i = start;
  while ( i < (str.length - 1 ) && str[i] != "<" ) {
    result += str[i];
    i++;
  }
  return result;
}

function strToDOM( str ) {
  let i = 0;
  let curContainer = new ContainerStr( null );

  while ( ( i = findNextTag(i , str) ) != -1 ) {
    let tag = getTag(i, str);
    i += tag.length + 2; //i.e. we see <h1>, then it's h1, <, >
    console.log( "tag: " + tag );
    //console.log( "curContainer tag: " + curContainer.getTag() );

    if ( tag == "/" + curContainer.getTag() ) { //closing a container
      let parent = curContainer.getParent();
      if ( parent == null ) {
        return curContainer;
      }
      curContainer = parent;
    } else if ( tag[0] != "/" ){
      switch ( tag ) {
        case "section":
          const sec = new SectionStr( curContainer );
          curContainer = sec;
          break;
        case "div":
          const div = new DivStr( curContainer );
          curContainer = div;
          break;
        case "ul":
          const list = new ListStr( curContainer );
          curContainer = list;
          break;
        case "h1":
          const h = new HeadingStr( getText(i, str), 1 );
          curContainer.addChild( h );
          break;
        case "h2":
          const hh = new HeadingStr( getText(i, str), 2 );
          curContainer.addChild( hh );
          break;
        case "li":
          const item = new ListItemStr( getText(i, str ) );
          curContainer.addChild( item );
          break;
        case "p":
          const p = new ParagraphStr( getText(i, str ) );
          curContainer.addChild( p );
          break;
        default:
          console.log("tag " + tag + " wasn't found");
      }

    }

  }
  return curContainer;
}
