# castle-search-api

これは、日本100名城を検索するAPIになります。
このAPIは[castle-search-api-environment](https://github.com/takenoko-gohan/castle-search-api-environment)で構築された環境で利用されることを前提にしています。

# 利用方法

APIを利用する下記のような形でリクエストします。
パラメーターは１つのみでも大丈夫です。

- keyword：検索時のキーワードを指定します。
- prefecture：絞り込みたい都道府県を指定します。

```sh
curl -XGET "http://localhost:8080/search?keyword=鶴ヶ城&prefecture=福島県"
```
