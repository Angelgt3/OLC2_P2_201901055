package environment

type TipoExpresion int

const (
	INTEGER   TipoExpresion = iota //0
	FLOAT                          //1
	STRING                         //2
	BOOLEAN                        //3
	NULL                           //4
	CHAR                           //5
	ARRAY                          //6
	A_INTEGER                      //7
	A_FLOAT                        //8
	A_STRING                       //9
	A_BOOLEAN                      //10
	A_NULL                         //11
	A_CHAR                         //12
	STRUCT                         //13
)
