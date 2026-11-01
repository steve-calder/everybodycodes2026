# Everybody Codes 2026

Go solutions for Everybody Codes. Requires Go 1.24.3 or newer.

## Configuration

From the project directory, copy the example configuration:

```sh
cp config.example.yaml config.yaml
```

Edit `config.yaml` with your event, account seed, and session token:

```yaml
EVENT: 2026
SEED: 83 # Replace with your account's seed.
SESSION: "your-session-token"
```

`EVENT` identifies the event to retrieve notes for. `SEED` is the `seed` value from your account's `/user/me` API response. `SESSION` is your Everybody Codes session cookie value.

`config.yaml` is ignored by Git. Run the app from the project directory because it reads `config.yaml` from the current working directory.

## Run

Download dependencies, then run all quests:

```sh
go mod download
go run .
```

To run a single quest, pass its number:

```sh
go run . 1
```

Event and seed settings come from `config.yaml`.

## Build

Build an executable and run it from the project directory:

```sh
go build -o bin/everybodycodes2026 .
./bin/everybodycodes2026
./bin/everybodycodes2026 1
```

## Input cache

On the first run, the app downloads the quest's JSON input notes and saves them to `cache/<event>-<quest>-<seed>.json`. Later runs read that file instead of making a web request. Delete the file to download it again. The `cache/` directory is ignored by Git.

## AI Usage

AI was used to develop the note retrieval, decryption, and caching. AI was not used for any quest answers.
