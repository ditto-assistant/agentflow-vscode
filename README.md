# agentflow-vscode README

## What is AgentFlow?

AgentFlow is a prompting language designed to generate LLM (Large Language Model) Agent code in the language of your choice. It provides a simple and intuitive syntax that allows developers to easily create prompts for AI agents without getting in the way of the creative process.

Key features of AgentFlow:
- Supports code generation for Python, TypeScript, and JavaScript (with more languages planned)
- Uses a .af file extension
- Employs an ultra-simple syntax focused on efficient prompting

## Current Extension Features

This VSCode extension currently provides:
- Syntax highlighting for AgentFlow (.af) files

**Note:** A Language Server Protocol (LSP) implementation is planned for future releases to provide more advanced features.

## AgentFlow Syntax Overview

AgentFlow uses a minimalist syntax:

1. Variables:
   ```
   <!variable_name>
   ```
   Defines a variable which becomes a string input in code generation.

2. Titles:
   ```
   .title This Is The Title Of The Section
   ```
   The rest of the line after `.title` becomes the title. Code generation automatically converts the title to camelCase or snake_case based on language conventions.

## Installation

[Provide installation instructions once the extension is published]

## Usage

1. Install the AgentFlow VSCode extension
2. Create a new file with the `.af` extension in VSCode
3. Start writing your AgentFlow prompts using the syntax described above
4. Enjoy syntax highlighting for your AgentFlow files

## Release Notes

### 0.0.1
- Initial release of AgentFlow syntax highlighting

## Planned Features

- Language Server Protocol (LSP) implementation for enhanced functionality
- Code generation integration within VSCode
- Support for additional programming languages

## Contributing

If you'd like to contribute to the development of this extension or the AgentFlow language, please [provide information on how to contribute].

## License

[Include your chosen license information here]

---

For more information on AgentFlow and its capabilities, please visit [your website or repository link].

**Enjoy using AgentFlow!**
