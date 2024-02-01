# Moloni API Go Client Library

This Go client library provides a convenient way to interact with the Moloni API.

## Roadmap

- ✅ Automatic authentication.
- ✅ Operations on Taxes.
- 🚧 Operations on Document Sets.
- 🚧 Operations on Products.
- 📅 Operations on more resources.
- 📅 Improve error handling.
- 📅 Improve testing.
- 📅 Allow to configure logger.

## Installation

To use this library, first ensure you have Go installed on your system. Install the package by running:

```sh
go get github.com/elabprosolutions/moloni-go
```

Then, you can import the package into your Go project:

```go
import "github.com/elabprosolutions/moloni-go"
```

## Usage

### Creating a Client

To start using the library, you need to create a `Client` instance. You can provide credentials directly or load them from environment variables:

```go
client, err := moloni.NewClient(
    moloni.WithCredentials(&moloni.Credentials{
        ClientID:     "moloni_client_id",
        ClientSecret: "moloni_client_secret",
        Username:     "moloni_username",
        Password:     "moloni_password",
    }),
)
if err != nil {
    log.Fatal(err)
}
```

Or using environment variables:

```go
client, err := moloni.NewClient(moloni.LoadCredentialsFromEnv())
if err != nil {
    log.Fatal(err)
}
```

The library gets the following environment variables:

- `MOLONI_CLIENT_ID`
- `MOLONI_CLIENT_SECRET`
- `MOLONI_USERNAME`
- `MOLONI_PASSWORD`

ℹ️ Please refer to the [Moloni API documentation](https://www.moloni.pt/dev) for detailed information on API usage, credentials, and setup.

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues.
