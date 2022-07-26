package routers

import (
	"io"
	"net/http"
	"os"

	"strings"

	"github.com/jarrieta31/twittor/bd"
	"github.com/jarrieta31/twittor/models"
)

/*SubirBanner sube el Banner al servidor*/
func SubirBanner(w http.ResponseWriter, r *http.Request) {

	/*Se procesa como un formulario de hatml, se captuar la variable banner del formulario
	 */
	file, handler, err := r.FormFile("banner")

	/*Creamos una variable para obtener la extensión del archivo, el nombre de archico
	viene en handler. Al hacer un Split obtenemos un vector y haciendo de esta manera
	se obtiene un string directamente con el valor*/
	var extension = strings.Split(handler.Filename, ".")[1]

	//esta es la ruta donde se guarda el archivo
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	/*Función del sistema operativo que sirve para abrir un archivo y devuelve 2 parámetros.
	Crea el espacio en disco para el archivo*/
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	/*Copia el archivo recibido y lo almacena en f (espacio en disco recien creado)
	En este punto el archivo se grabo en disco*/
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	/*Ahora debemos actualizar la base de datos agregando el nombre del banner en el usuario*/
	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension

	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el banner en la BD! "+err.Error(), http.StatusBadRequest)
		return
	}

	/*Si todo salió bien*/
	w.Header().Set("Content-Type", "applicaction/json")
	w.WriteHeader(http.StatusCreated)

}
