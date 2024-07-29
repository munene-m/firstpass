# firstpass

password generator CLI

## Installation

```sh
git clone https://github.com/munene-m/firstpass.git`
cd firstpass
make build
```

## Usage

Run the CLI using:

```
make start-app ARGS="[flags]"
```

### Flags

- `-l --length int`: Length of the password(default: 15)
- `-c --count int`: Number of passwords to generate
- `-L --lowercase bool`: Include lowercase characters
- `-U --uppercase bool`: Include uppercase characters
- `-n --numbers bool`: Include numeric values
- `-cp --copy bool`: Copy to clipboard
- `-s --special bool`: Include special characters

### Example

```
make start-app ARGS="-l 14 -U -L -n -s -c 2"
```

This command generates 2 passwords, each 14 characters long, including uppercase, lowercase, numbers, and special characters.

### Features

- [x] Generate strong passwords
- [x] Use flags to select attributes of your password
- [x] Generate multiple passwords at once
- [x] Copy generated passwords to clipboard
- [ ] Option to exclude similar looking passwords
- [ ] Save generated passwords to file
- [ ] Password strength meter
