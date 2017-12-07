var blob = `
[[parts]]
type = "valve"
id = "valve-1"
size = 1.2
rating = 4

[[parts]]
type = "valve"
id = "valve-2"
size = 2.1
rating = 5

[[parts]]
type = "pipe"
id = "pipe-1"
length = 2.1
diameter = 12

[[parts]]
type = "cable"
id = "cable-1"
length = 12
rating = 3.1
`
o := &order{}
err := Unmarshal([]byte(blob), o)
if err != nil {
    log.Fatal(err)
}

fmt.Println(len(o.parts))

for _, part := range o.parts {
    fmt.Println(part.Name())
}

// Code to implement UmarshalJSON.

// type order struct {
// 	// NOTE `order.parts` is a private slice of type `part` which is an
// 	// interface and may only be loaded from toml using the
// 	// UnmarshalTOML() method of the Umarshaler interface.
// 	parts parts
// }

// func (o *order) UnmarshalTOML(data interface{}) error {

// 	// NOTE the example below contains detailed type casting to show how
// 	// the 'data' is retrieved. In operational use, a type cast wrapper
// 	// may be prefered e.g.
// 	//
// 	// func AsMap(v interface{}) (map[string]interface{}, error) {
// 	// 		return v.(map[string]interface{})
// 	// }
// 	//
// 	// resulting in:
// 	// d, _ := AsMap(data)
// 	//

// 	d, _ := data.(map[string]interface{})
// 	parts, _ := d["parts"].([]map[string]interface{})

// 	for _, p := range parts {

// 		typ, _ := p["type"].(string)
// 		id, _ := p["id"].(string)

// 		// detect the type of part and handle each case
// 		switch p["type"] {
// 		case "valve":

// 			size := float32(p["size"].(float64))
// 			rating := int(p["rating"].(int64))

// 			valve := &valve{
// 				Type:   typ,
// 				ID:     id,
// 				Size:   size,
// 				Rating: rating,
// 			}

// 			o.parts = append(o.parts, valve)

// 		case "pipe":

// 			length := float32(p["length"].(float64))
// 			diameter := int(p["diameter"].(int64))

// 			pipe := &pipe{
// 				Type:     typ,
// 				ID:       id,
// 				Length:   length,
// 				Diameter: diameter,
// 			}

// 			o.parts = append(o.parts, pipe)

// 		case "cable":

// 			length := int(p["length"].(int64))
// 			rating := float32(p["rating"].(float64))

// 			cable := &cable{
// 				Type:   typ,
// 				ID:     id,
// 				Length: length,
// 				Rating: rating,
// 			}

// 			o.parts = append(o.parts, cable)

// 		}
// 	}

// 	return nil
// }

// type parts []part

// type part interface {
// 	Name() string
// }

// type valve struct {
// 	Type   string
// 	ID     string
// 	Size   float32
// 	Rating int
// }

// func (v *valve) Name() string {
// 	return fmt.Sprintf("VALVE: %s", v.ID)
// }

// type pipe struct {
// 	Type     string
// 	ID       string
// 	Length   float32
// 	Diameter int
// }

// func (p *pipe) Name() string {
// 	return fmt.Sprintf("PIPE: %s", p.ID)
// }

// type cable struct {
// 	Type   string
// 	ID     string
// 	Length int
// 	Rating float32
// }

// func (c *cable) Name() string {
// 	return fmt.Sprintf("CABLE: %s", c.ID)
// }
