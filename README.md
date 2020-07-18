#笔趣阁小说爬取

swag init 

http://localhost:8086/swagger/index.html

获取单本小说生成text

curl -X GET "http://localhost:8086/api/v1/novels?novelUrl=https%3A%2F%2Fwww.biquge.com.cn%2Fbook%2F44060%2F" -H "accept: application/json"