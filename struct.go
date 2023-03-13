package piscine

type Avion struct {
	Matricule         string
	N_de_passager_max int
	Duree_repo        int  //en minute
	En_Vol            bool //est true si l'avion est actuellement entrain d'effectuer un vol
}

type Vol struct {
	Avion         *Avion
	N_de_passager int
	VilleA        string
	VilleB        string
	Temps_de_vol  int //en minute
	Heure_Depart  int //en minute APRES le premier depart générale des avions
}

func C_Avion() []Avion {
	var ret []Avion
	var Avion1 Avion
	var Avion2 Avion
	var Avion3 Avion

	Avion1.Matricule = "B7A113"
	Avion1.N_de_passager_max = 122
	Avion1.Duree_repo = 24

	Avion2.Matricule = "P14F56"
	Avion2.N_de_passager_max = 87
	Avion2.Duree_repo = 15

	Avion3.Matricule = "MJK756"
	Avion3.N_de_passager_max = 178
	Avion3.Duree_repo = 43

	ret = append(ret, Avion1)
	ret = append(ret, Avion2)
	ret = append(ret, Avion3)

	return ret
}

func C_Vol() []Vol {
	var ret []Vol
	var Vol1 Vol
	var Vol2 Vol
	var Vol3 Vol
	var Vol4 Vol
	var Vol5 Vol
	var Vol6 Vol
	var Vol7 Vol
	var Vol8 Vol
	var Vol9 Vol
	var Vol10 Vol
	var Vol11 Vol
	var Vol12 Vol

	Vol1.Heure_Depart = 0
	Vol1.Temps_de_vol = 99
	Vol1.VilleA = "Paris"
	Vol1.VilleB = "Sofia"
	Vol1.N_de_passager = 91

	Vol2.Heure_Depart = 3
	Vol2.Temps_de_vol = 186
	Vol2.VilleA = "Paris"
	Vol2.VilleB = "Moscou"
	Vol2.N_de_passager = 120

	Vol3.Heure_Depart = 12
	Vol3.Temps_de_vol = 255
	Vol3.VilleA = "Paris"
	Vol3.VilleB = "New_York"
	Vol3.N_de_passager = 160

	Vol4.Heure_Depart = 80
	Vol4.Temps_de_vol = 50
	Vol4.VilleA = "Paris"
	Vol4.VilleB = "Dublin"
	Vol4.N_de_passager = 40

	Vol5.Heure_Depart = 190
	Vol5.Temps_de_vol = 56
	Vol5.VilleA = "Paris"
	Vol5.VilleB = "Madrid"
	Vol5.N_de_passager = 88

	Vol6.Heure_Depart = 315
	Vol6.Temps_de_vol = 135
	Vol6.VilleA = "Paris"
	Vol6.VilleB = "Alger"
	Vol6.N_de_passager = 60

	Vol7.Heure_Depart = 344
	Vol7.Temps_de_vol = 42
	Vol7.VilleA = "Paris"
	Vol7.VilleB = "Marseille"
	Vol7.N_de_passager = 90

	Vol8.Heure_Depart = 387
	Vol8.Temps_de_vol = 57
	Vol8.VilleA = "Paris"
	Vol8.VilleB = "Athene"
	Vol8.N_de_passager = 111

	Vol9.Heure_Depart = 414
	Vol9.Temps_de_vol = 140
	Vol9.VilleA = "Paris"
	Vol9.VilleB = "Le_Caire"
	Vol9.N_de_passager = 99

	Vol10.Heure_Depart = 452
	Vol10.Temps_de_vol = 48
	Vol10.VilleA = "Paris"
	Vol10.VilleB = "Londre"
	Vol10.N_de_passager = 22

	Vol11.Heure_Depart = 489
	Vol11.Temps_de_vol = 77
	Vol11.VilleA = "Paris"
	Vol11.VilleB = "Rome"
	Vol11.N_de_passager = 94

	Vol12.Heure_Depart = 501
	Vol12.Temps_de_vol = 49
	Vol12.VilleA = "Paris"
	Vol12.VilleB = "Copenhague"
	Vol12.N_de_passager = 47

	ret = append(ret, Vol1)
	ret = append(ret, Vol2)
	ret = append(ret, Vol3)
	ret = append(ret, Vol4)
	ret = append(ret, Vol5)
	ret = append(ret, Vol6)
	ret = append(ret, Vol7)
	ret = append(ret, Vol8)
	ret = append(ret, Vol9)
	ret = append(ret, Vol10)
	ret = append(ret, Vol11)
	ret = append(ret, Vol12)

	return ret
}

// infos importantes:
//		-la journée de travaille commence a 8h, et finis a 18h
//		-quand un avions a finis un traget, nous considererons qu'il est de nouveau oppérationelle a Paris apres son temps de repo
//		-un avion ne peut prendre que des vol avec un nombre de passager en dessous de son nombre max de passager
//		-La fonction C_Avion definis les 3 avions que vous avez a disposition
//		-La fonction C_Vol definis les 12 vol a effectuer dans la journée
//
// La consigne:
//		Atribuer des avions a la variable Avion de chaque vol afin que tout les vol puisse etre effectué dans la journer,
//		et se , en considerant le moins de retard ppossible
//
//		Trouver quelle avion a transporter le plus de passager a la fin de la journée
//		Trouver quelle avion a eu le plus de temps de vol
//
// direction a prendre pour aider:
// 		-créer une fonction qui change a variable En_Vol des avions quand ils decolent
//		-créer une boucle de d'execution qui par de 8h a 18h (0->600 minutes)
//		-tenter des choses !
//
//
//		IL N'Y A PAS DE BONNE REPONSE, JUSTE ESSAYEZ !
