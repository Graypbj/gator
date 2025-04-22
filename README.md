# Gator CLI

Gator is a command-line tool built in Go that connects to a PostgreSQL database and stores user preferences locally via a config file.

---

## 🔧 Requirements

Before using Gator, ensure you have the following installed:

- **Go** (v1.20 or later)  
  [Install Go](https://golang.org/dl/)

- **PostgreSQL**  
  [Install PostgreSQL](https://www.postgresql.org/download/)

---

## 🚀 Installing Gator CLI

You can install the Gator CLI using `go install`:

```bash
go install github.com/Graypbj/gator@latest
```

This will place the `gator` binary in your Go bin directory. Ensure your `$GOPATH/bin` or `$HOME/go/bin` is added to your `PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

---

## ⚙️ Configuration File Setup

Gator uses a local config file to store settings like the database URL and the current user. The config file is automatically read from your home directory:

```
~/.gatorconfig.json
```

### 🛠 Structure

The config file follows this structure:

```json
{
  "db_url": "postgres://username:password@localhost:5432/your_database",
  "current_user_name": "your_username"
}
```

Replace the `db_url` with your actual PostgreSQL connection string and set `current_user_name` as desired.

---

## 🧠 How It Works (Go Implementation)

Gator handles reading and writing this file using the following logic (from `config.go`):

```go
const configFileName = ".gatorconfig.json"

type Config struct {
    DBURL           string `json:"db_url"`
    CurrentUserName string `json:"current_user_name"`
}

// Saves the current user
func (cfg *Config) SetUser(userName string) error {
    cfg.CurrentUserName = userName
    return write(*cfg)
}

// Reads the config from ~/.gatorconfig.json
func Read() (Config, error) {
    fullPath, err := getConfigFilePath()
    ...
}

// Writes the config to ~/.gatorconfig.json
func write(cfg Config) error {
    ...
}
```

The file is stored at the root of your home directory and automatically managed by Gator.

---

## 🧪 Quick Start

1. **Create the config file** (if not already created by Gator):

```bash
touch ~/.gatorconfig.json
```

2. **Edit it with your database URL and username**:

```json
{
  "db_url": "postgres://postgres:password@localhost:5432/mydb",
  "current_user_name": "grayson"
}
```

3. **Run Gator!**  
Now you're ready to start using the CLI.

---

## 📄 License

MIT License
```

