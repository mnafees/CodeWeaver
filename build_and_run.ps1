# mit the symbol table, debug information and the DWARF symbol table by passing -s and -w go build -ldflags="-s -w" .
go build # -ldflags="-s -w" .

./codeweaver -h

./codeweaver -ignore="\.git.*,.+\.exe,codebase.md,excluded_paths.txt" -excluded-paths-file="excluded_paths.txt"

# "\.git.*,.+\.exe,.+\.[Ii][Cc][Oo],.+\.[Pp][Nn][Gg],.+\.[Jj][Pp][Ee]?[Gg],.+\.[Mm][Pp][34],.+\.[Aa][Vv][Ii],.+\.[Mm][Oo][Vv],.+\.[Ww][Mm][Aa],.+\.[Pp][Dd][Ff],.+\.[Dd][Oo][Cc][Xx]?,.+\.[Xx][Ll][Ss][Xx]?,.+\.[Pp][Pp][Tt][Xx]?,.+\.[Zz][Ii][Pp],.+\.[Rr][Aa][Rr],.+\.[7][Zz],.+\.[Ii][Ss][Oo],.+\.[Bb][Ii][Nn],.+\.[Dd][Aa][Tt],.+\.[Dd][Mm][Gg],.+\.[Gg][Ii][Ff],.+\.[Tt][Ii][Ff][Ff]?,.+\.[Pp][Ss][Dd],.+\.[Mm][Kk][Vv],.+\.[Ff][Ll][Aa][Cc],.+\.[Oo][Gg][Gg],.+\.[Ww][Ee][Bb][Pp],.+\.[Aa][Vv][Ii][Ff],.+\.[Hh][Ee][Ii][Cc],.+\.[Jj][Xx][Ll]"