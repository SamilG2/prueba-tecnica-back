# Prueba Técnica: CRUD Básico Ejercicio 1 Go

Este proyecto desarrolla un servicio backend RESTful utilizando **Go (Golang)** en el cual implementa un sistema de gestión de libros con autores.

## 📝 Resumen del Proyecto

* **API RESTful Completa:** Implementación de endpoints para crear, leer, actualizar y eliminar (CRUD) libros.
* **Relaciones de Base de Datos:** Manejo de relaciones **Uno a Muchos (1:N)** entre Autores y Libros.
* **Persistencia de Datos:** Conexión robusta a **MySQL** utilizando **GORM** como ORM.
* **Seed de Datos Automático:** El sistema crea con autores y libros de prueba automáticamente en la base de datos.

## 🛠 Tecnologías y Frameworks

| Tecnología | Propósito | Descripción |
| :--- | :--- | :--- |
| **Go (Golang)** | Lenguaje Core | Lenguaje principal del backend. |
| **Gin Gonic** | Web Framework | Framework HTTP para el enrutamiento y manejo de requests. |
| **GORM** | ORM | Librería para el mapeo objeto-relacional y manejo de SQL. |
| **MySQL** | Base de Datos | Motor de base de datos relacional. |
| **Gofakeit** | Testing/Seeding | Generación de datos aleatorios realistas para pruebas. |
| **Godotenv** | Configuración | Carga de variables de entorno desde archivos `.env`. |

## 📂 Estructura del Proyecto

La arquitectura sigue un diseño modular dentro de la carpeta `internal` para proteger la lógica de negocio:

```text
PRUEBA-TECNICA-BACK/
├── internal/
│   ├── models/           # Definición de Entidades
│   │   ├── autor.go      # Modelo de Autor
│   │   └── libro.go      # Modelo de Libro
│   ├── service/          # Lógica de Negocio
│   │   └── libro_service.go  
│   └── transport/        # Capa HTTP (Handlers)
│       └── libro_handler.go  
├── .env                  # Variables de entorno
├── go.sum           
├── go.mod                
├── main.go               # Punto de entrada
└── README.md            
```
---

## 🚀 Instrucciones de Uso

Sigue estos pasos para desplegar el entorno de desarrollo local.

### 📋 Requisitos Previos
* **Go:** 1.25.5.
* **MySQL:** Instancia local en el puerto 3306.
* **Git:** Para clonar el repositorio.
* **Cliente API:** Postman o de su preferencia.

### 👣 Pasos para Ejecutar

1.  **Clonar el repositorio**
    ```bash
    git clone https://github.com/SamilG2/prueba-tecnica-back.git
    cd prueba-tecnica-back
    ```

2.  **Configurar Variables de Entorno**
    ```properties
    MYSQL_USER=root
    MYSQL_PASSWORD=password
    MYSQL_HOST=localhost
    MYSQL_PORT=3306
    MYSQL_DATABASE=prueba_tecnica_bd
    ```

3.  **Instalar Dependencias**
    ```bash
    go mod tidy
    ```

4.  **Iniciar el Servidor**
    ```bash
    go run main.go
    ```
    > **Nota:** Al iniciar por primera vez, verás el mensaje: *"Base de datos sembrada con datos realistas"*. Esto indica que se han creado autores y libros de prueba automáticamente.

### 🧪 Probar Endpoints (Postman / cURL)

La API corre por defecto en `http://localhost:8080`.

| Método | Endpoint       | Descripción                           | 
| :----- | :------------- | :-----------------------------------  |  
| GET    | `/libros`      | Listar todos los libros (con Autor)   | 
| GET    | `/libros/:id`  | Obtener detalle de un libro por su ID | 
| POST   | `/libros`      | Crear un nuevo libro                  | 
| PUT    | `/libros/:id`  | Actualizar un libro existente         | 
| DELETE | `/libros/:id`  | Eliminar un libro                     | 