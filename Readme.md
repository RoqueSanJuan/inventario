Clean Architecture

Driven Design


- Primera Capa repositorio (Se comunica con la base de datos) (respository) No puede tener referencia a otra capa superior
    No se pueden importar elementos de las otras capas
    Golang evita este tipo de cosas (Circular Dependency)
    No se puede importar algo que dependa de ti

- Segunda Capa Business Logic (Capa de negocio) (Service)
    No se puede agregar un correo ya utilizado
    Un usuario puede tener 3 amigos agregados


- Presentacion Como mostrar los datos -> Rest API HTTP Service (API)
    Existen otras formas:
        Eventos
        Rabbit MQ 
        Kafka


## Carpetas
- Setting 
    Tiene las configuraciones 
    El puerto de la api
    Configuracion de la base (Creds)
    Guardo las configuraciones en el yaml settings

- Internal en golang tiene propiedades especiales
    Todo lo que sea sensible se recomienda meterlo ahi
    Golang oculta el contenido de internal
    Aca van las capas

Se Recomienda que el archivo tenga el mismo nombre del paquete

Base de datos en mariaDB
Docker para generar la base 

docker run -d --name mariadb --env MARIADB_ROOT_PASSWORD=rootroot  mariadb:10.7.4