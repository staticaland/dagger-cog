# Usage

```sh
dagger call cog --dir="." --file="README.md" --replace export --path="."
```

## Output from the command above

<!-- [[[cog
import subprocess

output = subprocess.check_output(['ls']).decode('utf-8')

cog.out(f"```sh\n{output}```")
]]] -->
```sh
LICENSE
README.md
dagger.gen.go
dagger.json
go.mod
go.sum
internal
main.go
```
<!-- [[[end]]] -->
