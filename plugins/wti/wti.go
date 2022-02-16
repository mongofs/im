// 构建tag 标签的这个功能，主要目标希望将用户分类起来，1000000 * 8 /1024 / 1024  = 7M的内存，额外使用tag包进行封装的话，
// 预计使用一百万用户的内存需要大概7~8M，这个范围是可以接受的。
package wti

import (
	"github.com/mongofs/im/client"
	"sync"
)

type tg struct {
	mp map[string] *Group // wti => []string
	rw *sync.RWMutex
}


func newwti() WTI {
	return &tg{
		mp: map[string]*Group{},
		rw: &sync.RWMutex{},
	}
}


// 给用户设置标签
func (t *tg) SetTAG(cli *client.Cli, tags ...string) {
	if len(tags)== 0 {
		return
	}
	t.rw.Lock()
	defer t.rw.Unlock()
	for _,tag := range tags{
		if group,ok:= t.mp[tag];!ok { // wti not exist
			 t.mp[tag]= NewGroup()
			 t.mp[tag].addCli(cli)
		}else { // wti exist
			group.addCli(cli)
		}
	}
}


// 给某一个标签的群体进行广播
func (t *tg) BroadCast(content []byte,tags ...string) {
	if len(tags)== 0 {
		return
	}
	t.rw.RLock()
	defer t.rw.RUnlock()

	for _,tag := range tags{
		if group,ok := t.mp[tag];ok{
			group.broadCast(content)
		}
	}
}


// 通知所有组进行自查某个用户，并删除
func (t *tg)Update(token ...string){
	t.rw.RLock()
	defer t.rw.RUnlock()
	for _,v := range t.mp {
		v.Update(token... )
	}
}


// 调用这个就可以分类广播，可能出现不同的targ 需要不同的内容
func(t *tg)BroadCastByTarget(targetAndContent map[string][]byte){
	if len(targetAndContent) == 0{ return }
	for target ,content := range targetAndContent {
		t.BroadCast(content,target)
	}
}


// 获取到用户token的所有TAG，时间复杂度是O(m) ,m是所有的房间
func (t *tg)GetClienterTAGs(token string)[]string{
	var res []string
	t.rw.RLock()
	defer t.rw.RUnlock()
	for k,v:= range t.mp{
		exsit:= v.isExsit(token)
		if exsit {
			res = append(res,k )
		}
	}
	return res
}

// 如果创建时间为0 ，表示没有这个房间
func (t *tg) GetTAGCreateTime(tag string) int64 {
	t.rw.RLock()
	defer t.rw.RUnlock()
	if v,ok:=t.mp[tag];ok{
		return v.createTime
	}
	return 0
}

// 获取到tag的总人数
func (t *tg) GetTAGClients(tag string) int64 {
	t.rw.RLock()
	defer t.rw.RUnlock()
	if v,ok:=t.mp[tag];ok{
		return v.Counter()
	}
	return 0
}

// 删除tag
func (t *tg) RecycleTAG(tag string) {
	t.rw.Lock()
	defer t.rw.Unlock()
	delete(t.mp,tag)
}