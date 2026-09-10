```
//// sliitmozilla.org                                                                        
//// ██████  ██████  ███████ ███████  ███████████████  ████     ██████  ████  ████ ████ ████
//// ██████  ██████  ███████ ███████  ███████████████  ████   ██████    ████  ████ ████ ████
//// █  ████  █████  ████       ████          ███████  ████  █████      ████       ████
//// ██  ████  ████  ████       ████      ███████      ████ ████        ████       ████
//// ███  ███  ████  ████       ████   ██████          ████  █████      ████       ████
//// ████  █   ████  ███████ ███████  ███████████████  ████   ██████    ████       ████
//// ████      ████  ███████ ███████  ███████████████  ████     ██████  ████       ████
```

Ditch the messy bash scripts. Apply presets, configure tools, and install packages entirely from your keyboard, wrapped in a beautiful terminal interface.

## Features

- **Self-Contained:** A single Go binary. No databases, no hidden services, no nonsense.
- **Keyboard Native:** Keep your hands on the home row. Navigate your entire setup without ever reaching for the mouse.
- **Declarative:** All menus, packages, and presets live in clean, readable TOML files embedded right into the build.
- **Hackable:** Run the baked-in defaults out of the box, or bring your own external TOML configs at runtime.
- **Smooth as Butter:** Built on the Charm Bubbletea v2 ecosystem for a flawless, Elm-style terminal experience.

## How it looks

<p align="center">
  <img width="1366" height="768" alt="Mozkit Demo Screenshot" src="https://github.com/user-attachments/assets/39cbfd59-fb29-4e8d-a7e2-2ec0b93b720d" />
</p>

## Installation

Grab the code and build it yourself using the Go toolchain:
```bash
# Install globally via Go
go install github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/cmd/mozkit@latest
```

Or, build it from source:
```bash
# Clone and build
git clone https://github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/cmd/mozkit.git
cd mozkit
go build -o mozkit ./cmd/mozkit
```

## Contributing
Got a killer new script you want to share with the world? Hit up the [wiki](https://github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/wiki) for the rundown on how to write scripts using TOML, whip one up, and open a PR. We'd love to add it to the collection. If you're looking to squash bugs or add core engine features, check out our **[contributing guide](https://github.com/Mozilla-Campus-Club-of-SLIIT/mozkit?tab=contributing-ov)** to jump right in.

## Whatcha think?
We'd love to hear your thoughts on this project. Need help writing your own setup scripts or figuring out the architecture? We gotchu.
First, dive into the **[Mozkit Wiki](https://github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/wiki)** for the full breakdown and TOML schemas. Still stuck or just want to chat? You can find us on: **[Discord](https://discord.gg/VD4jwG7QF)**

## License
**[Mozilla Public License Version 2.0](https://github.com/Mozilla-Campus-Club-of-SLIIT/mozkit?tab=MPL-2.0-1-ov-file)**

Built with lots of coffee by the **[Mozilla Campus Club of SLIIT](https://sliitmozilla.org/)**
