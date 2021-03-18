package main

//Member 멤버인터페이스
type Member interface {
	getid() string
	getrank() string
}

//SKMember SK멤버
type SKMember struct {
	id, rank string
}

//SSMember SAMSUNG멤버
type SSMember struct {
	id, rank string
}

//SK member 타입에 대한 member 인터페이스 구현
func (skm SKMember) getid() string   { return "SK member id: " + skm.id }
func (skm SKMember) getrank() string { return "SK member rank: " + skm.rank }

//SAMSUNG member 타입에 대한 member 인터페이스 구현
func (ssm SSMember) getid() string   { return "SAMSUNG member id: " + ssm.id }
func (ssm SSMember) getrank() string { return "SAMSUNG member rank: " + ssm.rank }

func main() {
	skm := SKMember{"magic1", "1"}
	ssm := SSMember{"magic2", "1"}

	printMember(skm, ssm)
}

func printMember(members ...Member) {
	for _, m := range members {
		// interface 메소드를 호출함.
		id := m.getid()
		rank := m.getrank()

		println(id)
		println(rank)
	}
}
