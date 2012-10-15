package main

// t = top, b = bottom, l = left, r = right
type box struct {
	tl int
	tr int
	bl int
	br int
}

type point struct {
	x int
	y int
}

type event string // TODO

// Player
type player struct {
	name string
	// control    gameio
	cash     int
	loan     int
	tenants1 int
	tenants2 int
	tenants3 int
	tenants4 int
	tenants5 int
}

// Estate
type estate struct {
	name  string
	owner *player
	box   box
}

// Game map
type gamemap struct {
	name       string
	estates    []estate
	background string // TODO
}

// Family
type familyProto struct {
	life            int
	rentPart        float32
	rentMax         int
	tenancyPeriod   int
	stressThreshold int
}

type familyProduction int

const (
	familyProdInvalid familyProduction = iota
	familyProdTenant1
	familyProdTenant2
	familyProdTenant3
	familyProdTenant4
	familyProdTenant5
	familyProdWorker
	familyProdCadet
	familyProdFavor
)

type family struct {
	proto      familyProto
	happiness  int
	production familyProduction
	queue      []familyProduction
}

// Building
type buildingKind int

const (
	buildingKindInvalid buildingKind = iota
	buildingKindHQ
	buildingKindWoodYard
	buildingKindHouse1
	buildingKindHouse2
	buildingKindCementYard
)

type building struct {
	border     box
	box        box
	kind       buildingKind
	owner      *player
	occupation *player
	rentedBy   *family
}

// Game unit
type unitPosition struct {
	inside bool
	in     *building
	out    *point
}

type unitKind int

const (
	unitKindInvalid unitKind = iota
	unitKindWorker
	unitKindForeman
	unitKindGangster
	unitKindRepairman
	unitKindPoliceman
	unitKindDeliveryman
)

type unitState int

const (
	unitStateInvalid unitState = iota
	unitStateMoving
	unitStateActing
)

type unitAction string // FIXME

type unit struct {
	pos    unitPosition
	kind   unitKind
	loyal  *player
	state  unitState
	health int
	queue  []unitAction
}

// Game
type game struct {
	goal      event
	players   []player
	map_      gamemap
	units     []unit
	buildings []building
}
