package main

/***
给定一个整数数组 怒貌似和一个目标值 target，请你在数组中找出和为目标值的两个整数，并返回他们的数组下标
你可以根据假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例：
	给定nums =[2,7,11,15], target=9
	因为nums[0]+num[1]=2+7=9
	所以返回[0,1]
 */

func TwoSum(nums []int, target int ) (bool, []int){
	v := make(map[int]int)
	for i:=0;i<len(nums);i++ {
		if value, has := v[target-nums[i]]; has {
			return true, []int{i,value}
		}
		v[nums[i]]=i
	}
	return false, nil
}

func main(){

	nums := []int{ 1,2,3,45,12}
	ret, val := TwoSum(nums, 13)

	if ret {
		print("result:")
		print(val[0], val[1])
	}else {
		print("result: null")
	}

}
