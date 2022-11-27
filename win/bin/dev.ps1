# This script mounts the current directory to a container with my dev setup

function currentDirectory {Split-Path -leaf -path (Get-Location)}

podman run --rm -it `
  --hostname pluto `
  --volume "./:/home/$(currentDirectory)" `
  --workdir "/home/$(currentDirectory)" `
  docker.io/hikib/pluto:1.0.0

