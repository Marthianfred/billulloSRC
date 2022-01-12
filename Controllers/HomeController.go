package Controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `
<h1>Tu Billullo API!</h1>
<p>Todos los derechos reservados</p>
`)
}
