# Sistema de Gestión de Libros Electrónicos

Proyecto desarrollado en Go. Permite registrar, listar, buscar y eliminar libros electrónicos mediante una API REST, con autenticación segura basada en JWT.

Proyecto modificado hasta el 2025-06-29

## Introducción
Aplicación REST desarrollada en Go 1.22 para registrar, listar, buscar y eliminar libros electrónicos, con autenticación basada en JWT y persistencia JSON.

## Estructura
```
sistema-gestion-ebooks/
├── cmd/server/           # Servidor HTTP (main.go)
├── internal/
│   ├── book/             # Entidad Book
│   ├── user/             # Entidad User
│   ├── auth/             # Autenticación (JWT)
│   └── handler/          # Handlers y rutas
├── pkg/storage/          # Persistencia JSON
├── data/                 # Archivos de datos (libros.json)
├── tests/                # Pruebas unitarias
└── README.md```

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
En la fase siguiente, de ser posible, se migrará la persistencia a PostgreSQL y se implementarán mejoras para avanzar

Autor
Cristina Zambrano
Proyecto académico – Universidad Internacional del Ecuador
2025
https://github.com/crzambrano/sistema-gestion-ebooks
