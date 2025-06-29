# Sistema de Gestión de Libros Electrónicos

Proyecto integrador – versión 2025-06-29

## Introducción
Aplicación REST desarrollada en Go 1.22 para registrar, listar, buscar y eliminar libros electrónicos, con autenticación basada en JWT y persistencia JSON.

## Estructura
```
sistema-gestion-ebooks/
├── cmd/server/        # entrypoint
├── internal/          # dominio + adaptadores
├── pkg/storage/       # utilidades de persistencia
├── data/              # archivos JSON
└── tests/             # pruebas unitarias
```

## Endpoints (8)
| Método | Ruta                 | Descripción                      |
|--------|----------------------|----------------------------------|
| POST   | /users/register      | Registro de usuario             |
| POST   | /users/login         | Login y obtención de JWT        |
| GET    | /books               | Listar libros (auth)            |
| POST   | /books               | Crear libro (auth)              |
| GET    | /books/{id}        | Obtener libro por ID (auth)     |
| DELETE | /books/{id}        | Eliminar libro (auth)           |
| GET    | /books/search?title= | Buscar libros por título (auth) |
| GET    | /health              | Comprobar salud (sin auth)      |

## Ejecución
```bash
go mod tidy
go run ./cmd/server
```

## Visualización del futuro
En la fase siguiente se migrará la persistencia a PostgreSQL, se containerizará con Docker y se desplegará en Kubernetes con CI/CD GitHub Actions.
