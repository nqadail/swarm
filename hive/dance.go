package hive

type Dance struct {
	length int
	code   int
	bee    *Bee
}

func NewDance(b *Bee, length, code int) *Dance {
	d := new(Dance)
	d.length = length
	d.code = code
	d.bee = b
	return d
}

func (d *Dance) Bee() *Bee   { return d.bee }
func (d *Dance) Code() int   { return d.code }
func (d *Dance) Length() int { return d.length }
