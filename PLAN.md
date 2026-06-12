# Mozkit — Final Vision

Mozkit is a TUI wizard that helps you bootstrap, configure, and provision a fresh Linux/macOS system in minutes. It is designed to be run once (or very rarely) — not as a daily driver.

---

## Phase 1 — First-Run Detection & System Discovery

When Mozkit starts without a local state file (e.g. `~/.config/mozkit/state.json`), it enters the **first-run wizard** rather than jumping straight into the script browser.

### 1.1 OS Detection

Mozkit probes the underlying OS by reading (in order of priority):

1. `/etc/os-release` (Linux) — parses `ID`, `ID_LIKE`, `VERSION_ID`
2. `sw_vers` output (macOS) — product name and version
3. `uname -s` / `uname -r` — fallback kernel-level detection

The resolved OS tuple (family, distro, version) is stored in memory and persisted to the state file so this step never repeats.

### 1.2 Package Manager Discovery

Mozkit scans the `$PATH` for known package managers and records which ones are present and functional. The detection order and supported managers:

| Priority | Binary       | Typical OS                     |
|--------- |------------- |------------------------------- |
| 1        | `pacman`     | Arch, Manjaro, EndeavourOS    |
| 2        | `apt-get`    | Debian, Ubuntu, Pop!_OS       |
| 3        | `dnf`        | Fedora, RHEL, CentOS          |
| 4        | `zypper`     | openSUSE                      |
| 5        | `brew`       | macOS, Linux (Homebrew)       |
| 6        | `apk`        | Alpine Linux                  |
| 7        | `snap`       | Any (Snap daemon)             |
| 8        | `flatpak`    | Any (Flatpak daemon)          |

Mozkit also checks for **AUR helpers** (`yay`, `paru`) on Arch-based systems, and for **Nix** (`nix-env`) as an optional cross-platform manager.

The final list of usable managers is stored; each `install` action in scripts will be dispatched through the highest-priority available manager unless the script explicitly targets a specific one.

### 1.3 First-Run Confirmation Prompt

Before touching the system at all, Mozkit presents a TUI screen that shows:

```
Operating System:    Arch Linux (btw)
Available managers:  pacman, yay, brew, snap

Initial steps Mozkit will take:
  1. Run system update (pacman -Syu / apt-get update && apt-get upgrade / etc.)
  2. Install base prerequisites (curl, git, base-devel / build-essential)

[ ENTER to continue ]  [ ESC to skip all and go to manual mode ]
```

- **ENTER** — Mozkit runs the update + prerequisite install, then proceeds to the script browser.
- **ESC** — Mozkit skips all automated steps and lands directly in the script browser. The user can manually run update/install packs later from the "packs" collection.

The user's choice is persisted so the prompt never reappears (unless state is deleted).

---

## Phase 2 — Script Browser (Current UI, Enhanced)

After first-run setup (or on subsequent launches), Mozkit drops into the existing TUI browser:

### Navigation
- **↑/↓** or **j/k** — move selection
- **/** — fuzzy-filter the current list
- **Enter** — descend into a collection or open a script detail page
- **Esc** — go back to the parent collection
- **Ctrl+C** — quit

### Collections (from `collection.toml`)
- **Presets** — opinionated, multi-step system configurations (DE setup, dev environment, etc.)
- **Configs** — individual tool/dotfile configuration
- **Packs** — software package groups to install

### Script Detail Page
Shows the list of actions the script will take. The user confirms by pressing **Enter 3 times** (safety gate). Once confirmed, the script executes inline with a live output viewport.

---

## Phase 3 — Script Execution Engine

Scripts are TOML files with a defined set of action types:

| Tool         | Behaviour                                                                 |
|------------- |-------------------------------------------------------------------------- |
| `echo`       | Prints text to stdout (cosmetic / info step).                             |
| `ask`        | Prompts the user for input (stored in a variable for later actions).       |
| `install`    | Installs one or more packages through the detected system package manager. |
| `shell`      | Runs an arbitrary shell command and streams output live.                   |
| `file`       | Writes or appends content to a file (e.g. dotfiles, configs).              |
| `service`    | Enables/starts/restarts a systemd service (or macOS launchd equivalent).   |

### Execution Model
Actions run **sequentially** in order of appearance. If any action fails (non-zero exit), Mozkit stops and shows the error with a suggestion. There is **no rollback** — the user is warned about this on the confirmation screen.

### Live Output
While a script runs, the viewport streams stdout/stderr in real time. A progress indicator shows `[3/7] Installing neovim...` at the bottom.

---

## Phase 4 — State & Idempotency

Mozkit maintains a state file at `~/.config/mozkit/state.json` containing:

```json
{
  "os": {"family": "arch", "distro": "Arch Linux", "version": "2025.06.01"},
  "managers": ["pacman", "yay"],
  "first_run_completed": true,
  "completed_scripts": ["presets/dev-base.toml", "packs/essentials.toml"],
  "variables": {"best_distro": "Arch Linux"}
}
```

- Completed scripts show a ✓ checkmark in the browser.
- The user can re-run any completed script (with a warning).
- Variables captured via `ask` persist across sessions.

---

## Phase 5 — Package Install Intelligence

When a TOML action says:

```toml
tool = "install"
arguments = ["neovim", "nvim"]
```

Mozkit tries each argument as a package name against the available package managers, in priority order, until one succeeds. This means a single pack TOML works across **Arch, Debian, Fedora, and macOS** without modification — the argument list serves as a fallback chain.

For AUR-only packages, the script can target a specific manager:

```toml
tool = "install::aur"
arguments = ["visual-studio-code-bin"]
```

---

## Phase 6 — Exit & Summary

When the user quits (Ctrl+C or after running scripts), Mozkit prints a summary:

```
✨ Done! Mozkit ran 3 scripts and installed 14 packages.
   Completed: dev-base, essentials, neovim-config
   Next step: reboot? run mozkit again? grab a coffee?
```

---

## Non-Goals (Scope Boundaries)

- No daemon/background service — Mozkit is a run-once-and-exit tool.
- No remote script fetching — all scripts are embedded at build time via Go's `embed.FS`.
- No uninstall/rollback — simplicity over complexity.
- No Windows support (target: Linux + macOS only).
