# Usar una imagen base de Python 3.10
FROM python:3.10

# Establecer un directorio de trabajo dentro de la imagen
WORKDIR /app

# Copiar el archivo de requisitos al directorio de trabajo
COPY requirements.txt .

# Instalar dependencias de Python
RUN pip install --no-cache-dir -r requirements.txt

# Copiar el código de la aplicación al directorio de trabajo
COPY . .

# Exponer el puerto 8000
EXPOSE 8000

# Establecer la variable de entorno para indicar el entorno de producción
ENV DJANGO_SETTINGS_MODULE=apiPython.settings

# Comando para ejecutar el servidor Django usando Gunicorn
CMD ["gunicorn", "--bind", "0.0.0.0:8000", "apiPython.wsgi:application"]
