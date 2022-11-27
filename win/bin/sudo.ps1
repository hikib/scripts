# Run script as admin
$adminscript = $args[0]
if($adminscript) {
  Start-Process powershell -verb runas -ArgumentList "-file $adminscript"
} else{ echo "missing command/script" }
