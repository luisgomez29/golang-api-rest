# API REST GO (golang)

Simple API REST con Go, Echo (framework), GORM (ORM), PostgreSQL para CRUD de dos tablas, usuario y producto.

## Instalación

1. Clonar el proyecto.

2. Crear archivo `.env` con las siguientes variables:
```
API_PORT =

# DATABASE CONFIG

HOST =
DB_PORT =
DB_NAME =
DB_USER =
DB_PWD =
```

3. Configurar la base de datos.

4. Ejecutar el servidor local:
```bash
# -rt para crear las tablas e insertar los datos de prueba
go run main.go -rt

# Ejecutar servidor sin crear tablas e insertar datos
go run main.go
```
5. Ingresar en el navegador a la siguiente URL (revisar que el puerto sea igual al definido en el archivo `.env`):


    http://localhost:8000/api/v1/users

## Autor

**Luis Guillermo Gómez**  
- [Github](https://github.com/luisgomez29)

`Gracias!.`