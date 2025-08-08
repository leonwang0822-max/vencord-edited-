# Vencord - Modified Edition

![](https://img.shields.io/github/package-json/v/Vendicated/Vencord?style=for-the-badge&logo=github&logoColor=d3869b&label=&color=1d2021&labelColor=282828)
[![Custom Badge Server](https://img.shields.io/badge/Custom%20Badge%20Server-Self%20Hosted-brightgreen?style=for-the-badge)](#badge-system-setup)
[![Enhanced Installer](https://img.shields.io/badge/Enhanced%20Installer-Available-blue?style=for-the-badge)](#enhanced-installer)

A modified version of Vencord with self-hosted badge server integration and an enhanced installer UI.

![](https://github.com/user-attachments/assets/3fac98c0-c411-4d2a-97a3-13b7da8687a2)

## 🚀 What's Modified

This version includes several enhancements over the original Vencord:

### Self-Hosted Badge Server
- **Custom Badge Endpoint**: Connect to your own badge server at `localhost:4000/badges.json`
- **Enhanced Badge System**: Additional badge types and custom styling
- **Real-time Updates**: Badges are fetched dynamically from your custom server
- **Easy Setup**: Use the provided badge system repository for quick deployment

### Enhanced Installer UI
- **Modern Discord Theme**: Beautiful Discord-inspired color palette
- **Smooth Animations**: Fade-in/fade-out effects and smooth transitions
- **Enhanced Buttons**: Modern styling with hover effects and shadows
- **Improved Typography**: Better font hierarchy and spacing
- **Glass Morphism Effects**: Subtle background effects for visual appeal
- **Better Error Handling**: Enhanced error modals with improved UX

## 📋 Prerequisites

Before installing, make sure you have:

- **Node.js** (v18 or higher) - [Download here](https://nodejs.org/)
- **pnpm** package manager - Install with `npm install -g pnpm`
- **Discord** application (must be closed during installation)
- **Git** (for building from source)

## 🏷️ Badge System Setup

To set up your own badge server for custom badges:

### Quick Setup

1. **Clone the Badge System Repository**:
   ```bash
   git clone https://github.com/leonwang0822-max/badge-system-for-my-vencord.git
   cd badge-system-for-my-vencord
   ```

2. **Install Dependencies**:
   ```bash
   npm install
   ```

3. **Configure Your Badges**:
   - Edit `badges.json` to add your custom badges
   - Follow the format: `{"userId": [{"badge": "badge_url", "tooltip": "Badge Description"}]}`

4. **Start the Badge Server**:
   ```bash
   npm start
   ```
   The server will run on `http://localhost:4000`

5. **Verify Setup**:
   - Visit `http://localhost:4000/badges.json` to see your badges
   - The Vencord modification will automatically connect to this endpoint

### Custom Configuration

- **Port Configuration**: Edit the port in the badge server if needed
- **Badge Format**: Supports image URLs and custom tooltips
- **User IDs**: Use Discord user IDs to assign badges to specific users
- **Real-time Updates**: Changes to `badges.json` are reflected immediately

## 🔧 Installation Methods

### Method 1: Enhanced Installer (Recommended)

1. **Download the Enhanced Installer**:
   ```bash
   # Clone this repository
   git clone https://github.com/leonwang0822-max/vencord-edited-.git
   cd vencord-edited-
   ```

2. **Build the Enhanced Installer**:
   ```bash
   cd Installer-main
   chmod +x build_enhanced.sh
   ./build_enhanced.sh
   ```

3. **Run the Enhanced Installer**:
   - **macOS**: `./VencordInstallerEnhanced`
   - **Linux**: `./VencordInstallerEnhanced`
   - **Windows**: `VencordInstallerEnhanced.exe`

4. **Follow the GUI**:
   - Select your Discord installation
   - Click "Install" to install Vencord
   - Enjoy the enhanced UI experience!

### Method 2: Build and Install from Source

1. **Clone and Setup**:
   ```bash
   git clone https://github.com/leonwang0822-max/vencord-edited-.git
   cd vencord-edited-
   pnpm install
   ```

2. **Build Vencord**:
   ```bash
   pnpm build
   ```

3. **Install to Discord**:
   ```bash
   node scripts/runInstaller.mjs
   ```

### Method 3: Manual Installation

1. **Build the project** (follow steps 1-2 from Method 2)

2. **Locate your Discord installation**:
   - **Windows**: `%LOCALAPPDATA%/Discord/app-x.x.x/resources/`
   - **macOS**: `/Applications/Discord.app/Contents/Resources/`
   - **Linux**: `/opt/discord/resources/` or `~/.local/share/discord/resources/`

3. **Backup and replace**:
   ```bash
   # Backup original (recommended)
   cp app.asar app.asar.backup

   # Copy built Vencord
   cp /path/to/vencord-edited-/dist/app.asar ./app.asar
   ```

## 🛠️ Development Workflow

For developers who want to modify and test:

```bash
# Install dependencies
pnpm install

# Build and install in one command
pnpm build && node scripts/runInstaller.mjs

# For development with auto-rebuild
pnpm dev
```

## ✨ Features

All original Vencord features plus:

### Original Vencord Features
- Easy to install
- [100+ built-in plugins](https://vencord.dev/plugins)
- Fairly lightweight despite the many inbuilt plugins
- Excellent Browser Support
- Works on any Discord branch (Stable, Canary, PTB)
- Custom CSS and Themes support
- Privacy friendly (blocks Discord analytics)
- Settings sync between devices

### Enhanced Features
- **Custom Badge Integration**: Connected to custom badge server
- **Enhanced Installer UI**: Modern, Discord-themed installer
- **Improved Visual Effects**: Smooth animations and transitions
- **Better Error Handling**: Enhanced error messages and recovery

## 🐛 Troubleshooting

### Common Issues

**"ERR_MODULE_NOT_FOUND" during build**:
```bash
# Reinstall dependencies
rm -rf node_modules pnpm-lock.yaml
pnpm install
pnpm build
```

**"Permission denied" on installer**:
```bash
# Make installer executable (macOS/Linux)
chmod +x VencordInstallerEnhanced
```

**Discord won't start after installation**:
1. Close Discord completely
2. Restore backup: `cp app.asar.backup app.asar`
3. Try installation again

**Badge server connection issues**:
- Check internet connection
- Ensure your badge server is running: `curl http://localhost:4000/badges.json`
- Verify the badge server is started (see Badge System Setup section)
- Restart Discord if badges don't appear

**Enhanced installer build fails**:
```bash
# Ensure Go is installed
go version

# Install dependencies
cd Installer-main
go mod tidy
./build_enhanced.sh
```

### Getting Help

If you encounter issues:
1. Check the [original Vencord documentation](https://vencord.dev)
2. Ensure all prerequisites are installed
3. Try the troubleshooting steps above
4. Create an issue in this repository

## 🔄 Updating

To update to the latest version:

```bash
# Pull latest changes
git pull origin main

# Reinstall dependencies and rebuild
pnpm install
pnpm build

# Reinstall to Discord
node scripts/runInstaller.mjs
```

## 🚫 Uninstalling

### Using the Enhanced Installer
1. Run the enhanced installer
2. Click "Uninstall"
3. Restart Discord

### Manual Uninstall
```bash
# Restore backup
cp app.asar.backup app.asar

# Or reinstall Discord
```

## 📜 Credits

- **Based on**: Original Vencord project
- **Enhanced Installer**: Custom UI improvements and modern styling
- **Self-Hosted Badge System**: Integration and custom badge system
- **Modified by**: [leonwang0822-max](https://github.com/leonwang0822-max)

## ⚖️ Disclaimer

Discord is trademark of Discord Inc. and solely mentioned for the sake of descriptivity.
Mention of it does not imply any affiliation with or endorsement by Discord Inc.

This is a modified version of Vencord. Client modifications may be against Discord's Terms of Service. Use at your own discretion.

## 🌟 Support

If you find this modified version helpful:
- ⭐ Star this repository
- 🍴 Fork and contribute
- 🐛 Report issues
- 💡 Suggest improvements

---

**Note**: This is a modified version of Vencord. For the original version, visit [Vencord.dev](https://vencord.dev)
