package core

import (
	"strings"
)

type Handle func(c *GContent)
type WebRouter struct {
	groups     map[string]*WebRouter
	methods    map[string][]Handle
	middlefunc []Handle
}

func newRootRouter(args ...Handle) *WebRouter {
	return &WebRouter{
		groups:     map[string]*WebRouter{},
		methods:    map[string][]Handle{},
		middlefunc: args,
	}
}

func (r *WebRouter) Group(path string, args ...Handle) *WebRouter {
	for strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	path = strings.ToLower(path)
	nr := &WebRouter{
		groups:     map[string]*WebRouter{},
		methods:    map[string][]Handle{},
		middlefunc: append(r.middlefunc, args...),
	}
	r.groups[path] = nr
	return nr
}
func (r *WebRouter) Get(path string, hand Handle, args ...Handle) {
	r.addmethods("get", path, hand, args...)
}
func (r *WebRouter) Post(path string, hand Handle, args ...Handle) {
	r.addmethods("post", path, hand, args...)
}
func (r *WebRouter) Delete(path string, hand Handle, args ...Handle) {
	r.addmethods("delete", path, hand, args...)
}
func (r *WebRouter) Put(path string, hand Handle, args ...Handle) {
	r.addmethods("put", path, hand, args...)
}
func (r *WebRouter) Options(path string, hand Handle, args ...Handle) {
	r.addmethods("options", path, hand, args...)
}
func (r *WebRouter) Head(path string, hand Handle, args ...Handle) {
	r.addmethods("head", path, hand, args...)
}
func (r *WebRouter) Any(path string, hand Handle, args ...Handle) {
	r.addmethods("any", path, hand, args...)
}
func (r *WebRouter) addmethods(methodname, path string, hand Handle, args ...Handle) {
	for strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	path = strings.ToLower(path)
	if item, ok := r.groups[path]; ok {
		item.methods[strings.ToUpper(methodname)] = append([]Handle{hand}, args...)
	} else {
		nr := &WebRouter{
			groups:     map[string]*WebRouter{},
			methods:    map[string][]Handle{strings.ToUpper(methodname): append([]Handle{hand}, args...)},
			middlefunc: r.middlefunc,
		}
		r.groups[path] = nr
	}

}
func (r *WebRouter) FindHandle(methodname, path string) (bool, []Handle) {
	for strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	methodname = strings.ToUpper(methodname)
	path = strings.ToLower(path)
	ps := strings.Split(path, "/")
	var rt *WebRouter
	funcs := r.middlefunc
	for _, k := range ps {
		if rt == nil {
			rt = r.findPath(k)
			if rt == nil {
				break
			}
		} else {
			funcs = rt.middlefunc
			rt = rt.findPath(k)
			if rt == nil {
				break
			}
		}
	}
	if rt == nil {
		return false, funcs
	}
	if ffs, on := rt.methods[methodname]; on {
		return true, ffs
	} else if ffs, on := rt.methods["ANY"]; on {
		return true, ffs
	}
	return false, nil
}
func (r *WebRouter) findPath(name string) *WebRouter {
	if rr, ok := r.groups[name]; ok {
		return rr
	}
	return nil
}
