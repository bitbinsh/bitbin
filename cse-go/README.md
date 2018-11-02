# cse-go

CSE-go is a golang package which implements the Bitbin CSE format. CSE is
implemented as a decorated Reader Writer pair.

## Format

CSE-go has the following format:

### CSE Version Independent Header

This header must be included at the beginning of a CSE file:

File Magic / Identifying header (10 bytes): `BITBIN CSE\0` - `42 49 54 42 49 4E
20 43 53 45 00`

Version: `0` - `00`

### CSE Version 0 Header

This header must be located directly after the CSE Version Independent Header:

Decryption Header Count (Varint)
Decryption Headers[Decyption Header Count]

### CSE Version 0 Decryption Header

Decryption Type (1 byte)

| Decryption ID | Decryption Type |
| ------------- | --------------- |
| 0             | PBKDF2 based Password |
| 1             | Public Key |

Decryption Payload Size (Varint)

Decryption Payload (Payload Size)

### CSE Version 0 Decryption Payload

A decryption payload is of variable size encrypted, however unencrypted using
its various Public/Private key mechanism the decryption payload is of constant
size. We support (for whatever reason) RSA 1024bit keys if the user so chooses,
so we will always have a maximum payload size of 86 bytes.

Our eventual payload is encrypted using AES-GCM-256. This makes our key 256 bits
long or 32 bytes. We also store inter-compression settings, such as whether or
not we compressed our payload before encryption and what we used.

AES Key
