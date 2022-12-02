package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis interface {
	//Set 命令用于设置给定 key 的值。如果 key 已经存储其他值， SET 就覆写旧值，且无视类型。	//SET KEY_NAME VALUE
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	//SetNx 命令在指定的 key 不存在时，为 key 设置指定的值。 //SETNX KEY_NAME VALUE
	SetNx(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	//Get 命令用于获取指定 key 的值。如果 key 不存在，返回 nil 。如果key 储存的值不是字符串类型，返回一个错误。	//GET KEY_NAME
	Get(ctx context.Context, key string) (string, error)
	//Del 命令用于删除已存在的键。不存在的 key 会被忽略。	//DEL KEY_NAME
	Del(ctx context.Context, key string) (int64, error)
	//MGet 命令返回所有(一个或多个)给定 key 的值。 //MGET KEY1 KEY2 .. KEYN
	MGet(ctx context.Context, keys ...string) ([]interface{}, error)
	//MSet :命令用于同时设置一个或多个 key-value 对。//MSET key1 value1 key2 value2 .. keyN valueN
	MSet(ctx context.Context, values ...interface{}) (string, error)
	//MSetNX 命令用于所有给定 key 都不存在时，同时设置一个或多个 key-value 对。//MSETNX key1 value1 key2 value2 ..
	MSetNX(ctx context.Context, values ...interface{}) (bool, error)
	//Exists 命令用于检查给定 key 是否存在。	//EXISTS KEY_NAME
	Exists(ctx context.Context, key string) (int64, error)
	//Incr 命令将 key 中储存的数字值增一,如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCR 操作 //INCR KEY_NAME
	Incr(ctx context.Context, key string) (int64, error)
	//Decr decr:命令将 key 中储存的数字值减一。//DECR KEY_NAME
	Decr(ctx context.Context, key string) (int64, error)
	//IncrBy 命令将 key 中储存的数字值增加任意数值
	IncrBy(ctx context.Context, key string, value int64) (int64, error)
	//DecrBy decrBy:命令将 key 所储存的值减去指定的减量值。	//DECRBY KEY_NAME DECREMENT_AMOUNT
	DecrBy(ctx context.Context, key string, value int64) (int64, error)
	//Expire 命令用于设置 key 的过期时间，key 过期后将不再可用。单位以秒计。//Expire KEY_NAME TIME_IN_SECONDS
	Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)
	// GetSet 命令用于设置指定 key 的值，并返回 key 的旧值。
	GetSet(ctx context.Context, key string, value interface{}) (string, error)
	// Append 命令用于为指定的 key 追加值。
	Append(ctx context.Context, key string, value string) (int64, error)

	//HGet 命令用于返回哈希表中指定字段的值。	//HGET KEY_NAME FIELD_NAME
	HGet(ctx context.Context, key string, value string) (string, error)
	//HSet 用于为哈希表中的字段赋值,如果哈希表不存在,一个新的哈希表被创建并进行HSET操作。字段已经存在,旧值被覆盖。HSET KEY_NAME FIELD VALUE
	HSet(ctx context.Context, key string, field string, data interface{}) (bool, error)
	// HSetNX 命令用于为哈希表中不存在的的字段赋值 。如果字段已经存在于哈希表中，操作无效。
	HSetNX(ctx context.Context, key string, field string, value interface{}) (bool, error)
	//HGetAll 命令用于返回哈希表中，所有的字段和值。//HGETALL KEY_NAME
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	// HDel 命令用于删除哈希表 key 中的一个或多个指定字段，不存在的字段将被忽略。
	HDel(ctx context.Context, key string, field ...string) (int64, error)
	// HExists 命令用于查看哈希表的指定字段是否存在。
	HExists(ctx context.Context, key string, field string) (bool, error)
	//Hkeys 命令用于获取哈希表中的所有域（field）
	Hkeys(ctx context.Context, key string) ([]string, error)
	//Hlen 命令用于获取哈希表中字段的数量。
	Hlen(ctx context.Context, key string) (int64, error)
	//HMGet 命令用于返回哈希表中，一个或多个给定字段的值。如果指定的字段不存在于哈希表，那么返回一个 nil 值。
	HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error)
	// HVals 命令返回哈希表所有的值。
	HVals(ctx context.Context, key string) ([]string, error)

	//SAdd 命令将一个或多个成员元素加入到集合中，已经存在于集合的成员元素将被忽略。//SADD KEY_NAME VALUE1..VALUEN
	SAdd(ctx context.Context, key string, members ...interface{}) (int64, error)
	//SRem 命令用于移除集合中的一个或多个成员元素，不存在的成员元素会被忽略。//SREM KEY MEMBER1..MEMBERN
	SRem(ctx context.Context, key string, members ...interface{}) (int64, error)
	//SMembers 命令返回集合中的所有的成员。 不存在的集合 key 被视为空集合。//SMEMBERS key
	SMembers(ctx context.Context, key string) ([]string, error)
	//SIsMembers :命令判断成员元素是否是集合的成员。//SISMEMBER KEY VALUE
	SIsMembers(ctx context.Context, key string, member string) (bool, error)
	//SRandMemberN 命令用于返回集合中的一个随机元素。//SRANDMEMBER KEY [count]
	SRandMemberN(ctx context.Context, key string, count int64) ([]string, error)

	//Llen 命令用于返回列表的长度	//LLEN KEY_NAME
	Llen(ctx context.Context, key string) (int64, error)
	//Lpush 命令将一个或多个值插入到列表头部。 如果 key 不存在，一个空列表会被创建并执行 LPUSH 操作。
	Lpush(ctx context.Context, key string, value ...interface{}) (int64, error)
	//LPushX 将一个值插入到已存在的列表头部，列表不存在时操作无效。
	LPushX(ctx context.Context, key string, value ...interface{}) (int64, error)
	//Rpush 命令用于将一个或多个值插入到列表的尾部(最右边)。//RPUSH KEY_NAME VALUE1..VALUEN
	Rpush(ctx context.Context, key string, value string) (int64, error)
	//LRange 返回列表中指定区间内的元素，区间以偏移量 START 和 END 指定。 其中 0 表示列表的第一个元素， 1表示列表的第二个元素，以此类推。
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	//RPop 命令用于移除列表的最后一个元素，返回值为移除的元素。	//RPOP KEY_NAME
	RPop(ctx context.Context, key string) (string, error)
	// BRpop 命令移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	BRpop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)
	//LPop 命令用于移除并返回列表的第一个元素。//Lpop KEY_NAME
	LPop(ctx context.Context, key string) (string, error)
	//LIndex 命令用于通过索引获取列表中的元素。你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
	LIndex(ctx context.Context, key string, index int64) (string, error)
	// LInsertBefore 命令用于在列表的元素前或者后插入元素。当指定元素不存在于列表中时，不执行任何操作。 如果 key 不是列表类型，返回一个错误。
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) (int64, error)
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) (int64, error)
	// LRem 根据参数 COUNT 的值，移除列表中与参数 VALUE 相等的元素。count>0从头开始移除count，count<0从尾移除count，count=0，移除所有
	LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error)
	// LSet 通过索引来设置元素的值。 当索引参数超出范围，或对一个空列表进行 LSET 时，返回一个错误。 LSET KEY_NAME INDEX VALUE
	LSet(ctx context.Context, key string, index int64, value interface{}) (string, error)
	//LTrim 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
	LTrim(ctx context.Context, key string, start, stop int64) (string, error)

	//Dump 命令用于序列化给定 key ，并返回被序列化的值。
	Dump(ctx context.Context, key string) (string, error)
	//Keys 命令用于查找所有符合给定模式 pattern 的 key   key*  *
	Keys(ctx context.Context, pattern string) ([]string, error)
	//Persist 命令用于移除给定 key 的过期时间，使得 key 永不过期。
	Persist(ctx context.Context, key string) (bool, error)
	//FlushDB 删除当前数据库所有 key
	FlushDB(ctx context.Context) (string, error)
	//RandomKey 命令从当前数据库中随机返回一个 key 。
	RandomKey(ctx context.Context) (string, error)
	//Rename 命令用于修改 key 的名称 。
	Rename(ctx context.Context, key, newkey string) (string, error)
	//Type 命令用于返回 key 所储存的值的类型。
	Type(ctx context.Context, key string) (string, error)
	//GetRange 命令用于获取存储在指定 key 中字符串的子字符串。字符串的截取范围由 start 和 end 两个偏移量决定(包括 start 和 end 在内)。
	GetRange(ctx context.Context, key string, start, end int64) (string, error)

	//Zadd 命令用于将一个或多个成员元素及其分数值加入到有序集当中。
	Zadd(ctx context.Context, key string, members ...*redis.Z) (int64, error)
	// ZRevrange Zrevrange 命令返回有序集中，指定区间内的成员。 其中成员的位置按分数值递减(从大到小)来排列。
	ZRevrange(ctx context.Context, key string, start, stop int64) ([]string, error)
	// ZRange 命令返回有序集中，指定区间内的成员。 其中成员的位置按分数值递增(从小到大)来排列。
	ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	//ZRank 命令可以获得成员按分数值递增(从小到大)排列的排名。
	ZRank(ctx context.Context, key, member string) (int64, error)
	//ZRevRank 命令返回有序集中成员的排名。其中有序集成员按分数值递减(从大到小)排序。 排名以 0 为底，也就是说， 分数值最大的成员排名为 0 。
	ZRevRank(ctx context.Context, key, member string) (int64, error)

	//Lock 分布式宏锁
	Lock(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	UnLock(ctx context.Context, keys []string, args ...interface{}) (interface{}, error)
}
