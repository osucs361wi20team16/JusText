# JusText
JusText text editor Agile project for CS361

## Building

In the main folder run:

```
$ go build -o justext cmd/justext/main.go
```

This creates the `justext` executable.

## Usage

 |-------------------------------------- \
 \| File   Edit                          \
 |-------------------------------------- \
 |                                       \
 |                                       \
 |         < File editing area >         \
 |                                       \
 |-------------------------------------- \
 \| < Status messages >                  \
 |-------------------------------------- \

### Menu bar

Press the `esc` key to switch context to the menu bar. Use the `tab` key to alternate between tabs. Press `enter` or an arrow, and arrow to the option you wish to select, and press enter.

To exit the menu bar, press `esc` again.

### Files
Open a file with `justext <filename>` or start a new file with `justext`.

### Editing
Edit the file in the File editing area just like any other editing application.

### Saving
Save your file using the file menu. 
