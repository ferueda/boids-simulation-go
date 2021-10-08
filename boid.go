package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	id       int
	position Vector2D
	velocity Vector2D
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddV((viewRadius)), b.position.AddV(-viewRadius)
	avgPos, avgVel, sep := Vector2D{0, 0}, Vector2D{0, 0}, Vector2D{0, 0}
	var count float64

	rwLock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
					count++
					avgVel = avgVel.Add(boids[otherBoidId].velocity)
					avgPos = avgPos.Add(boids[otherBoidId].position)
					sep = sep.Add(b.position.Subtract(boids[otherBoidId].position).DivideV(dist))
				}
			}
		}
	}
	rwLock.RUnlock()

	accel := Vector2D{b.borderBounce(b.position.x, screenWidth), b.borderBounce(b.position.y, screenHeight)}
	if count > 0 {
		avgPos, avgVel = avgPos.DivideV(count), avgVel.DivideV(count)
		accelAlignment := avgVel.Subtract(b.velocity).MultiplyV(adjRate)
		accelCohesion := avgPos.Subtract(b.position).MultiplyV(adjRate)
		accelSep := sep.MultiplyV(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSep)
	}
	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne() {
	accel := b.calcAcceleration()
	rwLock.Lock()
	b.velocity = b.velocity.Add(accel).Limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	nextPos := b.position.Add(b.velocity)
	b.position = nextPos
	boidMap[int(b.position.x)][int(b.position.y)] = b.id

	rwLock.Unlock()
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
		velocity: Vector2D{x: (rand.Float64() * 2) - 1, y: (rand.Float64() * 2) - 1},
	}

	boids[id] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	go b.start()
}
