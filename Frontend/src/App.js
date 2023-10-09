import './App.css';
import { Editor } from "@monaco-editor/react"
import { useState } from 'react'

function App() {

  const url = 'http://localhost:5000/ejecutar';

  const [contenido, setContenido] = useState("");
  const [resultado, setResultado] = useState("");
  
  

  function createPost() {
    //console.log(contenido);
    fetch(url,{
      method: 'POST',
      body: contenido
    })
    .then((response) => response.json()) 
    .then((response) => setResultado(response.result))

  }

  return (
    <div>
      <div id="title">
        <div id="opciones">
          <button onClick={createPost} className="button">
            Ejecutar
          </button>
        </div>
      </div>

      <div id="I1"></div>
      
      <div id ="M1">
        <Editor
          theme='vs-dark'
          defaultLanguage="Swift" //lenguaje de ayuda monaco
          value = {contenido} //el valor inicial del monaco
          onChange={(value) => setContenido(value)}
        />
      </div>

      <div id="I2"></div>
      
      <div id ="M2">
        <Editor
          theme='vs-dark'
          defaultLanguage="Plain Text" //lenguaje de ayuda monaco
          value = {resultado}  //el valor inicial del monaco
        />
      </div>

    </div>
    
  )
}

export default App;