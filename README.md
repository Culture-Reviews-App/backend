# backend
Backend for the culture reviews app.

## how to run
- type `sudo docker compose up` and server will start running on `0.0.0.0:5000`. 
- You can see the swagger at `0.0.0.0:5000/swagger`.

### important commands
- `swag init` to update the swagger.
- `sudo docker compose build --no-cache` to build without cache.
- `sudo docker exec -it mongodb mongosh` to connect to mongodb running on docker.
- `use culture-reviews` to select the related db on mongosh.
- `db.auth('mongoadmin', 'secret')` to authenticate on mongosh.
- `db.users.find()` to see all users.
- `sudo docker volume rm backend_mongo_data` to remove volume and re init mongodb.
- `sudo docker exec -it redis redis-cli` to connect to redis.

### tools used to create project structure
- https://github.com/create-go-app/fiber-go-template
- https://github.com/create-go-app/cli
- https://github.com/gofiber/swagger

### push to docker hub
- find container ID with `sudo docker container ls`
- commit with new name `sudo docker container commit <id> beratdalsuna/backend-app:latest`
- push to hub `sudo docker image push beratdalsuna/backend-app:latest`

### pull from docker
- latest version can be pulled from registry `sudo docker pull beratdalsuna/backend-app:latest`