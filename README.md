# Go Version Manager (GVM) #

## build/release status ##

![Release artifact](https://github.com/Eldius/gvm/workflows/Release%20artifact/badge.svg)
![Go](https://github.com/Eldius/gvm/workflows/Go/badge.svg)

## commands ##

- gvm:
  - ls: `gvm ls` list installed versions
  - ls-remote: `gvm ls-remote` list versions available to install
  - install: `gvm install <version>` install available version
  - use: `gvm use <version>` setup version as active
  - hooks: `gvm hooks` list hooks
    - add: `gvm hooks add <script>` add a script as hook (executed when you change the active version)
