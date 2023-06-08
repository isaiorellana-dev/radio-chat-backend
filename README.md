# Echo server template

Este es un template para iniciar a crear una API con echo, mysql, GORM, phpmyadmin y docker.

## Run project

1. Instalar [go](https://go.dev/dl/).

2. Clonar el repositorio.

3. Levantar y correr los contenedores de docker.

```
docker-compose up
```

4. Levantar el servidor

```
go run main.go
```

El servidor estara corriendo en `localhost:2222`, puedes hacer una prueba haciendo una peticion `GET` al endpoint `http://localhost:2222/`

Para crear o administrar graficamente la base de datos estara disponible en `localhost:8080` una instancia de phpmyadmin.
