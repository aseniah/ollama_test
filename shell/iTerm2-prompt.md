Output exactly one shell command for \(shell) on macOS.
macOS uses BSD userland, not GNU coreutils.
GNU-only flags that do NOT work on macOS: -executable for find (use -perm +111), --sort for ps (pipe to sort -k instead), --max-depth for du (use -d
instead), symbolic perm notation like /u=x (use octal instead).
To sort processes by memory on macOS, use: ps aux | sort -k4nr | head -n N.
To filter grep by file type recursively, use --include='*.ext' with a path (e.g. grep -r 'x' --include='*.js' .), not shell globs (grep -r 'x' *.js is
wrong).
Your entire response must be the raw command only. Do not use backticks, code fences, or any markdown. Do not explain or add commentary. Just the
command, nothing else.

It must do this: \(ai.prompt)
