package data

// AssemblyMemberData contains the names and email addresses of California State Assembly Members by district.
// Fetched November 2024 from https://admin.cdn.sos.ca.gov/ca-roster/2025/state-assembly.pdf
// Ought to be updated after elections

// District represents a California State Assembly District, which range from 1 to 80.
type District int

type AssemblyMember struct {
	Name  string
	Email string
}

var AssemblyMemberData = map[District]AssemblyMember{
	1: {
		Name:  "Heather Hadwick",
		Email: "assemblymember.hadwick@asm.ca.gov",
	},
	2: {
		Name:  "Chris Rogers",
		Email: "assemblymember.rogers@asm.ca.gov",
	},
	3: {
		Name:  "James Gallagher",
		Email: "assemblymember.gallagher@asm.ca.gov",
	},
	4: {
		Name:  "Cecilia Aguiar-Curry",
		Email: "assemblymember.aguiar-curry@asm.ca.gov",
	},
	5: {
		Name:  "Joe Patterson",
		Email: "assemblymember.joepatterson@asm.ca.gov",
	},
	6: {
		Name:  "Maggy Krell",
		Email: "assemblymember.krell@asm.ca.gov",
	},
	7: {
		Name:  "Josh Hoover",
		Email: "assemblymember.hoover@asm.ca.gov",
	},
	8: {
		Name:  "David Tangipa",
		Email: "assemblymember.tangipa@asm.ca.gov",
	},
	9: {
		Name:  "Heath Flora",
		Email: "assemblymember.flora@asm.ca.gov",
	},
	10: {
		Name:  "Stephanie Nguyen",
		Email: "assemblymember.nguyen@asm.ca.gov",
	},
	11: {
		Name:  "Lori Wilson",
		Email: "assemblymember.wilson@asm.ca.gov",
	},
	12: {
		Name:  "Damon Connolly",
		Email: "assemblymember.connolly@asm.ca.gov",
	},
	13: {
		Name:  "Rhodesia Ransom",
		Email: "assemblymember.ransom@asm.ca.gov",
	},
	14: {
		Name:  "Buffy Wicks",
		Email: "assemblymember.wicks@asm.ca.gov",
	},
	15: {
		Name:  "Anamarie Avila Farias",
		Email: "assemblymember.faris@asm.ca.gov",
	},
	16: {
		Name:  "Rebecca Bauer-Kahan",
		Email: "assemblymember.bauer-kahan@asm.ca.gov",
	},
	17: {
		Name:  "Matt Haney",
		Email: "assemblymember.haney@asm.ca.gov",
	},
	18: {
		Name:  "Mia Bonta",
		Email: "assemblymember.bonta@asm.ca.gov",
	},
	19: {
		Name:  "Catherine Stefani",
		Email: "assemblymember.stefani@asm.ca.gov",
	},
	20: {
		Name:  "Liz Ortega",
		Email: "assemblymember.ortega@asm.ca.gov",
	},
	21: {
		Name:  "Diane Papan",
		Email: "assemblymember.papan@asm.ca.gov",
	},
	22: {
		Name:  "Juan Alanis",
		Email: "assemblymember.alanis@asm.ca.gov",
	},
	23: {
		Name:  "Marc Berman",
		Email: "assemblymember.berman@asm.ca.gov",
	},
	24: {
		Name:  "Alex Lee",
		Email: "assemblymember.lee@asm.ca.gov",
	},
	25: {
		Name:  "Ash Kalra",
		Email: "assemblymember.kalra@asm.ca.gov",
	},
	26: {
		Name:  "Aherns Patrick",
		Email: "assemblymember.aherns@asm.ca.gov",
	},
	27: {
		Name:  "Esmeralda Soria",
		Email: "assemblymember.soria@asm.ca.gov",
	},
	28: {
		Name:  "Gail Pellerin",
		Email: "assemblymember.pellerin@asm.ca.gov",
	},
	29: {
		Name:  "Robert Rivas",
		Email: "assemblymember.rivas@asm.ca.gov",
	},
	30: {
		Name:  "Dawn Addis",
		Email: "assemblymember.addis@asm.ca.gov",
	},
	31: {
		Name:  "Dr. Joaquin Arambula",
		Email: "assemblymember.arambula@asm.ca.gov",
	},
	32: {
		Name:  "Stan Ellis",
		Email: "assemblymember.ellis@asm.ca.gov",
	},
	33: {
		Name:  "Alexandria Macedo",
		Email: "assemblymember.macedo@asm.ca.gov",
	},
	34: {
		Name:  "Tom Lackey",
		Email: "assemblymember.lackey@asm.ca.gov",
	},
	35: {
		Name:  "Dr. Jasmeet Kaur Bains",
		Email: "assemblymember.bains@asm.ca.gov",
	},
	36: {
		Name:  "Jeff Gonzalez",
		Email: "assemblymember.jeffgonzalez@asm.ca.gov",
	},
	37: {
		Name:  "Gregg Hart",
		Email: "assemblymember.hart@asm.ca.gov",
	},
	38: {
		Name:  "Steve Bennett",
		Email: "assemblymember.bennett@asm.ca.gov",
	},
	39: {
		Name:  "Juan Carrillo",
		Email: "assemblymember.juancarillo@asm.ca.gov",
	},
	40: {
		Name:  "Pilar Schiavo",
		Email: "assemblymember.pilar@asm.ca.gov",
	},
	41: {
		Name:  "Matt Harabedian",
		Email: "assemblymember.harabedian@asm.ca.gov",
	},
	42: {
		Name:  "Jacqui Irwin",
		Email: "assemblymember.irwin@asm.ca.gov",
	},
	43: {
		Name:  "Celeste Rodriguez",
		Email: "assemblymember.celesterodriguez@asm.ca.gov",
	},
	44: {
		Name:  "Nick Schultz",
		Email: "assemblymember.schultz@asm.ca.gov",
	},
	45: {
		Name:  "James Ramos",
		Email: "assemblymember.ramos@asm.ca.gov",
	},
	46: {
		Name:  "Jesse Gabriel",
		Email: "assemblymember.gabriel@asm.ca.gov",
	},
	47: {
		Name:  "Greg Wallis",
		Email: "assemblymember.wallis@asm.ca.gov",
	},
	48: {
		Name:  "Blanca Rubio",
		Email: "assemblymember.rubio@asm.ca.gov",
	},
	49: {
		Name:  "Mike Fong",
		Email: "assemblymember.mikefong@asm.ca.gov",
	},
	50: {
		Name:  "Robert Garcia",
		Email: "assemblymember.garcia@asm.ca.gov",
	},
	51: {
		Name:  "Rick Chavez Zbur",
		Email: "assemblymember.zbur@asm.ca.gov",
	},
	52: {
		Name:  "Jessica Caloza",
		Email: "assemblymember.caloza@asm.ca.gov",
	},
	53: {
		Name:  "Michelle Rodriguez",
		Email: "assemblymember.michellerodriguez@asm.ca.gov",
	},
	54: {
		Name:  "Mark González",
		Email: "assemblymember.markgonzalez@asm.ca.gov",
	},
	55: {
		Name:  "Isaac Bryan",
		Email: "assemblymember.bryan@asm.ca.gov",
	},
	56: {
		Name:  "Lisa Calderon",
		Email: "assemblymember.calderon@asm.ca.gov",
	},
	57: {
		Name:  "Sade Elhawary",
		Email: "assemblymember.elhawary@asm.ca.gov",
	},
	58: {
		Name:  "Leticia Castillo",
		Email: "assemblymember.castillo@asm.ca.gov",
	},
	59: {
		Name:  "Phillip Chen",
		Email: "assemblymember.chen@asm.ca.gov",
	},
	60: {
		Name:  "Dr. Corey Jackson",
		Email: "assemblymember.jackson@asm.ca.gov",
	},
	61: {
		Name:  "Tina McKinnor",
		Email: "assemblymember.mckinnor@asm.ca.gov",
	},
	62: {
		Name:  "Jr. Jose Luis Solache",
		Email: "assemblymember.solache@asm.ca.gov",
	},
	63: {
		Name:  "Natasha Johnson",
		Email: "assemblymember.johnson@assembly.ca.gov",
	},
	64: {
		Name:  "Blanca Pacheco",
		Email: "assemblymember.pacheco@asm.ca.gov",
	},
	65: {
		Name:  "Mike A. Gipson",
		Email: "assemblymember.gipson@asm.ca.gov",
	},
	66: {
		Name:  "Al Muratsuchi",
		Email: "assemblymember.muratsuchi@assembly.ca.gov",
	},
	67: {
		Name:  "Sharon Quirk-Silva",
		Email: "assemblymember.quirk-silva@asm.ca.gov",
	},
	68: {
		Name:  "Avelino Valencia",
		Email: "assemblymember.valencia@asm.ca.gov",
	},
	69: {
		Name:  "Josh Lowenthal",
		Email: "assemblymember.lowenthal@asm.ca.gov",
	},
	70: {
		Name:  "Tri Ta",
		Email: "assemblymember.ta@asm.ca.gov",
	},
	71: {
		Name:  "Kate Sanchez",
		Email: "assemblymember.sanchez@asm.ca.gov",
	},
	72: {
		Name:  "Diane Dixon",
		Email: "assemblymember.dixon@asm.ca.gov",
	},
	73: {
		Name:  "Cottie Petrie-Norris",
		Email: "assemblymember.petrienorris@asm.ca.gov",
	},
	74: {
		Name:  "Laurie Davies",
		Email: "assemblymember.davies@asm.ca.gov",
	},
	75: {
		Name:  "Carl DeMaio",
		Email: "assemblymember.demaio@asm.ca.gov",
	},
	76: {
		Name:  "Dr. Darshana Patel",
		Email: "assemblymember.patel@asm.ca.gov",
	},
	77: {
		Name:  "Tasha Boerner",
		Email: "assemblymember.boerner@asm.ca.gov",
	},
	78: {
		Name:  "Christopher Ward",
		Email: "assemblymember.ward@asm.ca.gov",
	},
	79: {
		Name:  "Dr. Lashae Sharp-Collins",
		Email: "assemblymember.sharp-collins@asm.ca.gov",
	},
	80: {
		Name:  "David Alvarez",
		Email: "assemblymember.alvarez@asm.ca.gov",
	},
}
