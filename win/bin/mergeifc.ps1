# Merge multiple IFC files
#
# Requires python module ifcopenshell
# and a local copy of the ifcopenshell repo (hardcoded path)

# TODO:
#   - help text
#   - Dynamic ifcopenshell repo location (or on common drive?)
#   - handle non-ifc files
#   - exit if ifcrepo does not exist

$here = "$(Get-Location)"
$merged = "$here\merged.ifc"
New-Item -Path "$merged" -ItemType File -Force > $null

$ifcrepo = "${env:userprofile}/repos/github.com/IfcOpenshell/IfcOpenshell/src/ifcpatch"
cd "$ifcrepo"

function patch { param([string[]] $paramargs)
  python -m ifcpatch `
    -i "$thisdir/$($paramargs[0])" `
    -o "$merged" `
    -r "MergeProject" `
    -a "$thisdir/$($paramargs[0])" #>/dev/null 2>&1
}

Write-Host "Merging $($args[0]) and $($args[1]) ..."
patch $args[0..1]

if( $args.Length -ge 3 )
{
  for ($i=2; $i -lt $args.Length; $i++)
  {
    Write-Host "Adding" $args[$i] "..."
    patch [$($args[$i]), $merged]
  }
  Write-Host "Merged $($args.Length) files"
}

cd $here
Write-Host "$merged"
