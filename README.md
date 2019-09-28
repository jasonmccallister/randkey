# randkey

This project provides a simple API endpoint that generates a key of a specific length. This code is organized to be cloud agnostic.

The "business logic" is stored in `pkg/key` and only exposes one function, `Generate` which takes a length. Keeping the business logic seperate from the HTTP layer allows vendor specific handlers the ability to wrap their requests around whats important to this project, generating a key of a specific length.
