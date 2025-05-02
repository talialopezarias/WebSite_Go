//Separar la lógica del servidor de la representación visual de los datos. Facilita la modularidad del código.

package models

import "html/template"

//Se utilizara para almacenar la información que se pasa del servidor a las plantillas HTML.
type PageData struct {
	Title         string
	Author        string
	Welcome       string
	ErrorCode     int
	ErrorMessage  string
	HeadContent   template.HTML
	NavbarContent template.HTML
}
