package main

import(
	"fmt"
)
type elemType int

func bubble(src []elemType) {
	var t elemType
	for i:=0;i<len(src)-1;i++{
		for j:=i+1;j<len(src);j++{
			if src[i]>src[j]{
				t=src[i]
				src[i]=src[j]
				src[j]=t
			}
		}
	}
	fmt.Println(src)
}

func quick(src []elemType){
	if len(src)<2{
		return
	}
	var p = src[0]
	var i,j = 0,len(src)-1
	for i<j{
		if src[j] >= p {
			j--	
		}
		src[i] = src[j]
		if i<j && src[i] <=p{
			i++		
		}
		src[j]=src[i]
	}
	src[i]=p
	quick(src[:i])
	quick(src[i+1:])
}

func insert(src []elemType){
	var l=len(src)
	var t elemType
	for i:=1;i<l;i++{
		for j:=0;j<i;j++{
		    if src[i]>=src[j]{
		    	j++
				continue
		    }
		    t=src[i]
		    src[i]=src[j]
		    src[j]=t
		}
	}
}

var index int=-1
func bsearch(src []elemType,target elemType,start int){
	var l=len(src)
	
	if l<2{
		if target ==src[0]{
			index=start
		}
		return
	}
	if target > src[l/2]{
		bsearch(src[l/2+1:],target,start+l/2+1)
	}else if target < src[l/2]{
		bsearch(src[:l/2],target,start)
	}else{
		index=l/2+start
	}
	return
}

func main() {
	src:=[]elemType{20,8,16,3,9,25}
	//bubble(src)
	//quick(src)
	insert(src)
	fmt.Println(src)
	bsearch(src,11,0)
	fmt.Println(index)
}


