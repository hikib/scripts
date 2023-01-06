# Requires bash for windows
param($file)
$file = $file -replace “\\”, “/” -replace “ “, “\ “
bash -c “vim $File”
