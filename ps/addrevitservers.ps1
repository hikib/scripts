# This script adds server addresses to Revit Server configuration files.

$next_year = [int](Get-Date -Format "yyyy") + 1
$servers = Get-Content (Join-Path $PSScriptRoot servers.txt)

for ($year = 2016; $year -le $next_year; $year++) {

  $directory = "C:\ProgramData\Autodesk\Revit Server {0}\Config" -f "$year"
  $directory_exists = Test-Path -LiteralPath $directory -PathType Container

  if (-not $directory_exists) {
    $msg = "FAIL: The folder <{0}> does not exist." -f $directory
    Write-Host $msg
    continue
  }

  $file = Join-Path $directory RSN.ini

  foreach ($server in $servers) {

    $file_exists = Test-Path -LiteralPath $file -PathType Leaf

    if ( $file_exists -And ((Get-Content $file) -contains $server)){
      $msg = "FAIL: Revit Server {0} already contains <{1}>." -f $year, $server
      Write-Host $msg
      continue
    }

    Add-Content -Path $file -Value $server
    $msg = "PASS: Added < {0} > to Revit Server {1}." -f $server, $year
    Write-Host $msg
  }
}

Read-Host -Prompt "Press Enter to exit"

