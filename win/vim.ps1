# Requires bash for windows
$file = $args -replace "\\", "/" -replace " ", "\ "
bash -c "vim $file"
