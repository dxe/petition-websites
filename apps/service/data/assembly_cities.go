package data

type Municipality string

// File was created by script at scripts/create_city_assembly_map/main.py

// MunicipalityDistrictPercent maps municipality -> district_num -> percent.
// Municipalities are from US Census Designated Places and Cities in California.
// Districts are California State Assembly districts (1-80)
// Percent is the percentage of the municipality's population in that district.

// If a municipality sums to 200% it indicates there are two cities of that name
// Percentages of 0 can happen if a very fractional part of the population is in
// that district and the percentage rounds to 0. These are kept for completeness
// but unused.

var MunicipalityDistrictPercent = map[Municipality]map[District]float64{
	"Acalanes Ridge": {
		16: 100,
	},

	"Acampo": {
		9: 100,
	},

	"Acton": {
		34: 100,
	},

	"Adelanto": {
		34: 0,
		39: 100,
	},

	"Adin": {
		1: 100,
	},

	"Agoura Hills": {
		42: 100,
	},

	"Agua Dulce": {
		34: 100,
	},

	"Aguanga": {
		36: 100,
	},

	"Ahwahnee": {
		8: 100,
	},

	"Airport": {
		22: 100,
	},

	"Alameda": {
		18: 100,
	},

	"Alamo": {
		16: 100,
	},

	"Albany": {
		14: 100,
	},

	"Albion": {
		2: 100,
	},

	"Alderpoint": {
		2: 100,
	},

	"Alhambra": {
		49: 100,
	},

	"Alhambra Valley": {
		15: 100,
	},

	"Aliso Viejo": {
		72: 100,
	},

	"Alleghany": {
		1: 100,
	},

	"Allendale": {
		11: 100,
	},

	"Allensworth": {
		33: 100,
	},

	"Almanor": {
		1: 100,
	},

	"Alondra Park": {
		66: 100,
	},

	"Alpaugh": {
		33: 100,
	},

	"Alpine": {
		75: 100,
	},

	"Alpine Village": {
		1: 100,
	},

	"Alta": {
		1: 100,
	},

	"Alta Sierra": {
		1:  100,
		32: 100,
	},

	"Altadena": {
		41: 100,
	},

	"Alto": {
		12: 100,
	},

	"Alturas": {
		1: 100,
	},

	"Alum Rock": {
		24: 7,
		25: 93,
	},

	"Amador City": {
		1: 100,
	},

	"Amador Pines": {
		1: 100,
	},

	"American Canyon": {
		4: 100,
	},

	"Amesti": {
		29: 100,
	},

	"Anaheim": {
		59: 19,
		67: 41,
		68: 40,
	},

	"Anchor Bay": {
		2: 100,
	},

	"Anderson": {
		1: 100,
	},

	"Angels": {
		8: 100,
	},

	"Angwin": {
		4: 100,
	},

	"Antelope": {
		6: 100,
	},

	"Antioch": {
		15: 100,
	},

	"Anza": {
		36: 100,
	},

	"Apple Valley": {
		34: 100,
	},

	"Aptos": {
		30: 100,
	},

	"Aptos Hills-Larkin Valley": {
		29: 100,
	},

	"Arbuckle": {
		4: 100,
	},

	"Arcadia": {
		49: 100,
	},

	"Arcata": {
		2: 100,
	},

	"Arden-Arcade": {
		6: 100,
	},

	"Armona": {
		33: 100,
	},

	"Arnold": {
		8: 100,
	},

	"Aromas": {
		29: 100,
	},

	"Arroyo Grande": {
		30: 100,
	},

	"Artesia": {
		67: 100,
	},

	"Artois": {
		3: 100,
	},

	"Arvin": {
		35: 100,
	},

	"Ashland": {
		20: 100,
	},

	"Aspen Springs": {
		8: 100,
	},

	"Atascadero": {
		30: 100,
	},

	"Atherton": {
		21: 3,
		23: 97,
	},

	"Atwater": {
		27: 100,
	},

	"Auberry": {
		8: 100,
	},

	"Auburn": {
		5: 100,
	},

	"Auburn Lake Trails": {
		5: 100,
	},

	"August": {
		13: 100,
	},

	"Avalon": {
		69: 100,
	},

	"Avenal": {
		33: 100,
	},

	"Avery": {
		8: 100,
	},

	"Avila Beach": {
		30: 100,
	},

	"Avocado Heights": {
		56: 100,
	},

	"Azusa": {
		48: 100,
	},

	"Baker": {
		34: 100,
	},

	"Bakersfield": {
		32: 43,
		35: 57,
	},

	"Bakersfield Country Club": {
		32: 100,
	},

	"Baldwin Park": {
		48: 100,
	},

	"Ballard": {
		37: 100,
	},

	"Ballico": {
		22: 100,
	},

	"Bangor": {
		3: 100,
	},

	"Banning": {
		47: 100,
	},

	"Barstow": {
		34: 100,
	},

	"Bass Lake": {
		8: 100,
	},

	"Bay Point": {
		15: 100,
	},

	"Bayview": {
		2:  100,
		14: 100,
	},

	"Baywood Park": {
		21: 100,
	},

	"Beale AFB": {
		3: 100,
	},

	"Bear Creek": {
		27: 100,
	},

	"Bear Valley": {
		1: 100,
		8: 100,
	},

	"Bear Valley Springs": {
		32: 100,
	},

	"Beaumont": {
		47: 100,
	},

	"Beckwourth": {
		1: 100,
	},

	"Belden": {
		1: 100,
	},

	"Bell": {
		64: 100,
	},

	"Bell Canyon": {
		46: 100,
	},

	"Bell Gardens": {
		64: 100,
	},

	"Bella Vista": {
		1: 100,
	},

	"Bellflower": {
		62: 100,
	},

	"Belmont": {
		21: 100,
	},

	"Belvedere": {
		12: 100,
	},

	"Ben Lomond": {
		28: 100,
	},

	"Benbow": {
		2: 100,
	},

	"Bend": {
		3: 100,
	},

	"Benicia": {
		11: 100,
	},

	"Benton": {
		8: 100,
	},

	"Benton Park": {
		35: 100,
	},

	"Berkeley": {
		14: 100,
	},

	"Bermuda Dunes": {
		47: 100,
	},

	"Berry Creek": {
		3: 100,
	},

	"Bertsch-Oceanview": {
		2: 100,
	},

	"Bethel Island": {
		11: 100,
	},

	"Beverly Hills": {
		51: 100,
	},

	"Bieber": {
		1: 100,
	},

	"Big Bear City": {
		34: 100,
	},

	"Big Bear Lake": {
		34: 100,
	},

	"Big Bend": {
		1: 100,
	},

	"Big Creek": {
		8: 100,
	},

	"Big Lagoon": {
		2: 100,
	},

	"Big Pine": {
		8: 100,
	},

	"Big River": {
		36: 100,
	},

	"Biggs": {
		3: 100,
	},

	"Biola": {
		27: 100,
	},

	"Bishop": {
		8: 100,
	},

	"Black Point-Green Point": {
		12: 100,
	},

	"Blackhawk": {
		16: 100,
	},

	"Blacklake": {
		37: 100,
	},

	"Blairsden": {
		1: 100,
	},

	"Bloomfield": {
		12: 100,
	},

	"Bloomington": {
		50: 100,
	},

	"Blue Lake": {
		2: 100,
	},

	"Bluewater": {
		36: 100,
	},

	"Blythe": {
		36: 100,
	},

	"Bodega": {
		2: 100,
	},

	"Bodega Bay": {
		2: 100,
	},

	"Bodfish": {
		32: 100,
	},

	"Bolinas": {
		12: 100,
	},

	"Bombay Beach": {
		36: 100,
	},

	"Bonadelle Ranchos": {
		8: 100,
	},

	"Bonita": {
		79: 17,
		80: 83,
	},

	"Bonny Doon": {
		28: 100,
	},

	"Bonsall": {
		75: 100,
	},

	"Boonville": {
		2: 100,
	},

	"Bootjack": {
		8: 100,
	},

	"Boron": {
		34: 100,
	},

	"Boronda": {
		29: 100,
	},

	"Borrego Springs": {
		75: 100,
	},

	"Bostonia": {
		75: 100,
	},

	"Boulder Creek": {
		28: 100,
	},

	"Boulevard": {
		75: 100,
	},

	"Bowles": {
		31: 100,
	},

	"Boyes Hot Springs": {
		4: 100,
	},

	"Bradbury": {
		41: 100,
	},

	"Bradley": {
		30: 100,
	},

	"Brawley": {
		36: 100,
	},

	"Brea": {
		59: 100,
	},

	"Brentwood": {
		15: 100,
	},

	"Bret Harte": {
		22: 100,
	},

	"Bridgeport": {
		8: 100,
	},

	"Brisbane": {
		21: 100,
	},

	"Broadmoor": {
		19: 100,
	},

	"Brookdale": {
		28: 100,
	},

	"Brooks": {
		4: 100,
	},

	"Brooktrails": {
		2: 100,
	},

	"Buck Meadows": {
		8: 100,
	},

	"Buckhorn": {
		1: 100,
	},

	"Bucks Lake": {
		1: 100,
	},

	"Buellton": {
		37: 100,
	},

	"Buena Park": {
		67: 100,
	},

	"Buena Vista": {
		9: 100,
	},

	"Burbank": {
		26: 100,
		44: 100,
	},

	"Burlingame": {
		21: 100,
	},

	"Burney": {
		1: 100,
	},

	"Burnt Ranch": {
		2: 100,
	},

	"Butte Creek Canyon": {
		3: 100,
	},

	"Butte Meadows": {
		3: 100,
	},

	"Butte Valley": {
		3: 100,
	},

	"Buttonwillow": {
		35: 100,
	},

	"Byron": {
		11: 100,
	},

	"Bystrom": {
		22: 100,
	},

	"C-Road": {
		1: 100,
	},

	"Cabazon": {
		47: 100,
	},

	"Calabasas": {
		42: 100,
	},

	"Calexico": {
		36: 100,
	},

	"California City": {
		34: 100,
	},

	"California Hot Springs": {
		32: 100,
	},

	"California Pines": {
		1: 100,
	},

	"California Polytechnic State University": {
		30: 100,
	},

	"Calimesa": {
		47: 100,
	},

	"Calipatria": {
		36: 100,
	},

	"Calistoga": {
		4: 100,
	},

	"Callender": {
		37: 100,
	},

	"Calpella": {
		2: 100,
	},

	"Calpine": {
		1: 100,
	},

	"Calwa": {
		31: 100,
	},

	"Camanche North Shore": {
		9: 100,
	},

	"Camanche Village": {
		9: 100,
	},

	"Camarillo": {
		38: 53,
		42: 47,
	},

	"Cambria": {
		30: 100,
	},

	"Cambrian Park": {
		28: 100,
	},

	"Cameron Park": {
		5: 100,
	},

	"Camino": {
		1: 100,
	},

	"Camino Tassajara": {
		16: 100,
	},

	"Camp Nelson": {
		32: 100,
	},

	"Camp Pendleton Mainside": {
		74: 100,
	},

	"Camp Pendleton South": {
		74: 100,
	},

	"Campbell": {
		23: 100,
	},

	"Campo": {
		75: 100,
	},

	"Camptonville": {
		3: 100,
	},

	"Canby": {
		1: 100,
	},

	"Cantua Creek": {
		27: 100,
	},

	"Canyon Lake": {
		63: 100,
	},

	"Canyondam": {
		1: 100,
	},

	"Capitola": {
		30: 100,
	},

	"Caribou": {
		1: 100,
	},

	"Carlsbad": {
		77: 100,
	},

	"Carmel Valley Village": {
		30: 100,
	},

	"Carmel-by-the-Sea": {
		30: 100,
	},

	"Carmet": {
		2: 100,
	},

	"Carmichael": {
		6: 28,
		7: 72,
	},

	"Carnelian Bay": {
		1: 100,
	},

	"Carpinteria": {
		37: 100,
	},

	"Carrick": {
		1: 100,
	},

	"Carson": {
		65: 28,
		69: 72,
	},

	"Cartago": {
		8: 100,
	},

	"Caruthers": {
		27: 100,
	},

	"Casa Conejo": {
		42: 100,
	},

	"Casa Loma": {
		35: 100,
	},

	"Casa de Oro-Mount Helix": {
		75: 100,
	},

	"Casmalia": {
		37: 100,
	},

	"Caspar": {
		2: 100,
	},

	"Cassel": {
		1: 100,
	},

	"Castaic": {
		40: 100,
	},

	"Castella": {
		1: 100,
	},

	"Castle Hill": {
		16: 100,
	},

	"Castro Valley": {
		20: 100,
	},

	"Castroville": {
		29: 100,
	},

	"Cathedral City": {
		47: 100,
	},

	"Catheys Valley": {
		8: 100,
	},

	"Cayucos": {
		30: 100,
	},

	"Cazadero": {
		2: 100,
	},

	"Cedar Flat": {
		1: 100,
	},

	"Cedar Ridge": {
		8: 100,
	},

	"Cedar Slope": {
		32: 100,
	},

	"Cedarville": {
		1: 100,
	},

	"Centerville": {
		1: 100,
		8: 100,
	},

	"Ceres": {
		22: 100,
	},

	"Cerritos": {
		67: 100,
	},

	"Chalfant": {
		8: 100,
	},

	"Challenge-Brownsville": {
		3: 100,
	},

	"Channel Islands Beach": {
		38: 100,
	},

	"Charleston View": {
		8: 100,
	},

	"Charter Oak": {
		48: 100,
	},

	"Cherokee": {
		3: 100,
	},

	"Cherokee Strip": {
		35: 100,
	},

	"Cherry Valley": {
		47: 100,
	},

	"Cherryland": {
		20: 100,
	},

	"Chester": {
		1: 100,
	},

	"Chico": {
		3: 100,
	},

	"Chilcoot-Vinton": {
		1: 100,
	},

	"China Lake Acres": {
		32: 100,
	},

	"Chinese Camp": {
		8: 100,
	},

	"Chino": {
		53: 89,
		59: 11,
	},

	"Chino Hills": {
		59: 100,
	},

	"Choctaw Valley": {
		32: 100,
	},

	"Chowchilla": {
		27: 100,
	},

	"Chualar": {
		29: 100,
	},

	"Chula Vista": {
		80: 100,
	},

	"Citrus": {
		48: 100,
	},

	"Citrus Heights": {
		7: 100,
	},

	"Claremont": {
		41: 100,
	},

	"Clarksburg": {
		4: 100,
	},

	"Clay": {
		9: 100,
	},

	"Clayton": {
		15: 100,
	},

	"Clear Creek": {
		1: 100,
	},

	"Clearlake": {
		4: 100,
	},

	"Clearlake Oaks": {
		4: 100,
	},

	"Clearlake Riviera": {
		4: 100,
	},

	"Cleone": {
		2: 100,
	},

	"Clio": {
		1: 100,
	},

	"Clipper Mills": {
		3: 100,
	},

	"Cloverdale": {
		2: 100,
	},

	"Clovis": {
		8:  98,
		31: 2,
	},

	"Clyde": {
		15: 100,
	},

	"Coachella": {
		36: 100,
	},

	"Coalinga": {
		27: 100,
	},

	"Coarsegold": {
		8: 100,
	},

	"Cobb": {
		4: 100,
	},

	"Coffee Creek": {
		2: 100,
	},

	"Cohasset": {
		3: 100,
	},

	"Cold Springs": {
		5: 100,
		8: 100,
	},

	"Coleville": {
		8: 100,
	},

	"Colfax": {
		1: 100,
	},

	"College City": {
		4: 100,
	},

	"Collierville": {
		9: 100,
	},

	"Colma": {
		19: 100,
	},

	"Coloma": {
		5: 100,
	},

	"Colton": {
		50: 100,
	},

	"Columbia": {
		8: 100,
	},

	"Colusa": {
		4: 100,
	},

	"Commerce": {
		54: 100,
	},

	"Comptche": {
		2: 100,
	},

	"Compton": {
		65: 100,
	},

	"Concord": {
		15: 100,
	},

	"Concow": {
		3: 100,
	},

	"Contra Costa Centre": {
		16: 100,
	},

	"Copperopolis": {
		9: 100,
	},

	"Corcoran": {
		33: 100,
	},

	"Corning": {
		3: 100,
	},

	"Corona": {
		58: 52,
		63: 48,
	},

	"Coronado": {
		77: 100,
	},

	"Coronita": {
		58: 100,
	},

	"Corralitos": {
		29: 100,
	},

	"Corte Madera": {
		12: 100,
	},

	"Costa Mesa": {
		73: 100,
	},

	"Cotati": {
		12: 100,
	},

	"Coto de Caza": {
		71: 100,
	},

	"Cottonwood": {
		1:  100,
		35: 100,
	},

	"Coulterville": {
		8: 100,
	},

	"Country Club": {
		13: 100,
	},

	"Courtland": {
		9: 100,
	},

	"Covelo": {
		2: 100,
	},

	"Covina": {
		48: 100,
	},

	"Cowan": {
		22: 100,
	},

	"Crane Creek": {
		8: 100,
	},

	"Crescent City": {
		2: 100,
	},

	"Crescent Mills": {
		1: 100,
	},

	"Cressey": {
		27: 100,
	},

	"Crest": {
		75: 100,
	},

	"Crestline": {
		34: 100,
	},

	"Creston": {
		30: 100,
	},

	"Crockett": {
		15: 100,
	},

	"Cromberg": {
		1: 100,
	},

	"Crowley Lake": {
		8: 100,
	},

	"Crows Landing": {
		22: 100,
	},

	"Cudahy": {
		64: 100,
	},

	"Culver City": {
		55: 100,
	},

	"Cupertino": {
		26: 100,
	},

	"Cutler": {
		33: 100,
	},

	"Cutten": {
		2: 100,
	},

	"Cuyama": {
		37: 100,
	},

	"Cypress": {
		67: 100,
	},

	"Dales": {
		3: 100,
	},

	"Daly City": {
		19: 100,
	},

	"Dana Point": {
		74: 100,
	},

	"Danville": {
		16: 100,
	},

	"Daphnedale Park": {
		1: 100,
	},

	"Darwin": {
		8: 100,
	},

	"Davenport": {
		28: 100,
	},

	"Davis": {
		4: 100,
	},

	"Day Valley": {
		30: 100,
	},

	"Deer Park": {
		4: 100,
	},

	"Del Aire": {
		61: 100,
	},

	"Del Dios": {
		76: 100,
	},

	"Del Mar": {
		77: 100,
	},

	"Del Monte Forest": {
		30: 100,
	},

	"Del Rey": {
		31: 100,
	},

	"Del Rey Oaks": {
		30: 100,
	},

	"Del Rio": {
		9: 100,
	},

	"Delano": {
		35: 100,
	},

	"Delft Colony": {
		33: 100,
	},

	"Delhi": {
		27: 100,
	},

	"Delleker": {
		1: 100,
	},

	"Denair": {
		22: 100,
	},

	"Derby Acres": {
		32: 100,
	},

	"Descanso": {
		75: 100,
	},

	"Desert Center": {
		36: 100,
	},

	"Desert Edge": {
		47: 100,
	},

	"Desert Hot Springs": {
		47: 100,
	},

	"Desert Palms": {
		36: 100,
	},

	"Desert Shores": {
		36: 100,
	},

	"Desert View Highlands": {
		34: 100,
	},

	"Di Giorgio": {
		35: 100,
	},

	"Diablo": {
		16: 100,
	},

	"Diablo Grande": {
		22: 100,
	},

	"Diamond Bar": {
		56: 100,
	},

	"Diamond Springs": {
		5: 100,
	},

	"Dillon Beach": {
		12: 100,
	},

	"Dinuba": {
		33: 100,
	},

	"Discovery Bay": {
		11: 100,
	},

	"Dixon": {
		11: 100,
	},

	"Dixon Lane-MeadowCreek": {
		8: 100,
	},

	"Dobbins": {
		3: 100,
	},

	"Dogtown": {
		9: 100,
	},

	"Dollar Point": {
		1: 100,
	},

	"Dorrington": {
		8: 100,
	},

	"Dorris": {
		1: 100,
	},

	"Dos Palos": {
		27: 100,
	},

	"Dos Palos Y": {
		27: 100,
	},

	"Douglas City": {
		2: 100,
	},

	"Downey": {
		64: 100,
	},

	"Downieville": {
		1: 100,
	},

	"Doyle": {
		1: 100,
	},

	"Drytown": {
		1: 100,
	},

	"Duarte": {
		41: 0,
		48: 100,
	},

	"Dublin": {
		16: 61,
		20: 39,
	},

	"Ducor": {
		33: 100,
	},

	"Dunnigan": {
		4: 100,
	},

	"Dunsmuir": {
		1: 100,
	},

	"Durham": {
		3: 100,
	},

	"Dustin Acres": {
		32: 100,
	},

	"Dutch Flat": {
		1: 100,
	},

	"Eagleville": {
		1: 100,
	},

	"Earlimart": {
		33: 100,
	},

	"East Bakersfield": {
		35: 100,
	},

	"East Foothills": {
		24: 42,
		25: 58,
	},

	"East Hemet": {
		60: 100,
	},

	"East Los Angeles": {
		52: 100,
	},

	"East Nicolaus": {
		3: 100,
	},

	"East Niles": {
		35: 100,
	},

	"East Oakdale": {
		9: 100,
	},

	"East Orosi": {
		33: 100,
	},

	"East Palo Alto": {
		21: 100,
	},

	"East Pasadena": {
		49: 100,
	},

	"East Porterville": {
		32: 0,
		33: 100,
	},

	"East Quincy": {
		1: 100,
	},

	"East Rancho Dominguez": {
		65: 100,
	},

	"East Richmond Heights": {
		14: 100,
	},

	"East San Gabriel": {
		49: 100,
	},

	"East Shore": {
		1: 100,
	},

	"East Sonora": {
		8: 100,
	},

	"East Tulare Villa": {
		33: 100,
	},

	"East Whittier": {
		64: 100,
	},

	"Eastern Goleta Valley": {
		37: 100,
	},

	"Easton": {
		31: 100,
	},

	"Eastvale": {
		58: 52,
		63: 48,
	},

	"Edgewood": {
		1: 100,
	},

	"Edison": {
		35: 100,
	},

	"Edmundson Acres": {
		35: 100,
	},

	"Edna": {
		30: 100,
	},

	"Edwards AFB": {
		34: 100,
	},

	"El Adobe": {
		35: 100,
	},

	"El Cajon": {
		78: 31,
		79: 69,
	},

	"El Centro": {
		36: 100,
	},

	"El Centro Naval Air Facility": {
		36: 100,
	},

	"El Cerrito": {
		14: 100,
		58: 100,
	},

	"El Dorado Hills": {
		5: 100,
	},

	"El Granada": {
		23: 100,
	},

	"El Macero": {
		4: 100,
	},

	"El Monte": {
		49: 60,
		56: 40,
	},

	"El Monte Mobile Village": {
		33: 100,
	},

	"El Nido": {
		27: 100,
	},

	"El Paso de Robles (Paso Robles)": {
		30: 100,
	},

	"El Portal": {
		8: 100,
	},

	"El Rancho": {
		33: 100,
	},

	"El Rio": {
		38: 100,
	},

	"El Segundo": {
		66: 100,
	},

	"El Sobrante": {
		14: 100,
		63: 100,
	},

	"El Verano": {
		4: 100,
	},

	"Eldridge": {
		4: 100,
	},

	"Elfin Forest": {
		76: 100,
	},

	"Elizabeth Lake": {
		34: 100,
	},

	"Elk Creek": {
		3: 100,
	},

	"Elk Grove": {
		10: 100,
	},

	"Elkhorn": {
		29: 100,
	},

	"Elmira": {
		11: 100,
	},

	"Elverta": {
		6: 100,
	},

	"Emerald Lake Hills": {
		21: 100,
	},

	"Emeryville": {
		18: 100,
	},

	"Empire": {
		22: 100,
	},

	"Encinitas": {
		77: 100,
	},

	"Escalon": {
		9: 100,
	},

	"Escondido": {
		76: 100,
	},

	"Esparto": {
		4: 100,
	},

	"Etna": {
		1: 100,
	},

	"Eucalyptus Hills": {
		75: 100,
	},

	"Eureka": {
		2: 100,
	},

	"Exeter": {
		32: 100,
	},

	"Fair Oaks": {
		7: 100,
	},

	"Fairbanks Ranch": {
		76: 100,
	},

	"Fairfax": {
		12: 100,
		35: 100,
	},

	"Fairfield": {
		11: 100,
	},

	"Fairhaven": {
		2: 100,
	},

	"Fairmead": {
		27: 100,
	},

	"Fairview": {
		20: 100,
	},

	"Fall River Mills": {
		1: 100,
	},

	"Fallbrook": {
		75: 100,
	},

	"Farmersville": {
		33: 100,
	},

	"Farmington": {
		9: 100,
	},

	"Fellows": {
		32: 100,
	},

	"Felton": {
		28: 100,
	},

	"Ferndale": {
		2: 100,
	},

	"Fetters Hot Springs-Agua Caliente": {
		4: 100,
	},

	"Fiddletown": {
		1: 100,
	},

	"Fieldbrook": {
		2: 100,
	},

	"Fields Landing": {
		2: 100,
	},

	"Fillmore": {
		38: 100,
	},

	"Firebaugh": {
		27: 100,
	},

	"Fish Camp": {
		8: 100,
	},

	"Florence-Graham": {
		57: 100,
	},

	"Florin": {
		10: 100,
	},

	"Floriston": {
		1: 100,
	},

	"Flournoy": {
		3: 100,
	},

	"Folsom": {
		7: 100,
	},

	"Fontana": {
		45: 21,
		50: 79,
	},

	"Foothill Farms": {
		7: 100,
	},

	"Forbestown": {
		3: 100,
	},

	"Ford City": {
		32: 100,
	},

	"Forest Meadows": {
		8: 100,
	},

	"Forest Ranch": {
		3: 100,
	},

	"Foresthill": {
		1: 100,
	},

	"Forestville": {
		2: 100,
	},

	"Fort Bidwell": {
		1: 100,
	},

	"Fort Bragg": {
		2: 100,
	},

	"Fort Dick": {
		2: 100,
	},

	"Fort Hunter Liggett": {
		30: 100,
	},

	"Fort Irwin": {
		34: 100,
	},

	"Fort Jones": {
		1: 100,
	},

	"Fort Washington": {
		8: 100,
	},

	"Fortuna": {
		2: 100,
	},

	"Foster City": {
		21: 100,
	},

	"Fountain Valley": {
		70: 100,
	},

	"Fowler": {
		31: 100,
	},

	"Franklin": {
		9:  100,
		27: 100,
	},

	"Frazier Park": {
		32: 100,
	},

	"Freedom": {
		29: 100,
	},

	"Freeport": {
		10: 100,
	},

	"Fremont": {
		24: 100,
	},

	"French Camp": {
		13: 100,
	},

	"French Gulch": {
		1: 100,
	},

	"French Valley": {
		36: 100,
	},

	"Fresno": {
		8:  27,
		27: 11,
		31: 62,
	},

	"Friant": {
		8: 100,
	},

	"Fruitdale": {
		26: 100,
	},

	"Fruitridge Pocket": {
		10: 100,
	},

	"Fuller Acres": {
		35: 100,
	},

	"Fullerton": {
		59: 36,
		67: 64,
	},

	"Fulton": {
		2: 100,
	},

	"Furnace Creek": {
		8: 100,
	},

	"Galt": {
		9: 100,
	},

	"Garberville": {
		2: 100,
	},

	"Garden Acres": {
		9: 100,
	},

	"Garden Farms": {
		30: 100,
	},

	"Garden Grove": {
		70: 100,
	},

	"Gardena": {
		61: 25,
		66: 75,
	},

	"Garey": {
		37: 100,
	},

	"Garnet": {
		47: 100,
	},

	"Gasquet": {
		2: 100,
	},

	"Gazelle": {
		1: 100,
	},

	"Georgetown": {
		5: 100,
	},

	"Gerber": {
		3: 100,
	},

	"Geyserville": {
		2: 100,
	},

	"Gilroy": {
		29: 100,
	},

	"Glen Ellen": {
		12: 100,
	},

	"Glendale": {
		44: 60,
		52: 40,
	},

	"Glendora": {
		48: 100,
	},

	"Glennville": {
		32: 100,
	},

	"Gold Mountain": {
		1: 100,
	},

	"Gold River": {
		7: 100,
	},

	"Golden Hills": {
		32: 100,
	},

	"Goleta": {
		37: 100,
	},

	"Gonzales": {
		29: 100,
	},

	"Good Hope": {
		60: 100,
	},

	"Goodmanville": {
		32: 100,
	},

	"Goodyears Bar": {
		1: 100,
	},

	"Goshen": {
		32: 100,
	},

	"Graeagle": {
		1: 100,
	},

	"Grand Terrace": {
		58: 100,
	},

	"Grangeville": {
		33: 100,
	},

	"Granite Bay": {
		5: 100,
	},

	"Granite Hills": {
		75: 100,
	},

	"Graniteville": {
		1: 100,
	},

	"Grass Valley": {
		1: 100,
	},

	"Graton": {
		2: 100,
	},

	"Grayson": {
		22: 100,
	},

	"Greeley Hill": {
		8: 100,
	},

	"Green Acres": {
		60: 100,
	},

	"Green Valley": {
		11: 100,
		34: 100,
	},

	"Greenacres": {
		32: 100,
	},

	"Greenfield": {
		29: 100,
		35: 100,
	},

	"Greenhorn": {
		1: 100,
	},

	"Greenview": {
		1: 100,
	},

	"Greenville": {
		1: 100,
	},

	"Grenada": {
		1: 100,
	},

	"Gridley": {
		3: 100,
	},

	"Grimes": {
		4: 100,
	},

	"Grizzly Flats": {
		1: 100,
	},

	"Groveland": {
		8: 100,
	},

	"Grover Beach": {
		30: 100,
	},

	"Guadalupe": {
		37: 100,
	},

	"Guerneville": {
		2: 100,
	},

	"Guinda": {
		4: 100,
	},

	"Gustine": {
		22: 100,
	},

	"Hacienda Heights": {
		56: 100,
	},

	"Half Moon Bay": {
		23: 100,
	},

	"Hamilton Branch": {
		1: 100,
	},

	"Hamilton City": {
		3: 100,
	},

	"Hanford": {
		33: 100,
	},

	"Happy Camp": {
		1: 100,
	},

	"Happy Valley": {
		1: 100,
	},

	"Harbison Canyon": {
		75: 100,
	},

	"Hardwick": {
		33: 100,
	},

	"Harmony Grove": {
		76: 100,
	},

	"Hartland": {
		32: 100,
	},

	"Hartley": {
		11: 100,
	},

	"Hasley Canyon": {
		40: 100,
	},

	"Hat Creek": {
		1: 100,
	},

	"Hawaiian Gardens": {
		67: 100,
	},

	"Hawthorne": {
		61: 100,
	},

	"Hayfork": {
		2: 100,
	},

	"Hayward": {
		20: 100,
	},

	"Healdsburg": {
		2: 100,
	},

	"Heber": {
		36: 100,
	},

	"Hemet": {
		36: 51,
		60: 49,
	},

	"Herald": {
		9: 100,
	},

	"Hercules": {
		14: 100,
	},

	"Herlong": {
		1: 100,
	},

	"Hermosa Beach": {
		66: 100,
	},

	"Hesperia": {
		34: 48,
		39: 51,
		41: 1,
	},

	"Hickman": {
		9: 100,
	},

	"Hidden Hills": {
		42: 100,
	},

	"Hidden Meadows": {
		75: 100,
	},

	"Hidden Valley Lake": {
		4: 100,
	},

	"Highgrove": {
		58: 100,
	},

	"Highland": {
		34: 23,
		45: 56,
		47: 21,
	},

	"Highlands": {
		21: 100,
	},

	"Hillcrest": {
		35: 100,
	},

	"Hillsborough": {
		21: 100,
	},

	"Hilmar-Irwin": {
		22: 100,
	},

	"Hiouchi": {
		2: 100,
	},

	"Hollister": {
		29: 100,
	},

	"Holtville": {
		36: 100,
	},

	"Home Garden": {
		33: 100,
	},

	"Home Gardens": {
		58: 100,
	},

	"Homeland": {
		60: 100,
	},

	"Homestead Valley": {
		34: 100,
	},

	"Homewood Canyon": {
		8: 100,
	},

	"Honcut": {
		3: 100,
	},

	"Hood": {
		9: 100,
	},

	"Hoopa": {
		2: 100,
	},

	"Hopland": {
		2: 100,
	},

	"Hornbrook": {
		1: 100,
	},

	"Hornitos": {
		8: 100,
	},

	"Hughson": {
		9: 100,
	},

	"Humboldt Hill": {
		2: 100,
	},

	"Huntington Beach": {
		70: 10,
		72: 90,
	},

	"Huntington Park": {
		62: 100,
	},

	"Huron": {
		27: 100,
	},

	"Hyampom": {
		2: 100,
	},

	"Hydesville": {
		2: 100,
	},

	"Hypericum": {
		32: 100,
	},

	"Idlewild": {
		32: 100,
	},

	"Idyllwild-Pine Cove": {
		47: 100,
	},

	"Igo": {
		1: 100,
	},

	"Imperial": {
		36: 100,
	},

	"Imperial Beach": {
		80: 100,
	},

	"Independence": {
		8: 100,
	},

	"Indian Falls": {
		1: 100,
	},

	"Indian Wells": {
		47: 100,
	},

	"Indianola": {
		2: 100,
	},

	"Indio": {
		36: 100,
	},

	"Indio Hills": {
		36: 100,
	},

	"Industry": {
		56: 100,
	},

	"Inglewood": {
		61: 100,
	},

	"Interlaken": {
		29: 100,
	},

	"Inverness": {
		12: 100,
	},

	"Inyokern": {
		32: 100,
	},

	"Ione": {
		9: 100,
	},

	"Iron Horse": {
		1: 100,
	},

	"Irvine": {
		73: 100,
	},

	"Irwindale": {
		48: 100,
	},

	"Isla Vista": {
		37: 100,
	},

	"Isleton": {
		9: 100,
	},

	"Ivanhoe": {
		33: 100,
	},

	"Jackson": {
		1: 100,
	},

	"Jacumba": {
		75: 100,
	},

	"Jamestown": {
		8: 100,
	},

	"Jamul": {
		75: 100,
	},

	"Janesville": {
		1: 100,
	},

	"Jenner": {
		2: 100,
	},

	"Johannesburg": {
		34: 100,
	},

	"Johnson Park": {
		1: 100,
	},

	"Johnstonville": {
		1: 100,
	},

	"Johnsville": {
		1: 100,
	},

	"Jones Valley": {
		1: 100,
	},

	"Joshua Tree": {
		47: 100,
	},

	"Jovista": {
		33: 100,
	},

	"Julian": {
		75: 100,
	},

	"Junction City": {
		2: 100,
	},

	"June Lake": {
		8: 100,
	},

	"Jurupa Valley": {
		58: 100,
	},

	"Keddie": {
		1: 100,
	},

	"Keeler": {
		8: 100,
	},

	"Keene": {
		32: 100,
	},

	"Kelly Ridge": {
		3: 100,
	},

	"Kelseyville": {
		4: 100,
	},

	"Kennedy": {
		13: 100,
	},

	"Kennedy Meadows": {
		32: 100,
	},

	"Kensington": {
		14: 100,
	},

	"Kentfield": {
		12: 100,
	},

	"Kenwood": {
		12: 100,
	},

	"Kep'el": {
		2: 100,
	},

	"Kerman": {
		27: 100,
	},

	"Kernville": {
		32: 100,
	},

	"Keswick": {
		1: 100,
	},

	"Kettleman City": {
		33: 100,
	},

	"Keyes": {
		22: 100,
	},

	"King City": {
		29: 100,
	},

	"Kings Beach": {
		1: 100,
	},

	"Kingsburg": {
		33: 100,
	},

	"Kingvale": {
		1: 100,
	},

	"Kirkwood": {
		1: 100,
	},

	"Klamath": {
		2: 100,
	},

	"Knights Ferry": {
		9: 100,
	},

	"Knights Landing": {
		4: 100,
	},

	"Knightsen": {
		11: 100,
	},

	"La Cañada Flintridge": {
		41: 100,
	},

	"La Crescenta-Montrose": {
		44: 100,
	},

	"La Cresta": {
		35: 100,
	},

	"La Grange": {
		9: 100,
	},

	"La Habra": {
		64: 100,
	},

	"La Habra Heights": {
		56: 100,
	},

	"La Honda": {
		23: 100,
	},

	"La Mesa": {
		79: 100,
	},

	"La Mirada": {
		64: 100,
	},

	"La Palma": {
		67: 100,
	},

	"La Porte": {
		1: 100,
	},

	"La Presa": {
		79: 100,
	},

	"La Puente": {
		56: 100,
	},

	"La Quinta": {
		47: 100,
	},

	"La Riviera": {
		6: 100,
	},

	"La Selva Beach": {
		30: 100,
	},

	"La Verne": {
		41: 100,
	},

	"La Vina": {
		27: 100,
	},

	"Ladera": {
		23: 100,
	},

	"Ladera Heights": {
		55: 100,
	},

	"Ladera Ranch": {
		71: 100,
	},

	"Lafayette": {
		16: 100,
	},

	"Laguna Beach": {
		72: 100,
	},

	"Laguna Hills": {
		72: 100,
	},

	"Laguna Niguel": {
		74: 100,
	},

	"Laguna Woods": {
		72: 100,
	},

	"Lagunitas-Forest Knolls": {
		12: 100,
	},

	"Lake Almanor Country Club": {
		1: 100,
	},

	"Lake Almanor Peninsula": {
		1: 100,
	},

	"Lake Almanor West": {
		1: 100,
	},

	"Lake Arrowhead": {
		34: 100,
	},

	"Lake California": {
		3: 100,
	},

	"Lake City": {
		1: 100,
	},

	"Lake Davis": {
		1: 100,
	},

	"Lake Don Pedro": {
		8: 100,
	},

	"Lake Elsinore": {
		63: 100,
	},

	"Lake Forest": {
		72: 100,
	},

	"Lake Hughes": {
		34: 100,
	},

	"Lake Isabella": {
		32: 100,
	},

	"Lake Los Angeles": {
		39: 100,
	},

	"Lake Mathews": {
		63: 100,
	},

	"Lake Nacimiento": {
		30: 100,
	},

	"Lake Riverside": {
		36: 100,
	},

	"Lake San Marcos": {
		76: 100,
	},

	"Lake Shastina": {
		1: 100,
	},

	"Lake Sherwood": {
		42: 100,
	},

	"Lake Wildwood": {
		1: 100,
	},

	"Lake of the Pines": {
		1: 100,
	},

	"Lake of the Woods": {
		32: 100,
	},

	"Lakehead": {
		1: 100,
	},

	"Lakeland Village": {
		63: 100,
	},

	"Lakeport": {
		4: 100,
	},

	"Lakeside": {
		32: 100,
		75: 100,
	},

	"Lakeview": {
		60: 100,
	},

	"Lakewood": {
		62: 100,
	},

	"Lamont": {
		35: 100,
	},

	"Lanare": {
		31: 100,
	},

	"Lancaster": {
		34: 33,
		39: 67,
	},

	"Larkfield-Wikiup": {
		2: 100,
	},

	"Larkspur": {
		12: 100,
	},

	"Las Flores": {
		3:  100,
		71: 100,
	},

	"Las Lomas": {
		29: 100,
	},

	"Lathrop": {
		9: 100,
	},

	"Laton": {
		31: 100,
	},

	"Lawndale": {
		61: 100,
	},

	"Laytonville": {
		2: 100,
	},

	"Le Grand": {
		27: 100,
	},

	"Lebec": {
		32: 100,
	},

	"Lee Vining": {
		8: 100,
	},

	"Leggett": {
		2: 100,
	},

	"Lemon Cove": {
		32: 100,
	},

	"Lemon Grove": {
		79: 100,
	},

	"Lemon Hill": {
		10: 100,
	},

	"Lemoore": {
		33: 100,
	},

	"Lemoore Station": {
		33: 100,
	},

	"Lennox": {
		61: 100,
	},

	"Lenwood": {
		34: 100,
	},

	"Leona Valley": {
		34: 100,
	},

	"Lewiston": {
		2: 100,
	},

	"Lexington Hills": {
		28: 100,
	},

	"Likely": {
		1: 100,
	},

	"Lincoln": {
		3: 0,
		5: 100,
	},

	"Lincoln Village": {
		13: 100,
	},

	"Linda": {
		3: 100,
	},

	"Lindcove": {
		32: 100,
	},

	"Linden": {
		9: 100,
	},

	"Lindsay": {
		33: 100,
	},

	"Linnell Camp": {
		33: 100,
	},

	"Litchfield": {
		1: 100,
	},

	"Little Grass Valley": {
		1: 100,
	},

	"Little River": {
		2: 100,
	},

	"Little Valley": {
		1: 100,
	},

	"Littlerock": {
		39: 100,
	},

	"Live Oak": {
		3:  100,
		30: 100,
	},

	"Livermore": {
		16: 100,
	},

	"Livingston": {
		27: 100,
	},

	"Lockeford": {
		9: 100,
	},

	"Lockwood": {
		1:  100,
		30: 100,
	},

	"Lodi": {
		9: 100,
	},

	"Lodoga": {
		4: 100,
	},

	"Loleta": {
		2: 100,
	},

	"Loma Linda": {
		50: 100,
	},

	"Loma Mar": {
		23: 100,
	},

	"Loma Rica": {
		3: 100,
	},

	"Lomita": {
		66: 100,
	},

	"Lompico": {
		28: 100,
	},

	"Lompoc": {
		37: 100,
	},

	"London": {
		33: 100,
	},

	"Lone Pine": {
		8: 100,
	},

	"Long Barn": {
		8: 100,
	},

	"Long Beach": {
		65: 16,
		69: 84,
	},

	"Lookout": {
		1: 100,
	},

	"Loomis": {
		5: 100,
	},

	"Los Alamitos": {
		70: 100,
	},

	"Los Alamos": {
		37: 100,
	},

	"Los Altos": {
		23: 100,
	},

	"Los Altos Hills": {
		23: 100,
	},

	"Los Angeles": {
		40: 6,
		42: 2,
		43: 12,
		44: 7,
		46: 13,
		51: 9,
		52: 8,
		54: 11,
		55: 11,
		57: 11,
		61: 4,
		65: 6,
		66: 1,
	},

	"Los Banos": {
		27: 100,
	},

	"Los Berros": {
		37: 100,
	},

	"Los Gatos": {
		28: 100,
	},

	"Los Molinos": {
		3: 100,
	},

	"Los Olivos": {
		37: 100,
	},

	"Los Osos": {
		30: 100,
	},

	"Los Ranchos": {
		30: 100,
	},

	"Lost Hills": {
		35: 100,
	},

	"Lower Lake": {
		4: 100,
	},

	"Loyalton": {
		1: 100,
	},

	"Loyola": {
		23: 100,
	},

	"Lucas Valley-Marinwood": {
		12: 100,
	},

	"Lucerne": {
		4: 100,
	},

	"Lucerne Valley": {
		34: 100,
	},

	"Lynwood": {
		62: 100,
	},

	"Lytle Creek": {
		41: 100,
	},

	"Mabie": {
		1: 100,
	},

	"Macdoel": {
		1: 100,
	},

	"Mad River": {
		2: 100,
	},

	"Madeline": {
		1: 100,
	},

	"Madera": {
		27: 100,
	},

	"Madera Acres": {
		27: 100,
	},

	"Madera Ranchos": {
		8: 100,
	},

	"Madison": {
		4: 100,
	},

	"Magalia": {
		3: 100,
	},

	"Malaga": {
		31: 100,
	},

	"Malibu": {
		42: 100,
	},

	"Mammoth Lakes": {
		8: 100,
	},

	"Manchester": {
		2: 100,
	},

	"Manhattan Beach": {
		66: 100,
	},

	"Manila": {
		2: 100,
	},

	"Manteca": {
		9: 100,
	},

	"Manton": {
		3: 100,
	},

	"March ARB": {
		60: 62,
		63: 38,
	},

	"Maricopa": {
		32: 100,
	},

	"Marin City": {
		12: 100,
	},

	"Marina": {
		30: 100,
	},

	"Marina del Rey": {
		61: 100,
	},

	"Mariposa": {
		8: 100,
	},

	"Markleeville": {
		1: 100,
	},

	"Martell": {
		1: 100,
	},

	"Martinez": {
		15: 100,
	},

	"Marysville": {
		3: 100,
	},

	"Matheny": {
		33: 100,
	},

	"Mather": {
		7: 100,
	},

	"Maxwell": {
		4: 100,
	},

	"Mayfair": {
		31: 100,
	},

	"Mayflower Village": {
		48: 100,
	},

	"Maywood": {
		62: 100,
	},

	"McArthur": {
		1: 100,
	},

	"McClellan Park": {
		7: 100,
	},

	"McClenney Tract": {
		32: 100,
	},

	"McCloud": {
		1: 100,
	},

	"McFarland": {
		35: 100,
	},

	"McGee Creek": {
		8: 100,
	},

	"McKinleyville": {
		2: 100,
	},

	"McKittrick": {
		32: 100,
	},

	"McSwain": {
		27: 100,
	},

	"Mead Valley": {
		60: 100,
	},

	"Meadow Valley": {
		1: 100,
	},

	"Meadow Vista": {
		5: 100,
	},

	"Meadowbrook": {
		63: 100,
	},

	"Mecca": {
		36: 100,
	},

	"Meiners Oaks": {
		38: 100,
	},

	"Mendocino": {
		2: 100,
	},

	"Mendota": {
		27: 100,
	},

	"Menifee": {
		63: 100,
	},

	"Menlo Park": {
		21: 19,
		23: 81,
	},

	"Mentone": {
		45: 73,
		47: 27,
	},

	"Merced": {
		27: 100,
	},

	"Meridian": {
		3: 100,
	},

	"Mesa": {
		8: 100,
	},

	"Mesa Verde": {
		36: 100,
	},

	"Mesa Vista": {
		1: 100,
	},

	"Mettler": {
		32: 100,
	},

	"Mexican Colony": {
		35: 100,
	},

	"Meyers": {
		1: 100,
	},

	"Mi-Wuk Village": {
		8: 100,
	},

	"Middletown": {
		4: 100,
	},

	"Midpines": {
		8: 100,
	},

	"Midway City": {
		70: 100,
	},

	"Milford": {
		1: 100,
	},

	"Mill Valley": {
		12: 100,
	},

	"Millbrae": {
		21: 100,
	},

	"Millerton": {
		8: 100,
	},

	"Millville": {
		1: 100,
	},

	"Milpitas": {
		24: 100,
	},

	"Mineral": {
		3: 100,
	},

	"Minkler": {
		8: 100,
	},

	"Mira Monte": {
		38: 100,
	},

	"Miranda": {
		2: 100,
	},

	"Mission Canyon": {
		37: 100,
	},

	"Mission Hills": {
		37: 100,
	},

	"Mission Viejo": {
		71: 100,
	},

	"Modesto": {
		22: 100,
	},

	"Modjeska": {
		71: 100,
	},

	"Mohawk Vista": {
		1: 100,
	},

	"Mojave": {
		34: 100,
	},

	"Mokelumne Hill": {
		8: 100,
	},

	"Monmouth": {
		31: 100,
	},

	"Mono City": {
		8: 100,
	},

	"Mono Vista": {
		8: 100,
	},

	"Monrovia": {
		41: 99,
		48: 1,
	},

	"Monson": {
		33: 100,
	},

	"Montague": {
		1: 100,
	},

	"Montalvin Manor": {
		14: 100,
	},

	"Montara": {
		23: 100,
	},

	"Montclair": {
		53: 100,
	},

	"Monte Rio": {
		2: 100,
	},

	"Monte Sereno": {
		28: 100,
	},

	"Montebello": {
		54: 100,
	},

	"Montecito": {
		37: 100,
	},

	"Monterey": {
		30: 100,
	},

	"Monterey Park": {
		49: 100,
	},

	"Monterey Park Tract": {
		22: 100,
	},

	"Montgomery Creek": {
		1: 100,
	},

	"Monument Hills": {
		4: 100,
	},

	"Moorpark": {
		42: 100,
	},

	"Morada": {
		9: 100,
	},

	"Moraga": {
		16: 100,
	},

	"Moreno Valley": {
		60: 100,
	},

	"Morgan Hill": {
		28: 100,
	},

	"Morongo Valley": {
		47: 100,
	},

	"Morro Bay": {
		30: 100,
	},

	"Moskowite Corner": {
		4: 100,
	},

	"Moss Beach": {
		23: 100,
	},

	"Moss Landing": {
		30: 100,
	},

	"Mount Hebron": {
		1: 100,
	},

	"Mount Hermon": {
		28: 100,
	},

	"Mount Laguna": {
		75: 100,
	},

	"Mount Shasta": {
		1: 100,
	},

	"Mountain Center": {
		47: 100,
	},

	"Mountain Gate": {
		1: 100,
	},

	"Mountain House": {
		13: 100,
	},

	"Mountain Meadows": {
		32: 100,
	},

	"Mountain Mesa": {
		32: 100,
	},

	"Mountain Ranch": {
		8: 100,
	},

	"Mountain View": {
		15: 100,
		23: 100,
	},

	"Mountain View Acres": {
		39: 100,
	},

	"Mt. Bullion": {
		8: 100,
	},

	"Muir Beach": {
		12: 100,
	},

	"Murphys": {
		8: 100,
	},

	"Murrieta": {
		71: 100,
	},

	"Muscoy": {
		45: 100,
	},

	"Myers Flat": {
		2: 100,
	},

	"Myrtletown": {
		2: 100,
	},

	"Napa": {
		4: 100,
	},

	"National City": {
		80: 100,
	},

	"Needles": {
		36: 100,
	},

	"Nevada City": {
		1: 100,
	},

	"New Cuyama": {
		37: 100,
	},

	"New Pine Creek": {
		1: 100,
	},

	"Newark": {
		24: 100,
	},

	"Newcastle": {
		5: 100,
	},

	"Newell": {
		1: 100,
	},

	"Newman": {
		22: 100,
	},

	"Newport Beach": {
		72: 100,
	},

	"Nicasio": {
		12: 100,
	},

	"Nice": {
		4: 100,
	},

	"Nicolaus": {
		3: 100,
	},

	"Niland": {
		36: 100,
	},

	"Nipinnawasee": {
		8: 100,
	},

	"Nipomo": {
		37: 100,
	},

	"Norco": {
		63: 100,
	},

	"Nord": {
		3: 100,
	},

	"Norris Canyon": {
		16: 100,
	},

	"North Auburn": {
		5: 100,
	},

	"North Edwards": {
		34: 100,
	},

	"North El Monte": {
		49: 100,
	},

	"North Fair Oaks": {
		21: 100,
		23: 0,
	},

	"North Fork": {
		8: 100,
	},

	"North Gate": {
		16: 100,
	},

	"North Highlands": {
		7: 100,
	},

	"North Lakeport": {
		4: 100,
	},

	"North Richmond": {
		14: 100,
	},

	"North San Juan": {
		1: 100,
	},

	"North Shore": {
		36: 100,
	},

	"North Tustin": {
		59: 100,
	},

	"Norwalk": {
		64: 100,
	},

	"Novato": {
		12: 100,
	},

	"Nubieber": {
		1: 100,
	},

	"Nuevo": {
		60: 100,
	},

	"Oak Glen": {
		47: 100,
	},

	"Oak Hills": {
		41: 100,
	},

	"Oak Park": {
		42: 100,
	},

	"Oak Run": {
		1: 100,
	},

	"Oak Shores": {
		30: 100,
	},

	"Oak View": {
		38: 100,
	},

	"Oakdale": {
		9: 100,
	},

	"Oakhurst": {
		8: 100,
	},

	"Oakland": {
		14: 11,
		18: 89,
		20: 0,
	},

	"Oakley": {
		11: 100,
	},

	"Oakville": {
		4: 100,
	},

	"Oasis": {
		36: 100,
	},

	"Occidental": {
		2: 100,
	},

	"Oceano": {
		30: 100,
	},

	"Oceanside": {
		74: 100,
	},

	"Ocotillo": {
		36: 100,
	},

	"Oildale": {
		32: 100,
	},

	"Ojai": {
		38: 100,
	},

	"Olancha": {
		8: 100,
	},

	"Old Fig Garden": {
		8: 100,
	},

	"Old River": {
		32: 100,
	},

	"Old Station": {
		1: 100,
	},

	"Old Stine": {
		35: 100,
	},

	"Olde Stockdale": {
		32: 100,
	},

	"Olivehurst": {
		3: 100,
	},

	"Ono": {
		1: 100,
	},

	"Ontario": {
		50: 13,
		53: 87,
	},

	"Onyx": {
		32: 100,
	},

	"Orange": {
		59: 40,
		68: 60,
	},

	"Orange Blossom": {
		9: 100,
	},

	"Orange Cove": {
		31: 100,
	},

	"Orangevale": {
		7: 100,
	},

	"Orcutt": {
		37: 100,
	},

	"Orick": {
		2: 100,
	},

	"Orinda": {
		16: 100,
	},

	"Orland": {
		3: 100,
	},

	"Orosi": {
		33: 100,
	},

	"Oroville": {
		3: 100,
	},

	"Oroville East": {
		3: 100,
	},

	"Oxnard": {
		38: 100,
	},

	"Pacheco": {
		15: 100,
	},

	"Pacific Grove": {
		30: 100,
	},

	"Pacifica": {
		23: 100,
	},

	"Pajaro": {
		29: 100,
	},

	"Pajaro Dunes": {
		30: 100,
	},

	"Pala": {
		75: 100,
	},

	"Palermo": {
		3: 100,
	},

	"Palm Desert": {
		47: 100,
	},

	"Palm Springs": {
		47: 100,
	},

	"Palmdale": {
		34: 28,
		39: 72,
	},

	"Palo Alto": {
		23: 100,
	},

	"Palo Cedro": {
		1: 100,
	},

	"Palo Verde": {
		36: 100,
	},

	"Palos Verdes Estates": {
		66: 100,
	},

	"Panorama Heights": {
		32: 100,
	},

	"Paradise": {
		3: 100,
		8: 100,
	},

	"Paradise Park": {
		28: 100,
	},

	"Paramount": {
		62: 100,
	},

	"Parklawn": {
		22: 100,
	},

	"Parksdale": {
		27: 100,
	},

	"Parkway": {
		10: 100,
	},

	"Parkwood": {
		27: 100,
	},

	"Parlier": {
		31: 100,
	},

	"Pasadena": {
		41: 100,
	},

	"Pasatiempo": {
		28: 100,
	},

	"Paskenta": {
		3: 100,
	},

	"Patterson": {
		22: 100,
	},

	"Patterson Tract": {
		33: 100,
	},

	"Patton Village": {
		1: 100,
	},

	"Paynes Creek": {
		3: 100,
	},

	"Pearsonville": {
		8: 100,
	},

	"Penn Valley": {
		1: 100,
	},

	"Penngrove": {
		12: 100,
	},

	"Penryn": {
		5: 100,
	},

	"Pepperdine University": {
		42: 100,
	},

	"Perris": {
		60: 100,
	},

	"Pescadero": {
		23: 100,
	},

	"Petaluma": {
		12: 100,
	},

	"Petaluma Center": {
		12: 100,
	},

	"Peters": {
		9: 100,
	},

	"Phelan": {
		41: 100,
	},

	"Phillipsville": {
		2: 100,
	},

	"Philo": {
		2: 100,
	},

	"Phoenix Lake": {
		8: 100,
	},

	"Pico Rivera": {
		56: 100,
	},

	"Piedmont": {
		14: 100,
	},

	"Pierpoint": {
		32: 100,
	},

	"Pike": {
		1: 100,
	},

	"Pine Canyon": {
		29: 100,
	},

	"Pine Flat": {
		32: 100,
	},

	"Pine Grove": {
		1: 100,
	},

	"Pine Hills": {
		2: 100,
	},

	"Pine Mountain Club": {
		32: 100,
	},

	"Pine Mountain Lake": {
		8: 100,
	},

	"Pine Valley": {
		75: 100,
	},

	"Pinole": {
		14: 100,
	},

	"Pioneer": {
		1: 100,
	},

	"Piru": {
		38: 100,
	},

	"Pismo Beach": {
		30: 100,
	},

	"Pittsburg": {
		15: 100,
	},

	"Pixley": {
		33: 100,
	},

	"Piñon Hills": {
		41: 100,
	},

	"Placentia": {
		59: 100,
	},

	"Placerville": {
		5: 100,
	},

	"Plainview": {
		33: 100,
	},

	"Planada": {
		27: 100,
	},

	"Platina": {
		1: 100,
	},

	"Pleasant Hill": {
		15: 100,
	},

	"Pleasanton": {
		16: 80,
		20: 20,
	},

	"Pleasure Point": {
		30: 100,
	},

	"Plumas Eureka": {
		1: 100,
	},

	"Plumas Lake": {
		3: 100,
	},

	"Plymouth": {
		1: 100,
	},

	"Point Arena": {
		2: 100,
	},

	"Point Reyes Station": {
		12: 100,
	},

	"Pollock Pines": {
		1: 100,
	},

	"Pomona": {
		53: 100,
	},

	"Ponderosa": {
		32: 100,
	},

	"Poplar-Cotton Center": {
		33: 100,
	},

	"Port Costa": {
		15: 100,
	},

	"Port Hueneme": {
		38: 100,
	},

	"Porterville": {
		33: 100,
	},

	"Portola": {
		1: 100,
	},

	"Portola Valley": {
		23: 100,
	},

	"Posey": {
		32: 100,
	},

	"Poso Park": {
		32: 100,
	},

	"Post Mountain": {
		2: 100,
	},

	"Potomac Park": {
		35: 100,
	},

	"Potrero": {
		75: 100,
	},

	"Potter Valley": {
		2: 100,
	},

	"Poway": {
		75: 100,
	},

	"Prattville": {
		1: 100,
	},

	"Princeton": {
		4: 100,
	},

	"Proberta": {
		3: 100,
	},

	"Prunedale": {
		29: 100,
	},

	"Pumpkin Center": {
		35: 100,
	},

	"Quartz Hill": {
		34: 100,
	},

	"Quincy": {
		1: 100,
	},

	"Rackerby": {
		3: 100,
	},

	"Rail Road Flat": {
		8: 100,
	},

	"Rainbow": {
		75: 100,
	},

	"Raisin City": {
		27: 100,
	},

	"Ramona": {
		75: 100,
	},

	"Rancho Calaveras": {
		9: 100,
	},

	"Rancho Cordova": {
		7: 100,
	},

	"Rancho Cucamonga": {
		41: 21,
		45: 0,
		50: 79,
	},

	"Rancho Mirage": {
		47: 100,
	},

	"Rancho Mission Viejo": {
		71: 100,
	},

	"Rancho Murieta": {
		9: 100,
	},

	"Rancho Palos Verdes": {
		66: 100,
	},

	"Rancho San Diego": {
		75: 100,
	},

	"Rancho Santa Fe": {
		76: 100,
	},

	"Rancho Santa Margarita": {
		71: 100,
	},

	"Rancho Tehama Reserve": {
		3: 100,
	},

	"Randsburg": {
		34: 100,
	},

	"Red Bluff": {
		3: 100,
	},

	"Red Corral": {
		1: 100,
	},

	"Redcrest": {
		2: 100,
	},

	"Redding": {
		1: 100,
	},

	"Redlands": {
		45: 53,
		47: 14,
		50: 33,
	},

	"Redondo Beach": {
		66: 100,
	},

	"Redway": {
		2: 100,
	},

	"Redwood City": {
		21: 100,
	},

	"Redwood Valley": {
		2: 100,
	},

	"Reedley": {
		31: 0,
		33: 100,
	},

	"Reliez Valley": {
		15: 100,
	},

	"Rexland Acres": {
		35: 100,
	},

	"Rialto": {
		45: 100,
		50: 0,
	},

	"Richfield": {
		3: 100,
	},

	"Richgrove": {
		33: 100,
	},

	"Richmond": {
		14: 100,
	},

	"Richvale": {
		3: 100,
	},

	"Ridgecrest": {
		32: 100,
	},

	"Ridgecrest Heights": {
		32: 100,
	},

	"Ridgemark": {
		29: 100,
	},

	"Rio Dell": {
		2: 100,
	},

	"Rio Linda": {
		6: 100,
	},

	"Rio Oso": {
		3: 100,
	},

	"Rio Vista": {
		11: 100,
	},

	"Rio del Mar": {
		30: 100,
	},

	"Ripley": {
		36: 100,
	},

	"Ripon": {
		9: 100,
	},

	"River Pines": {
		1: 100,
	},

	"Riverbank": {
		9: 100,
	},

	"Riverdale": {
		31: 100,
	},

	"Riverdale Park": {
		22: 100,
	},

	"Rivergrove": {
		32: 100,
	},

	"Riverside": {
		58: 73,
		60: 2,
		63: 24,
	},

	"Robbins": {
		3: 100,
	},

	"Robinson Mill": {
		3: 100,
	},

	"Rocklin": {
		5: 100,
	},

	"Rodeo": {
		14: 100,
	},

	"Rodriguez Camp": {
		33: 100,
	},

	"Rohnert Park": {
		12: 100,
	},

	"Rolling Hills": {
		8:  100,
		66: 100,
	},

	"Rolling Hills Estates": {
		66: 100,
	},

	"Rollingwood": {
		14: 100,
	},

	"Romoland": {
		60: 100,
	},

	"Rosamond": {
		34: 100,
	},

	"Rose Hills": {
		56: 100,
	},

	"Rosedale": {
		32: 100,
	},

	"Rosemead": {
		49: 100,
	},

	"Rosemont": {
		7: 100,
	},

	"Roseville": {
		5: 100,
	},

	"Ross": {
		12: 100,
	},

	"Rossmoor": {
		70: 100,
	},

	"Rough and Ready": {
		1: 100,
	},

	"Round Mountain": {
		1: 100,
	},

	"Round Valley": {
		8: 100,
	},

	"Rouse": {
		22: 100,
	},

	"Rowland Heights": {
		56: 100,
	},

	"Rumsey": {
		4: 100,
	},

	"Running Springs": {
		34: 100,
	},

	"Ruth": {
		2: 100,
	},

	"Rutherford": {
		4: 100,
	},

	"Sacramento": {
		6:  61,
		10: 39,
	},

	"Sage": {
		36: 100,
	},

	"Salida": {
		9: 100,
	},

	"Salinas": {
		29: 100,
	},

	"Salmon Creek": {
		2: 100,
	},

	"Salton City": {
		36: 100,
	},

	"Salton Sea Beach": {
		36: 100,
	},

	"Salyer": {
		2: 100,
	},

	"Samoa": {
		2: 100,
	},

	"San Andreas": {
		8: 100,
	},

	"San Anselmo": {
		12: 100,
	},

	"San Antonio Heights": {
		41: 100,
	},

	"San Ardo": {
		29: 100,
	},

	"San Bernardino": {
		45: 96,
		50: 4,
	},

	"San Bruno": {
		19: 4,
		21: 96,
	},

	"San Buenaventura (Ventura)": {
		38: 100,
	},

	"San Carlos": {
		21: 100,
	},

	"San Clemente": {
		74: 100,
	},

	"San Diego": {
		75: 3,
		76: 11,
		77: 21,
		78: 34,
		79: 20,
		80: 10,
	},

	"San Diego Country Estates": {
		75: 100,
	},

	"San Dimas": {
		41: 100,
	},

	"San Fernando": {
		43: 100,
	},

	"San Francisco": {
		12: 0,
		17: 59,
		18: 0,
		19: 41,
	},

	"San Gabriel": {
		49: 100,
	},

	"San Geronimo": {
		12: 100,
	},

	"San Jacinto": {
		47: 1,
		60: 99,
	},

	"San Joaquin": {
		27: 100,
	},

	"San Jose": {
		23: 6,
		24: 11,
		25: 47,
		26: 12,
		28: 24,
	},

	"San Juan Bautista": {
		29: 100,
	},

	"San Juan Capistrano": {
		74: 100,
	},

	"San Leandro": {
		20: 100,
	},

	"San Lorenzo": {
		20: 100,
	},

	"San Lucas": {
		29: 100,
	},

	"San Luis Obispo": {
		30: 100,
	},

	"San Marcos": {
		76: 100,
	},

	"San Marino": {
		49: 100,
	},

	"San Martin": {
		29: 100,
	},

	"San Mateo": {
		21: 100,
	},

	"San Miguel": {
		16: 100,
		30: 100,
	},

	"San Pablo": {
		14: 100,
	},

	"San Pasqual": {
		41: 100,
	},

	"San Rafael": {
		12: 100,
	},

	"San Ramon": {
		16: 100,
	},

	"San Simeon": {
		30: 100,
	},

	"Sand City": {
		30: 100,
	},

	"Sanger": {
		31: 100,
	},

	"Santa Ana": {
		68: 83,
		70: 17,
	},

	"Santa Barbara": {
		37: 100,
	},

	"Santa Clara": {
		26: 100,
	},

	"Santa Clarita": {
		40: 100,
	},

	"Santa Cruz": {
		28: 100,
	},

	"Santa Fe Springs": {
		64: 100,
	},

	"Santa Margarita": {
		30: 100,
	},

	"Santa Maria": {
		37: 100,
	},

	"Santa Monica": {
		51: 100,
	},

	"Santa Nella": {
		27: 100,
	},

	"Santa Paula": {
		38: 100,
	},

	"Santa Rosa": {
		2:  65,
		12: 35,
	},

	"Santa Rosa Valley": {
		42: 100,
	},

	"Santa Susana": {
		42: 100,
	},

	"Santa Venetia": {
		12: 100,
	},

	"Santa Ynez": {
		37: 100,
	},

	"Santee": {
		75: 100,
	},

	"Saranap": {
		16: 100,
	},

	"Saratoga": {
		23: 100,
	},

	"Saticoy": {
		38: 100,
	},

	"Sattley": {
		1: 100,
	},

	"Sausalito": {
		12: 100,
	},

	"Scotia": {
		2: 100,
	},

	"Scotts Valley": {
		28: 100,
	},

	"Sea Ranch": {
		2: 100,
	},

	"Seacliff": {
		30: 100,
	},

	"Seal Beach": {
		70: 28,
		72: 72,
	},

	"Searles Valley": {
		34: 100,
	},

	"Seaside": {
		30: 100,
	},

	"Sebastopol": {
		2: 100,
	},

	"Seeley": {
		36: 100,
	},

	"Selma": {
		31: 100,
	},

	"Sequoia Crest": {
		32: 100,
	},

	"Sereno del Mar": {
		2: 100,
	},

	"Seville": {
		33: 100,
	},

	"Shafter": {
		35: 100,
	},

	"Shandon": {
		30: 100,
	},

	"Shasta": {
		1: 100,
	},

	"Shasta Lake": {
		1: 100,
	},

	"Shaver Lake": {
		8: 100,
	},

	"Shell Ridge": {
		16: 100,
	},

	"Shelter Cove": {
		2: 100,
	},

	"Sheridan": {
		3: 100,
	},

	"Shingle Springs": {
		5: 100,
	},

	"Shingletown": {
		1: 100,
	},

	"Shoshone": {
		8: 100,
	},

	"Sierra Brooks": {
		1: 100,
	},

	"Sierra City": {
		1: 100,
	},

	"Sierra Madre": {
		41: 100,
	},

	"Sierra Village": {
		8: 100,
	},

	"Sierraville": {
		1: 100,
	},

	"Signal Hill": {
		69: 100,
	},

	"Silver Lakes": {
		34: 100,
	},

	"Silverado": {
		71: 100,
	},

	"Silverado Resort": {
		4: 100,
	},

	"Simi Valley": {
		42: 100,
	},

	"Sisquoc": {
		37: 100,
	},

	"Sky Valley": {
		36: 100,
	},

	"Sleepy Hollow": {
		12: 100,
	},

	"Smartsville": {
		3: 100,
	},

	"Smith Corner": {
		35: 100,
	},

	"Smith River": {
		2: 100,
	},

	"Snelling": {
		22: 100,
	},

	"Soda Bay": {
		4: 100,
	},

	"Soda Springs": {
		1: 100,
	},

	"Solana Beach": {
		77: 100,
	},

	"Soledad": {
		29: 100,
	},

	"Solvang": {
		37: 100,
	},

	"Somis": {
		38: 100,
	},

	"Sonoma": {
		4: 100,
	},

	"Sonoma State University": {
		12: 100,
	},

	"Sonora": {
		8: 100,
	},

	"Soquel": {
		30: 100,
	},

	"Soulsbyville": {
		8: 100,
	},

	"South Dos Palos": {
		27: 100,
	},

	"South El Monte": {
		56: 100,
	},

	"South Gate": {
		62: 100,
	},

	"South Lake Tahoe": {
		1: 100,
	},

	"South Monrovia Island": {
		48: 100,
	},

	"South Oroville": {
		3: 100,
	},

	"South Pasadena": {
		49: 100,
	},

	"South San Francisco": {
		19: 50,
		21: 50,
	},

	"South San Gabriel": {
		49: 100,
	},

	"South San Jose Hills": {
		48: 100,
	},

	"South Taft": {
		32: 100,
	},

	"South Whittier": {
		64: 100,
	},

	"Spaulding": {
		1: 100,
	},

	"Spreckels": {
		29: 100,
	},

	"Spring Garden": {
		1: 100,
	},

	"Spring Valley": {
		4:  100,
		79: 100,
	},

	"Spring Valley Lake": {
		34: 100,
	},

	"Springville": {
		32: 100,
	},

	"Squaw Valley": {
		8: 100,
	},

	"Squirrel Mountain Valley": {
		32: 100,
	},

	"St. Helena": {
		4: 100,
	},

	"Stallion Springs": {
		32: 100,
	},

	"Stanford": {
		23: 100,
	},

	"Stanton": {
		70: 100,
	},

	"Stebbins": {
		35: 100,
	},

	"Stevenson Ranch": {
		40: 100,
	},

	"Stevinson": {
		22: 100,
	},

	"Stinson Beach": {
		12: 100,
	},

	"Stirling City": {
		3: 100,
	},

	"Stockton": {
		13: 100,
	},

	"Stones Landing": {
		1: 100,
	},

	"Stonyford": {
		4: 100,
	},

	"Stratford": {
		33: 100,
	},

	"Strathmore": {
		33: 100,
	},

	"Strawberry": {
		8:  100,
		12: 100,
	},

	"Sugarloaf Saw Mill": {
		32: 100,
	},

	"Sugarloaf Village": {
		32: 100,
	},

	"Suisun City": {
		11: 100,
	},

	"Sultana": {
		33: 100,
	},

	"Summerland": {
		37: 100,
	},

	"Sun Village": {
		39: 100,
	},

	"Sunny Slopes": {
		8: 100,
	},

	"Sunnyside": {
		31: 100,
	},

	"Sunnyside-Tahoe City": {
		1: 100,
	},

	"Sunnyvale": {
		26: 100,
	},

	"Sunol": {
		24: 100,
	},

	"Susanville": {
		1: 100,
	},

	"Sutter": {
		3: 100,
	},

	"Sutter Creek": {
		1: 100,
	},

	"Swall Meadows": {
		8: 100,
	},

	"Taft": {
		32: 100,
	},

	"Taft Heights": {
		32: 100,
	},

	"Taft Mosswood": {
		13: 100,
	},

	"Tahoe Vista": {
		1: 100,
	},

	"Tahoma": {
		1: 100,
	},

	"Talmage": {
		2: 100,
	},

	"Tamalpais-Homestead Valley": {
		12: 100,
	},

	"Tancred": {
		4: 100,
	},

	"Tara Hills": {
		14: 100,
	},

	"Tarina": {
		32: 100,
	},

	"Tarpey Village": {
		8:  82,
		31: 18,
	},

	"Taylorsville": {
		1: 100,
	},

	"Tecopa": {
		8: 100,
	},

	"Tehachapi": {
		32: 100,
	},

	"Tehama": {
		3: 100,
	},

	"Temecula": {
		71: 100,
	},

	"Temelec": {
		12: 100,
	},

	"Temescal Valley": {
		63: 100,
	},

	"Temple City": {
		49: 100,
	},

	"Templeton": {
		30: 100,
	},

	"Tennant": {
		1: 100,
	},

	"Terminous": {
		9: 100,
	},

	"Terra Bella": {
		33: 100,
	},

	"Teviston": {
		33: 100,
	},

	"Thermal": {
		36: 100,
	},

	"Thermalito": {
		3: 100,
	},

	"Thornton": {
		9: 100,
	},

	"Thousand Oaks": {
		42: 100,
	},

	"Thousand Palms": {
		47: 100,
	},

	"Three Rivers": {
		32: 100,
	},

	"Three Rocks": {
		27: 100,
	},

	"Tiburon": {
		12: 100,
	},

	"Timber Cove": {
		2: 100,
	},

	"Tipton": {
		33: 100,
	},

	"Tobin": {
		1: 100,
	},

	"Tomales": {
		12: 100,
	},

	"Tonyville": {
		32: 100,
	},

	"Tooleville": {
		32: 100,
	},

	"Topanga": {
		42: 100,
	},

	"Topaz": {
		8: 100,
	},

	"Toro Canyon": {
		37: 100,
	},

	"Torrance": {
		66: 100,
	},

	"Trabuco Canyon": {
		71: 100,
	},

	"Tracy": {
		13: 100,
	},

	"Tranquillity": {
		27: 100,
	},

	"Traver": {
		33: 100,
	},

	"Tres Pinos": {
		29: 100,
	},

	"Trinidad": {
		2: 100,
	},

	"Trinity Center": {
		2: 100,
	},

	"Trinity Village": {
		2: 100,
	},

	"Trona": {
		8: 100,
	},

	"Trowbridge": {
		3: 100,
	},

	"Truckee": {
		1: 100,
	},

	"Tulare": {
		33: 100,
	},

	"Tulelake": {
		1: 100,
	},

	"Tuolumne City": {
		8: 100,
	},

	"Tupman": {
		32: 100,
	},

	"Turlock": {
		22: 100,
	},

	"Tustin": {
		59: 0,
		73: 100,
	},

	"Tuttle": {
		27: 100,
	},

	"Tuttletown": {
		8: 100,
	},

	"Twain": {
		1: 100,
	},

	"Twain Harte": {
		8: 100,
	},

	"Twentynine Palms": {
		34: 100,
	},

	"Twin Lakes": {
		8:  100,
		30: 100,
	},

	"Ukiah": {
		2: 100,
	},

	"Union City": {
		20: 100,
	},

	"University of California-Davis": {
		4: 100,
	},

	"University of California-Santa Barbara": {
		37: 100,
	},

	"Upland": {
		41: 51,
		53: 49,
	},

	"Upper Lake": {
		4: 100,
	},

	"Vacaville": {
		11: 100,
	},

	"Val Verde": {
		40: 100,
	},

	"Valinda": {
		48: 100,
	},

	"Valle Vista": {
		47: 100,
	},

	"Vallecito": {
		8: 100,
	},

	"Vallejo": {
		11: 100,
	},

	"Valley Acres": {
		32: 100,
	},

	"Valley Center": {
		75: 100,
	},

	"Valley Ford": {
		2: 100,
	},

	"Valley Home": {
		9: 100,
	},

	"Valley Ranch": {
		1: 100,
	},

	"Valley Springs": {
		8: 100,
	},

	"Vandenberg AFB": {
		37: 100,
	},

	"Vandenberg Village": {
		37: 100,
	},

	"Verdi": {
		1: 100,
	},

	"Vernon": {
		54: 100,
	},

	"Victor": {
		9: 100,
	},

	"Victorville": {
		34: 16,
		39: 84,
	},

	"View Park-Windsor Hills": {
		55: 100,
	},

	"Villa Park": {
		59: 100,
	},

	"Vina": {
		3: 100,
	},

	"Vincent": {
		48: 100,
	},

	"Vine Hill": {
		15: 100,
	},

	"Vineyard": {
		10: 100,
	},

	"Virginia Lakes": {
		8: 100,
	},

	"Visalia": {
		32: 79,
		33: 21,
	},

	"Vista": {
		74: 100,
	},

	"Vista Santa Rosa": {
		36: 100,
	},

	"Volcano": {
		1: 100,
	},

	"Volta": {
		27: 100,
	},

	"Walker": {
		8: 100,
	},

	"Wallace": {
		9: 100,
	},

	"Walnut": {
		56: 100,
	},

	"Walnut Creek": {
		15: 0,
		16: 100,
	},

	"Walnut Grove": {
		9: 100,
	},

	"Walnut Park": {
		62: 100,
	},

	"Warm Springs": {
		63: 100,
	},

	"Warner Valley": {
		1: 100,
	},

	"Wasco": {
		35: 100,
	},

	"Washington": {
		1: 100,
	},

	"Waterford": {
		9: 100,
	},

	"Waterloo": {
		9: 100,
	},

	"Watsonville": {
		29: 100,
		30: 0,
	},

	"Waukena": {
		33: 100,
	},

	"Wautec": {
		2: 100,
	},

	"Wawona": {
		8: 100,
	},

	"Weaverville": {
		2: 100,
	},

	"Weed": {
		1: 100,
	},

	"Weedpatch": {
		35: 100,
	},

	"Weitchpec": {
		2: 100,
	},

	"Weldon": {
		32: 100,
	},

	"Weott": {
		2: 100,
	},

	"West Athens": {
		61: 100,
	},

	"West Bishop": {
		8: 100,
	},

	"West Carson": {
		65: 100,
	},

	"West Covina": {
		48: 100,
	},

	"West Goshen": {
		33: 100,
	},

	"West Hollywood": {
		51: 100,
	},

	"West Menlo Park": {
		23: 100,
	},

	"West Modesto": {
		22: 100,
	},

	"West Park": {
		31: 100,
	},

	"West Point": {
		8: 100,
	},

	"West Puente Valley": {
		48: 100,
	},

	"West Rancho Dominguez": {
		65: 100,
	},

	"West Sacramento": {
		4: 100,
	},

	"West Whittier-Los Nietos": {
		56: 100,
	},

	"Westhaven-Moonstone": {
		2: 100,
	},

	"Westlake Village": {
		42: 100,
	},

	"Westley": {
		22: 100,
	},

	"Westminster": {
		70: 100,
	},

	"Westmont": {
		61: 100,
	},

	"Westmorland": {
		36: 100,
	},

	"Westside": {
		27: 100,
	},

	"Westwood": {
		1: 100,
	},

	"Wheatland": {
		3: 100,
	},

	"Whitehawk": {
		1: 100,
	},

	"Whitewater": {
		47: 100,
	},

	"Whitley Gardens": {
		30: 100,
	},

	"Whitmore": {
		1: 100,
	},

	"Whittier": {
		56: 100,
	},

	"Wildomar": {
		71: 100,
	},

	"Wilkerson": {
		8: 100,
	},

	"Williams": {
		4: 100,
	},

	"Williams Canyon": {
		71: 100,
	},

	"Willits": {
		2: 100,
	},

	"Willow Creek": {
		2: 100,
	},

	"Willowbrook": {
		65: 100,
	},

	"Willows": {
		3: 100,
	},

	"Wilsonia": {
		32: 100,
	},

	"Wilton": {
		9: 100,
	},

	"Winchester": {
		36: 100,
	},

	"Windsor": {
		2: 100,
	},

	"Winter Gardens": {
		75: 100,
	},

	"Winterhaven": {
		36: 100,
	},

	"Winters": {
		4: 100,
	},

	"Winton": {
		27: 100,
	},

	"Wofford Heights": {
		32: 100,
	},

	"Woodacre": {
		12: 100,
	},

	"Woodbridge": {
		9: 100,
	},

	"Woodcrest": {
		63: 100,
	},

	"Woodlake": {
		33: 100,
	},

	"Woodland": {
		4: 100,
	},

	"Woodlands": {
		37: 100,
	},

	"Woodside": {
		23: 100,
	},

	"Woodville": {
		33: 100,
	},

	"Woodville Farm Labor Camp": {
		33: 100,
	},

	"Woody": {
		32: 100,
	},

	"Wrightwood": {
		41: 100,
	},

	"Yankee Hill": {
		3: 100,
	},

	"Yermo": {
		34: 100,
	},

	"Yettem": {
		33: 100,
	},

	"Yolo": {
		4: 100,
	},

	"Yorba Linda": {
		59: 100,
	},

	"Yosemite Lakes": {
		8: 100,
	},

	"Yosemite Valley": {
		8: 100,
	},

	"Yosemite West": {
		8: 100,
	},

	"Yountville": {
		4: 100,
	},

	"Yreka": {
		1: 100,
	},

	"Yuba City": {
		3: 100,
	},

	"Yucaipa": {
		47: 100,
	},

	"Yucca Valley": {
		47: 100,
	},

	"Zayante": {
		28: 100,
	},
}
