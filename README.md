# WASAPhoto

This project was made for the *Web and Software Architecture* course in my bachelor's degree in Computer Science.

## APIs

The APIs are available inside the `doc/api.yaml` file.

## Build

### Backend

To build the backend without embedding the WebUI into the final executable, run

```sh
go build ./cmd/webapi/
```

if you want to embed the WebUI, instead run

```sh
./open-npm.sh
npm run build-embed
exit # exiting out of the npm container
go build -tags webui ./cmd/webapi/
```

and if you want to embed it in release mode, run

```sh
./open-npm.sh
npm run build-prod
exit # exiting out of the npm container
go build -tags webui ./cmd/webapi/
```

### Frontend

To build (and run) the frontend in development mode, run

```sh
./open-npm.sh
npm run dev
```

if you need the release mode, instead run

```sh
./open-npm.sh
npm run preview
```

## Containers

### Backend

To build the backend Docker container, run

```sh
docker build -t wasa-photos-backend:latest -f Dockerfile.backend .
```

and to run the container image, run

```sh
docker run -it --rm -p 3000:3000 wasa-photos-backend:latest
```

### Frontend

To build the frontend Docker container, run

```sh
docker build -t wasa-photos-frontend:latest -f Dockerfile.frontend .
```

and to run the container image, run

```sh
docker run -it --rm -p 8081:80 wasa-photos-frontend:latest
```

