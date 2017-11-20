package data

// PlayerRecord defines the basic structure
// for a record in the collection of
// registered players
type PlayerRecord struct {
	ID       int
	Nick     string
	Password string
}

// PlayerRecords defines the type
// for a list of player records
type PlayerRecords []PlayerRecord

// Players is an instance of the
// PlayerRecords type and defines
// a few players to work with from
// the get-go
var Players = PlayerRecords{
	{
		ID:       1223145,
		Nick:     "phreak",
		Password: "c0mm4nd0",
	},
	{
		// Already connected
		ID:       8535253,
		Nick:     "fatal1ty",
		Password: "Quake4ever",
	},
	{
		// Already connected
		ID:       5457676,
		Nick:     "How4rd",
		Password: "Noriko<3",
	},
	{
		// Already connected
		ID:       9464779,
		Nick:     "Kr4zed",
		Password: "3328425913",
	},
	{
		// Already connected
		ID:       6735772,
		Nick:     "Sir3n",
		Password: "Teh n00b!",
	},
	{
		ID:       1159734,
		Nick:     "Kamikaze",
		Password: "Get'em",
	},
	{
		ID:       4648464,
		Nick:     "Lone_Hunter",
		Password: "SniperFtw",
	},
	{
		ID:       6992112,
		Nick:     "ne0phyte",
		Password: "star4748",
	},
	{
		// Already connected
		ID:       6433858,
		Nick:     "CrimsonDawn",
		Password: "necr0mancer",
	},
	{
		ID:       5747548,
		Nick:     "TheDamned1",
		Password: "f4llen1",
	},
	{
		ID:       1321878,
		Nick:     "SoulScorcher",
		Password: "Burn'em_all!",
	},
	{
		ID:       5723425,
		Nick:     "LittlePony",
		Password: "pink",
	},
	{
		// Already connected
		ID:       4148994,
		Nick:     "JigSaw",
		Password: "pure_3vil",
	},
	{
		// Already connected
		ID:       9661327,
		Nick:     "Camping_Gaz",
		Password: "somepassw0rd",
	},
	{
		// Already connected
		ID:       1412491,
		Nick:     "Dead3y3",
		Password: "5t4lk3r",
	},
	{
		// Already connected
		ID:       8712722,
		Nick:     "Tweety",
		Password: "Silvester",
	},
	{
		ID:       4219691,
		Nick:     "Mikky",
		Password: "Come|Get|Some",
	},
}
