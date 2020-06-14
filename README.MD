# Touch

This is a simple program to provide a utility similar to the Linux touch command
for Windows users. I commonly find myself needing this command but forgetting
what it is called in Powershell or CMD so I decided to make it into a simple go
program.

## Installation

I recommend just building this program from source. It has no external
dependencies and relies only on the standard lib as well as being quite small.
Install it with "go install" and now you can use it just like you would on
Linux.

## Other Notes

-   I have used this app in go 1.13.4, it should work on any recent version.

## Usage

A basic CLI has been implemented for this program. Using the -h option will
provide a usage message like this one:

-   Usage: `touch <name> [content]`
    -   name: the name of the file to create
    -   content: initial content to put in the new file
    -   This command creates a new file. If content is provided it is placed in
    -   the new file. If the path is absolute then the given path is unmodified
    -   and the file is attempted to be created using that path. Otherwise the
    -   path is prepended with the current working directory path and this new
    -   path is used to create the file. If the file already exists its original
    -   content is erased.
