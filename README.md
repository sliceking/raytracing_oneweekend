# raytracing in one weekend

[book](https://raytracing.github.io/books/RayTracingInOneWeekend.html)

## how to do it

This is my first time using c++, so i tried following the book, which said to use cmake.
This made a bunch of folders and files after creating an empty cmakelist file.
I didnt see an executable so what i did instead was use g++ and that worked.

Made using g++ (GCC) 15.2.1 20251112

`cd src`
`g++ main.cc -o rayray`
`./rayray >> ok.ppm`

there are some output files i've included, they can be viewed in gimp or anything
else i did experiment with [other implementations](https://github.com/eduardohenriquearnold/raytracer) for using multiple cores, and
using cuda, to see how things ran on my laptops gpu.

The standalone run took hours to render the final image. Using the multicore
setup using openmp it took a couple minutes, and the cuda setup with the
[correct flags](https://gt3rs.medium.com/compile-with-nvcc-3566fbdfdbf) for my GPU took seconds. I didnt measure the runs but my
description of how long it took has enough context.

The final run was: running it, falling asleep, waking up in the middle
of the night to see it still running, waking up in the morning
to it being completed. Compared to the GPU run where it finished in under
5 seconds, wow.

There are a lot more keywords and syntax in c++ when compared to go,
which is what I'm more comfortable using. Hunting down performance
optimizations and secret flags does make you feel smart. I've
always been scared to touch c or c++ because i've heard about its
lack of tooling, the comforts of a garbage collector and test
frameworks built into the language are absent. This experience of rendering
3d scenes and watching the performance boosts is amazing and not
as scary as I imagined.
