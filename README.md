# API REST GO (golang)

Simple API REST con Go, Echo (framework), GORM (ORM), PostgreSQL para CRUD de dos tablas, usuario y producto.

## Instalación

1. Clonar el proyecto.

2. Crear archivo **.env** con las siguientes variables:
```
API_PORT =

# DATABASE CONFIG

HOST =
DB_NAME =
DB_USER =
DB_PWD =
```

3. Configurar la base de datos.

4. Ejecutar el servidor local:
```
// -rt para crear las tablas e insertar los datos de prueba
go run server.go -rt

// Ejecutar servidor sin crear tablas e insertar datos
go run server.go
```

## Autor

**Luis Guillermo Gómez**  
- [Github](https://github.com/luisgomez29)

```
Gracias!.
```