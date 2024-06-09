FROM python:3.7-alpine

WORKDIR /build
COPY requirements.txt .
RUN pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple/
COPY . .

EXPOSE 5001
CMD ["python", "app.py"]
