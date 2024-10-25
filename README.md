# 1. Como compilar y ejecutar.
Para compilar y ejecutar:
- Poner el archivo "fuerza_bruta.go" en la misma carpeta que el archivo "archive.pdf.gpg".
- Para compilar ejecutar el siguiente comando: "go build fuerza_bruta.go" (esto generará un ejecutable llamado "fuerza_bruta").
- Para ejecutar el programa usar: "./fuerza_bruta".
- Cambiar el valor de la variable "pass_max_length" para aumentar el tamaño de la posible contraseña (por defecto ccomo máximo 6 caracteres).

# 2. ¿Por qué Go?
He escogido Go como lenguaje ya que en unas pruebas preliminares, obtuve los siguientes resultados:
Al probar con lenguajes como python y c++ obtenía tiempos algo superiores a Go. Tanto en la ejecución de un solo proceso al probar de descifrar un archivo, como al momento de incluir varios procesos o hilos a la ejecución. Probé 10 hilos y que cada hilo descifrase unas 100 veces. Python tardaba unos segundos más que los otros. En cambio c++, era mas lento por centésimas. En estos programas más o menos se suelen probar la mitad de las posibles contraseñas mas una para encontrar la correcta, 26*26*26*26 en el caso de 4 caracteres, lo que nos deja 456976 contraseñas a probar. Al ser tantas (más si el número de caracteres aumenta) me decidí por Go para que en el peor de los casos la ejecución tarde menos.

# 3. No diccionarios.
No he usado ningún diccionario de contraseñas descargado de internet, ya que la mayoría de estas contraseñas no cumplían una característica que se nos da en el enunciado, "clave tenía sólo caracteres en minúsculas". Esto descarta muchas de las contraseñas y no quise incluir los diccionarios porque probarían contraseñas que sabemos de antemano que son incorrectas. Se podría haber incluido un filtro para hacer que solo se probasen contraseñas con esta restricción pero me parecía que en tema de rendimiento para el programa era mejor desechar esta idea. 

# 4. Resultados.
Como resultados obtenidos tras la ejecución del código fueron los siguientes:
  - Contraseña encontrada: "jlsl".
  - Tiempo tardado: 14751.13 segundos (Un poco más de 4 horas).
  - Contraseñass probadas: algo más de 344000.
  - Archivo descifrado: posible enunciado práctica 3.
