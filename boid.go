package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	id       int
	position Vector2D
	velocity Vector2D
}

func (b *Boid) moveOne() {
	nextPos := b.position.Add(b.velocity)
	b.position = nextPos

	// If next position is bigger than screen size, change velocity direction
	if nextPos.x >= screenWidth || nextPos.x < 0 {
		b.velocity = Vector2D{x: -b.velocity.x, y: b.velocity.y}
	}

	if nextPos.y >= screenHeight || nextPos.y < 0 {
		b.velocity = Vector2D{x: b.velocity.x, y: -b.velocity.y}
	}
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(id int) {
	b := Boid{
		id:       id,
		position: Vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		velocity: Vector2D{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
	}

	boids[id] = &b
	go b.start()
}
