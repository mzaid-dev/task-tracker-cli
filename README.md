<div align="center">
  <img src="https://capsule-render.vercel.app/api?type=soft&color=0:BEE3F8,50:00ADD8,100:2B6CB0&height=200&section=header&text=Task%20Tracker%20CLI&fontSize=70&fontAlignY=40&fontColor=FFFFFF&desc=A%20Lightweight%20Go-Powered%20Task%20Manager&descAlignY=65&descSize=20&fontUrl=https://fonts.googleapis.com/css2?family=Outfit:wght@700" alt="Task Tracker Banner" width="100%" />

  <a href="https://git.io/typing-svg">
    <img src="https://readme-typing-svg.demolab.com?font=Outfit&weight=500&size=24&pause=1000&color=00ADD8&center=true&vCenter=true&width=800&lines=Robust+Command-Line+Interface;Zero-Dependency+Standard+Library;Instant+JSON+Persistence;Streamlined+Workflow" alt="Typing Animation" />
  </a>

<br>
  <p>
    <a href="#"><img src="https://img.shields.io/badge/Language-Go%201.21+-00ADD8?style=social&logo=go" alt="Go" /></a>&nbsp;&nbsp;
    <a href="#"><img src="https://img.shields.io/badge/License-MIT-4caf50?style=social&logo=opensourceinitiative" alt="License" /></a>
  </p>

</div>



## 📋 Overview

**Task Tracker CLI** is a professional-grade command-line tool developed in **Go**. It provides a focus-driven environment for managing your daily tasks directly from your terminal, without the bloat of external dependencies or complex databases.

Built with a commitment to **clean architecture** and **minimalism**, it serves as both a powerful utility and a reference for standard library Go development.

<br>

## ✨ Core Features

<table border="0">
  <tr>
    <td width="50%" valign="top">
      <h3>🚀 Fast Execution</h3>
      <p>Compiled Go binary ensures near-instant response times for all CLI interactions.</p>
    </td>
    <td width="50%" valign="top">
      <h3>📁 Precise Logging</h3>
      <p>Automatic timestamping and ID generation for every task you create.</p>
    </td>
  </tr>
  <tr>
    <td width="50%" valign="top">
      <h3>🔍 Smart Filtering</h3>
      <p>List tasks by status (todo, in-progress, done) to keep your focus sharp.</p>
    </td>
    <td width="50%" valign="top">
      <h3>🛠️ Zero Setup</h3>
      <p>Simply run the binary. No database to configure or external libraries to install.</p>
    </td>
  </tr>
</table>


## 🧰 Technical Stack

<div align="center">
  <br>
  <img src="https://skillicons.dev/icons?i=go,git,github,bash,linux&theme=light" height="50" alt="Tech Stack" />
  
  <br>

  | Component | Implementation |
  | :--- | :--- |
  | **Language** | Go 1.21+ (Standard Library) |
  | **Persistence** | Structured JSON Data |
  | **Packaging** | Go Modules |

</div>

<br>

## 📂 Project Structure

```bash
task-tracker-cli/
├── cmd/
│   └── root.go           # CLI handler: parses commands and calls services
├── models/
│   └── task.go           # Task struct and TaskData struct for JSON storage
├── services/
│   └── task_service.go   # Business logic: add, update, delete, mark, list tasks
├── storage/
│   └── file.go           # JSON file handling: load and save tasks
├── go.mod                # Go module definition and dependencies
├── main.go               # Entry point: executes CLI handler
└── README.md             # Project documentation and usage instructions

```

<br>

## 🚀 Installation & Usage

1.  **Clone** and Navigate:
    ```bash
    git clone https://github.com/mzaid-dev/task-tracker-cli.git
    cd task-tracker-cli
    ```
2.  **Quick Run**:
    ```bash
    # Add a task
    go run main.go add "Finish the Go project"

    # List all tasks
    go run main.go list

    # Update status
    go run main.go update 1 "done"
    ```
3.  **Command Reference**:
    - `add` - Create new task entry
    - `list` - View all tasks by status
    - `update` - Modify description/status
    - `delete` - Remove task from storage

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
