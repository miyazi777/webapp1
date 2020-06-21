動作環境構築用のGoアプリケーションです

# 開発環境セットアップ

```
cd <project_root>/app
docker-compose up
```

# 本番用イメージのビルド
以下でイメージ作成
```
cd <project_root>/app
docker build -t webapp1 -f Dockerfile.prod .
```

動作確認用コンテナ起動
```
dokcer run --name webapp1 -p 3030:3030 -it -d webapp1
```

動作確認
```
❯ curl localhost:3030/
{"count":42}
```

コンテナ停止
```
docker stop webapp1
```


## 参考
https://qiita.com/takasp/items/c6288d4836e79801bb19
