---
sidebar_position: 1
---
# Anatomy of a Cosmos SDK Application

:::note Synopsis


This document describes the core parts of a Cosmos SDK application, represented throughout the document as a placeholder application named `app`.


:::

## Node Client
デーモン、または [フルノードクライアント](../core/03-node.md) は、Cosmos SDKベースのブロックチェーンのコアプロセスです。ネットワークの参加者はこのプロセスを実行して、自分のステートマシンを初期化し、他のフルノードと接続し、新しいブロックが入ってくると自分のステートマシンを更新します。

```javascript
// app/app.go
関数 main() {
// アプリケーションの初期化
app := initApp()

// アプリケーションの実行
app.run()
}

```

