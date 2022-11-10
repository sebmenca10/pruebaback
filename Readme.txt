Iniciar en la terminal el proyecto con el comando --> go run main.go
en otra terminal ejecutar el endpoint --> curl --request POST 127.0.1:8080/tickets
luego con el siguiente comando se crean los tickets en formato JSON segun los campos creados --> --data-raw '{"id":1,"usuario":"sebas","fcreacion":"1/11/2022","factualizacion:"1/11/2022","estatus":"abierto"}'
luego para verificar ejecutar el endpoint --> curl --request GET 127.0.1:8080/tickets