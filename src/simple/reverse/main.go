package main

/**
给出一个32位的有符号整数，你需要这个整数中每位上的数字进行反转。
示例1：
	输入：123
	输出：321

示例2：
	输入：-123
	输出：-321

示例3：
	输入：120
	输出：21

注意：假设我们的环境只能存储得下32位的有符号整数，则其数值取值范围为[-2^31, 2^31-1]。请根据这个假设，如果反转后整数
溢出那么就返回0.
 */

func reverse(x int) int{
	val := x
	result:=0
	for ;val !=0;{
		result = result*10+val%10
		if !(-(1<<31)<=result && result<=(1<<31)-1){
			return 0
		}
		val = val/10
	}

	return result
}

func main(){
	print( reverse( 120))
}
