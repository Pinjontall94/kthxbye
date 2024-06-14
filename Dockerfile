FROM python:3
WORKDIR /usr/src/app
EXPOSE 8000

# Install poetry and deps
# Exporting to requirements.txt will avoid hassle with poetry venv
COPY pyproject.toml ./
COPY poetry.lock ./
RUN pip install poetry
RUN poetry export --with dev -f requirements.txt --output requirements.txt
RUN pip install -r requirements.txt

# Copy project
COPY manage.py manage.py
COPY kthxbye kthxbye
COPY . .

# Collect static files, pass dummy DJANGO_SECRET_KEY
# or else ./manage.py will fail
# since DJANGO_SECRET_KEY is not available during the build step
RUN DJANGO_SECRET_KEY=dummy ./manage.py collectstatic

# Run server
CMD [ "poetry", "run", "gunicorn", "-b", "0.0.0.0:8000", "kthxbye.wsgi" ]
