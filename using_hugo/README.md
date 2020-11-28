## Installation
See the [releases](https://github.com/gohugoio/hugo/releases)
Where you can get the hugo application for your OS.

## Quickstart
```bash
hugo new site sitename
cd sitename
git init
git submodule add https://github.com/budparr/gohugo-theme-ananke.git themes/ananke
echo 'theme = "ananke"' >> config.toml

# add some content
hugo new posts/my-first-post.md

# or we can change the theme
git submodule add https://github.com/vaga/hugo-theme-m10c.git themes/m10c.git themes/m10c

hugo server -D
```

## Examples using hugo
See the hugo [showcase](https://gohugo.io/showcase/)
