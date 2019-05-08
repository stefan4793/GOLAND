package main

import "fmt"
import "math"

type Fractie struct {
	Numarator int
	Numitor int
}

func cmmdc (x int, y int) int{
	for x % y != 0{
		r:= x%y
		x=y
		y=r
	}
	return y
}

func(f *Fractie) simplif(){
	c:=cmmdc(f.Numarator, f.Numitor)
	f.Numarator /= c
	f.Numitor /= c
}

func (f Fractie) suma(f1 Fractie) Fractie{
	s:= Fractie{Numitor: f1.Numitor*f.Numarator + f.Numitor * f1.Numarator, Numarator: f.Numitor* f1.Numitor}
	s.simplif()
	return s
}

func (f Fractie) dif(f1 Fractie) Fractie{
	s:= Fractie{Numitor: f1.Numitor*f.Numarator - f.Numitor * f1.Numarator, Numarator: f.Numitor* f1.Numitor}
	s.simplif()
	return s
}

func (f Fractie) prod(f1 Fractie) Fractie{
	s:= Fractie{Numitor: f.Numitor * f1.Numitor, Numarator: f.Numarator * f.Numarator}
	s.simplif()
	return s
}

func (f Fractie) cat(f1 Fractie) Fractie{
	s:= Fractie{Numitor: f.Numitor * f1.Numarator, Numarator: f.Numarator * f1.Numitor}
	s.simplif()
	return s
}

func (f Fractie) getValue() float64{
	return float64(f.Numarator)/float64(f.Numitor)
}

func fact(n int) int{
	p:=1
	for i:=1; i<=n;{
		p*=i
		i++
	}
	return p
}

func e (eps float64) float64{
	t:= 1.0
	n:= 1
	for true {
		t1:= t
		t+= Fractie{Numarator: 1, Numitor: fact(n)}.getValue()
		if math.Abs(t - t1) < eps{
			return t
		}
		n++
	}
	return -1
}

func main(){
	//f = new (Fractie)
	f:=Fractie{Numarator: 5 , Numitor: 10 }
	f.simplif()
	//fmt.Println(f)
	g:= Fractie{Numitor: 4, Numarator: 6}
	fmt.Println(f.suma(g))
	fmt.Println(f.dif(g))
	fmt.Println(f.prod(g))
	fmt.Println(f.cat(g).getValue())
	fmt.Println(e(1e-9))
}



