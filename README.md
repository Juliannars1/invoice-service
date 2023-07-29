# Invoice-Service

Invoice-Service es un proyecto que proporciona un servicio para gestionar facturas utilizando una base de datos MongoDB y una aplicación desarrollada en Golang.

## Requisitos previos

Antes de comenzar, asegúrate de tener instalado Docker y Docker Compose en tu sistema. Puedes descargar Docker desde su sitio web oficial: [https://www.docker.com/get-started](https://www.docker.com/get-started)

## Inicialización y creación de contenedores

Para iniciar el servicio, sigue estos pasos:

1. Clona este repositorio en tu máquina local:
   git clone https://github.com/Juliannars1/invoice-service.git
   _cd invoice-service_

2. Construye y levanta los contenedores utilizando Docker Compose:
   _docker-compose up -d_
   Esto creará y pondrá en funcionamiento los contenedores de MongoDB y Golang definidos en el archivo `docker-compose.yml`. El servicio de Invoice-Service estará ahora accesible.

3. Verifica que los contenedores están en ejecución:
   _docker-compose ps_
   Deberías ver una salida que muestre que tanto el contenedor de MongoDB como el de Golang están "Up".

## Uso del servicio

Accede al servicio desde tu navegador web o utilizando herramientas como Postman:
[http://localhost:8080/invoices](http://localhost:8080/invoices)

Realiza una solicitud HTTP POST para agregar una nueva factura:
**POST** http://localhost:8080/invoices

Body:
{
"Number": "INV001",
"Customer": "Juan",
"Items": [{"Name": "arroz", "Quantity": 3, "Price": 200}],
"TotalAmount": 500000
}
Puedes utilizar herramientas como Postman o incluso realizar la solicitud desde tu propio código.

Nota: Asegúrate de tener el servicio en ejecución (docker-compose up -d) antes de realizar la solicitud POST.

#### Obtener una factura específica (GET)

Realiza una solicitud HTTP GET a la siguiente ruta para obtener una factura específica por su "Number":
**GET** http://localhost:8080/invoices/{Number}
Reemplaza {Number} con el número de la factura que deseas obtener. Por ejemplo, para obtener la factura con el número "INV001":
GET http://localhost:8080/invoices/INV001

####Actualizar una factura (UPDATE)
Realiza una solicitud HTTP PUT a la siguiente ruta para actualizar una factura específica por su "Number":
**PUT** http://localhost:8080/invoices/{Number}

Reemplaza {Number} con el número de la factura que deseas actualizar. Incluye en el cuerpo de la solicitud los campos que deseas modificar. Por ejemplo, para actualizar la factura con el número "INV001":
PUT http://localhost:8080/invoices/INV001
Body:
{
"Number": "INV001",
"Customer": "Pedro",
"Items": [{"Name": "arroz", "Quantity": 5, "Price": 200}],
"TotalAmount": 1000
}
####Eliminar una factura (DELETE)
Realiza una solicitud HTTP DELETE a la siguiente ruta para eliminar una factura específica por su "Number":
**DELETE** http://localhost:8080/invoices/{Number}

Reemplaza {Number} con el número de la factura que deseas eliminar. Por ejemplo, para eliminar la factura con el número "INV001":
DELETE http://localhost:8080/invoices/INV001

##Personalización
Si deseas personalizar la configuración del servicio, puedes modificar el archivo `docker-compose.yml` y ajustar los parámetros según tus necesidades. Además, puedes agregar nuevas funcionalidades y características al servicio Invoice-Service desarrollando en el lenguaje de programación Golang.

##Contribución
Si deseas contribuir a este proyecto, por favor, realiza un fork del repositorio y crea una rama para tu contribución. Luego, envía un pull request cuando estés listo para que revisemos tus cambios.
