#PF-Assgn-42
def find_factors(num):
    factors = []
    for i in range(2,(num+1)):
        if(num%i==0):
            factors.append(i)
    return factors

def is_prime(num, i):
    if(i==1):
        return True
    elif(num%i==0):
        return False;
    else:
        return(is_prime(num,i-1))

def find_largest_prime_factor(list_of_factors):
    max_fac = list_of_factors[0]
    for i in list_of_factors:
        if is_prime(i) and i>max_fac:
            max_fac = i
            

# def find_f(num):
   
# def find_g(num):
print()
print(find_g(10))