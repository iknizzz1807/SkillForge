FROM python:3.11.2-slim-bullseye

WORKDIR /app


COPY requirements.txt ./
RUN pip install --no-cache-dir --upgrade pip
RUN pip install --no-cache-dir torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cpu
RUN pip install --no-cache-dir -r ./requirements.txt

COPY . .

CMD ["fastapi", "run", "main.py", "--port", "5000"]