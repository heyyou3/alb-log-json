# alb-log-json

## What

- S3 に保存されている ALB のログファイル(gz)を展開、結合し、正規表現で属性を抽出し、JSON 文字列として出力する CLI ツール
- 出力される JSON KEY 一覧
    - 詳細は[公式ドキュメント](https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/application/load-balancer-access-logs.html)を参照

| KEY名                     |
|--------------------------|
| type                     |
| timestamp                |
| elb                      |
| client                   |
| target                   |
| request_processing_time  |
| target_processing_time   |
| response_processing_time |
| elb_status_code          |
| target_status_code       |
| received_bytes           |
| sent_bytes               |
| method                   |
| http_version             |
| protocol                 |
| host                     |
| port                     |
| uri                      |
| user_agent               |
| ssl_cipher               |
| ssl_protocol             |
| target_group_arn         |
| trace_id                 |
| domain_name              |
| chosen_cert_arn          |
| matched_rule_priority    |
| request_creation_time    |
| actions_executed         |
| redirect_url             |
| error_reason             |
| target:port_list         |
| target_status_code_list  |

## How

1. Release ページからバイナリをダウンロード

- Mac(arm64)

```shell
curl -sLO https://github.com/heyyou3/alb-log-json/releases/download/v0.0.1-beta-darwin-arm64/alblogjson && chmod +x ./alblogjson
```

- Mac(amd64)

```shell
curl -sLO https://github.com/heyyou3/alb-log-json/releases/download/v0.0.1-beta-darwin-amd64/alblogjson && chmod +x ./alblogjson
```

- Linux(amd64)
- 
```shell
curl -sLO https://github.com/heyyou3/alb-log-json/releases/download/v0.0.1-beta-linux-amd64/alblogjson && chmod +x ./alblogjson
```

2. [./example/alblogjson-config.json](./example/alblogjson-config.toml)を `alblogjson` コマンドと同階層に配置し、パラメータを入力

パラメータ説明

| パラメータ名           | 説明                                     |
|------------------|----------------------------------------|
| aws_profile_name | AWS プロファイル名                            |
| bucket_name      | ALBログが保存されているバケット名                     |
| s3_key           | ALBログが保存されているオブジェクトパス(日のディレクトリまでを指定する) |

3. 実行

標準出力に JSON 文字列が出力される

```shell
./alblogjson
```

リダイレクトし、JSON 文字列をファイル化

```shell
./alblogjson > alblogYYYYMMDD.json
```

## Tips(Columnq)

[Columnq](https://github.com/roapi/roapi/blob/main/columnq-cli/README.md)と組み合わせて使用すると JSON ファイルに対して SQL を実行、結果を出力させることができ、便利です。

ローカル環境を汚さないため、alblogjson と同様、GitHub Release からバイナリをダウンロードすることをおすすめします。

使用例:

```shell
# ./alblogjson > alblogYYYYMMDD.json
# ホスト毎のリクエスト数をカウントし、降順、上位10件を表示
# ./columnq sql --table alblogYYYYMMDD.json "SELECT host, COUNT(*) as cnt FROM alblogYYYYMMDD GROUP BY host ORDER BY cnt DESC LIMIT 10"
```
