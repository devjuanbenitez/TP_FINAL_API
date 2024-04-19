# Crear y Activar enviroment
python3 -m venv env
source env/bin/activate

# instalar django y postgrescli
pip install django djangorestframework psycopg2-binary

# sync DB
python manage.py makemigrations
python manage.py migrate

# ejecutar
python manage.py runserver
