---
sidebar_position: 1
---

# Anatomy of a Cosmos SDK Application

:::note Synopsis
This document describes the core parts of a Cosmos SDK application, represented throughout the document as a placeholder application named `app`.
:::

## Node Client

The Daemon, or [Full-Node Client](../core/03-node.md), is the core process of a Cosmos SDK-based blockchain. Participants in the network run this process to initialize their state-machine, connect with other full-nodes, and update their state-machine as new blocks come in.

```javascript

// app/app.go
function main() {
  // Initialize the application
  app := initApp()

  // Run the application
  app.run()
}
```


```