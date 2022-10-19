func main{
	//分析，素数的定义,//判断一个数是不是素数
	
	a := 17
	b := 2

	for b<a {
		if b%a == 0{
			break
		}
		b ++
	}

	if b==a{
		fmt.printf("YES")
	}
	else{
		fmt.printf("FALSE")
	}
	return 0 
}