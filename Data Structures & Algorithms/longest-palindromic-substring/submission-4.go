func longestPalindrome(s string) string {
    // Reverse thinking
	// Say the biggest string is of n:=len(s)
	// now starting with i:=n, i run an algorithm which finds a palindrome of size i in the string

	// start at a point and then compare
	n:=len(s)
	maxL:=0
	a,b:=0,0
	// We need to output the string
	for i:=0;i<n;i++{
		l,r:=i-1,i+1
		for l>=0&&r<n&&s[l]==s[r]{
			if maxL<=len(s[l:r+1]){
				maxL=len(s[l:r+1])
				a=l
				b=r
			}
			
			l--
			r++
		}
		l,r=i,i+1
		for l>=0&&r<n&&s[l]==s[r]{
			if maxL<=len(s[l:r+1]){
				maxL=len(s[l:r+1])
				a=l
				b=r
			}
			
			l--
			r++
		}
	}
return s[a:b+1]
}
