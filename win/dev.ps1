# Mounts the current directory to one of my dev containers

function currentDirectory {Split-Path -leaf -path (Get-Location)}

if ($args) { $devImage = "docker.io/hikib/dev.$args" }
else { $devImage = "docker.io/hikib/dev" }

podman run --rm --interactive --tty `
  --hostname dev `
  --volume "./:/home/hikib/$(currentDirectory)" `
  --workdir "/home/hikib/$(currentDirectory)" `
  $devImage
