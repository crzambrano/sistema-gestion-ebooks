# sistema-gestion-ebooks
Sistema de gestión de libros electrónicos

# Sistema de Gestión de Libros Electrónicos

Este proyecto está desarrollado en Go con enfoque funcional. Permite registrar y consultar libros electrónicos desde consola y guardarlos en un archivo JSON.

## Estructura
- `cmd/`: Punto de entrada
- `internal/book/`: Lógica de libros
- `pkg/storage/`: Persistencia
- `data/`: Archivo JSON

## Instalación local
```bash
git clone https://github.com/TU_USUARIO/sistema-gestion-ebooks.git
cd sistema-gestion-ebooks
go run ./cmd
