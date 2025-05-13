# Installing MySQL on macOS M1

If you don't already have Homebrew installed, you can install it with:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

[Homebrew website](https://brew.sh/)

Then use it to install MySQL:

```bash
brew install mysql
```

To start MySQL server, use:

```bash
brew services start mysql
```

Then secure it with:

```bash
mysql_secure_installation
```

Use can connect to the MySQL server with:

```bash
mysql -uroot -p
```

