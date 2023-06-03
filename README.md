# Go_Mysql_Docker_Air

ローカル時ホットリロードGoとMysql\
デプロイ時マルチステージ

```bash
# local
docker compose build
docker-compose build --no-cache
docker compose up -d

# deploy
cd go
docker build -t name:latest --target deploy .

# コンテナ一覧
docker container ls -a
docker ps

# コンテナに入る
docker exec -it <コンテナ名 or コンテナID> bash
----or----
docker exec -it <コンテナ名 or コンテナID> sh

# コンテナのログ
docker logs <コンテナ名 or コンテナID>

# コンテナ削除
docker container rm <コンテナID>

# イメージ一覧
docker image ls
docker images

# イメージの削除
docker rm <イメージID or イメージ名>
# コンテナから参照されておらずタグもないイメージを全て削除
docker image prune
# コンテナに関連付けられていないイメージを全て削除
docker image prune -a
```