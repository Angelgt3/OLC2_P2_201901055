package reports

import (
	"Backend/environment"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func RErrors(ast environment.AST) {
	//se realiza el contenido
	var contenido = "digraph Errores{ \n"
	contenido += "node [ shape=none fontname=Helvetica] \n"
	contenido += "n1 [ label = <<table><tr> \n"
	contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">No</td>\n"
	contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">Descripcio</td>\n"
	contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">Ambito</td>\n"
	contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">Linea</td>\n"
	contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">Columna</td>\n"
	contenido += "</tr>\n"
	var errores []environment.Error
	errores = ast.GetErrors()
	for error := range errores {
		contenido += "<tr><td colspan=\"2\" bgcolor=\"#ccccff\">" + strconv.Itoa((error)) + "</td>\n"
		contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">" + errores[error].Error + "</td>\n"
		contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">" + errores[error].Entorno + "</td>\n"
		contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">" + strconv.Itoa((errores[error].Lin)) + "</td>\n"
		contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">" + strconv.Itoa((errores[error].Col)) + " </td></tr>\n"
	}
	contenido += "</table> > ]; \n }"

	//la ruta el archivo
	rutaActual, err := os.Getwd()
	name := " Reporte_Errores"
	ruta := rutaActual + "/Reportes_Creados/" + name + ".dot"
	//se crear el archivo dot
	b := []byte(contenido)
	err = ioutil.WriteFile(ruta, b, 0755)
	if err != nil {
		panic(err)
	}
	//ejecutar comando
	err = exec.Command("dot", "-Tjpg", ruta, "-o", rutaActual+"/Reportes_Creados/"+name+".jpg").Run()
	if err != nil {
		panic(err)
	}

}
