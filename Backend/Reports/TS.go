package reports

import (
	"Backend/environment"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func Rts(ast environment.AST) {
	//se realiza el contenido
	var contenido = "digraph TS{ \n"
	contenido += "node [ shape=none fontname=Helvetica] \n"
	contenido += "n1 [ label = <<table><tr> \n"
	contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">Id</td>\n"
	contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">Tipo símbolo</td>\n"
	contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">Tipo dato</td>\n"
	contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">Ámbito</td>\n"
	contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">Línea </td>\n"
	contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">Columna</td>\n"
	contenido += "</tr>\n"

	var Ts []environment.RS
	Ts = ast.GetRTS()

	for simbolo := range Ts {
		var temp_tipo = ""
		if Ts[simbolo].Tipo_dato == environment.INTEGER {
			temp_tipo = "Int"
		} else if Ts[simbolo].Tipo_dato == environment.FLOAT {
			temp_tipo = "Float"
		} else if Ts[simbolo].Tipo_dato == environment.STRING {
			temp_tipo = "String"
		} else if Ts[simbolo].Tipo_dato == environment.CHAR {
			temp_tipo = "Char"
		} else if Ts[simbolo].Tipo_dato == environment.BOOLEAN {
			temp_tipo = "Boolean"
		} else {
			temp_tipo = ""
		}

		contenido += "<tr><td colspan=\"2\" bgcolor=\"#ccccff\">" + Ts[simbolo].Id + "</td>\n"
		contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">" + Ts[simbolo].Tipo_simbolo + "</td>\n"
		contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">" + temp_tipo + "</td>\n"
		contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">" + Ts[simbolo].Entorno + "</td>\n"
		contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">" + strconv.Itoa((Ts[simbolo].Lin)) + "</td>\n"
		contenido += "<td colspan=\"2\" bgcolor=\"#ccccff\">" + strconv.Itoa((Ts[simbolo].Col)) + "</td></tr>\n"

	}
	contenido += "</table> > ]; \n }"

	//la ruta el archivo
	rutaActual, err := os.Getwd()
	name := " Reporte_TS"
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
