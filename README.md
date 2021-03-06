# FileServer

# Author: Yohan Reyes

# Descripción
Este proyecto consiste en la implementación de un sistema de transferencia de archivos mediante un
protocolo personalizado, este protocolo está basado en TCP. El sistema consiste en 2 programas
un cliente y un servidor, el cliente puede subscribirse para recibir archivos en un canal o puede enviar un archivo a todos los clientes subscritos a un canal.

# Protocolo
- client subcribe

    {Canal} [3 bytes]
- client post

    {Canal} [3 bytes]
    
    {Nombre de Archivo} [256 bytes]
    
    {Data} [* bytes]
- client get

    {Nombre de Archivo} [256 bytes]
    
    {Data}  [* bytes]

# Uso
- Servidor
    El servidor se inicializa corriendo el ejecutable. No son necesarios comandos o banderas adicionales. Luego de ser inicializado, presionar cualquier tecla finalizara la aplicación.
- Cliente
    El cliente consta de 2 comandos, 'subscribe' y "post", estos deben ser llamados mientras el servidor está activo y se usan de la siguiente forma
    * ./client subscribe {canal}
    
    Donde el canal es un número entre 1 y 200
    
    Al recibir archivos, estos serán guardados en el mismo directorio en que se encuentre el cliente
    * ./client post {Nombre de Archivo} {Canal}
    
    El archivo debe estar ubicado en el mismo directorio que el ejecutable del cliente
    El canal debe ser un número entre 1 y 200
