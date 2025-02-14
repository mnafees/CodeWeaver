$version = "dev" 
$commit = git rev-parse --short HEAD
$date = Get-Date -format "yyyy-MM-ddTHH:mm:ss"

go build -ldflags "-X 'main.version=$version' -X 'main.commit=$commit' -X 'main.date=$date'" .

# goreleaser build --clean

./CodeWeaver -version
./CodeWeaver -ignore="\.git.*,.+\.exe,codebase.md,excluded_paths.txt" -excluded-paths-file="excluded_paths.txt"