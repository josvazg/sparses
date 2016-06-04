# sparses
Quick and dirty Go Command Line tool to find sparse files

Sparses walks a given path recursively, checking the size in blocks of each file and the actual blocks used.
Sparse files will have a bigger file size than the actual number of blocks x block size.
Those will be returned in the output.

This can be used to find any sparse files in your sparse capabel file system, like half downloded files or torrents.

## Usage

```
sparses {path}
```
