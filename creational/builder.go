package creational

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetStructure().SetWheels().SetSeats()
}
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

type VehicleProduct struct {
	Wheels int
	Seats int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "car"
	return c
}
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}
func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 3
	return c
}
func (c *BikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "bike"
	return c
}
func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BusBuilder struct {
	v VehicleProduct
}
func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 4*2
	return b
}
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "bus"
	return b
}
func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}