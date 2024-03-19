package gos7

import (
	"fmt"
	"sort"

	"github.com/snple/kirara/pb"
)

func groupTags(tags map[string]*Tag) map[string][]readGroup {
	tags2 := make([]*Tag, 0, len(tags))

	for _, tag := range tags {
		tags2 = append(tags2, tag)
	}

	sort.Sort(tagsSort(tags2))

	return group(tags2)
}

type Tag struct {
	raw  *pb.Tag
	addr S7Addr
}

func newTag(tag *pb.Tag) *Tag {
	return &Tag{
		raw: tag,
	}
}

func (t *Tag) parse() {
	t.addr = ParseTagAddress(t.raw.Address)
}

// sort
type tagsSort []*Tag

func (s tagsSort) Len() int {
	return len(s)
}

func (s tagsSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s tagsSort) Less(i, j int) bool {
	return s[i].addr.Address < s[j].addr.Address
}

type readGroup struct {
	area    int
	db      int
	start   int
	end     int
	tagList []*Tag
}

const block = 256

func group(tags []*Tag) map[string][]readGroup {
	groupMap := make(map[string][]readGroup)

	for _, tag := range tags {
		key := fmt.Sprintf("%v:%v", tag.addr.Area, tag.addr.DB)

		if groups, ok := groupMap[key]; ok {
			if (tag.addr.Address + tag.addr.Size - groups[len(groups)-1].start) < block {
				groups[len(groups)-1].end = tag.addr.Address + tag.addr.Size - 1
				groups[len(groups)-1].tagList = append(groups[len(groups)-1].tagList, tag)
			} else {
				group := readGroup{
					area:    tag.addr.Area,
					db:      tag.addr.DB,
					start:   tag.addr.Address,
					end:     tag.addr.Address + tag.addr.Size - 1,
					tagList: []*Tag{tag},
				}

				groupMap[key] = append(groupMap[key], group)
			}
		} else {
			group := readGroup{
				area:    tag.addr.Area,
				db:      tag.addr.DB,
				start:   tag.addr.Address,
				end:     tag.addr.Address + tag.addr.Size - 1,
				tagList: []*Tag{tag},
			}

			groupMap[key] = []readGroup{group}
		}
	}

	return groupMap
}

// type writeQueue struct {
// 	lock  sync.Mutex
// 	queue queue.Queue[writeTask]
// }

// func newWriteQueue() *writeQueue {
// 	return &writeQueue{
// 		queue: queue.New[writeTask](),
// 	}
// }

// func (q *writeQueue) Push(task writeTask) {
// 	q.lock.Lock()
// 	defer q.lock.Unlock()

// 	q.queue.Push(task)
// }

// func (q *writeQueue) Pop() types.Option[writeTask] {
// 	q.lock.Lock()
// 	defer q.lock.Unlock()

// 	return q.queue.Pop()
// }

// type writeTask struct {
// 	Id    string
// 	Value string
// 	Retry int
// }
