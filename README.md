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

El servidor estara corriendo en `localhost:5050`, puedes hacer una prueba haciendo una peticion `GET` al endpoint `http://localhost:5050/api/v1/hello`

Para crear o administrar graficamente la base de datos estara disponible en `localhost:8080` una instancia de phpmyadmin.

## Production

1. Crear una red de docker

```shell
docker network create echo-api-network
```

2. Crear y ejecutar los contenedores de la base de datos y phpmyadmin

```shell
docker-compose up
```

3. Build al Dockerfile

```shell
docker build . -t echo-api-app
```

4. Crear y ejecutar el contenedor

```shell
docker run --name echo-api-rest -p 5050:5050 echo-api-app
```

5. Conectar los contenedores a la red.

```shell
docker network connect echo-api-network echo-api-template-db-server-1
docker network connect echo-api-network echo-api-template-phpmyadmin-1
docker network connect echo-api-network echo-api-rest
```
