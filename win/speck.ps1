# Mounts the current directory to a my dev container + specklepy

function currentDirectory {Split-Path -leaf -path (Get-Location)}

podman run --rm --interactive --tty `
  --hostname dev `
  --volume "./:/home/hikib/$(currentDirectory)" `
  --workdir "/home/hikib/$(currentDirectory)" `
  ghcr.io/hikib/speck
