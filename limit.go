package rate

import (
	"sync"
	"time"
)

// 计数器
// 在一段时间内，对请求进行计数，与阀值进行比较判断是否需要限流，到时间后计数清零
// 缺陷：无法处理临界值得问题，比如一分钟限制访问1000次，这1000次可以在一秒内访问成功

type LimitRate struct {
	rate  int           //阀值
	begin time.Time     //计数开始时间
	cycle time.Duration //计数周期
	count int           //收到的请求数
	lock  sync.Mutex    //锁
}

func (limit *LimitRate) Allow() bool {
	limit.lock.Lock()
	defer limit.lock.Unlock()

	// 判断收到请求数是否达到阀值
	if limit.count == limit.rate-1 {
		now := time.Now()
		// 达到阀值后，判断是否是请求周期内
		if now.Sub(limit.begin) >= limit.cycle {
			limit.Reset(now)
			return true
		}
		return false
	} else {
		limit.count++
		return true
	}
}

func (limit *LimitRate) Set(rate int, cycle time.Duration) {
	limit.rate = rate
	limit.begin = time.Now()
	limit.cycle = cycle
	limit.count = 0
}

func (limit *LimitRate) Reset(begin time.Time) {
	limit.begin = begin
	limit.count = 0
}
