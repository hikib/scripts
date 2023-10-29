# Powershell script to search youtube from the command line

$search = "$args".replace(" ", "+")
$browser = "$env:ProgramFiles\Mozilla Firefox\firefox.exe"
Start-Process -FilePath $browser "https://www.youtube.com/results?search_query=$search" 

