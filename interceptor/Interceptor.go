package interceptor

import (
	"github.com/gentwolf-shen/gohelper/logger"
	"regexp"
)

type Interceptor struct {
	pathPatterns        []*regexp.Regexp
	excludePathPatterns []*regexp.Regexp
}

func NewInterceptor() *Interceptor {
	obj := &Interceptor{}
	obj.pathPatterns = make([]*regexp.Regexp, 0)
	obj.excludePathPatterns = make([]*regexp.Regexp, 0)
	return obj
}

func (this *Interceptor) AddPathPatterns(paths ...string) *Interceptor {
	for _, path := range paths {
		this.pathPatterns = append(this.pathPatterns, regexp.MustCompile(path))
	}

	return this
}

func (this *Interceptor) ExcludePathPatterns(paths ...string) *Interceptor {
	for _, path := range paths {
		this.excludePathPatterns = append(this.excludePathPatterns, regexp.MustCompile(path))
	}

	return this
}

func (this *Interceptor) IsMustAuthorize(path []byte) bool {
	bl := this.isMatched(path, this.pathPatterns)
	logger.Debugf("isMatched [include] %s -> %v", string(path), bl)
	if bl {
		tmp := this.isMatched(path, this.excludePathPatterns)
		logger.Debugf("isMatched [exclude] %s -> %v", string(path), tmp)
		bl = !tmp
	}

	return bl
}

func (this *Interceptor) isMatched(path []byte, items []*regexp.Regexp) bool {
	bl := false
	for _, ptn := range items {
		bl = ptn.Match(path)
		if bl {
			break
		}
	}

	return bl
}
