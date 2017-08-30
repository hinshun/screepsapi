package screepstype

type BodyPart struct {
	Type BodyPartType `json:"type"`
	Hits int          `json:"hits"`
}

type Controller struct {
	Object
	UserOwned

	Level             int `json:"level"`
	Progress          int `json:"progress"`
	DowngradeTime     int `json:"downgradeTime"`
	SafeMode          int `json:"safeMode"`
	SafeModeAvailable int `json:"safeModeAvailable"`
}

type Creep struct {
	Object
	Destructible
	EnergyStore
	UserOwned

	Name      string         `json:"name"`
	Body      []BodyPart     `json:"body"`
	Spawning  bool           `json:"spawning"`
	Fatigue   int            `json:"fatigue"`
	AgeTime   int            `json:"ageTime"`
	ActionLog CreepActionLog `json:"actionLog"`
}

type CreepActionLog struct {
	Attacked          *Point `json:"attacked"`
	Healed            *Point `json:"healed"`
	Attack            *Point `json:"attack"`
	RangedAttack      *Point `json:"rangedAttacked"`
	RangedMassAttack  *Point `json:"rangedMassAttacked"`
	RangedHeal        *Point `json:"rangedHeal"`
	Harvest           *Point `json:"harvest"`
	Heal              *Point `json:"heal"`
	Repair            *Point `json:"repair"`
	Build             *Point `json:"build"`
	Say               *Point `json:"say"`
	UpgradeController *Point `json:"upgradeController"`
	ReserveController *Point `json:"reserveController"`
}

type Destructible struct {
	NotifyWhenAttacked bool `json:"notifyWhenAttacked"`
	Hits               int  `json:"hits"`
	HitsMax            int  `json:"hitsMax"`
}

type Energy struct {
	Energy int `json:"energy"`
	// TODO(hinshun): figure out if this is deprecated/unnecessary
	ResourceType string `json:"resourceType"`
}

type EnergyStore struct {
	Energy         int `json:"energy"`
	EnergyCapacity int `json:"energyCapacity"`
}

type Extension struct {
	Object
	Destructible
	EnergyStore
	UserOwned

	Off bool `json:"off"`
}

type Mineral struct {
	Object

	Density       int         `json:"density"`
	MineralType   MineralType `json:"mineralType"`
	MineralAmount int         `json:"mineralAmount"`
}

type Object struct {
	ID   string     `json:"_id"`
	Room string     `json:"room"`
	Type ObjectType `json:"type"`
	X    int        `json:"x"`
	Y    int        `json:"y"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Road struct {
	Object
	Destructible

	NextDecayTime int `json:nextDecayTime`
}

type Source struct {
	Object
	EnergyStore

	TicksToRegeneration  int `json:"ticksToRegeneration"`
	InvaderHarvested     int `json:"invaderHarvested"`
	NextRegenerationTime int `json:"nextRegenerationTime"`
}

type Spawn struct {
	Object
	Destructible
	EnergyStore
	UserOwned

	Name     string `json:"name"`
	Spawning bool   `json:"spawning"`
	Off      bool   `json:"off"`
}

type Storage struct {
	Object
	Destructible
	EnergyStore
	UserOwned
}

type Tower struct {
	Object
	Destructible
	EnergyStore
	UserOwned

	ActionLog TowerActionLog `json:"actionLog"`
}

type TowerActionLog struct {
	Attack *Point `json:"attack"`
	Heal   *Point `json:"heal"`
	Repair *Point `json:"repair"`
}

type Wall struct {
	Object
	Destructible
}

type UserOwned struct {
	User string `json:"user"`
}
