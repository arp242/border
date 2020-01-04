border is a simple commandline tool to add a border around PNG images.

When posting screenshots it's often useful to add a little border around the
edges so it's easier to see where the screenshot ends. This tool does exactly
that.

Install with `go get arp242.net/border`, which will put the binary in
`~/go/bin/border`.

Use `-border` to control the border width (default: 2) and `-color` to control
the colour (default: #999). The result is written to the original filename with
the value of `-write` inserted before the extension default: `_border`). Use
`-write ''` to overwrite the original.

Note: images are always saved as RGBA; you'll have to manually regenerate the
palette if the input was indexed or greyscale.

Alternatively, you can use ImageMagick:

    $ convert -border 2x2 -bordercolor '#999999' example.png example_border.png

Example:

![example.png](https://raw.githubusercontent.com/arp242/border/master/example.png)
![example_border.png](https://raw.githubusercontent.com/arp242/border/master/example_border.png)
