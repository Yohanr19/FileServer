# FileServer

# Author: Yohan Reyes

# Descripción
Este proyecto consiste en la implementación de un sistema de transferencia de archivos mediante un
protocolo personalizado, este protocolo está basado en TCP. El sistema consiste en 2 programas
un cliente y un servidor, el cliente puede subscribirse para recibir archivos en un canal o puede enviar un archivo a todos los clientes subscritos a un canal.

# Protocolo
- client subcribe
{Canal}
- client post
{Canal}
{Nombre de Archivo}
{Data}
- client get
{Nombre de Archivo}
{Data}

# Uso
- Servidor
    El servidor se inicializa corriendo el ejecutable. No son necesarios comandos o banderas adicionales. Luego de ser inicializado, presionar cualquier tecla finalizara la aplicación.
- Cliente
    El cliente consta de 2 comandos, 'subscribe' y "post", estos deben ser llamados mientras el servidor está activo y se usan de la siguiente forma
    * ./client subscribe {canal} -keep
    Donde el canal es un número entre 1 y 200
    Keep es una bandera, si se utiliza, el cliente restablecerá la conexión luego de recibir un archivo, permitiéndole recibir archivos indeterminadamente, si no se utiliza, el cliente terminará luego de recibir un archivo 
    Al recibir archivos, estos serán guardados en el mismo directorio en que se encuentre el cliente
    * ./client post {Nombre de Archivo} {Canal}
    El nombre de archivo debe ser un archivo que esté ubicado en el mismo directorio que el ejecutable
    El canal debe ser un número entre 1 y 200
