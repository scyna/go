git tag $1
git push origin $1
GOPROXY=proxy.golang.org go list -m github.com/scyna/go@$1
