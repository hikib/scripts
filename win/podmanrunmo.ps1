# Runs mkdocs with my mome directory in browser

# startup directory -> windows+r 'shell:startup'
# Add the following line to the startup.cmd script
# #startup.cmd
# PowerShell %USERPROFILE%\.local\bin\podmanrunmo.ps1 >> "%TEMP%\StartupLog.txt" 4>&1

podman run -d --rm -it -p 8000:8000 -v ${HOME}/repos/github.com/hikib/mome:/docs docker.io/squidfunk/mkdocs-material
# give it a second to start up. you have time.
start-sleep -seconds 3
firefox localhost:8000
