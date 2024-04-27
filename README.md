# IlluminateMyWords

IlluminateMyWords is a dynamic keyword highlighting tool that improves readability and facilitates error monitoring in terminal outputs. This standalone utility color-codes predefined keywords, making it easier to spot critical information in dense text flows. It is environment-agnostic, capable of being integrated into various shells and applications.

## Features

- [x] **Independent Utility**:
  Functions independently of the terminal environment, providing flexibility across different platforms.
- [x] Configurable using Yaml:
  Allow a custom mapping between keywords and colors, by types.
- [x] **Dynamic Highlighting**:
  Automatically highlights words based on a configurable set of rules.
- [x] **Customizable Patterns**:
  Users can define their own keywords and associated colors through a simple configuration file.
- [ ] **Hot Reload**:
  Modification of the configuration is applied without needing to relaunch the tool.

## Installation

1. **Clone the repository**:

  ```bash
  git clone https://github.com/d0tiks/IlluminateMyWords.git
  cd IlluminateMyWords
  ```

2. Compile the sources

  ```bash
  go build -o ilmw /src/main.go
  ```

3. Configure your Keyword Mapping

  ```yaml
  types:
  - name: errors
    color: red
    keywords:
      - error(s)?
      - fail(ed)?
      - ko
      - not found
      - invalid
  - name: warnings
    color: orange
    keywords:
      - warning(s)?
      - alert(s)?
  - name: informations
    color: cyan
    keywords:
      - info(rmation(s)?)?
      - notice
      - verbose
  - name: successes
    color: green
    keywords:
      - success
      - succeded
      - ok
      - valid
  - name: debug
    color: yellow
    keywords:
      - debug
  ```

4. Run the tool

  ```bash
  ./ilmw -conf your_config.yaml
  ```
