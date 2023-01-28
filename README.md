# panda-sale-ape

## Introduction

This project is a simple tool that allows you to interact directly with a pandasale presale. 

## Installation

### Requirements
golang 1.16+
.env file with the following variables:

```bash
PRIVATE_KEY={your private key}
```

### Build

```bash
git clone
cd panda-sale-ape
go build -o panda-sale-ape
```

## Usage

### Presale

```bash
./panda-sale-ape [Presale Address] [Token Amount ETH]
```

# Example
```bash
./panda-sale-ape 0x9c8d8d9a9f5a5f9e9d9d9d9d9d9d9d9d9d9d9d9d  0.1
```