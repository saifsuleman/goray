package ray

import (
	"github.com/saifsuleman/goray/vec"
)

type Ray struct {
	Origin    vec.Vector
	Direction vec.Vector
}

func NewRay(origin vec.Vector, direction vec.Vector) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction.Normalize(),
	}
}

func (r *Ray) Cast(scene *Scene) *RayHit {
	var closest *RayHit

	for _, entity := range scene.Entities {
		intersection, hit := entity.CalculateIntersection(r)
		if !hit {
			continue
		}

		if closest == nil || r.Origin.Distance(intersection) < closest.Position.Distance(*closest.Entity.Position) {
			closest = &RayHit{
				Position: intersection,
				Normal:   entity.GetNormalAt(intersection),
				Entity:   entity,
				Ray:      r,
			}
		}
	}

	return closest
}

type RayHit struct {
	Position vec.Vector
	Normal   vec.Vector
	Entity   *Entity
	Ray      *Ray
}
