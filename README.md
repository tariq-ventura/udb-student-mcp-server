# UDB MCP Server

Un servidor [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) escrito en Go diseñado para interactuar con plataformas Moodle y sistemas de reserva de bibliotecas (MRBS/Koha). Este servidor expone herramientas para que los agentes LLM puedan consultar cursos, listar recursos, descargar documentos PDF directamente desde Moodle y realizar reservas de cubículos.

## Características (Herramientas MCP)

El servidor expone las siguientes herramientas:

*   **`get_my_courses`**: Obtiene la lista de cursos en los que el estudiante está inscrito.
*   **`get_course_pdfs`**: Obtiene la lista de PDFs disponibles dentro de un curso específico.
*   **`download_moodle_pdf`**: Descarga un PDF desde Moodle hacia el entorno local (carpeta `downloads/`) y devuelve la ruta absoluta para que el agente pueda leerlo con sus herramientas nativas.
*   **`reserve_library_room`**: Reserva un cubículo en la biblioteca (sistema MRBS) especificando fecha, hora de inicio, hora de fin y el ID del cubículo.

## Requisitos Previos

*   [Go](https://golang.org/) 1.24 o superior.

## Configuración

El proyecto requiere variables de entorno para su configuración y autenticación tanto de Moodle como de la Biblioteca. Puedes definir estas variables en tu entorno o usar un archivo `.env` basándote en `.env.example`.

Debes configurar las siguientes variables:

```sh
# Configuración general
MCP_TYPE=mcp-go

# Configuración de Moodle
AUTH_TYPE=moodle
MOODLE_BASE_URL=https://tu-plataforma-moodle.com
MOODLE_USERNAME=tu_usuario
MOODLE_PASSWORD=tu_contraseña

# Configuración de Biblioteca (Koha/MRBS)
BIBLIO_TYPE=koha
KOHA_BASE_URL=https://tu-sistema-koha.com
KOHA_USERNAME=tu_usuario_koha
KOHA_PASSWORD=tu_contraseña_koha
```