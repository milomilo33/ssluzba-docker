FROM python:3.10-alpine

WORKDIR /app

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

EXPOSE 8000

RUN chmod +x /app/skripta.sh

ENTRYPOINT [ "/app/skripta.sh" ]