# Boids simulation in Go (Golang)

[Boids](https://en.wikipedia.org/wiki/Boids) is an artificial life program, developed by Craig Reynolds in 1986, which simulates the flocking behaviour of birds.

### How do boids work?
Each bird attempts to maintain a reasonable amount of distance between itself and any nearby birds, to prevent overcrowding. Birds try to change their position so that it corresponds with the average alignment of other nearby birds. Every bird attempts to move towards the average position of other nearby birds.

[Source](https://cs.stanford.edu/people/eroberts/courses/soco/projects/2008-09/modeling-natural-systems/boids.html)