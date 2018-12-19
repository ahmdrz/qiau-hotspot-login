# QIAU Hotspot Login 🤦🏻‍

Nothing just a simple http client to log us in with less pain and effort.


## Installation

```bash
    sudo wget https://github.com/ahmdrz/qiau-hotspot-login/releases/download/0.0.1/qiau-login-linux -O /usr/local/bin/qiau-hotspot-login
    sudo chmod +x /usr/local/bin/qiau-hotspot-login
    sudo chown $(whoami) /usr/local/bin/qiau-hotspot-login

    # Do not forget to set environment variables in your .bashrc|.zshrc .
```

## Usage

Set following environment variables and run the binary.

```bash
QIAU_USERNAME
QIAU_PASSWORD
```

## Related projects

- [Captive Login](https://github.com/authq/captive-login) - Captive-portal login utility for headless environments written in pure bash.

## Licence

Whatever. 😑

##### Forked from [Reyhoon Hotspot Login](https://github.com/mamal72/reyhoon-hotspot-login)