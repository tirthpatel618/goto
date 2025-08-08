goto is a cmd line tool to go to any path instantly, and build aliases easily

To use it for yourself, once you add all the files, run:
go build -o ~/bin/goto
and then add this into your ~/.zshrc file:
function goto() {
  local dest=$(~/bin/goto "$@")
  if [ -d "$dest" ]; then
    cd "$dest"
  else
    echo "$dest"
  fi
}
