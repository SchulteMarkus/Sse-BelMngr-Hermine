[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-pink.svg)](https://conventionalcommits.org)
[![Conventional Changelog](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-conventional--changelog-e10079.svg?style=flat)](https://github.com/conventional-changelog/conventional-changelog)
[![Renovate enabled](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://renovatebot.com/)
[![Build](https://github.com/SchulteMarkus/Sse-BelMngr-Hermine/actions/workflows/build.yml/badge.svg)](https://github.com/SchulteMarkus/Sse-BelMngr-Hermine/actions/workflows/build.yml)
[![Test](https://github.com/SchulteMarkus/Sse-BelMngr-Hermine/actions/workflows/test.yml/badge.svg)](https://github.com/SchulteMarkus/Sse-BelMngr-Hermine/actions/workflows/test.yml)
[![Lint](https://github.com/SchulteMarkus/Sse-BelMngr-Hermine/actions/workflows/lint.yml/badge.svg)](https://github.com/SchulteMarkus/Sse-BelMngr-Hermine/actions/workflows/lint.yml)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=SchulteMarkus_Sse-BelMngr-Hermine&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=SchulteMarkus_Sse-BelMngr-Hermine)

# SteuerSparErklärung BelegManger Invoice Importer "Hermine"

**Sse-BelMngr-Hermine** is a command-line tool designed to analyze local documents using
[Azure® AI Document Intelligence](https://azure.microsoft.com/en-us/products/ai-services/ai-document-intelligence)
and seamlessly import the processed data into the BelegManager of
[SteuerSparErklärung®](https://www.steuertipps.de).

```mermaid
sequenceDiagram
    participant User as 👤 User
    participant Hermine as 🪄 Sse-BelMngr-Hermine 🪄
    participant AzureAI as Azure® AI Document Intelligence
    participant BelegManager as Sse® BelegManager
    User ->> Hermine: Run for specific local documents
    Hermine ->> AzureAI: Analyze documents
    AzureAI ->> Hermine: Return analyzed data
    Hermine ->> BelegManager: Import documents
```

# Table of contents

* [🚀 Key Features](#-key-features)
    * [Supported File Types](#supported-file-types)
    * [Advanced Features](#advanced-features)
* [🛠️ Installation](#-installation)
* [🌟 Usage](#-usage)
    * [Command-Line Quickstart](#command-line-quickstart)
    * [Command-Line Flags](#command-line-flags)
* [⚙️ Configuration File](#-configuration-file)
* [🎯 Workflow](#-workflow)
* [📝 Examples](#-examples)
    * [Example Run](#example-run)
    * [Outcome](#outcome)
* [🛡️ Error Handling](#-error-handling)
* [🖥️ Project Structure](#-project-structure)
* [📚 Dependencies](#-dependencies)
* [📜 License](#-license)
* [Disclaimer](#disclaimer)
* [💬 Feedback](#-feedback)
* [Additional Resources](#additional-resources)
* [Project Name Inspiration](#project-name-inspiration)

---

## 🚀 Key Features

- **Document Analysis**  
  Harnesses Azure AI Document Intelligence to extract information (vendor, total, VAT, etc.) from
  PDF, JPG, PNG, TIF/TIFF documents.

- **Smooth Import**  
  Effortlessly imports processed documents into the BelegManager tool.

- **Flexible Configuration**  
  Integrates seamlessly with command-line arguments, environment variables, and configuration files.

- **Custom File Selection**  
  Uses glob patterns to filter and choose only the files you need for analysis.

- **Robust Logging**  
  Configurable logging levels—from detailed debug logs to concise summaries.

### Supported File Types

- **Input Files**: `jpg`, `pdf`, `png`, `tif`, `tiff`
- **Document Intelligence Compatible Types**: Includes additional types like `jfif`, `jp(e)g`, and
  more.

### Advanced Features

- **Data Backup**:
  Automatically backs up the BelegManager SQLite database before processing.

- **Parallel File Processing**:
  Supports concurrent processing of multiple files for efficiency.

- **Error Handling**:
  Gracefully handles missing files, invalid configurations, and failed imports.

---

## 🛠️ Installation

1. **Prerequisites**
    - Go (Golang) SDK 1.23 or later.
    - An Azure Cognitive Services account with Document Intelligence enabled.

2. **Install Application**

```shell script
go install github.com/SchulteMarkus/sse-belmngr-hermine@latest
```

## 🌟 Usage

### Command-Line Quickstart

```shell script
sse-belmngr-hermine --di-key <Azure_AI_key> --di-endpoint <Azure_AI_endpoint>
```

### Command-Line Flags

| Flag                             | Shorthand | Description                                                                                                                             | Required | Default Value                                             |
|----------------------------------|-----------|-----------------------------------------------------------------------------------------------------------------------------------------|----------|:----------------------------------------------------------|
| `--config`                       | `-c`      | Path to the configuration file (optional).                                                                                              | No       | *None*                                                    |
| `--di-key`                       |           | Azure Document Intelligence API key. Use this to authenticate against Azure services.                                                   | Yes      | *None*                                                    |
| `--di-endpoint`                  |           | Azure Document Intelligence endpoint URL.                                                                                               | Yes      | *None*                                                    |
| `--files-to-import-glob`         | `-f`      | Glob pattern to locate the input document files (supports wildcards). Defaults to user documents directory under `BelegManager-Import`. | No       | Documents/BelegManager-Import/**/*.{jpg,pdf,png,tif,tiff} |
| `--beleg-manager-data-directory` |           | Specify the root directory for BelegManager data (default: the `Documents/BelegManager-Daten` folder in the user's home directory).     | No       | Documents/BelegManager-Daten                              |
| `--log-level`                    | `-l`      | Specify the logging level (trace, debug, info, warn, error, fatal, panic). Defaults to `info`.                                          | No       | `info`                                       ~~~~         |

---

## ⚙️ Configuration File

Optionally, you can manage settings via a configuration file (e.g., `config.yml`):

```yaml
di-key: "your-azure-ai-key"
di-endpoint: "https://<your-endpoint>.cognitiveservices.azure.com/"
files-to-import-glob: "C:/Users/<your-user-name>/Documents/BelegManager-Import/**/*.pdf"
beleg-manager-data-directory: "C:/Users/<your-user-name>/Documents/BelegManager-Daten"
log-level: "debug"
```

Then run:

```shell script
sse-belmngr-hermine -c config.yaml
```

---

## 🎯 Workflow

1. **Analyze Documents**
    - Scans local files using the specified `--files-to-import-glob`.
    - Extracts invoice data (vendor, total, VAT, items) via Azure AI Document Intelligence.

2. **Process Data**
    - Validates database compatibility.
    - Structures and readies extracted info for BelegManager.

3. **Import to BelegManager**
    - Inserts discovered information into the BelegManager database.
    - Creates a backup of the BelegManager database before any changes.

4. **Logging & Summaries**
    - Outputs a processed file report (CSV) detailing import status for each document.

---

## 📝 Examples

### Example Run

```shell script
sse-belmngr-hermine run \
  --files-to-import-glob "~/Documents/BelegManager-Import/**/*.pdf" \
  --beleg-manager-data-directory "~/Documents/BelegManager-Daten" \
  --di-key "<your-azure-ai-key>" \
  --di-endpoint "<your-azure-ai-endpoint>" \
  --log-level "debug"
```

### Outcome

- Documents are imported and linked in BelegManager.
- A CSV log file is generated with status and any encountered errors:

```
~/Documents/BelegManager-Daten/_import-log-<timestamp>.csv
```

---

## 🛡️ Error Handling

- **Missing Database**: Alerts if the BelegManager database is not found.
- **Unsupported Document**: Skips files if format mismatches or duplicates exist.
- **Azure Failures**: Employs retries and logs any network or API issues.

---

## 🖥️ Project Structure

- **cli/**  
  Houses CLI logic such as flag handling and command execution.
- **hermine/**  
  Contains the core functionality for document analysis and database interactions.

---

## 📚 Dependencies

This application makes use of the following key libraries/packages:

- [Cobra](https://github.com/spf13/cobra) for CLI command building.
- [Viper](https://github.com/spf13/viper) for configuration management.
- [Logrus](https://github.com/sirupsen/logrus) for structured logging.
- [sqlx](https://github.com/jmoiron/sqlx) for database querying.
- [doublestar](https://github.com/bmatcuk/doublestar) for glob pattern matching.
- Azure AI for Document Intelligence.

---

## 📜 License

This project is released under the [GPLv3 License](./LICENSE).
Please ensure you comply with this license when using the application.

---

## Disclaimer

This software is currently in the test phase. No further liability is assumed. Use at your own risk.

---

## 💬 Feedback

If you have any questions, issues, or suggestions, feel free to open an issue in
the [GitHub Issues Section](https://github.com/SchulteMarkus/Sse-BelMngr-Hermine/issues).

--- 

## Additional Resources

- [Azure Document Intelligence Documentation](https://learn.microsoft.com/en-us/azure/applied-ai-services/form-recognizer/overview)
- [SteuerSparErklärung BelegManager](https://www.steuertipps.de/steuererklaerung/software/steuer-spar-erklaerung)

--- 

## Project Name Inspiration

Project name "Hermine" is inspired by the resourceful and knowledgeable character
[Hermione Granger](https://en.wikipedia.org/wiki/Hermione_Granger) from the Harry
Potter® series.
