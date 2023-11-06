# Upgrading NodeJs to Latest LTS Version on Ubuntu

First install `npm`:

```bash
sudo apt update
sudo apt install npm
```

Install the `n` module:

```bash
sudo npm install -g n
```

Install the latest LTS version of node with:

```bash
sudo n stable
```

Reload your shell and you will see the new version with:

```bash
node -v
```

