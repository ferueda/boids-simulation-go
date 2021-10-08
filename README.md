# Boids simulation in Go (Golang)

[Boids](https://en.wikipedia.org/wiki/Boids) is an artificial life program, developed by Craig Reynolds in 1986, which simulates the flocking behaviour of birds.

## Implementation - Model boids using goroutines
* Using concurrent programming and goroutines, I'm modelling the flocking behavior of birds when they are in large groups. 
* Each goroutine models the behavior of a single bird.
* Each goroutine (bird) interact with each other according to specific rules.
* Graphics (2D) are implemented with the open source library [Ebiten](https://ebiten.org) (also developed in Go) by [Hajime Hoshi](https://github.com/hajimehoshi).

Check this [YouTube video](https://www.youtube.com/watch?v=V4f_1_r80RY) to better understand what the model is trying to simulate.

https://www.youtube.com/watch?v=V4f_1_r80RY


## How do boids work?
Each bird attempts to maintain a reasonable amount of distance between itself and any nearby birds, to prevent overcrowding. Birds try to change their position so that it corresponds with the average alignment of other nearby birds. Every bird attempts to move towards the average position of other nearby birds.



[Source](https://cs.stanford.edu/people/eroberts/courses/soco/projects/2008-09/modeling-natural-systems/boids.html)