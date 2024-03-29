# Mounts the current directory to one of my dev containers

function currentDirectory {Split-Path -leaf -path (Get-Location)}

if ($args) { $devImage = "ghcr.io/hikib/dev.$args" }
else { $devImage = "ghcr.io/hikib/dev" }

podman run --rm --interactive --tty `
  --hostname dev `
  --volume "$env:MOME/:/home/hikib/mome" `
  --volume "./:/home/hikib/$(currentDirectory)" `
  --workdir "/home/hikib/$(currentDirectory)" `
  $devImage
