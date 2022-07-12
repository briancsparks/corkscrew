
# TODO

## Mandelbrot

Status:

* Just put `both` library in, and hooked up the `mandelboth` command
  * MUCH faster (0.25 sec vs. 3.5sec)
  * Works much better for the Work/Display duality
* ~~Lost axis~~
* Lost work splitting
  * Look at both.split()

Features:

* Look into using https://github.com/kvartborg/vector
  * For vectors
* ~~Concurrent-ize the compute loop~~
* Zoom

Clean up:

* More general way to send the tiles to Joe
  * Right now, Joe has one hard-coded 'main' tile, and a list of tiles,
    the field also has a main function to call to draw, and then a 'final' one.
* Stop using the p5 global object - use one that is allocated and started by us.

Fix:

*

