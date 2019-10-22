package piscine

func TrimAtoi(s string) int {
	count:=0
	znak:=1
	res:=0
	for _, k:=range s {
        if k=='-' && count==0 {
        	znak=-1
        }

        if k=='+' && count==0 {
        	znak=1
        }

		if ('0'<=k && k<='9') {
			count++

			res=res*10+int(k-'0')
		}

	}

res=res*znak	
return res

}