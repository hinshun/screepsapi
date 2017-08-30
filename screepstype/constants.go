package screepstype

type ActiveName string

const (
	ActiveNameWorld ActiveName = "activeWorld"
	ActiveNameSim              = "activeSim"
)

type BodyPartType string

const (
	BodyPartTypeMove   BodyPartType = "move"
	BodyPartTypeWork                = "work"
	BodyPartTypeCarry               = "carry"
	BodyPartTypeAttack              = "attack"
	// TODO(hinshun): confirm camelCase or snake_case
	BodyPartTypeRangedAttack = "rangedAttack"
	BodyPartTypeHeal         = "heal"
	BodyPartTypeClaim        = "claim"
	BodyPartTypeTough        = "tough"
)

type Color int

const (
	ColorNone Color = iota
	ColorRed
	ColorPurple
	ColorBlue
	ColorCyan
	ColorGreen
	ColorYellow
	ColorOrange
	ColorBrown
	ColorGrey
	ColorWhite
)

// CompressionType is an internal constant used to differentiate compression
// from gzipped data from API and zlib-gzipped data from the websocket API.
type CompressionType string

const (
	CompressionTypeGzip CompressionType = "gzip"
	CompressionTypeZlib                 = "zlib"
)

type Error string

const (
	ErrorNone          Error = ""
	ErrorInvalidParams       = "invalid params"
	ErrorInvalidStatus       = "invalid status"
)

type MarketOrderType string

const (
	MarketOrderTypeBuy  MarketOrderType = "buy"
	MarketOrderTypeSell                 = "sell"
)

type MessageDirection string

const (
	MessageDirectionIn  MessageDirection = "in"
	MessageDirectionOut                  = "out"
)

type MineralType string

const (
	MineralTypeHydrogen MineralType = "H"
)

type MoneyOrderType string

const (
	// TODO(hinshun): verify values
	MoneyOrderTypeNew         MoneyOrderType = "new"
	MoneyOrderTypeExtended                   = "extended"
	MoneyOrderTypeFulfilled                  = "fulfilled"
	MoneyOrderTypePriceChange                = "priceChange"
)

type ObjectIntent string

const (
	ObjectIntentRemove            ObjectIntent = "remove"
	ObjectIntentSuicideCreep                   = "suicide"
	ObjectIntentUnclaimController              = "unclaim"
	ObjectIntentDestroyStructure               = "destroyStructure"
)

type ObjectType string

const (
	ObjectTypeController ObjectType = "controller"
	ObjectTypeCreep                 = "creep"
	ObjectTypeEnergy                = "energy"
	ObjectTypeExtension             = "extension"
	ObjectTypeMineral               = "mineral"
	ObjectTypeRoad                  = "road"
	ObjectTypeSource                = "source"
	ObjectTypeSpawn                 = "spawn"
	ObjectTypeStorage               = "storage"
	ObjectTypeTower                 = "tower"
	ObjectTypeWall                  = "constructedWall"
)

type RoomStatus string

const (
	RoomStatusNormal       RoomStatus = "normal"
	RoomStatusOutOfBorders            = "out of borders"
)

type StatName string

const (
	StatNameNone               StatName = ""
	StatNameOwner                       = "owner0"
	StatNameClaim                       = "claim0"
	StatNameCreepsLost                  = "creepsLost"
	StatNameCreepsProduced              = "creepsProduced"
	StatNameEnergyConstruction          = "energyConstruction"
	StatNameEnergyControl               = "energyControl"
	StatNameEnergyCreeps                = "energyCreeps"
	StatNameEnergyHarvested             = "energyHarvested"
)

type StatsPeriod string

const (
	StatsPeriodNone    StatsPeriod = ""
	StatsPeriod1Hour               = "8"
	StatsPeriod24Hours             = "180"
	StatsPeriod7Days               = "1440"
)

type Structure string

const (
	StructureSpawn       Structure = "spawn"
	StructureExtension             = "extension"
	StructureRoad                  = "road"
	StructureWall                  = "constructedWall"
	StructureRampart               = "rampart"
	StructureKeeper_lair           = "keeperLair"
	StructurePortal                = "portal"
	StructureController            = "controller"
	StructureLink                  = "link"
	StructureStorage               = "storage"
	StructureTower                 = "tower"
	StructureObserver              = "observer"
	StructurePower_bank            = "powerBank"
	StructurePower_spawn           = "powerSpawn"
	StructureExtractor             = "extractor"
	StructureLab                   = "lab"
	StructureTerminal              = "terminal"
	StructureContainer             = "container"
	StructureNuker                 = "nuker"
)

type Terrain string

const (
	TerrainPlain    Terrain = "plain"
	TerrainWall             = "wall"
	TerrainSwamp            = "swamp"
	TerrainAlsoWall         = "also wall"
)

type UpsertMessageDataType string

const (
	UpsertMessageDataTypeBuffer UpsertMessageDataType = "Buffer"
)

type UniqueObjectName string

const (
	UniqueObjectNameFlag  UniqueObjectName = "flag"
	UniqueObjectNameSpawn                  = "spawn"
)

type WorldStatus string

const (
	WorldStatusEmpty  WorldStatus = "empty"
	WorldStatusNormal             = "normal"
)
