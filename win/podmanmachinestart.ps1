# Starts the podman machine called 'dev'

# startup directory -> windows+r 'shell:startup'
# Add the following line to the startup.cmd script
# #startup.cmd
# PowerShell %USERPROFILE%\.local\bin\podmanmachinestart.ps1 >> "%TEMP%\StartupLog.txt" 4>&1

podman machine start dev
