package main

func main() {

}
func isIsomorphic(s string, t string) bool {
	s2t := [128]byte{}
	t2s := [128]byte{}
	for i, end := 0, len(s); i < end; i++ {
		if s2t[s[i]] != t2s[t[i]] {
			return false
		}
		s2t[s[i]], t2s[t[i]] = byte(i+1), byte(i+1)
	}
	return true
}

/*


badc
baba

b=b
a=a
d=b
c=a



concat(olar, denis)
return concat

output:
y/N?
Y>?N/?
= hotties?
y/N?
y
= coupulies?
y/N?
ERROR: probably is not input use y/N
N

output:

|----
|   |
|   o
|  /O|
|   Л
    П

*/
