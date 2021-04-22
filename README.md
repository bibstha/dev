# dev

An opinionated command to clone and change directory locally. Inspired from the
`dev` command used at Shopify.

Supports `dev clone` and `dev cd`.

## Installation

```
git clone https://github.com/bibstha/dev
cd dev
go build
cp dev ~/bin # copy dev binary to $PATH
```

Add `scripts/dev.fish` to your fish config file.
```
# Path: ~/.config/fish/config.fish 
if test -f ~/path/to/dev/scripts/dev.fish
  source ~/path/to/dev/scripts/dev.fish
end
```

Note that it is lacking support for `bash` or `zsh` but that should be easy to
add.

## Usage

### dev clone

Clones a repository from github.

The root folder is hardcoded as `~/src/github.com`. All projects are cloned
inside this folder in the `~/src/github.com/username/projectname` format.

The following command will clone [https://github.com/rails/rails](https://github.com/rails/rails)
to `~/src/github.com/rails/rails` folder.

```
dev clone rails/rails
```

### dev cd

Easily `cd` into the right directory inside `~/src/github.com/`.
It uses a fuzzy finder to find the best match for given name.

For example, given multiple folders below, it picks the right folder.

```
~ $ tree -L 2 ~/src/github.com
/home/bibek/src/github.com
├── bibstha
│   ├── dev
│   ├── dotfiles
│   ├── project1
│   ├── project2
│   └── project3
└── rails
    └── rails

~ $ dev cd proj1 # picks bibstha/project1
~/s/g/b/project1 $ pwd
/home/bibek/src/github.com/bibstha/project1

~/s/g/b/project1 $ dev cd rail
~/s/g/r/rails $ pwd
/home/bibek/src/github.com/rails/rails
```

As you can see, `proj1` correctly points to `project1` and `rail` points to
`rails/rails`.
