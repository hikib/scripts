New-Item -ItemType SymbolicLink -Path "$home\.local\bin" -Target "$PSScriptRoot\bin" -Force

Read-Host -Prompt "Press Enter to exit"

