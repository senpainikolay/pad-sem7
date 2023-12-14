FROM python:3.8

WORKDIR /app

COPY . /app

ENV PORT=$port

RUN pip install -r requirements.txt

EXPOSE $port

CMD ["python", "main.py"]