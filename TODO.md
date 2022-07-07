
# TODO

## Mandelbrot

Features:

* Concurrent-ize the compute loop
* Zoom

Clean up:

* More general way to send the tiles to Joe
  * Right now, Joe has one hard-coded 'main' tile, and a list of tiles,
    the field also has a main function to call to draw, and then a 'final' one.
* Stop using the p5 global object - use one that is allocated and started by us.

Fix:

*

