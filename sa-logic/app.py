from flask import Flask, request, jsonify
from flask_cors import CORS
from snownlp import SnowNLP

app = Flask(__name__)
CORS(app)  # 在应用上下文中启用CORS

@app.route('/analyse/sentiment', methods=['POST'])
def analyse_sentiment():
    print("接收到 POST 请求")
    data = request.json
    print("请求数据:", data)
    if 'content' not in data:
        print("错误：请求数据中未找到 content 字段")
        return jsonify(error='Content not found'), 400
    s = data['content']
    print("要进行情感分析的内容:", s)
    sentence = SnowNLP(s)
    print("情感分析结果:", sentence.sentiments)
    return jsonify(
        code=200,
        data={
            'sentiments': sentence.sentiments
        }
    )

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)

