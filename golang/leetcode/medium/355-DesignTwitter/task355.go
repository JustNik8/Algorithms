package main

import "container/heap"

func main() {

}

type Tweet struct {
	TweetId int
	Count   int
}

type TweetHeap []Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x any)        { *h = append(*h, x.(Tweet)) }
func (h *TweetHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Followee map[int]struct{}

type Twitter struct {
	Follows map[int]Followee
	Tweets  map[int][]Tweet
	Count   int
}

func Constructor() Twitter {
	return Twitter{
		Follows: make(map[int]Followee),
		Tweets:  make(map[int][]Tweet),
		Count:   0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Count++
	tweet := Tweet{TweetId: tweetId, Count: this.Count}
	this.Tweets[userId] = append(this.Tweets[userId], tweet)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	newsFeed := make([]Tweet, 0)
	for followee, _ := range this.Follows[userId] {
		newsFeed = append(newsFeed, this.Tweets[followee]...)
	}
	newsFeed = append(newsFeed, this.Tweets[userId]...)

	maxHeap := TweetHeap(newsFeed)
	heap.Init(&maxHeap)

	ansSize := min(10, len(maxHeap))
	ans := make([]int, ansSize)
	for i := 0; i < ansSize; i++ {
		ans[i] = heap.Pop(&maxHeap).(Tweet).TweetId
	}

	return ans
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	_, hasFollows := this.Follows[followerId]
	if !hasFollows {
		this.Follows[followerId] = make(Followee)
	}

	this.Follows[followerId][followeeId] = struct{}{}

}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.Follows[followerId], followeeId)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
