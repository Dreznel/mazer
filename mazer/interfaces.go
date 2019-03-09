package mazer

type Room interface {
	DoChallenge(Character) bool
	//RewardTreasure()
}

type Character interface {
	Attack() int
	UseSkill() int
	UseMagic() int
	//Flee() int
	GetTreasure(amount int)
	IsDead() bool
	ChangeHp(amount int)
	ChangeStamina(amount int)
	PrintStats()
	GetName() string
}
