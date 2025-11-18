# TCP Chat (NetCat Clone)

A lightweight TCP-based group chat inspired by the NetCat (`nc`) utility.  
This project recreates core NetCat behavior with a server–client architecture that supports multiple simultaneous users.

---

## Features

- TCP server with **1 → many** connections  
- Maximum **10 clients**
- Clients must provide a **non-empty username**
- Messages include **timestamp, username, and content**
- **Empty messages are ignored**
- New clients receive the **full chat history**
- Join/leave notifications for all clients
- Messages are broadcast to all connected clients
- Server keeps running even if clients disconnect
- Default port: **8989**
- Wrong input shows usage message

---

## Requirements

- Written in **Go**
- Uses **goroutines**
- Uses **channels or mutexes**
- Clean code structure
- Error handling on server and client
- Recommended: unit tests

**Allowed packages:**  
`io`, `log`, `os`, `fmt`, `net`, `sync`, `time`, `bufio`, `errors`, `strings`, `reflect`

---

## Usage

### Start Server

Default port:

```sh
go run .
