package main

import "fmt"

func main()  {
 chest := Chest{"Small chest"}
 enemy := Enemy{name: "Scary small skeleton"}
 adaptedEnemy := AdaptedEnemy{Enemy: &enemy}

 destryableUnitList := []DistroyableUnit{&chest, &adaptedEnemy}

 for _, unit := range destryableUnitList {
  unit.Disappear()
 }

}

type DistroyableUnit interface {
 Disappear()
}

type Enemy struct {
 name string
}

func (s *Enemy) Die() {
 fmt.Printf("%s dies \n", s.name)
}

type Chest struct {
 name string
}

func (c *Chest) Disappear() {
 fmt.Printf("%s disappears \n", c.name)
}

type AdaptedEnemy struct {
 *Enemy
}

func (ae *AdaptedEnemy) Disappear() {
 ae.Die()
}