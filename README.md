# SO1_P1_201901849

## Fernando Mauricio Gómez Santos
## 201901849

Segunda práctica del curso de Sistemas Operativos 1. El sistema consiste en una aplicación que realiza despliega un dashboard con la información del uso de la memoria RAM, CPU y un listado de los procesos que se están llevando a cabo. La información es obtenida a través de los comandos 

## Antes de iniciar

Se recomienda que el repositorio sea usado únicamente para estudiar el código que compone cada parte del sistema. Si el lector desea desplegar la aplicación, entonces Docker es una mejor opción. En __Docker__ se ampliará la información.

## Tecnologías usadas

* React
* JS
* Node 16
* Go 1.19
* MySQL

## Frontend
El frontend fue desarrollado usando JS y React. Para los estilos se utilizó el framework Bootstrap. Los componentes del frontend son:

  * __Usages:__ Contiene dos gráficas de pie que muestran los porcentajes utilizados de la memoria RAM y de CPU.
  * __Process:__ Despliega una tabla que contiene a todos procesos registrados. La información que es mostrada de cada proceso es el PID, nombre, propietario y estado.

## Backend
Para el backend se utiliza un servidor creado con Go el cual se encuentra conectado a la base de datos de MongoDB. Los principales paquetes a utilizar son fiber para crear el servidor y los drivers de sql y MySQL.

## Base de Datos
Se utiliza una base de datos de MySQL en la nube de Google Cloud. Se registran las medidas de RAM y CPU junto con su fecha.
## Docker
Frontend: fernandomgs/frontend_p2
Backend: fernandomgs/backend_p2_test
