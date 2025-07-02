<div align="center">
  <h1>🌟 Altalune</h1>
  <p><em>Navigate your epics among the stars</em></p>

  ![Go Version](https://img.shields.io/badge/go-%3E%3D1.24.3-00ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
  ![Vue.js](https://img.shields.io/badge/vue.js-%2335495e.svg?style=for-the-badge&logo=vuedotjs&logoColor=%234FC08D)
  ![License](https://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge)
  ![Build Status](https://img.shields.io/github/actions/workflow/status/Fuabioo/altalune/release.yml?style=for-the-badge)
</div>

## ✨ What is Altalune?

Altalune is a beautiful, modern Jira epic management tool that transforms how you visualize and analyze your project epics. With stunning data visualizations powered by D3.js and an intuitive Vue.js interface, Altalune brings clarity to complex project hierarchies and helps teams navigate their epics among the stars.

### 🚀 Key Features

- **📊 Epic Analytics** - Deep insights into epic progress, completion rates, and team performance
- **🎨 Beautiful Visualizations** - Interactive charts and graphs powered by D3.js
- **🌐 Modern Web Interface** - Clean, responsive UI built with Vue.js 3 and Pinia
- **🔧 Easy Configuration** - Simple setup with environment variables or config files
- **📈 Real-time Progress Tracking** - Live updates on epic and story completion
- **🎯 Epic Management** - Save, organize, and monitor your most important epics
- **🔒 Secure Integration** - Direct connection to Jira Cloud API with personal access tokens
- **⚡ Static Binary** - Single executable with embedded frontend assets

## 📦 Installation

### 🍺 Homebrew (macOS/Linux)

```bash
# Add the tap
brew tap fuabioo/homebrew-tap

# Install altalune
brew install altalune

# Upgrade to latest version
brew upgrade altalune
```

### 📦 APT (Linux Debian/Ubuntu)

```bash
# Add the repository
head to apt.fuabioo.com

# Install altalune
sudo apt update
sudo apt install altalune

# Upgrade to latest version
sudo apt upgrade altalune
```

### 📥 Download Pre-built Binaries

Visit the [releases page](https://github.com/Fuabioo/altalune/releases) and download the appropriate binary for your platform:

```bash
# Linux (x64)
wget https://github.com/Fuabioo/altalune/releases/latest/download/altalune_Linux_amd64.tar.gz
tar -xzf altalune_Linux_amd64.tar.gz
sudo mv altalune /usr/local/bin/

# macOS (Apple Silicon)
wget https://github.com/Fuabioo/altalune/releases/latest/download/altalune_Darwin_arm64.tar.gz
tar -xzf altalune_Darwin_arm64.tar.gz
sudo mv altalune /usr/local/bin/

# Windows (PowerShell)
Invoke-WebRequest -Uri "https://github.com/Fuabioo/altalune/releases/latest/download/altalune_Windows_amd64.zip" -OutFile "altalune.zip"
Expand-Archive -Path "altalune.zip" -DestinationPath "."
```

### 🔨 Build from Source

```bash
git clone https://github.com/Fuabioo/altalune.git
cd altalune
just build  # or: go build -o bin/altalune .
```

## 🎯 Quick Start

### 1. Get Your Jira API Credentials

You'll need these details from your Jira instance:

- **Jira Host**: Your Atlassian domain (e.g., `company.atlassian.net`)
- **Email**: Your Jira account email
- **API Token**: Generate one at [Atlassian Account Settings](https://id.atlassian.com/manage-profile/security/api-tokens)
- **Workspace**: Your project key or workspace identifier

### 2. Configure Altalune

Create a `.env` file or set environment variables:

```bash
# .env file
JIRA_EPIC_HOST=your-domain.atlassian.net
JIRA_EPIC_EMAIL=your-email@example.com
JIRA_EPIC_TOKEN=your-jira-api-token
JIRA_EPIC_WORKSPACE=your-workspace-id
```

### 3. Start Altalune

```bash
altalune --verbose
```

The server will start on `http://localhost:3002` by default.

### 4. Access the Web Interface

Open your browser and navigate to `http://localhost:3002` to start managing your epics!

## 🖥️ Web Interface Features

### 🏠 **Home Dashboard**
- Quick overview of saved epics
- Epic statistics and progress tracking
- Beautiful data visualizations

### 📋 **Epic Management**
- Save and organize your most important epics
- View detailed epic information and progress
- Interactive epic navigation

### ⚙️ **Admin Setup**
- Configure Jira connection settings
- Manage workspace and project settings
- Test API connectivity

### 📊 **Epic Analytics**
- Detailed progress charts and graphs
- Completion rate analysis
- Team performance insights
- Interactive D3.js visualizations

## 🛠️ Configuration

### Command Line Options

```bash
altalune [flags]

Flags:
      --email string          Jira email address
  -h, --help                  Help for altalune
      --host string           Jira host (e.g., company.atlassian.net)
      --server-host string    Server host (default "0.0.0.0")
      --server-port int       Server port (default 3002)
      --super-debug          Enable super debug logging
      --token string          Jira API token
      --verbose              Enable verbose logging
  -v, --version              Version information

Commands:
      completion              Generate shell completion scripts
      help                    Help about any command
```

### Configuration File

Create a `config.yaml` file in one of these locations:
- Current directory (`.`)
- Home config directory (`$HOME/.config/`)
- System config directory (`/etc/`)

```yaml
host: "your-domain.atlassian.net"
email: "your-email@example.com"
token: "your-jira-api-token"
workspace: "your-workspace-id"
server-host: "0.0.0.0"
server-port: 3002
verbose: false
super-debug: false
```

### Environment Variables

All configuration options can be set via environment variables with the `JIRA_EPIC_` prefix:

```bash
export JIRA_EPIC_HOST="your-domain.atlassian.net"
export JIRA_EPIC_EMAIL="your-email@example.com"
export JIRA_EPIC_TOKEN="your-jira-api-token"
export JIRA_EPIC_WORKSPACE="your-workspace-id"
export JIRA_EPIC_SERVER_HOST="0.0.0.0"
export JIRA_EPIC_SERVER_PORT="3002"
export JIRA_EPIC_VERBOSE="false"
export JIRA_EPIC_SUPER_DEBUG="false"
```

## 🖥️ Development

### Prerequisites

- **Go** 1.24.3 or higher
- **Bun** (for frontend development)
- **Just** (task runner) - `cargo install just`

### Development Setup

```bash
# Clone the repository
git clone https://github.com/Fuabioo/altalune.git
cd altalune

# Install dependencies
just deps

# Run in development mode (backend + frontend)
just run

# Or run components separately:
just run-be  # Backend only
just run-fe  # Frontend only
```

### Available Just Commands

```bash
just build        # Build everything (frontend + backend)
just build-fe     # Build frontend only
just run          # Run both frontend and backend
just run-be       # Run backend server only
just run-fe       # Run frontend dev server only
just snapshot     # Generate GoReleaser snapshot
just deps         # Download all dependencies
just test-go      # Run Go tests
just clean        # Clean build artifacts
just help-cli     # Show CLI help
just help         # Show all available commands
```

### Tech Stack

- **Backend**: Go with Cobra CLI framework
- **Frontend**: Vue.js 3 with Composition API
- **State Management**: Pinia
- **Routing**: Vue Router 4
- **Styling**: SCSS with custom CSS variables
- **Charts**: D3.js for data visualizations
- **Build Tools**: Vite for frontend, Go build for backend
- **Bundling**: Embedded static assets using Go embed

## 🏗️ Architecture

Altalune follows a modern full-stack architecture:

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Vue.js 3 SPA  │───▶│  Go HTTP Server  │───▶│   Jira Cloud    │
│                 │    │                  │    │      API        │
│ • Pinia Store   │    │ • Cobra CLI      │    │                 │
│ • Vue Router    │    │ • Embedded Assets│    │ • REST API      │
│ • D3.js Charts  │    │ • Static Binary  │    │ • Authentication│
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

### Key Components

- **Single Binary Deployment**: Frontend assets are embedded into the Go binary
- **RESTful API**: Clean API design for frontend-backend communication
- **Real-time Updates**: Live data synchronization with Jira
- **Responsive Design**: Mobile-friendly interface
- **Static Linking**: No external dependencies required

## 🤝 Contributing

We welcome contributions! Here's how to get started:

1. **Fork** the repository
2. **Create** a feature branch: `git checkout -b feature/amazing-feature`
3. **Make** your changes and add tests
4. **Test** your changes: `just test-go`
5. **Build** to ensure everything works: `just build`
6. **Commit** your changes: `git commit -m 'feat: add amazing feature'`
7. **Push** to the branch: `git push origin feature/amazing-feature`
8. **Open** a Pull Request

### Development Guidelines

- Follow Go best practices and use `gofmt`
- Use conventional commit messages
- Write tests for new functionality
- Update documentation for user-facing changes
- Ensure the frontend builds without errors
- Test on multiple platforms when possible

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Charm](https://charm.sh/) for beautiful CLI libraries (Cobra, Lipgloss, Log)
- [Vue.js](https://vuejs.org/) for the reactive frontend framework
- [D3.js](https://d3js.org/) for powerful data visualizations
- [Vite](https://vitejs.dev/) for lightning-fast development builds
- [GoReleaser](https://goreleaser.com/) for automated releases
- [Pinia](https://pinia.vuejs.org/) for intuitive state management

## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=Fuabioo/altalune&type=Date)](https://star-history.com/#Fuabioo/altalune&Date)

---

<div align="center">
  <p><strong>Made with ❤️ and ☕</strong></p>
  <p><em>Navigate your epics among the stars</em> ⭐</p>
</div>
