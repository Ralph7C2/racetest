package main

type thing struct{}

var tt = &thing{}

func (t *thing) DoIt() {}

func GlobalCaller() {
	tt.DoIt()
}

func main() {
	for i := 0; i < 10000; i++ {
		i := i
		go func() {
			if i%2 == 0 {
				tt = &thing{}
			} else {
				GlobalCaller()
			}
		}()
	}
}
