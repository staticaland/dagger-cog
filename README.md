# Wuz

<!-- :::cog
import subprocess

output = subprocess.check_output(['ls']).decode('utf-8')

cog.out(f"```sh\n{output}```")
::: -->
```sh
README.md
```
<!-- :::end::: -->

## How to use Cog

<!-- :::cog
import subprocess

output = subprocess.run(['cog', '-h'], stderr=subprocess.PIPE, universal_newlines=True).stderr

cog.out(f"```sh\n{output}```")
::: -->
```sh
cog - generate content with inlined Python code.

cog [OPTIONS] [INFILE | @FILELIST | &FILELIST] ...

INFILE is the name of an input file, '-' will read from stdin.
FILELIST is the name of a text file containing file names or
other @FILELISTs.

For @FILELIST, paths in the file list are relative to the working
directory where cog was called.  For &FILELIST, paths in the file
list are relative to the file list location.

OPTIONS:
    -c          Checksum the output to protect it against accidental change.
    -d          Delete the generator code from the output file.
    -D name=val Define a global string available to your generator code.
    -e          Warn if a file has no cog code in it.
    -I PATH     Add PATH to the list of directories for data files and modules.
    -n ENCODING Use ENCODING when reading and writing files.
    -o OUTNAME  Write the output to OUTNAME.
    -p PROLOGUE Prepend the generator source with PROLOGUE. Useful to insert an
                import line. Example: -p "import math"
    -P          Use print() instead of cog.outl() for code output.
    -r          Replace the input file with the output.
    -s STRING   Suffix all generated output lines with STRING.
    -U          Write the output with Unix newlines (only LF line-endings).
    -w CMD      Use CMD if the output file needs to be made writable.
                    A %s in the CMD will be filled with the filename.
    -x          Excise all the generated output without running the generators.
    -z          The end-output marker can be omitted, and is assumed at eof.
    -v          Print the version of cog and exit.
    --check     Check that the files would not change if run again.
    --markers='START END END-OUTPUT'
                The patterns surrounding cog inline instructions. Should
                include three values separated by spaces, the start, end,
                and end-output markers. Defaults to '[[[cog ]]] [[[end]]]'.
    --verbosity=VERBOSITY
                Control the amount of output. 2 (the default) lists all files,
                1 lists only changed files, 0 lists no files.
    -h          Print this help.
```
<!-- :::end::: -->
