# Runs mkdocs with my mome directory in browser

# startup directory -> windows+r 'shell:startup'
# Add the following line to the startup.cmd script

# #startup.cmd
# PowerShell -WindowStyle hidden -Command "Set-ExecutionPolicy Unrestricted" >> "%TEMP%\StartupLog.txt" 2>&1
# PowerShell -WindowStyle hidden %USERPROFILE%\.local\bin\podmanmachinestart.ps1 >> "%TEMP%\StartupLog.txt" 4>&1
# PowerShell -WindowStyle hidden %USERPROFILE%\.local\bin\podmanrunmo.ps1 "%TEMP%\StartupLog.txt" 4>&1

# give it a second to start up. you have time.
echo "y" | podman container prune
podman run -d --rm -it -p 8000:8000 -v ${HOME}/repos/github.com/hikib/mome:/docs docker.io/squidfunk/mkdocs-material
start-sleep -seconds 5
firefox localhost:8000
