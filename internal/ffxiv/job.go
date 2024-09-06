package ffxiv

type Job int

const (
	GNB Job = iota
	PLD
	GLD
	DRK
	WAR
	MRD
	SCH
	ACN // Arcanist
	SGE
	AST
	WHM
	CNJ
	SAM
	DRG
	NIN
	MNK
	RPR
	VPR
	BRD
	MCH
	DNC
	BLM
	BLU
	SMN
	PCT
	RDM
	LNC
	PUG
	ROG
	THM
	ARC // Archer
	Unknown
)

func JobFromAbbreviation(abbreviation string) Job {
	switch abbreviation {
	case "GNB":
		return GNB
	case "PLD":
		return PLD
	case "GLD":
		return GLD
	case "DRK":
		return DRK
	case "WAR":
		return WAR
	case "MRD":
		return MRD
	case "SCH":
		return SCH
	case "ACN":
		return ACN
	case "SGE":
		return SGE
	case "AST":
		return AST
	case "WHM":
		return WHM
	case "CNJ":
		return CNJ
	case "SAM":
		return SAM
	case "DRG":
		return DRG
	case "NIN":
		return NIN
	case "MNK":
		return MNK
	case "RPR":
		return RPR
	case "VPR":
		return VPR
	case "BRD":
		return BRD
	case "MCH":
		return MCH
	case "DNC":
		return DNC
	case "BLM":
		return BLM
	case "BLU":
		return BLU
	case "SMN":
		return SMN
	case "PCT":
		return PCT
	case "RDM":
		return RDM
	case "LNC":
		return LNC
	case "PUG":
		return PUG
	case "ROG":
		return ROG
	case "THM":
		return THM
	case "ARC":
		return ARC
	}
	return Unknown
}

func (j Job) Emoji() string {
	switch j {
	case GNB:
		return "<:GBR:1281650873684066304>"
	case PLD:
		return "<:PLD:1281651108737187912>"
	case GLD:
		return "<:GLD:1281650863802155054>"
	case DRK:
		return "<:DRK:1281650805945925713>"
	case WAR:
		return "<:WAR:1281651224671682631>"
	case MRD:
		return "<:MRD:1281651025140252824>"
	case SCH:
		return "<:SCH:1281651086100267110>"
	case ACN:
		return "<:ACN:1281650544372486154>"
	case SGE:
		return "<:SGE:1281651153611788308>"
	case AST:
		return "<:AST:1281650670772031558>"
	case WHM:
		return "<:WHM:1281651252303888414>"
	case CNJ:
		return "<:CNJ:1281650789718425630>"
	case SAM:
		return "<:SAM:1281651097370366044>"
	case DRG:
		return "<:DRG:1281650816259981433>"
	case NIN:
		return "<:NIN:1281650916503846934>"
	case MNK:
		return "<:MNK:1281650902117122059>"
	case RPR:
		return "<:RPR:1281651133244244063>"
	case VPR:
		return "<:VPR:1281651198348496999>"
	case BRD:
		return "<:BRD:1281650709577597112>"
	case MCH:
		return "<:MCH:1281650892705370247>"
	case DNC:
		return "<:DNC:1281650853006278690>"
	case BLM:
		return "<:BLM:1281650765944979520>"
	case BLU:
		return "<:BLU:1281650731086118922>"
	case SMN:
		return "<:SMN:1281651161996197968>"
	case RDM:
		return "<:RDM:1281651075803250688>"
	case PCT:
		return "<:PCT:1281651054265503806>"
	case LNC:
		return "<:LNC:1281650884161568788>"
	case PUG:
		return "<:PGL:1281651066076926042>"
	case ROG:
		return "<:ROG:1281651143084343318>"
	case THM:
		return "<:THM:1281651187619201209>"
	case ARC:
		return "<:ARC:1281650658738438265>"
	}
	return "<:DOH:1281650835629281472>"
}
