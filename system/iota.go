package system

//注意iota逐行增加
const (
	a = 1<<iota		//a==1 (iota ==0)
	b=  1<<iota		//b==2 (iota ==1)
	c=	3			//c==3 (iota==2,unused)
	d=	1<<iota		//d==8 (iota==3)
)

