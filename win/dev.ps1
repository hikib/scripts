# This script mounts the current directory to a container with my dev setup

function currentDirectory {Split-Path -leaf -path (Get-Location)}

if ($args) { $devImage = "docker.io/hikib/dev.$args" }
else { $devImage = "docker.io/hikib/dev" }

podman run --rm --interactive --tty `
  --hostname dev `
  --volume "./:/home/$(currentDirectory)" `
  --workdir "/home/$(currentDirectory)" `
  $devImage
