<div align="center">
  <img src="https://capsule-render.vercel.app/api?type=soft&color=0:BEE3F8,50:00ADD8,100:2B6CB0&height=200&section=header&text=Task%20Tracker%20CLI&fontSize=70&fontAlignY=40&fontColor=FFFFFF&desc=A%20Lightweight%20Go-Powered%20Task%20Manager&descAlignY=65&descSize=20&fontUrl=https://fonts.googleapis.com/css2?family=Outfit:wght@700" alt="Task Tracker Banner" width="100%" />

  <a href="https://git.io/typing-svg">
    <img src="https://readme-typing-svg.demolab.com?font=Outfit&weight=500&size=24&pause=1000&color=00ADD8&center=true&vCenter=true&width=800&lines=Robust+Command-Line+Interface;Zero-Dependency+Standard+Library;Instant+JSON+Persistence;Modular+Go+Architecture" alt="Typing Animation" />
  </a>

<br>
  <p>
    <a href="#"><img src="https://img.shields.io/badge/Language-Go%201.21+-00ADD8?style=social&logo=go" alt="Go" /></a>&nbsp;&nbsp;
    <a href="#"><img src="https://img.shields.io/badge/License-MIT-4caf50?style=social&logo=opensourceinitiative" alt="License" /></a>
  </p>

</div>

<br>

## 📋 Overview

**Task Tracker CLI** is a professional terminal-based utility developed in **Go**. It provides a high-performance, modular environment for managing tasks with full persistence and status tracking.

Designed with **Clean Architecture** principles, the project demonstrates a robust separation of concerns between CLI commands, business logic, and data storage—all while using only the Go standard library.

<br>

## ✨ Core Features

<table border="0">
  <tr>
    <td width="50%" valign="top">
      <h3>🚀 Faster Workflow</h3>
      <p>Compiled Go binary for near-zero latency command execution.</p>
    </td>
    <td width="50%" valign="top">
      <h3>📁 Modular Design</h3>
      <p>Clean separation between <code>cmd</code>, <code>services</code>, and <code>storage</code> layers.</p>
    </td>
  </tr>
  <tr>
    <td width="50%" valign="top">
      <h3>🔍 Status-Based Listing</h3>
      <p>Filter tasks instantly by their status: <code>todo</code>, <code>in-progress</code>, or <code>done</code>.</p>
    </td>
    <td width="50%" valign="top">
      <h3>💾 Robust Persistence</h3>
      <p>Structured JSON storage with automatic ID management and timestamping.</p>
    </td>
  </tr>
</table>

<br>

## 🧰 Technical Stack

<div align="center">
  <br>
  <img src="https://skillicons.dev/icons?i=go,git,github&theme=light" height="50" alt="Tech Stack" />
  
  <br><br>

  | Component | Implementation |
  | :--- | :--- |
  | **Core Language** | Go 1.21+ (Standard Library) |
  | **Architecture** | Feature-First / Modular |
  | **CLI Parsing** | Native <code>os.Args</code> implementation |
  | **Data Format** | Validated JSON Persistence |

</div>

<br>

## 📂 Project Architecture

A clean, modular structure ensures maintainability and scalability.

```bash
.
├── cmd/               # ⌨️ CLI Command Handlers
├── models/            # 📦 Data Structures (Task, TaskData)
├── services/          # 🧠 Business Logic (CRUD Operations)
├── storage/           # 💾 Data Persistence Layer
├── main.go            # 🏁 Application Entry Point
└── tasks.json         # 🗄️ Local Data Storage
```

<br>

## 🚀 Installation & Usage

1.  **Clone** and Navigate:
    ```bash
    git clone https://github.com/mzaid-dev/task-tracker-cli.git
    cd task-tracker-cli
    ```
2.  **Usage Syntax**:
    ```bash
    go run main.go [command] [args]
    ```
3.  **Command Reference**:
    | Command | Example | Description |
    | :--- | :--- | :--- |
    | `add` | `add "Task Description"` | Create a new task |
    | `list` | `list [status]` | View all or filtered tasks |
    | `update` | `update <id> <status>` | Change task status |
    | `delete` | `delete <id>` | Remove task from storage |

<br>

---

<div align="center">
  <br>
  
  <h3>👨‍💻 Crafted with ❤️ by Muhammad Zaid</h3>
  
  <p>
    <a href="https://github.com/mzaid-dev"><img src="https://img.shields.io/badge/GitHub-Message_Me-181717?style=social&logo=github" alt="GitHub"></a>&nbsp;&nbsp;
    <a href="https://linkedin.com/in/mzaid-dev"><img src="https://img.shields.io/badge/LinkedIn-Connect-0077B5?style=social&logo=linkedin&logoColor=0077B5" alt="LinkedIn"></a>&nbsp;&nbsp;
    <a href="mailto:dev.mzaid@gmail.com"><img src="https://img.shields.io/badge/Email-Contact_Me-D14836?style=social&logo=gmail" alt="Email"></a>
  </p>

</div>
