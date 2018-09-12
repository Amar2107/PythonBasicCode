def find_ten_substring(num_str):
	lst=[]
	for i in range(0,len(num_str)):
		sum_val =0
		substr=""
		for j in range(i,len(num_str)):
			substr+=num_str[j]
			sum_of_substr=list(map(int,substr))
			sum_val=sum(sum_of_substr)
			if sum_val == 10:
				lst.append(substr)
	return  lst

num_str="3523014"
print("The number is:",num_str)
result_list=find_ten_substring(num_str)
print(result_list)