# UDB MCP Server

Un servidor [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) escrito en Go diseñado para interactuar con plataformas Moodle. Este servidor expone herramientas para que los agentes LLM puedan consultar cursos, listar recursos y descargar documentos PDF directamente desde Moodle.

## Características (Herramientas MCP)

El servidor expone las siguientes herramientas:

*   **`get_my_courses`**: Obtiene la lista de cursos en los que el estudiante está inscrito.
*   **`get_course_pdfs`**: Obtiene la lista de PDFs disponibles dentro de un curso específico.
*   **`download_moodle_pdf`**: Descarga un PDF desde Moodle hacia el entorno local (carpeta `downloads/`) y devuelve la ruta absoluta para que el agente pueda leerlo con sus herramientas nativas.

## Requisitos Previos

*   [Go](https://golang.org/) 1.24 o superior.

## Configuración

El proyecto requiere variables de entorno para su configuración y autenticación. Puedes definir estas variables en tu entorno o usar un archivo `.env` (que ya está excluido en el [`.gitignore`](.gitignore)).

Debes configurar las siguientes variables:

```sh
# Configuración general
MCP_TYPE=mcp-go
AUTH_TYPE=moodle

# Configuración de Moodle
MOODLE_BASE_URL=https://tu-plataforma-moodle.com
MOODLE_USERNAME=tu_usuario
MOODLE_PASSWORD=tu_contraseña