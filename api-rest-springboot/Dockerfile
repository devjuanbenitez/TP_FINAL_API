# Usa la imagen de OpenJDK 17 como base
FROM openjdk:17-alpine

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia el archivo JAR generado por Maven al directorio de trabajo
COPY target/inventario-0.0.1-SNAPSHOT.jar .

# Expone el puerto 8080
EXPOSE 8080

# Ejecuta la aplicación Spring Boot cuando se inicie el contenedor
CMD ["java", "-jar", "inventario-0.0.1-SNAPSHOT.jar"]