# Compress
A module to standize compression.

## Table of Contents
1. [Usage](#usage)

### Usage
At the moment, the only compression algorithm supported is `gzip`. To compress data the function `Gzip` is conveniently defined. This function takes in a slice of bytes and either returns a compressed slice of bytes or an `IOError`. To uncompress data, the function `Gunzip` is conveniently defined. This function takes in a compressed slice of bytes and either returns an uncompressed slice of bytes or an `IOError`.